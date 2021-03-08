package main

import (
  "context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"strings"
  "bufio"
  "errors"
  "os"

  "github.com/jackc/pgx"
  "github.com/msteinert/pam"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/dgrijalva/jwt-go"
)

var DATABASE_URL = "postgresql://root:password@localhost/postgres"

//Global Database Connection
var databaseConnection *pgx.Conn

//Session Initialization//
//UserSession tracks user sessions and permissions
type UserSession struct {
	Username		string
	Authenticated	bool
}

//Request Struct - For each request filled by form
type Request struct {
	Name   string `json:"name"`
	NetID  string `json:"netid"`
	Email  string `json:"email"`
	Course string `json:"course"`
	Status string `json:"status"` // Unresolved | Accepted | Rejected | Archived
	Date   string `json:"date"`
}

//Instance Struct - For each existing Linux Instance
type Instance struct {
	Name	string	`json:"name"`
	NetID  string `json:"netid"`
	Status string `json:"status"` // Online | Offline | Error
	Date   string `json:"date"`
}

//Task struct for task requests sent by the frontend
type Task struct {
	Requestor	string	`json:"requestor"`
	NetID 		string	`json:"netid"`
	Task  		string 	`json:"task"`
}

type TokenResponse struct {
	Token	string `json:"token"`
}
//CORS for fetch across different ports
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// getUser returns a user from session s
// on error returns an empty user
// TODO: Session Management
func getUser(s *sessions.Session) UserSession {
	val := s.Values["user"]
	var user = UserSession{}
	user, ok := val.(UserSession)
	if !ok {
		return UserSession{Authenticated: false}
	}
	return user
}

// Returns false if the credentials were not found in the PAM stack
// Returns true otherwise
func authenticate(user string, password string) bool {
  t, err := pam.StartFunc("", user, func(s pam.Style, msg string) (string,error) {
  switch s {
      case pam.PromptEchoOff:
        return password, nil
      case pam.PromptEchoOn:
        fmt.Print(msg + " ")
        input, err := bufio.NewReader(os.Stdin).ReadString('\n')
        if err != nil {
          return "", err
        }
        return input[:len(input)-1], nil
      case pam.ErrorMsg:
        log.Print(msg)
        return "", nil
      case pam.TextInfo:
        fmt.Println(msg)
        return "", nil
    }
    return "", errors.New("Unrecognized message style")
  })
  if err != nil {
    log.Fatalf("Start: %s", err.Error())
  }
  err = t.Authenticate(0)
  if err != nil {
    return false
  }
  return true
}

func checkAuthLevel(netid string, password string) int{
  authenticated := authenticate(netid, password)
  if !authenticated {
    return 0
  }
  if netid == "vpham" || netid == "newellz2" {
    return 3
  }
  return 1
}

func verifyToken(tokenString string) (jwt.Claims, error) {
	//TODO: Change Signing Key - source from .env
	signingKey := []byte("ecc-secret")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}


func authMiddlewareRequests(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.WriteHeader(http.StatusOK)
	//Check Authentication of User
	tokenString := r.Header.Get("Authorization")
	if len(tokenString) == 0 {
		w.Write([]byte("Missing Authorization Header"))
		return
	}
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	claims, err := verifyToken(tokenString)
	if err != nil {
		w.Write([]byte("Error verifying JWT token: " + err.Error()))
		fmt.Println("Error verifying JWT token: " + err.Error())
		return
	}	
	netid := claims.(jwt.MapClaims)["netid"].(string)
  pass := claims.(jwt.MapClaims)["password"].(string)
	level := checkAuthLevel(netid, pass)
	if level == 0 {
		w.Write([]byte("Unauthorized User Header"))
		fmt.Println("Unauthorized User Header")
		return
	}
	fmt.Printf("Instance Retrieval for %s : Level %d\n", netid, level)
	r.Header.Set("netid", netid)
	//Grab Real Instances Here
	requests := getRequests(level, netid)
	json.NewEncoder(w).Encode(requests)
}

