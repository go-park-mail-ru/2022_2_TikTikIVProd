package usecase_test

// import (
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/bxcodec/faker"
// 	authDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/auth/delivery"
// 	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/auth/usecase/mocks"
// 	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
// 	"github.com/labstack/echo/v4"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// )

// type TestCase struct {
// 	Data string
// 	Error error
// 	StatusCode int
// }

// func TestSignUp(t *testing.T) {
// 	var mockUser models.User

// 	err := faker.FakeData(&mockUser)
// 	assert.NoError(t, err)

// 	jsonUser, err := json.Marshal(mockUser)
// 	assert.NoError(t, err)

// 	mockUCase := new(mocks.UseCaseI)

// 	var mockCookie models.Cookie

// 	err = faker.FakeData(&mockCookie)
// 	assert.NoError(t, err)

// 	mockUCase.On("SignUp", &mockUser).Return(&mockCookie, nil)

// 	handler := authDelivery.Delivery{
// 		AuthUC: mockUCase,
// 	}

// 	e := echo.New()

// 	cases := map[string]TestCase {
// 		"success": {
// 			Data:   string(jsonUser),
// 			Error: nil,
// 			StatusCode: http.StatusCreated,
// 		},
// 		"bad_request": {
// 			Data:   "aaa",
// 			Error: &echo.HTTPError{
// 				Code: http.StatusBadRequest,
// 				Message: "bad request",
// 			},
// 		},
// 	}

// 	for name, test := range cases {
// 		t.Run(name, func(t *testing.T) {
// 			req := httptest.NewRequest(echo.POST, "/signup", strings.NewReader(test.Data))
// 			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 			rec := httptest.NewRecorder()
// 			c := e.NewContext(req, rec)
// 			c.SetPath("/signup")

// 			err = handler.SignUp(c)
// 			require.Equal(t, test.Error, err)

// 			if err == nil {
// 				assert.Equal(t, test.StatusCode, rec.Code)
// 			}

// 			mockUCase.AssertExpectations(t)
// 		})
// 	}
// }

// func TestSignIn(t *testing.T) {
// 	var mockUserSignIn models.UserSignIn

// 	err := faker.FakeData(&mockUserSignIn)
// 	assert.NoError(t, err)

// 	jsonUser, err := json.Marshal(mockUserSignIn)
// 	assert.NoError(t, err)

// 	mockUCase := new(mocks.UseCaseI)

// 	var mockCookie models.Cookie

// 	err = faker.FakeData(&mockCookie)
// 	assert.NoError(t, err)

// 	var mockUser models.User

// 	err = faker.FakeData(&mockUser)
// 	assert.NoError(t, err)

// 	mockUCase.On("SignIn", mockUserSignIn).Return(&mockUser, &mockCookie, nil)

// 	handler := authDelivery.Delivery{
// 		AuthUC: mockUCase,
// 	}

// 	e := echo.New()

// 	cases := map[string]TestCase {
// 		"success": {
// 			Data:   string(jsonUser),
// 			Error: nil,
// 			StatusCode: http.StatusOK,
// 		},
// 		"bad_request": {
// 			Data:   "aaa",
// 			Error: &echo.HTTPError{
// 				Code: http.StatusBadRequest,
// 				Message: "bad request",
// 			},
// 		},
// 	}

// 	for name, test := range cases {
// 		t.Run(name, func(t *testing.T) {
// 			req := httptest.NewRequest(echo.POST, "/signin", strings.NewReader(test.Data))
// 			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 			rec := httptest.NewRecorder()
// 			c := e.NewContext(req, rec)
// 			c.SetPath("/signin")

// 			err = handler.SignIn(c)
// 			require.Equal(t, test.Error, err)

// 			if err == nil {
// 				assert.Equal(t, test.StatusCode, rec.Code)
// 			}

// 			mockUCase.AssertExpectations(t)
// 		})
// 	}
// }

// func TestLogout(t *testing.T) {
// 	var valueCookie string

// 	err := faker.FakeData(&valueCookie)
// 	assert.NoError(t, err)

// 	mockUCase := new(mocks.UseCaseI)

// 	mockUCase.On("DeleteCookie", valueCookie).Return(nil)

// 	handler := authDelivery.Delivery{
// 		AuthUC: mockUCase,
// 	}

// 	cookie := &http.Cookie{
// 		Name:     "session_token",
// 		Value:    valueCookie,
// 		HttpOnly: true,
// 	}

// 	e := echo.New()

// 	cases := map[string]TestCase {
// 		"success": {
// 			Error: nil,
// 			StatusCode: http.StatusNoContent,
// 		},
// 		"unauthorized": {
// 			Error: &echo.HTTPError{
// 				Code: http.StatusUnauthorized,
// 				Message: http.ErrNoCookie.Error(),
// 			},
// 		},
// 	}

// 	for name, test := range cases {
// 		t.Run(name, func(t *testing.T) {
// 			req := httptest.NewRequest(echo.POST, "/logout", strings.NewReader(test.Data))
// 			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 			if name == "success" {
// 				req.AddCookie(cookie)
// 			}
			
// 			rec := httptest.NewRecorder()
// 			c := e.NewContext(req, rec)
// 			c.SetPath("/logout")

// 			err = handler.Logout(c)
// 			require.Equal(t, test.Error, err)

// 			if err == nil {
// 				assert.Equal(t, test.StatusCode, rec.Code)
// 			}

// 			mockUCase.AssertExpectations(t)
// 		})
// 	}
// }

