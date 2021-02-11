package main;

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	Name	string			`json:"name"`
	NetID 	string			`json:"netid"`
	Email   string          `json:"email"`
	Course	string			`json:"course"`
	Status	string			`json:"status"`
	Date	string			`json:"date"`
}

//CORS for fetch across different ports
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Remote Linux Request API")
	fmt.Println("Endpoint: homePage")
}

func testResponse(w http.ResponseWriter, r *http.Request){
	requests := []Request{
		Request{Name: "Zachary Newell",
			NetID: "newellz2",
			Email: "newellz2@nevada.unr.edu",
			Course: "",
			Status: "Archived",
			Date: "2/20/19"},
		Request{Name: "Andrew Mcintyre",
			NetID: "amcintyre",
			Email: "amcintyre@nevada.unr.edu",
			Course: "CS 202",
			Status: "Unresolved",
			Date: "9/20/20"},
		Request{Name: "Vincent Pham",
			NetID: "vpham",
			Email: "vpham@nevada.unr.edu",
			Course: "CS 202",
			Status: "Resolved",
			Date: "8/15/19"},
	}
	enableCors(&w);
	json.NewEncoder(w).Encode(requests)
	fmt.Println("Endpoint: testResponse")
}

func handleRequests() {
	http.HandleFunc("", homePage)
	http.HandleFunc("/test", testResponse)
	fmt.Println(http.ListenAndServe(":3001", nil))
}

func main() {
	handleRequests()
}