func getRequests(authLevel int, requestor string) []Request {
  var requests = []Request{}
  if authLevel < 1 {
    return requests
  } else if authLevel == 1 {
    rows, err := databaseConnection.Query(context.Background(), "select * from request where NETID=$1", requestor)
    if err != nil {
      log.Fatal(err)
    }
    for rows.Next() {
      var date time.Time
      var request Request
      err := rows.Scan(&request.Name, &request.NetID, &request.Email, &request.Course, &request.Status, &date)
      if err != nil {
        log.Fatal(err)
      }
      request.Date = date.Format("2006/02/03")
      requests = append(requests, request)
    }
  } else {
    rows, err := databaseConnection.Query(context.Background(), "select * from request")
    if err != nil {
      log.Fatal(err)
    }
    for rows.Next() {
      var date time.Time
      var request Request
      err := rows.Scan(&request.Name, &request.NetID, &request.Email, &request.Course, &request.Status, &date)
      if err != nil {
        log.Fatal(err)
      }
      request.Date = date.Format("2006/02/03")
      requests = append(requests, request)
    }
  }
  return requests
}

func getInstances(authLevel int, requestor string) []Instance {		
  var instances = []Instance{}
  if authLevel < 1{
    return instances
  } else if authLevel == 1{
    rows, err := databaseConnection.Query(context.Background(), "select * from instance where NETID=$1", requestor)
    if err != nil {
      log.Fatal(err)
    }
    for rows.Next() {
      var date time.Time
      var instance Instance
      err := rows.Scan(&instance.Name, &instance.NetID, &instance.Status, &date)
      if err != nil {
        log.Fatal(err)
      }
      instance.Date = date.Format("2006/02/03")
      instances = append(instances, instance)
    }
  } else {
    rows, err := databaseConnection.Query(context.Background(), "select * from instance")
    if err != nil {
      log.Fatal(err)
    }
    for rows.Next() {
      var date time.Time
      var instance Instance
      err := rows.Scan(&instance.Name, &instance.NetID, &instance.Status, &date)
      if err != nil {
        log.Fatal(err)
      }
      instance.Date = date.Format("2006/02/03")
      instances = append(instances, instance)
    }
  }
  return instances
}

func authMiddlewareInstances(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.WriteHeader(http.StatusOK)
	//Check Authentication of User
	tokenString := r.Header.Get("Authorization")
	fmt.Println("TokenString: ", tokenString)
	if len(tokenString) == 0 {
		w.Write([]byte("Missing Authorization Header"))
		return
	}
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	claims, err := verifyToken(tokenString)
	if err != nil {
		w.Write([]byte("Error verifying JWT token: " + err.Error()))
		fmt.Println("Error verifying JWT token: " + err.Error())
		return
	}
	netid := claims.(jwt.MapClaims)["netid"].(string)
  pass := claims.(jwt.MapClaims)["password"].(string)
	level := checkAuthLevel(netid, pass)
	if level == 0 {
		w.Write([]byte("Unauthorized User Header"))
		fmt.Println("Unauthorized User Header")
		return
	}
	fmt.Printf("Instance Retrieval for %s : Level %d", netid, level)
	r.Header.Set("netid", netid)
	//Grab Real Instances Here
	requests := getInstances(level, netid)
	json.NewEncoder(w).Encode(requests)
}

//Login Handler
func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: LoginHandler")
	r.ParseForm()
	fmt.Printf("Login: %s\n", r.FormValue("netid"))
	//TODO: heck user in PAM Stack

	authLevel := checkAuthLevel(r.FormValue("netid"), r.FormValue("password"))
	fmt.Printf("authLevel: %d\n", authLevel)

	w.Header().Set("Content-Type", "application/json")
	enableCors(&w)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["netid"] = r.FormValue("netid")
	claims["expiry"] = time.Now().Add(time.Minute * 30).Unix()
	claims["level"] = authLevel
	t, _ := token.SignedString([]byte("ecc-secret"))
	tObj := TokenResponse{Token: t}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tObj)
}

func actionHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Println("Endpoint: actionHandler")
	var instanceData Task
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&instanceData)
	requestor := instanceData.Requestor
	netid := instanceData.NetID
	task := instanceData.Task
	if requestor=="" || netid=="" || task=="" {
		return
	}
	fmt.Println("POST REQUESTOR: " + requestor)
	fmt.Println("POST NETID: " + netid)
	fmt.Println("POST TASK: " + task)
	type Response struct {
		Status	string  `json:"status"`
	}
	time.Sleep(2 * time.Second);
	json.NewEncoder(w).Encode(Response{Status: task+" finished!"})
}


func main() {
  //Database Driver Setup
  databaseConnection, _ = pgx.Connect(context.Background(), DATABASE_URL)


	router := mux.NewRouter()

	//Backend Paths
	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/login", loginHandler)
	api.HandleFunc("/instances", authMiddlewareInstances)
	api.HandleFunc("/requests", authMiddlewareRequests)
	api.HandleFunc("/action", actionHandler)

	//Server Parameters
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:3001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
