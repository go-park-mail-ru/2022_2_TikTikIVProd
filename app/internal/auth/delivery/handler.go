package delivery

import (
	"net/http"
	"time"
	"github.com/pkg/errors"

	authUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/auth/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/pkg"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/middleware"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

// type DeliveryI interface {
// 	SignUp(c echo.Context) error
// 	SignIn(c echo.Context) error
// 	Auth(c echo.Context) error
// 	Logout(c echo.Context) error
// }

type Delivery struct {
	AuthUC authUsecase.UseCaseI
}

// SignUp godoc
// @Summary      SignUp
// @Description  user sign up
// @Tags     auth
// @Accept	 application/json
// @Produce  application/json
// @Param    user body models.User true "user info"
// @Success 201 {object} pkg.Response{body=models.User} "user created"
// @Failure 405 {object} echo.HTTPError "Method Not Allowed"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 409 {object} echo.HTTPError "nickname already in use"
// @Failure 409 {object} echo.HTTPError "user with this email already exists"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Router   /signup [post]
func (del *Delivery) SignUp(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderAccessControlAllowMethods, echo.POST)

	var user models.User
	err := c.Bind(&user); if err != nil {
		c.Logger().Error(err)
		//return c.JSON(http.StatusBadRequest, pkg.Response{Body: user})
		//c.Response().Status = http.StatusBadRequest
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}
	
	if ok, err := isRequestValid(&user); !ok {
		c.Logger().Error(err)
		//c.Response().Status = http.StatusBadRequest
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}

	createdCookie, err := del.AuthUC.SignUp(&user)
	if err != nil {
		causeErr := errors.Cause(err)
		switch {
		case errors.Is(causeErr, models.ErrConflictNickname):
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusConflict, models.ErrConflictNickname.Error())
		case errors.Is(causeErr, models.ErrConflictEmail):
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusConflict, models.ErrConflictEmail.Error())
		case errors.Is(causeErr, models.ErrBadRequest):
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusConflict, models.ErrBadRequest.Error())
		default:
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	c.SetCookie(&http.Cookie{
		Name:     "session_token",
		Value:    createdCookie.SessionToken,
		Expires:  createdCookie.Expires,
		HttpOnly: true,
	})

	return c.JSON(http.StatusCreated, pkg.Response{Body: user})
}

// SignIn godoc
// @Summary      SignIn
// @Description  user sign in
// @Tags     auth
// @Accept	 application/json
// @Produce  application/json
// @Param    user body models.UserSignIn true "user info"
// @Success  200 {object} pkg.Response{body=models.User} "success sign in"
// @Failure 405 {object} echo.HTTPError "Method Not Allowed"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 404 {object} echo.HTTPError "user doesn't exist"
// @Failure 401 {object} echo.HTTPError "invalid password"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Router   /signin [post]
func (del *Delivery) SignIn(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderAccessControlAllowMethods, echo.POST)

	var user models.UserSignIn
	err := c.Bind(&user); if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}

	if ok, err := isRequestValid(&user); !ok {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}

	gotUser, createdCookie, err := del.AuthUC.SignIn(user)
	if err != nil {
		causeErr := errors.Cause(err)
		switch {
		case errors.Is(causeErr, models.ErrNotFound):
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusNotFound, models.ErrNotFound.Error())
		case errors.Is(causeErr, models.ErrInvalidPassword):
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusUnauthorized, models.ErrInvalidPassword.Error())
		default:
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	c.SetCookie(&http.Cookie{
		Name:     "session_token",
		Value:    createdCookie.SessionToken,
		Expires:  createdCookie.Expires,
		HttpOnly: true,
	})

	return c.JSON(http.StatusOK, pkg.Response{Body: gotUser})
}

// Logout godoc
// @Summary      Logout
// @Description  user logout
// @Tags     auth
// @Produce  application/json
// @Success  204 "success logout, body is empty"
// @Failure 405 {object} echo.HTTPError "Method Not Allowed"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Router   /logout [delete]
func (del *Delivery) Logout(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderAccessControlAllowMethods, echo.DELETE)

	cookie, err := c.Cookie("session_token")
	if err == http.ErrNoCookie {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	} else if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = del.AuthUC.DeleteCookie(cookie.Value)   //мб обрабатывать NotFound???
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	c.SetCookie(&http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now().AddDate(0, 0, -1),
	})

	return c.NoContent(http.StatusNoContent)
}

// Auth godoc
// @Summary      Auth
// @Description  check user auth
// @Tags     auth
// @Produce  application/json
// @Success  200 {object} pkg.Response{body=models.User} "success auth"
// @Failure 405 {object} echo.HTTPError "Method Not Allowed"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Router   /auth [get]
func (del *Delivery) Auth(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderAccessControlAllowMethods, echo.GET)

	cookie, err := c.Cookie("session_token")
	if err == http.ErrNoCookie {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	} else if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	gotUser, err := del.AuthUC.Auth(cookie.Value)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, pkg.Response{Body: gotUser})
}

func isRequestValid(user interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewDelivery(e *echo.Echo, au authUsecase.UseCaseI, authMid *middleware.Middleware) {
	handler := &Delivery{
		AuthUC: au,
	}

	e.POST("/signin", handler.SignIn)
	e.POST("/signup", handler.SignUp)
	e.GET("/auth", handler.Auth)
	e.DELETE("/logout", handler.Logout, authMid.Auth)
}
