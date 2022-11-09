package delivery

import (
	"net/http"
	"time"

	"github.com/pkg/errors"

	authUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/auth/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/pkg"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/pkg/csrf"
	"github.com/labstack/echo/v4"
	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/go-playground/validator.v9"
)

const session_name = "session_token"

type Delivery struct {
	AuthUC authUsecase.UseCaseI
}

// SignUp godoc
// @Summary      SignUp
// @Description  user sign up
// @Tags     auth
// @Accept	 application/json
// @Produce  application/json
// @Param    user body models.User true "user data"
// @Success 201 {object} pkg.Response{body=models.User} "user created"
// @Failure 405 {object} echo.HTTPError "Method Not Allowed"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 409 {object} echo.HTTPError "nickname already in use"
// @Failure 409 {object} echo.HTTPError "user with this email already exists"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Router   /signup [post]
func (del *Delivery) SignUp(c echo.Context) error {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	if ok, err := isRequestValid(&user); !ok {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	requestSanitizeSignUp(&user)

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
		default:
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, causeErr.Error())
		}
	}

	c.SetCookie(&http.Cookie{
		Name:     session_name,
		Value:    createdCookie.SessionToken,
		MaxAge:   createdCookie.MaxAge,
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
	var user models.UserSignIn
	err := c.Bind(&user)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	if ok, err := isRequestValid(&user); !ok {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	requestSanitizeSignIn(&user)

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
			return echo.NewHTTPError(http.StatusInternalServerError, causeErr.Error())
		}
	}

	c.SetCookie(&http.Cookie{
		Name:     session_name,
		Value:    createdCookie.SessionToken,
		MaxAge:   createdCookie.MaxAge,
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
// @Failure 403 {object} echo.HTTPError "invalid csrf"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Router   /logout [post]
func (del *Delivery) Logout(c echo.Context) error {
	cookie, err := c.Cookie(session_name)
	if err == http.ErrNoCookie {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	} else if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = del.AuthUC.DeleteCookie(cookie.Value)
	if err != nil {
		causeErr := errors.Cause(err)
		switch {
		case errors.Is(causeErr, models.ErrNotFound):
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusNotFound, models.ErrNotFound.Error())
		default:
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, causeErr.Error())
		}
	}

	c.SetCookie(&http.Cookie{
		Name:    session_name,
		Value:   "",
		Expires: time.Now().AddDate(0, 0, -1),
	})
	return c.NoContent(http.StatusNoContent)
}

// CreateCSRF godoc
// @Summary      CreateCSRF
// @Description  Get CSRF token
// @Tags         auth
// @Success      204    "success create csrf, body is empty"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 500 {object}  echo.HTTPError  "Internal server error"
// @Router /create_csrf [post]
func (del *Delivery) CreateCSRF(c echo.Context) error {
	cookie, err := c.Cookie(session_name)
	if err == http.ErrNoCookie {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	} else if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	csrfToken, err := csrf.CreateCSRF(cookie.Value)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set(echo.HeaderXCSRFToken, csrfToken)
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
	cookie, err := c.Cookie(session_name)
	if err == http.ErrNoCookie {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	} else if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	gotUser, err := del.AuthUC.Auth(cookie.Value)
	if err != nil {
		causeErr := errors.Cause(err)
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusUnauthorized, causeErr.Error())
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

func requestSanitizeSignUp(user *models.User) {
	sanitizer := bluemonday.UGCPolicy()

	user.FirstName = sanitizer.Sanitize(user.FirstName)
	user.LastName = sanitizer.Sanitize(user.LastName)
	user.NickName = sanitizer.Sanitize(user.NickName)
	user.Email = sanitizer.Sanitize(user.Email)
	user.Password = sanitizer.Sanitize(user.Password)
}

func requestSanitizeSignIn(user *models.UserSignIn) {
	sanitizer := bluemonday.UGCPolicy()

	user.Email = sanitizer.Sanitize(user.Email)
	user.Password = sanitizer.Sanitize(user.Password)
}

func NewDelivery(e *echo.Echo, uc authUsecase.UseCaseI) {
	handler := &Delivery{
		AuthUC: uc,
	}

	e.POST("/signin", handler.SignIn)
	e.POST("/signup", handler.SignUp)
	e.POST("/create_csrf", handler.CreateCSRF)
	e.POST("/logout", handler.Logout)
	e.GET("/auth", handler.Auth)
}
