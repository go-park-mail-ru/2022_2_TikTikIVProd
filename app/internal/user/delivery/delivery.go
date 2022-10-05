package delivery

import (
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/model"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/usecase"
)

type Response struct {
	Body interface{} `json:"body"`
}

// type MessageError struct {
// 	Message string `json:"error"`
// }

type DeliveryI interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
	Feed(w http.ResponseWriter, r *http.Request)
}

type delivery struct {
	uc usecase.UseCaseI
}

func New(uc usecase.UseCaseI) DeliveryI {
	return &delivery {
		uc:     uc,
	}
}

// SignIn godoc
// @Summary      SignUp
// @Description  user sign up
// @Tags     auth
// @Accept	 application/json
// @Produce  application/json
// @Param    user body model.User true "user info"
// @Success  200 {object} Response
// @Router   /signin [post]
func (del *delivery) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "incorrect http method", http.StatusMethodNotAllowed)
		return
	}

	user := model.User{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	createdUser, err := del.uc.CreateUser(user)
	if err.Error() == "nickname " + user.NickName + "already in use." {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	} else if err.Error() == "user with email " + user.Email + "already exists." {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	createdCookie, err := del.uc.CreateCookie(createdUser.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   createdCookie.SessionToken,
		Expires: createdCookie.Expires,
		HttpOnly: true,
	})

	err = json.NewEncoder(w).Encode(Response {
										Body: createdUser,
									})
	if err != nil {
		http.Error(w, "encode error", http.StatusInternalServerError)
		return
	}
}


// SignIn godoc
// @Summary      SignIn
// @Description  user sign in
// @Tags     auth
// @Accept	 application/json
// @Produce  application/json
// @Param    user body model.User true "user info"
// @Success  200 {object} model.User
// @Router   /signin [post]
func (del *delivery) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "incorrect http method", http.StatusMethodNotAllowed)
		return
	}
	
	user := model.User{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	gotUser, err := del.uc.SignIn(user)
	if err.Error() == "can't find user with email " + user.Email {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	} else if err.Error() == "incorrect password" {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	createdCookie, err := del.uc.CreateCookie(gotUser.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   createdCookie.SessionToken,
		Expires: createdCookie.Expires,
		HttpOnly: true,
	})
}

func (del *delivery) Auth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "incorrect http method", http.StatusMethodNotAllowed)
		return
	}

	cookie, err := r.Cookie("session_token")
	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = del.uc.SelectCookie(cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (del *delivery) Feed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Feed")
}

