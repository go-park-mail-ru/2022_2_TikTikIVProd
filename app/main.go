package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("SignIn")
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("SignUp")
}

func Feed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Feed")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/feed", Feed)
	r.HandleFunc("/signin", SignIn)
	r.HandleFunc("/signup", SignUp)

	log.Println("start serving :8080")
	http.ListenAndServe(":8080", r)
}
