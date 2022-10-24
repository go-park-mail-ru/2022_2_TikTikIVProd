package delivery

import (
	"encoding/json"
	"net/http"
	"time"

	authUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/auth/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/pkg"
	"github.com/labstack/echo/v4"
)

type DeliveryI interface {
	SignUp(c echo.Context) error
	SignIn(c echo.Context) error
	Auth(c echo.Context) error
	Logout(c echo.Context) error
}

type delivery struct {
	authUC authUsecase.UseCaseI
}

// SignUp godoc
// @Summary      SignUp
// @Description  user sign up
// @Tags     auth
// @Accept	 application/json
// @Produce  application/json
// @Param    user body models.User true "user info"
// @Success 201 {object} pkg.Response{body=models.User} "user created"
// @Failure 405 {object} pkg.Error "invalid http method"
// @Failure 400 {object} pkg.Error "bad request"
// @Failure 409 {object} pkg.Error "nickname already in use"
// @Failure 409 {object} pkg.Error "user with this email already exists"
// @Failure 500 {object} pkg.Error "internal server error"
// @Router   /signup [post]
func (del *delivery) SignUp(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderAccessControlAllowMethods, echo.POST)
	if c.Request().Method != http.MethodPost {
		c.Logger().Error("invalid http method")
		return pkg.ErrorResponse(c.Response(), http.StatusMethodNotAllowed, "invalid http method")
	}

	user := models.User{}

	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		c.Logger().Error(err.Error())
		return pkg.ErrorResponse(c.Response(), http.StatusBadRequest, "bad request")
	}

	createdUser, createdCookie, err := del.authUC.SignUp(user)
	if err != nil {
		switch err.Error() {
		case "nickname already in use":
			c.Logger().Error(err.Error())
			return pkg.ErrorResponse(c.Response(), http.StatusConflict, err.Error())
		case "user with such email already exists":
			c.Logger().Error(err.Error())
			return pkg.ErrorResponse(c.Response(), http.StatusConflict, err.Error())
		default:
			c.Logger().Error(err.Error())
			return pkg.ErrorResponse(c.Response(), http.StatusInternalServerError, err.Error())
		}
	}

	http.SetCookie(c.Response(), &http.Cookie{
		Name:     "session_token",
		Value:    createdCookie.SessionToken,
		Expires:  createdCookie.Expires,
		HttpOnly: true,
	})

	err = pkg.JSONresponse(c.Response(), http.StatusCreated, createdUser)
	if err != nil {
		c.Logger().Error(err.Error())
		return pkg.ErrorResponse(c.Response(), http.StatusInternalServerError, err.Error())
	}

	return nil
}

// SignIn godoc
// @Summary      SignIn
// @Description  user sign in
// @Tags     auth
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
func (del *delivery) SignIn(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderAccessControlAllowMethods, echo.POST)
	if c.Request().Method != http.MethodPost {
		return pkg.ErrorResponse(c.Response(), http.StatusMethodNotAllowed, "invalid http method")
	}

	user := models.UserSignIn{}

	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		return pkg.ErrorResponse(c.Response(), http.StatusBadRequest, "bad request")
	}

	gotUser, createdCookie, err := del.authUC.SignIn(user)
	if err != nil {
		switch err.Error() {
		case "can't find user with such email":
			c.Logger().Error(err.Error())
			return pkg.ErrorResponse(c.Response(), http.StatusNotFound, err.Error())
		case "invalid password":
			c.Logger().Error(err.Error())
			return pkg.ErrorResponse(c.Response(), http.StatusUnauthorized, err.Error())
		default:
			c.Logger().Error(err.Error())
			return pkg.ErrorResponse(c.Response(), http.StatusInternalServerError, err.Error())
		}
	}

	http.SetCookie(c.Response(), &http.Cookie{
		Name:     "session_token",
		Value:    createdCookie.SessionToken,
		Expires:  createdCookie.Expires,
		HttpOnly: true,
	})

	err = pkg.JSONresponse(c.Response(), http.StatusOK, gotUser)
	if err != nil {
		return pkg.ErrorResponse(c.Response(), http.StatusInternalServerError, err.Error())
	}
	return nil
}

// Logout godoc
// @Summary      Logout
// @Description  user logout
// @Tags     auth
// @Produce  application/json
// @Success  204 "success logout, body is empty"
// @Failure 405 {object} pkg.Error "invalid http method"
// @Failure 400 {object} pkg.Error "bad request"
// @Failure 401 {object} pkg.Error "no cookie"
// @Router   /logout [delete]
func (del *delivery) Logout(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderAccessControlAllowMethods, echo.DELETE)
	if c.Request().Method != http.MethodDelete {
		return pkg.ErrorResponse(c.Response(), http.StatusMethodNotAllowed, "invalid http method")
	}

	cookie, err := c.Request().Cookie("session_token")
	if err == http.ErrNoCookie {
		return pkg.ErrorResponse(c.Response(), http.StatusUnauthorized, err.Error())
	} else if err != nil {
		return pkg.ErrorResponse(c.Response(), http.StatusBadRequest, err.Error())
	}

	err = del.authUC.DeleteCookie(cookie.Value)
	if err != nil {
		return pkg.ErrorResponse(c.Response(), http.StatusUnauthorized, err.Error())
	}

	http.SetCookie(c.Response(), &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now().AddDate(0, 0, -1),
	})

	c.Response().WriteHeader(http.StatusNoContent)
	return nil
}

// Auth godoc
// @Summary      Auth
// @Description  check user auth
// @Tags     auth
// @Produce  application/json
// @Success  200 {object} pkg.Response{body=models.User} "success auth"
// @Failure 405 {object} pkg.Error "invalid http method"
// @Failure 400 {object} pkg.Error "bad request"
// @Failure 500 {object} pkg.Error "internal server error"
// @Failure 401 {object} pkg.Error "no cookie"
// @Router   /auth [get]
func (del *delivery) Auth(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderAccessControlAllowMethods, echo.GET)
	if c.Request().Method != http.MethodGet {
		return pkg.ErrorResponse(c.Response(), http.StatusMethodNotAllowed, "invalid http method")
	}

	cookie, err := c.Request().Cookie("session_token")
	if err == http.ErrNoCookie {
		return pkg.ErrorResponse(c.Response(), http.StatusUnauthorized, err.Error())
	} else if err != nil {
		return pkg.ErrorResponse(c.Response(), http.StatusBadRequest, err.Error())
	}

	gotUser, err := del.authUC.Auth(cookie.Value)
	if err != nil {
		return pkg.ErrorResponse(c.Response(), http.StatusUnauthorized, err.Error())
	}

	err = pkg.JSONresponse(c.Response(), http.StatusOK, gotUser)
	if err != nil {
		return pkg.ErrorResponse(c.Response(), http.StatusInternalServerError, err.Error())
	}
	return nil
}

func NewDelivery(e *echo.Echo, au authUsecase.UseCaseI) {
	handler := &delivery{
		authUC: au,
	}

	e.POST("/signin", handler.SignIn)
	e.POST("/signup", handler.SignUp)
	e.GET("/auth", handler.Auth)
	e.DELETE("/logout", handler.Logout)
}
