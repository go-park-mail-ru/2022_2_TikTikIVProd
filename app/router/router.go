package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router
	/*здесь будут данные какие-то*/
}

func NewForumRouter(/*будут переданы данные какие то*/) *Router {
	r := &Router {
		Router: mux.NewRouter(),
	}

	r.HandleFunc("/feed", Feed)
	r.HandleFunc("/signin", SignIn)
	r.HandleFunc("/signup", SignUp)
	return r
}

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

