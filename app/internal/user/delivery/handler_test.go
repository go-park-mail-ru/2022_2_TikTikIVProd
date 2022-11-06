package delivery_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	userDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/delivery"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/usecase/mocks"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestCaseGetProfile struct {
	ArgData string
	Error error
	StatusCode int
}

type TestCaseUpdateUser struct {
	ArgDataBody string
	ArgDataContext int
	Error error
	StatusCode int
}

func TestDeliveryGetProfile(t *testing.T) {
	var user models.User
	err := faker.FakeData(&user)
	assert.NoError(t, err)

	userIdBadRequest := "hgcv"

	mockUCase := new(mocks.UseCaseI)

	mockUCase.On("SelectUserById", user.Id).Return(&user, nil)

	handler := userDelivery.Delivery{
		UserUC: mockUCase,
	}

	e := echo.New()

	cases := map[string]TestCaseGetProfile {
		"success": {
			ArgData: strconv.Itoa(user.Id),
			Error: nil,
			StatusCode: http.StatusOK,
		},
		"bad_request": {
			ArgData: userIdBadRequest,
			Error: &echo.HTTPError{
				Code: http.StatusBadRequest,
				Message: models.ErrBadRequest.Error(),
			},
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(echo.POST, "/users/:id", strings.NewReader(""))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/users/:id")
			c.SetParamNames("id")
			c.SetParamValues(test.ArgData)

			err := handler.GetProfile(c)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.StatusCode, rec.Code)
			}
		})
	}

	mockUCase.AssertExpectations(t)
}

func TestDeliveryUpdateUser(t *testing.T) {
	var mockUser models.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)

	jsonUser, err := json.Marshal(mockUser)
	assert.NoError(t, err)

	mockUser.Id = 1

	mockUCase := new(mocks.UseCaseI)

	mockUCase.On("UpdateUser", mockUser).Return(nil)

	handler := userDelivery.Delivery{
		UserUC: mockUCase,
	}

	e := echo.New()

	cases := map[string]TestCaseUpdateUser {
		"success": {
			ArgDataBody: string(jsonUser),
			ArgDataContext: mockUser.Id,
			Error: nil,
			StatusCode: http.StatusNoContent,
		},
		"bad_request": {
			ArgDataBody: "sffvfb",
			Error: &echo.HTTPError{
				Code: http.StatusBadRequest,
				Message: models.ErrBadRequest.Error(),
			},
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(echo.POST, "/users/update", strings.NewReader(test.ArgDataBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/users/update")
			c.Set("user_id", test.ArgDataContext)

			err := handler.UpdateUser(c)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.StatusCode, rec.Code)
			}
		})
	}

	mockUCase.AssertExpectations(t)
}



// func TestDeliverySignUp(t *testing.T) {
// 	var mockUser models.User
// 	err := faker.FakeData(&mockUser)
// 	assert.NoError(t, err)

// 	var mockCookie models.Cookie
// 	err = faker.FakeData(&mockCookie)
// 	assert.NoError(t, err)

// 	jsonUser, err := json.Marshal(mockUser)
// 	assert.NoError(t, err)

// 	mockUCase := new(mocks.UseCaseI)

// 	mockUCase.On("SignUp", &mockUser).Return(&mockCookie, nil).Run(func(args mock.Arguments) {
// 		arg := args.Get(0).(*models.User)
// 		arg.Id = mockCookie.UserId
// 	})

// 	handler := authDelivery.Delivery{
// 		AuthUC: mockUCase,
// 	}

// 	userResponse := mockUser
// 	userResponse.Id = mockCookie.UserId
// 	response := pkg.Response {
// 		Body: userResponse,
// 	}
// 	jsonResponse, err := json.Marshal(response)
// 	assert.NoError(t, err)

// 	e := echo.New()

// 	cases := map[string]TestCaseSignUp {
// 		"success": {
// 			ArgData:   string(jsonUser),
// 			ExpectedResponse: string(jsonResponse) + "\n",
// 			Error: nil,
// 			StatusCode: http.StatusCreated,
// 		},
// 		"bad_request": {
// 			ArgData:   "aaa",
// 			Error: &echo.HTTPError{
// 				Code: http.StatusBadRequest,
// 				Message: models.ErrBadRequest.Error(),
// 			},
// 		},
// 	}

// 	for name, test := range cases {
// 		t.Run(name, func(t *testing.T) {
// 			req := httptest.NewRequest(echo.POST, "/signup", strings.NewReader(test.ArgData))
// 			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 			rec := httptest.NewRecorder()
// 			c := e.NewContext(req, rec)
// 			c.SetPath("/signup")

// 			err = handler.SignUp(c)
// 			require.Equal(t, test.Error, err)

// 			if err == nil {
// 				assert.Equal(t, test.StatusCode, rec.Code)
// 				assert.Equal(t, test.ExpectedResponse, rec.Body.String())
// 			}
// 		})
// 	}

// 	mockUCase.AssertExpectations(t)
// }


