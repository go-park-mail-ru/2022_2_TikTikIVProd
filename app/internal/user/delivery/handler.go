package delivery

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/pkg"
)

type DeliveryI interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
	Auth(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type delivery struct {
	uc usecase.UseCaseI
}

func New(uc usecase.UseCaseI) DeliveryI {
	return &delivery{
		uc: uc,
	}
}

// SignUp godoc
// @Summary      SignUp
// @Description  user sign up
// @Tags     users
// @Accept	 application/json
// @Produce  application/json
// @Param    user body models.User true "user info"
// @Success  201 {object} pkg.Response{body=models.User} "user created"
// @Failure 405 {object} pkg.Error "invalid http method"
// @Failure 400 {object} pkg.Error "bad request"
// @Failure 409 {object} pkg.Error "nickname already in use"
// @Failure 409 {object} pkg.Error "user with this email already exists"
// @Failure 500 {object} pkg.Error "internal server error"
// @Router   /signup [post]
func (del *delivery) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	if r.Method != http.MethodPost {
		pkg.ErrorResponse(w, http.StatusMethodNotAllowed, "invalid http method")
		return
	}

	user := models.User{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		pkg.ErrorResponse(w, http.StatusBadRequest, "bad request")
		return
	}

	createdUser, createdCookie, err := del.uc.SignUp(user)
	if err != nil {
		if err.Error() == "nickname already in use" {
			pkg.ErrorResponse(w, http.StatusConflict, err.Error())
			return
		} else if err.Error() == "user with such email already exists" {
			pkg.ErrorResponse(w, http.StatusConflict, err.Error())
			return
		}
		pkg.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    createdCookie.SessionToken,
		Expires:  createdCookie.Expires,
		HttpOnly: true,
	})

	err = pkg.JSONresponse(w, http.StatusCreated, createdUser)
	if err != nil {
		pkg.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
}

// SignIn godoc
// @Summary      SignIn
// @Description  user sign in
// @Tags     users
// @Accept	 application/json
// @Produce  application/json
// @Param    user body models.UserSignIn true "user info"
// @Success  200 {object} pkg.Response{body=models.User} "success sign in"
// @Failure 405 {object} pkg.Error "invalid http method"
// @Failure 400 {object} pkg.Error "bad request"
// @Failure 404 {object} pkg.Error "user doesn't exist"
// @Failure 401 {object} pkg.Error "invalid password"
// @Failure 500 {object} pkg.Error "internal server error"
// @Router   /signin [post]
func (del *delivery) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	if r.Method != http.MethodPost {
		pkg.ErrorResponse(w, http.StatusMethodNotAllowed, "invalid http method")
		return
	}

	user := models.UserSignIn{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		pkg.ErrorResponse(w, http.StatusBadRequest, "bad request")
		return
	}

	gotUser, createdCookie, err := del.uc.SignIn(user)
	if err != nil {
		if err.Error() == "can't find user with such email" {
			pkg.ErrorResponse(w, http.StatusNotFound, err.Error())
			return
		} else if err.Error() == "invalid password" {
			pkg.ErrorResponse(w, http.StatusUnauthorized, err.Error())
			return
		}
		pkg.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    createdCookie.SessionToken,
		Expires:  createdCookie.Expires,
		HttpOnly: true,
	})

	err = pkg.JSONresponse(w, http.StatusOK, gotUser)
	if err != nil {
		pkg.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
}

// Logout godoc
// @Summary      Logout
// @Description  user logout
// @Tags     users
// @Accept	 application/json
// @Produce  application/json
// @Success  200 "success logout"
// @Failure 405 {object} pkg.Error "invalid http method"
// @Failure 400 {object} pkg.Error "bad request"
// @Failure 401 {object} pkg.Error "no cookie"
// @Router   /logout [get]
func (del *delivery) Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	if r.Method != http.MethodGet {
		pkg.ErrorResponse(w, http.StatusMethodNotAllowed, "invalid http method")
		return
	}

	cookie, err := r.Cookie("session_token")
	if err == http.ErrNoCookie {
		pkg.ErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	} else if err != nil {
		pkg.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = del.uc.DeleteCookie(cookie.Value)
	if err != nil {
		pkg.ErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now().AddDate(0, 0, -1),
	})

	w.WriteHeader(http.StatusOK)
}

// Auth godoc
// @Summary      Auth
// @Description  check user auth
// @Tags     users
// @Accept	 application/json
// @Produce  application/json
// @Success  200 {object} pkg.Response{body=models.User} "success auth"
// @Failure 405 {object} pkg.Error "invalid http method"
// @Failure 400 {object} pkg.Error "bad request"
// @Failure 500 {object} pkg.Error "internal server error"
// @Failure 401 {object} pkg.Error "no cookie"
// @Router   /auth [get]
func (del *delivery) Auth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	if r.Method != http.MethodGet {
		pkg.ErrorResponse(w, http.StatusMethodNotAllowed, "invalid http method")
		return
	}

	cookie, err := r.Cookie("session_token")
	if err == http.ErrNoCookie {
		pkg.ErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	} else if err != nil {
		pkg.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	gotCookie, err := del.uc.SelectCookie(cookie.Value)
	if err != nil {
		pkg.ErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	gotUser, err := del.uc.SelectUserById(gotCookie.UserId)
	if err != nil {
		pkg.ErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	err = pkg.JSONresponse(w, http.StatusOK, gotUser)
	if err != nil {
		pkg.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
}