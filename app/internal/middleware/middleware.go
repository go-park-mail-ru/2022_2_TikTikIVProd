package middleware

import (
	"net/http"

	authUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/auth/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/pkg/csrf"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

const session_name = "session_token"

type Middleware struct {
	authUC authUsecase.UseCaseI
}

func NewMiddleware(authUC authUsecase.UseCaseI) *Middleware {
	return &Middleware{authUC: authUC}
}

func (m *Middleware) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().URL.Path == "/signup" || c.Request().URL.Path == "/signin" ||
															c.Request().URL.Path == "/auth" {
			return next(c)
		}

		cookie, err := c.Cookie(session_name)
		if err == http.ErrNoCookie {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		} else if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		user, err := m.authUC.Auth(cookie.Value)
		if err != nil {
			causeErr := errors.Cause(err)
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusUnauthorized, causeErr.Error())
		}

		c.Set("user_id", user.Id)

		return next(c)
	}
}

func (m *Middleware) CSRF(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().URL.Path == "/create_csrf" || c.Request().URL.Path == "/signup" ||
					c.Request().URL.Path == "/signin" || c.Request().URL.Path == "/auth" ||
													c.Request().Method == http.MethodGet {
			return next(c)
		}
		
		token := c.Request().Header.Get(echo.HeaderXCSRFToken)
		if token == "" {
			err := errors.New("empty csrf token")
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}

		sess, err := session.Get(session_name, c)
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest)
		} else if sess.IsNew {
			c.Logger().Error(models.ErrUnauthorized)
			return echo.NewHTTPError(http.StatusUnauthorized, models.ErrUnauthorized)
		}

		isTokenValid, err := csrf.CheckCSRF(sess.ID, token)
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		if !isTokenValid {
			err := errors.New("invalid csrf")
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}

		return next(c)
	}
}

