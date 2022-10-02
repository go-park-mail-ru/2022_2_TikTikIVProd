package router

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"

	usersRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/repository"
)

type Response struct {
	Body interface{} `json:"body"`
}

type MessageError struct {
	Message string `json:"error"`
}

type UserRouter struct {
	*mux.Router
	ur *usersRep.UsersRep
}

type Router struct {
	UserRouter
}

func NewRouter(ur *usersRep.UsersRep) *RouterUser {
	r := &Router {
		Router: mux.NewRouter(),
		ur: ur,
	}

	r.HandleFunc("/feed", r.Feed)
	r.HandleFunc("/signin", r.SignIn)
	r.HandleFunc("/signup", r.SignUp)
	return r
}


func (router *RouterUser) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed) //Надо ли?
		json.NewEncoder(w).Encode(MessageError {
			Message: "incorrect http method",
		})
		http.Error(w, "bad request", http.StatusMethodNotAllowed)
		return
	}

	user := usersRep.User{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //Надо ли?
		json.NewEncoder(w).Encode(MessageError {
			Message: err.Error(),
		})
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	createdUser, err := router.ur.CreateUser(user)
	if err.Error() == "nickname " + user.NickName + "already in use." {
		w.WriteHeader(http.StatusConflict) //Надо ли?
		json.NewEncoder(w).Encode(MessageError {
										Message: err.Error(),
									})
		http.Error(w, "nickname already in use", http.StatusConflict)
		return
	} else if err.Error() == "user with email " + user.Email + "already exists." {
		w.WriteHeader(http.StatusConflict) //Надо ли?
		json.NewEncoder(w).Encode(MessageError {
										Message: err.Error(),
									})
		http.Error(w, "user exists", http.StatusConflict)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError) //Надо ли?
		json.NewEncoder(w).Encode(MessageError {
			Message: err.Error(),
		})
		http.Error(w, "not created", http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusCreated)
	}

	err = json.NewEncoder(w).Encode(Response {
										Body: createdUser,
									})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) //Надо ли?
		http.Error(w, "encode error", http.StatusInternalServerError)
		return
	}
}



func (router *UserRouter) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed) //Надо ли?
		json.NewEncoder(w).Encode(MessageError {
			Message: "incorrect http method",
		})
		http.Error(w, "bad request", http.StatusMethodNotAllowed)
		return
	}
	
	user := usersRep.User{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(MessageError {
			Message: err.Error(),
		})
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	gotUser, err := router.ur.SignIn(user)
	if err.Error() == "can't find user with email " + user.Email {
		w.WriteHeader(http.StatusConflict) //Надо ли?
		json.NewEncoder(w).Encode(MessageError {
										Message: err.Error(),
									})
		http.Error(w, "user doesn't exist", http.StatusConflict)
		return
	} else if err.Error() == "incorrect password" {
		w.WriteHeader(http.StatusUnauthorized) //Надо ли?
		json.NewEncoder(w).Encode(MessageError {
										Message: err.Error(),
									})
		http.Error(w, "user doesn't exist", http.StatusUnauthorized)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(MessageError {
			Message: err.Error(),
		})
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
	
	err = json.NewEncoder(w).Encode(Response {
										Body: gotUser,
									})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) //Надо ли?
		http.Error(w, "encode error", http.StatusInternalServerError)
		return
	}
}

func (router *Router) Feed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Feed")
}

