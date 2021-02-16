package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/dgrijalva/jwt-go"
)

//Session Initialization//
//UserSession tracks user sessions and permissions
type UserSession struct {
	Username		string
	Authenticated	bool
}

//Request Struct - For each request filled by form
//TODO: Implement State Pattern Status
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
}

//Task struct for task requests sent by the frontend
type Task struct {
	NetID string `json:"netid"`
	Task  string `json:"task"`
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
func getUser(s *sessions.Session) UserSession {
	val := s.Values["user"]
	var user = UserSession{}
	user, ok := val.(UserSession)
	if !ok {
		return UserSession{Authenticated: false}
	}
	return user
}

func checkAuthLevel(netid string) int{
	//TODO: Server PAM Stack Authentication
	if netid == "vpham" || netid == "newellz2" {
		return 3;
	} else if netid == "sskidmore" {
		return 2;
	} else if netid == "jakobdellosantos" || netid == "prim" {
		return 1;
	} else {
		return 0;
	}
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

func getRequests(authLevel int, requestor string) []Request {		
	//TODO: Database Retrieval
	if authLevel > 1 {
		requests := []Request{
			Request{Name: "Zachary Newell",
				NetID:  "newellz2",
				Email:  "newellz2@nevada.unr.edu",
				Course: "",
				Status: "Archived",
				Date:   "2/20/19"},
			Request{Name: "Andrew Mcintyre",
				NetID:  "amcintyre",
				Email:  "amcintyre@nevada.unr.edu",
				Course: "CS 202",
				Status: "Unresolved",
				Date:   "9/20/20"},
			Request{Name: "Vincent Pham",
				NetID:  "vpham",
				Email:  "vpham@nevada.unr.edu",
				Course: "CS 202",
				Status: "Resolved",
				Date:   "8/15/19"},
		}
		return requests
	} else {
		requests := []Request{
			Request{Name: requestor,
			NetID: requestor,
			Email: "variable@email.com",
			Course: "TEST 101",
			Status: "Resolved",
			Date: "Apple",
			},
		}
		return requests
	}
}

func authMiddlewareRequests(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.WriteHeader(http.StatusOK)
	//Check Authentication of User
	tokenString := r.Header.Get("Authorization")
	fmt.Println("TokenString: ", tokenString)
	if len(tokenString) == 0 {
		w.Write([]byte("Missing Authorization Header"))
		fmt.Println("Missing Authorization Header")
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
	level := checkAuthLevel(netid)
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

func getInstances(authLevel int, requestor string) []Instance {		
	//TODO: Database Retrieval
	if authLevel > 1 {
		instances := []Instance{
			Instance{Name: "Zachary Newell",
				NetID:  "newellz2",
				Status: "Online",},
			Instance{Name: "Andrew Mcintyre",
				NetID:  "amcintyre",
				Status: "Offline",},
			Instance{Name: "Vincent Pham",
				NetID:  "vpham",
				Status: "Error",},
		}
		return instances
	} else {
		instances := []Instance{
			Instance{Name: requestor,
			NetID: requestor,
			Status: "Resolved",
			},
		}
		return instances
	}
}

func authMiddlewareInstances(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.WriteHeader(http.StatusOK)
	//Check Authentication of User
	tokenString := r.Header.Get("Authorization")
	fmt.Println("TokenString: ", tokenString)
	if len(tokenString) == 0 {
		w.Write([]byte("Missing Authorization Header"))
		fmt.Println("Missing Authorization Header")
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
	level := checkAuthLevel(netid)
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

	authLevel := checkAuthLevel(r.FormValue("netid"))
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
	netid := instanceData.NetID
	task := instanceData.Task
	//TODO: Server Side Tasks
	fmt.Println("POST NETID: " + netid)
	fmt.Println("POST TASK: " + task)
	type Response struct {
		Status	string  `json:"status"`
	}
	time.Sleep(2 * time.Second);
	json.NewEncoder(w).Encode(Response{Status: task+" finished!"})
}

func main() {
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
