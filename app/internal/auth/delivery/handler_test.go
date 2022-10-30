package delivery_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	authDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/auth/delivery"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/auth/usecase/mocks"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	//"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type TestCase struct {
	JsonUser string
	Error error
	StatusCode int
}

func TestSignUp(t *testing.T) {
	var mockUser models.User

	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)

	jsonUser, err := json.Marshal(mockUser)
	assert.NoError(t, err)

	mockUCase := new(mocks.UseCaseI)

	var mockCookie models.Cookie

	err = faker.FakeData(&mockCookie)
	assert.NoError(t, err)

	mockUCase.On("SignUp", &mockUser).Return(&mockCookie, nil)

	handler := authDelivery.Delivery{
		AuthUC: mockUCase,
	}

	e := echo.New()

	cases := []TestCase{
		{
			JsonUser:   string(jsonUser),
			Error: nil,
			StatusCode: http.StatusCreated,
		},
		{
			JsonUser:   "aaa",
			Error: &echo.HTTPError{
				Code: http.StatusBadRequest,
				Message: "bad request",
			},
			StatusCode: http.StatusBadRequest,
		},
	}
	for _, item := range cases {
		req := httptest.NewRequest(echo.POST, "/signup", strings.NewReader(item.JsonUser))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/signup")

		err = handler.SignUp(c)
		//e.DefaultHTTPErrorHandler)
		require.Equal(t, item.Error, err)

		// if err == nil {

		// }

		assert.Equal(t, item.StatusCode, rec.Code)
		mockUCase.AssertExpectations(t)
	}	
}


//assert.Equal(t, userJSON, rec.Body.String())












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
// 			JsonUser:   string(jsonUser),
// 			Error: nil,
// 			StatusCode: http.StatusCreated,
// 		},
// 		"bad_request": {
// 			JsonUser:   "aaa",
// 			Error: &echo.HTTPError{
// 				Code: http.StatusBadRequest,
// 				Message: "bad request",
// 			},
// 			StatusCode: http.StatusBadRequest,
// 		},
// 	}

// 	for name, test := range cases {
// 		t.Run(name, func(t *testing.T) {
// 			req := httptest.NewRequest(echo.POST, "/signup", strings.NewReader(test.JsonUser))
// 			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 			rec := httptest.NewRecorder()
// 			c := e.NewContext(req, rec)
// 			c.SetPath("/signup")

// 			err = handler.SignUp(c)
// 			require.Equal(t, test.Error, err)

// 			assert.Equal(t, test.StatusCode, rec.Code)
// 			mockUCase.AssertExpectations(t)
// 		})
// 	}	










// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/repository/localstorage"
// 	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/usecase"
// )

// type TestCase struct {
// 	User       string
// 	Method string
// 	StatusCode int
// }

// func TestDelivery_signUp(t *testing.T) {
// 	cases := []TestCase{
// 		{
// 			User:   `{"first_name":"Nastya", "last_name":"Kuznetsova", "nick_name":"kuzkus", "email":"aaa@gmail.com", "password":"password1"}`,
// 			Method: "POST",
// 			StatusCode: http.StatusCreated,
// 		},
// 		{
// 			User:       `{""}`,
// 			Method: "POST",
// 			StatusCode: http.StatusBadRequest,
// 		},
// 		{
// 			User:       `{""}`,
// 			Method: "GET",
// 			StatusCode: http.StatusMethodNotAllowed,
// 		},
// 	}
// 	for caseNum, item := range cases {
// 		url := "/signup"
// 		req := httptest.NewRequest(item.Method, url, strings.NewReader(item.User))
// 		w := httptest.NewRecorder()

// 		usersLocalStorage := localstorage.New()
// 		usersUC := usecase.New(usersLocalStorage)
// 		usersDeliver := New(usersUC)
// 		usersDeliver.SignUp(w, req)

// 		if w.Code != item.StatusCode {
// 			t.Errorf("[%d] wrong StatusCode: got %d, expected %d",
// 				caseNum, w.Code, item.StatusCode)
// 		}
// 	}
// }

// func TestDelivery_signInFailure(t *testing.T) {
// 	cases := []TestCase{
// 		{
// 			User:       `{""}`,
// 			Method: "POST",
// 			StatusCode: http.StatusBadRequest,
// 		},
// 		{
// 			User:       `{""}`,
// 			Method: "GET",
// 			StatusCode: http.StatusMethodNotAllowed,
// 		},
// 		{
// 			User:       `{"email":"aaa@gmail.com", "password":"password1"}`,
// 			Method: "POST",
// 			StatusCode: http.StatusNotFound,
// 		},
// 	}
// 	for caseNum, item := range cases {
// 		url := "http://89.208.197.127:8080/signin"
// 		req := httptest.NewRequest(item.Method, url, strings.NewReader(item.User))
// 		w := httptest.NewRecorder()

// 		usersLocalStorage := localstorage.New()
// 		usersUC := usecase.New(usersLocalStorage)
// 		usersDeliver := New(usersUC)
// 		usersDeliver.SignIn(w, req)

// 		if w.Code != item.StatusCode {
// 			t.Errorf("[%d] wrong StatusCode: got %d, expected %d",
// 				caseNum, w.Code, item.StatusCode)
// 		}
// 	}
// }

// func TestDelivery_LogoutFailure(t *testing.T) {
// 	cases := []TestCase{
// 		{
// 			User:       `{""}`,
// 			Method: "POST",
// 			StatusCode: http.StatusBadRequest,
// 		},
// 		{
// 			User:       `{""}`,
// 			Method: "GET",
// 			StatusCode: http.StatusMethodNotAllowed,
// 		},
// 		{
// 			User:       `{"email":"aaa@gmail.com", "password":"password1"}`,
// 			Method: "POST",
// 			StatusCode: http.StatusNotFound,
// 		},
// 	}
// 	for caseNum, item := range cases {
// 		url := "http://89.208.197.127:8080/signin"
// 		req := httptest.NewRequest(item.Method, url, strings.NewReader(item.User))
// 		w := httptest.NewRecorder()

// 		usersLocalStorage := localstorage.New()
// 		usersUC := usecase.New(usersLocalStorage)
// 		usersDeliver := New(usersUC)
// 		usersDeliver.SignIn(w, req)

// 		if w.Code != item.StatusCode {
// 			t.Errorf("[%d] wrong StatusCode: got %d, expected %d",
// 				caseNum, w.Code, item.StatusCode)
// 		}
// 	}
// }


