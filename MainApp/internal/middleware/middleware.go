package middleware

import (
	"net/http"

	authUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/auth/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/pkg/csrf"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

const session_name = "session_token"

type middleware struct {
	authUC authUsecase.UseCaseI
}

func NewMiddleware(authUC authUsecase.UseCaseI) *middleware {
	return &middleware{authUC: authUC}
}

func (m *middleware) Auth(next echo.HandlerFunc) echo.HandlerFunc {
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

func (m *middleware) CSRF(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().URL.Path == "/create_csrf" || c.Request().URL.Path == "/signup" ||
					c.Request().URL.Path == "/signin" || c.Request().URL.Path == "/auth" ||
													c.Request().Method == http.MethodGet {
			return next(c)
		}
		
		token := c.Request().Header.Get(echo.HeaderXCSRFToken)
		if token == "" {
			c.Logger().Error(models.ErrEmptyCsrf)
			return echo.NewHTTPError(http.StatusForbidden, models.ErrEmptyCsrf.Error())
		}

		cookie, err := c.Cookie(session_name)
		if err == http.ErrNoCookie {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		} else if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		isTokenValid, err := csrf.CheckCSRF(cookie.Value, token)
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		if !isTokenValid {
			c.Logger().Error(models.ErrInvalidCsrf)
			return echo.NewHTTPError(http.StatusForbidden, models.ErrInvalidCsrf.Error())
		}

		return next(c)
	}
}

