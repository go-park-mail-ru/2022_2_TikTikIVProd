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

func NewRouter(/*будут переданы данные какие то*/) *Router {
	r := &Router {
		Router: mux.NewRouter(),
	}

	r.HandleFunc("/feed", r.Feed)
	r.HandleFunc("/signin", r.SignIn)
	r.HandleFunc("/signup", r.SignUp)
	return r
}

func (router *Router) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("SignIn")
}

func (router *Router) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("SignUp")
}

func (router *Router) Feed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Feed")
}
