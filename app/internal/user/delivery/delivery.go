package delivery

import (
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/model"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/pkg"
)

type DeliveryI interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
	Auth(w http.ResponseWriter, r *http.Request)
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

// SignUp godoc
// @Summary      SignUp
// @Description  user sign up
// @Tags     users
// @Accept	 application/json
// @Produce  application/json
// @Param    user body model.User true "user info"
// @Success  201 {object} model.User "user created"
// @Failure 405 {object} pkg.Error "invalid http method"
// @Failure 400 {object} pkg.Error "bad request"
// @Failure 409 {object} pkg.Error "nickname already in use"
// @Failure 409 {object} pkg.Error "user with this email already exists"
// @Failure 500 {object} pkg.Error "internal server error"
// @Router   /signup [post]
func (del *delivery) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		pkg.ErrorResponse(w, http.StatusMethodNotAllowed, "invalid http method")
		return
	}

	user := model.User{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		pkg.ErrorResponse(w, http.StatusBadRequest, "bad request")
		return
	}

	createdUser, createdCookie, err := del.uc.SignUp(user)
	if err.Error() == "nickname " + user.NickName + "already in use." {
		pkg.ErrorResponse(w, http.StatusConflict, err.Error())
		return
	} else if err.Error() == "user with email " + user.Email + "already exists." {
		pkg.ErrorResponse(w, http.StatusConflict, err.Error())
		return
	} else if err != nil {
		pkg.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   createdCookie.SessionToken,
		Expires: createdCookie.Expires,
		HttpOnly: true,
	})

	err = pkg.JSONresponse(w, http.StatusCreated, pkg.Response {
												Body: createdUser,
											})
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
// @Param    user body model.UserSignIn true "user info"
// @Success  200 {object} model.User "success sign in"
// @Failure 405 {object} pkg.Error "invalid http method"
// @Failure 400 {object} pkg.Error "bad request"
// @Failure 404 {object} pkg.Error "user doesn't exist"
// @Failure 401 {object} pkg.Error "invalid password"
// @Failure 500 {object} pkg.Error "internal server error"
// @Router   /signin [post]
func (del *delivery) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		pkg.ErrorResponse(w, http.StatusMethodNotAllowed, "invalid http method")
		return
	}
	
	user := model.UserSignIn{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		pkg.ErrorResponse(w, http.StatusBadRequest, "bad request")
		return
	}

	gotUser, createdCookie, err := del.uc.SignIn(user)
	if err.Error() == "can't find user with email " + user.Email {
		pkg.ErrorResponse(w, http.StatusNotFound, err.Error())
		return
	} else if err.Error() == "invalid password" {
		pkg.ErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	} else if err != nil {
		pkg.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   createdCookie.SessionToken,
		Expires: createdCookie.Expires,
		HttpOnly: true,
	})

	err = pkg.JSONresponse(w, http.StatusOK, pkg.Response {
												Body: gotUser,
											 })
	if err != nil {
		pkg.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
}

// Logout godoc
// @Summary      Logout
// @Description  user log out
// @Tags     users
// @Accept	 application/json
// @Produce  application/json
// @Param    user body model.UserSignIn true "user info"
// @Success  200 {object} model.User "success sign in"
// @Failure 405 {object} pkg.Error "invalid http method"
// @Failure 400 {object} pkg.Error "bad request"
// @Failure 404 {object} pkg.Error "user doesn't exist"
// @Failure 401 {object} pkg.Error "invalid password"
// @Failure 500 {object} pkg.Error "internal server error"
// @Router   /signin [post]
func (del *delivery) Logout(w http.ResponseWriter, r *http.Request) {
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

	_, err = del.uc.SelectCookie(cookie.Value)
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
// @Success  200 {object} model.User "success auth"
// @Failure 405 {object} pkg.Error "invalid http method"
// @Failure 400 {object} pkg.Error "bad request"
// @Failure 401 {object} pkg.Error "no cookie"
// @Router   /auth [get]
func (del *delivery) Auth(w http.ResponseWriter, r *http.Request) {
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

	_, err = del.uc.SelectCookie(cookie.Value)
	if err != nil {
		pkg.ErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (del *delivery) Feed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Feed")
}

