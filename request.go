package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//spaHandler implements http.Handler interface so we can respond
type spaHandler struct {
	staticPath string
	indexPath  string
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

//CORS for fetch across different ports
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
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

	//Backend Paths
	router.HandleFunc("/api", apiMain)
	router.HandleFunc("/api/test", testResponse)
	router.HandleFunc("/api/action", actionHandler)

	//Frontend Paths
	buildHandler := http.FileServer(http.Dir("frontend/out"))
	router.PathPrefix("/").Handler(buildHandler)
	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("frontend/out/_next/static")))
	router.PathPrefix("/static/").Handler(staticHandler)

	//Server Parameters
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:3001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
