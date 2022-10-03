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

func NewRouter(ur *usersRep.UsersRep) *UserRouter {
	r := &UserRouter {
		Router: mux.NewRouter(),
		ur: ur,
	}

	r.HandleFunc("/feed", r.Feed)
	r.HandleFunc("/signin", r.SignIn)
	r.HandleFunc("/signup", r.SignUp)
	return r
}


func (router *UserRouter) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "incorrect http method", http.StatusMethodNotAllowed)
		return
	}

	user := usersRep.User{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	createdUser, err := router.ur.CreateUser(user)
	if err.Error() == "nickname " + user.NickName + "already in use." {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	} else if err.Error() == "user with email " + user.Email + "already exists." {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusCreated)
	}

	err = json.NewEncoder(w).Encode(Response {
										Body: createdUser,
									})
	if err != nil {
		http.Error(w, "encode error", http.StatusInternalServerError)
		return
	}
}

func (router *UserRouter) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "incorrect http method", http.StatusMethodNotAllowed)
		return
	}
	
	user := usersRep.User{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	gotUser, err := router.ur.SignIn(user)
	if err.Error() == "can't find user with email " + user.Email {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	} else if err.Error() == "incorrect password" {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}


	
	// Create a new random session token
	// we use the "github.com/google/uuid" library to generate UUIDs
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	// Set the token in the session map, along with the session information
	sessions[sessionToken] = session{
		username: creds.Username,
		expiry:   expiresAt,
	}

	// Finally, we set the client cookie for "session_token" as the session token we just generated
	// we also set an expiry time of 120 seconds
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt,
	})


	
	err = json.NewEncoder(w).Encode(Response {
										Body: gotUser,
									})
	if err != nil {
		http.Error(w, "encode error", http.StatusInternalServerError)
		return
	}
}

func (router *UserRouter) Feed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Feed")
}

