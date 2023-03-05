package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("My module")
	GetUser()

	// web router handler
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	// r.HandleFunc("/user", GetUser)
	log.Fatal(http.ListenAndServe(":4000", r))
}

func GetUser() {
	fmt.Println("Getting go users...")
}
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Go mod package tets handle a nd learning</h1>"))
}
