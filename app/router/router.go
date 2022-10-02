package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	usersRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/repository"
)

type Router struct {
	*mux.Router
	ur *usersRep.UsersRep
}

func NewRouter(ur *usersRep.UsersRep) *Router {
	r := &Router {
		Router: mux.NewRouter(),
		ur: ur,
	}

	r.HandleFunc("/feed", r.Feed)
	r.HandleFunc("/signin", r.SignIn)
	r.HandleFunc("/signup", r.SignUp)
	return r
}

func (router *Router) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("SignUp")
}

func (router *Router) Feed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Feed")
}

func (router *Router) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("SignIn")
}

