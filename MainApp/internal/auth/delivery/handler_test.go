package delivery_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/bxcodec/faker"
	authDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/auth/delivery"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/auth/usecase/mocks"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/pkg"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type TestCaseSignUp struct {
	ArgData string
	ExpectedResponse string
	Error error
	StatusCode int
}

type TestCase struct {
	ArgData string
	Error error
	StatusCode int
}

func TestDeliverySignUp(t *testing.T) {
	var mockUser models.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)

	var mockUserConflictNickName models.User
	err = faker.FakeData(&mockUserConflictNickName)
	assert.NoError(t, err)

	var mockUserConflictEmail models.User
	err = faker.FakeData(&mockUserConflictEmail)
	assert.NoError(t, err)

	var mockUserInternalErr models.User
	err = faker.FakeData(&mockUserInternalErr)
	assert.NoError(t, err)

	mockUserInvalid := models.User{}

	mockUser.Id = 1
	mockUser.CreatedAt = time.Date(2022, time.September, 5, 1, 12, 12, 12, time.Local)

	mockUserConflictNickName.Id = 2
	mockUserConflictNickName.CreatedAt = time.Date(2022, time.September, 5, 1, 12, 12, 12, time.Local)

	mockUserConflictEmail.Id = 3
	mockUserConflictEmail.CreatedAt = time.Date(2022, time.September, 5, 1, 12, 12, 12, time.Local)
	
	mockUserInternalErr.Id = 4
	mockUserInternalErr.CreatedAt = time.Date(2022, time.September, 5, 1, 12, 12, 12, time.Local)

	var mockCookie models.Cookie
	err = faker.FakeData(&mockCookie)
	assert.NoError(t, err)

	jsonUser, err := json.Marshal(mockUser)
	assert.NoError(t, err)

	jsonUserInvalid, err := json.Marshal(mockUserInvalid)
	assert.NoError(t, err)

	jsonUserConflictNickName, err := json.Marshal(mockUserConflictNickName)
	assert.NoError(t, err)

	jsonUserConflictEmail, err := json.Marshal(mockUserConflictEmail)
	assert.NoError(t, err)

	jsonUserInternalErr, err := json.Marshal(mockUserInternalErr)
	assert.NoError(t, err)

	mockUCase := mocks.NewUseCaseI(t)

	mockUCase.On("SignUp", &mockUser).Return(&mockCookie, nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*models.User)
		arg.Id = mockCookie.UserId
	})

	mockUCase.On("SignUp", &mockUserConflictNickName).Return(nil, models.ErrConflictNickname)

	mockUCase.On("SignUp", &mockUserConflictEmail).Return(nil, models.ErrConflictEmail)

	mockUCase.On("SignUp", &mockUserInternalErr).Return(nil, models.ErrInternalServerError)

	handler := authDelivery.Delivery{
		AuthUC: mockUCase,
	}

	userResponse := mockUser
	userResponse.Id = mockCookie.UserId
	response := pkg.Response {
		Body: userResponse,
	}
	jsonResponse, err := json.Marshal(response)
	assert.NoError(t, err)

	e := echo.New()
	authDelivery.NewDelivery(e, mockUCase)

	cases := map[string]TestCaseSignUp {
		"success": {
			ArgData:   string(jsonUser),
			ExpectedResponse: string(jsonResponse) + "\n",
			Error: nil,
			StatusCode: http.StatusCreated,
		},
		"bad_request": {
			ArgData:   "aaa",
			Error: &echo.HTTPError{
				Code: http.StatusBadRequest,
				Message: models.ErrBadRequest.Error(),
			},
		},
		"invalid_request": {
			ArgData:   string(jsonUserInvalid),
			Error: &echo.HTTPError{
				Code: http.StatusBadRequest,
				Message: models.ErrBadRequest.Error(),
			},
		},
		"conflict_nickname": {
			ArgData:   string(jsonUserConflictNickName),
			Error: &echo.HTTPError{
				Code: http.StatusConflict,
				Message: models.ErrConflictNickname.Error(),
			},
		},
		"conflict_email": {
			ArgData:   string(jsonUserConflictEmail),
			Error: &echo.HTTPError{
				Code: http.StatusConflict,
				Message: models.ErrConflictEmail.Error(),
			},
		},
		"internal_error": {
			ArgData:   string(jsonUserInternalErr),
			Error: &echo.HTTPError{
				Code: http.StatusInternalServerError,
				Message: models.ErrInternalServerError.Error(),
			},
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(echo.POST, "/signup", strings.NewReader(test.ArgData))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/signup")

			err = handler.SignUp(c)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.StatusCode, rec.Code)
				assert.Equal(t, test.ExpectedResponse, rec.Body.String())
			}
		})
	}

	mockUCase.AssertExpectations(t)
}

