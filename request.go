package main

import (
	"encoding/json"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"
	"html/template"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/gorilla/securecookie"
	"github.com/dgrijalva/jwt-go"
)


//Session Initialization//
//UserSession tracks user sessions and permissions
type UserSession struct {
	Username		string
	Authenticated	bool
}
//Stores Cookie Data
var store *sessions.CookieStore
//Holds parsed templates
var tpl *template.Template
func init() {
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)

	store = sessions.NewCookieStore(
		authKeyOne,
		encryptionKeyOne,
	)

	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 15,
		HttpOnly: true,
	}

	gob.Register(UserSession{})

	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}
//Request Struct
type Request struct {
	Name   string `json:"name"`
	NetID  string `json:"netid"`
	Email  string `json:"email"`
	Course string `json:"course"`
	Status string `json:"status"`
	Date   string `json:"date"`
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

//Index Handler
func indexHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user := getUser(session)
	tpl.ExecuteTemplate(w, "index.gohtml", user)
}

func checkAuthLevel(netid string) int{
	if netid == "vpham" {
		return 3;
	} else if netid == "prim" {
		return 2;
	} else {
		return 0;
	}
}

func verifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte("ecc-secret")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}

func authMiddleware(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	tokenString := r.Header.Get("Authorization")
	if len(tokenString) == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Missing Authorization Header"))
		return
	}
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	claims, err := verifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error verifying JWT token: " + err.Error()))
		return
	}
	netid := claims.(jwt.MapClaims)["netid"].(string)
	r.Header.Set("netid", netid)

	//Grab Real Instances Here
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
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(requests)
}

//Login Handler
func loginHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint: LoginHandler")
	r.ParseForm()
	fmt.Printf("Test: %s\n", r.FormValue("netid"))
	fmt.Printf("Test: %s\n", r.FormValue("password"))
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

func apiMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Remote Linux Request API")
	fmt.Println("Endpoint: API Home Page")
}

func actionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: action Handler")
	//body, err := ioutil.ReadAll(r.Body)
	//defer r.Body.Close()
	//if err != nil {
	//log.Printf("Error reading body: %v", err)
	//http.Error(w, "Cannot read POST body", http.StatusBadRequest)
	//}
	var instanceData Task
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&instanceData)
	netid := instanceData.NetID
	task := instanceData.Task
	fmt.Println("POST NETID: " + netid)
	fmt.Println("POST TASK: " + task)
	enableCors(&w)
}

func testResponse(w http.ResponseWriter, r *http.Request) {
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
	enableCors(&w)
	json.NewEncoder(w).Encode(requests)
	fmt.Println("Endpoint: testResponse")
}

func main() {
	router := mux.NewRouter()

	//Authentication Paths
	//router.HandleFunc("/", indexHandler)
	//router.HandleFunc("/login", loginHandler)
	//router.HandleFunc("/logout", logoutHandler)
	//router.HandleFunc("/forbidden", loginHandler)
	//router.HandleFunc("/admin", loginHandler)

	//Backend Paths
	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/", apiMain)
	api.HandleFunc("/login", loginHandler)
	api.HandleFunc("/instances", authMiddleware)
	api.HandleFunc("/action", actionHandler)

	//Frontend Paths
	//buildHandler := http.FileServer(http.Dir("frontend/out"))
	//router.PathPrefix("/").Handler(buildHandler)
	//staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("frontend/out/_next/static")))
	//router.PathPrefix("/static/").Handler(staticHandler)

	//Server Parameters
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:3001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
