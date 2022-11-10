package middleware_test

import (
	//"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/auth/usecase/mocks"
	middlewares "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/middleware"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/pkg/csrf"

	//"github.com/go-park-mail-ru/2022_2_TikTikIVProd/pkg"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	//"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type TestCase struct {
	ArgData string
	Error error
	StatusCode int
}

func handler(c echo.Context) error {
	return nil
}

func TestMiddlewareAuth(t *testing.T) {
	var valueCookie string
	err := faker.FakeData(&valueCookie)
	assert.NoError(t, err)

	var user models.User
	err = faker.FakeData(&user)
	assert.NoError(t, err)

	mockUCase := mocks.NewUseCaseI(t)

	mockUCase.On("Auth", valueCookie).Return(&user, nil)

	middleware := middlewares.NewMiddleware(mockUCase)

	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    valueCookie,
		HttpOnly: true,
	}

	e := echo.New()

	cases := map[string]TestCase {
		"success": {
			Error: nil,
			StatusCode: http.StatusOK,
		},
		"unauthorized": {
			Error: &echo.HTTPError{
				Code: http.StatusUnauthorized,
				Message: http.ErrNoCookie.Error(),
			},
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(echo.POST, "/users", strings.NewReader(test.ArgData))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			if name == "success" {
				req.AddCookie(cookie)
			}
			
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/users")

			err := middleware.Auth(handler)(c)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.StatusCode, rec.Code)
			}
		})
	}

	mockUCase.AssertExpectations(t)
}

func TestMiddlewareCSRF(t *testing.T) {
	var valueCookie string
	err := faker.FakeData(&valueCookie)
	assert.NoError(t, err)

	var user models.User
	err = faker.FakeData(&user)
	assert.NoError(t, err)

	mockUCase := mocks.NewUseCaseI(t)

	middleware := middlewares.NewMiddleware(mockUCase)

	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    valueCookie,
		HttpOnly: true,
	}

	e := echo.New()

	cases := map[string]TestCase {
		"success": {
			Error: nil,
			StatusCode: http.StatusOK,
		},
		"unauthorized": {
			Error: &echo.HTTPError{
				Code: http.StatusUnauthorized,
				Message: http.ErrNoCookie.Error(),
			},
		},
		"empty csrf": {
			Error: &echo.HTTPError{
				Code: http.StatusForbidden,
				Message: models.ErrEmptyCsrf.Error(),
			},
		},
	}

	csrfToken, err := csrf.CreateCSRF(valueCookie)
	assert.NoError(t, err)

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(echo.POST, "/users", strings.NewReader(test.ArgData))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			if name != "unauthorized" {
				req.AddCookie(cookie)
			}
			if name != "empty csrf" {
				req.Header.Add(echo.HeaderXCSRFToken, csrfToken)
			}
			
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/users")

			err := middleware.CSRF(handler)(c)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.StatusCode, rec.Code)
			}
		})
	}

	mockUCase.AssertExpectations(t)
}