func TestDeliverySignIn(t *testing.T) {
	var mockUserSignIn models.UserSignIn
	err := faker.FakeData(&mockUserSignIn)
	assert.NoError(t, err)

	jsonUser, err := json.Marshal(mockUserSignIn)
	assert.NoError(t, err)

	mockUserInvalid := models.UserSignIn{}

	jsonUserInvalid, err := json.Marshal(mockUserInvalid)
	assert.NoError(t, err)

	var mockUserInvalidPassword models.UserSignIn
	err = faker.FakeData(&mockUserInvalidPassword)
	assert.NoError(t, err)

	jsonUserInvalidPassword, err := json.Marshal(mockUserInvalidPassword)
	assert.NoError(t, err)

	var mockUserIternalErr models.UserSignIn
	err = faker.FakeData(&mockUserIternalErr)
	assert.NoError(t, err)

	jsonUserInternalErr, err := json.Marshal(mockUserIternalErr)
	assert.NoError(t, err)

	var mockUserNotFound models.UserSignIn
	err = faker.FakeData(&mockUserNotFound)
	assert.NoError(t, err)

	jsonUserNotFound, err := json.Marshal(mockUserNotFound)
	assert.NoError(t, err)

	mockUCase := mocks.NewUseCaseI(t)

	var mockCookie models.Cookie
	err = faker.FakeData(&mockCookie)
	assert.NoError(t, err)

	var mockUser models.User
	err = faker.FakeData(&mockUser)
	assert.NoError(t, err)

	mockUCase.On("SignIn", mockUserSignIn).Return(&mockUser, &mockCookie, nil)
	mockUCase.On("SignIn", mockUserInvalidPassword).Return(nil, nil, models.ErrInvalidPassword)
	mockUCase.On("SignIn", mockUserIternalErr).Return(nil, nil, models.ErrInternalServerError)
	mockUCase.On("SignIn", mockUserNotFound).Return(nil, nil, models.ErrNotFound)

	handler := authDelivery.Delivery{
		AuthUC: mockUCase,
	}

	e := echo.New()
	authDelivery.NewDelivery(e, mockUCase)

	cases := map[string]TestCase {
		"success": {
			ArgData:   string(jsonUser),
			Error: nil,
			StatusCode: http.StatusOK,
		},
		"bad_request": {
			ArgData:   "aaa",
			Error: &echo.HTTPError{
				Code: http.StatusBadRequest,
				Message: models.ErrBadRequest.Error(),
			},
		},
		"invalid_request": {
			ArgData:   string(jsonUserInvalid),
			Error: &echo.HTTPError{
				Code: http.StatusBadRequest,
				Message: models.ErrBadRequest.Error(),
			},
		},
		"not_found": {
			ArgData:   string(jsonUserNotFound),
			Error: &echo.HTTPError{
				Code: http.StatusNotFound,
				Message: models.ErrNotFound.Error(),
			},
		},
		"invalid_password": {
			ArgData:   string(jsonUserInvalidPassword),
			Error: &echo.HTTPError{
				Code: http.StatusUnauthorized,
				Message: models.ErrInvalidPassword.Error(),
			},
		},
		"internal_error": {
			ArgData:   string(jsonUserInternalErr),
			Error: &echo.HTTPError{
				Code: http.StatusInternalServerError,
				Message: models.ErrInternalServerError.Error(),
			},
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(echo.POST, "/signin", strings.NewReader(test.ArgData))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/signin")

			err = handler.SignIn(c)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.StatusCode, rec.Code)
			}
		})
	}

	mockUCase.AssertExpectations(t)
}

func TestDeliveryLogout(t *testing.T) {
	var valueCookie string
	err := faker.FakeData(&valueCookie)
	assert.NoError(t, err)

	var valueCookieNotFound string
	err = faker.FakeData(&valueCookieNotFound)
	assert.NoError(t, err)

	var valueCookieInternalErr string
	err = faker.FakeData(&valueCookieInternalErr)
	assert.NoError(t, err)

	mockUCase := mocks.NewUseCaseI(t)

	mockUCase.On("DeleteCookie", valueCookie).Return(nil)
	mockUCase.On("DeleteCookie", valueCookieNotFound).Return(models.ErrNotFound)
	mockUCase.On("DeleteCookie", valueCookieInternalErr).Return(models.ErrInternalServerError)

	handler := authDelivery.Delivery{
		AuthUC: mockUCase,
	}

	e := echo.New()
	authDelivery.NewDelivery(e, mockUCase)

	cases := map[string]TestCase {
		"success": {
			ArgData: valueCookie,
			Error: nil,
			StatusCode: http.StatusNoContent,
		},
		"unauthorized": {
			Error: &echo.HTTPError{
				Code: http.StatusUnauthorized,
				Message: http.ErrNoCookie.Error(),
			},
		},
		"not_found": {
			ArgData:   valueCookieNotFound,
			Error: &echo.HTTPError{
				Code: http.StatusNotFound,
				Message: models.ErrNotFound.Error(),
			},
		},
		"internal_error": {
			ArgData:   valueCookieInternalErr,
			Error: &echo.HTTPError{
				Code: http.StatusInternalServerError,
				Message: models.ErrInternalServerError.Error(),
			},
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(echo.POST, "/logout", strings.NewReader(""))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			if name != "unauthorized" {
				cookie := &http.Cookie{
					Name:     "session_token",
					Value:    test.ArgData,
					HttpOnly: true,
				}

				req.AddCookie(cookie)
			}
			
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/logout")

			err := handler.Logout(c)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.StatusCode, rec.Code)
			}
		})
	}

	mockUCase.AssertExpectations(t)
}

func TestDeliveryAuth(t *testing.T) {
	var valueCookie string
	err := faker.FakeData(&valueCookie)
	assert.NoError(t, err)

	var valueCookieUnauthorized string
	err = faker.FakeData(&valueCookieUnauthorized)
	assert.NoError(t, err)

	var user models.User
	err = faker.FakeData(&user)
	assert.NoError(t, err)

	mockUCase := mocks.NewUseCaseI(t)

	mockUCase.On("Auth", valueCookie).Return(&user, nil)
	mockUCase.On("Auth", valueCookieUnauthorized).Return(nil, models.ErrNotFound)

	handler := authDelivery.Delivery{
		AuthUC: mockUCase,
	}

	e := echo.New()
	authDelivery.NewDelivery(e, mockUCase)

	cases := map[string]TestCase {
		"success": {
			ArgData: valueCookie,
			Error: nil,
			StatusCode: http.StatusOK,
		},
		"unauthorized_no_cookie": {
			Error: &echo.HTTPError{
				Code: http.StatusUnauthorized,
				Message: http.ErrNoCookie.Error(),
			},
		},
		"unauthorized_no_user": {
			ArgData: valueCookieUnauthorized,
			Error: &echo.HTTPError{
				Code: http.StatusUnauthorized,
				Message: models.ErrNotFound.Error(),
			},
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(echo.GET, "/auth", strings.NewReader(""))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			if name != "unauthorized_no_cookie" {
				cookie := &http.Cookie{
					Name:     "session_token",
					Value:    test.ArgData,
					HttpOnly: true,
				}
				req.AddCookie(cookie)
			}
			
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/auth")

			err = handler.Auth(c)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.StatusCode, rec.Code)
			}
		})
	}

	mockUCase.AssertExpectations(t)
}

func TestDeliveryCreateCSRF(t *testing.T) {
	var valueCookie string
	err := faker.FakeData(&valueCookie)
	assert.NoError(t, err)

	mockUCase := mocks.NewUseCaseI(t)

	handler := authDelivery.Delivery{
		AuthUC: mockUCase,
	}

	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    valueCookie,
		HttpOnly: true,
	}

	e := echo.New()
	authDelivery.NewDelivery(e, mockUCase)

	cases := map[string]TestCase {
		"success": {
			Error: nil,
			StatusCode: http.StatusNoContent,
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
			req := httptest.NewRequest(echo.POST, "/create_csrf", strings.NewReader(test.ArgData))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			if name == "success" {
				req.AddCookie(cookie)
			}
			
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/create_csrf")

			err := handler.CreateCSRF(c)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.StatusCode, rec.Code)
				assert.NotEqual(t, c.Response().Header().Get(echo.HeaderXCSRFToken), "")
			}
		})
	}

	mockUCase.AssertExpectations(t)
}

