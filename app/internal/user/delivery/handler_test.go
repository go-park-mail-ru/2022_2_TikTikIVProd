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

type TestCaseGetUsers struct {
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

	mockUCase := mocks.NewUseCaseI(t)

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

	mockUCase := mocks.NewUseCaseI(t)

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

func TestDeliveryGetUsers(t *testing.T) {
	users := make([]models.User, 0, 10)
	err := faker.FakeData(&users)
	assert.NoError(t, err)

	mockUCase := mocks.NewUseCaseI(t)

	mockUCase.On("SelectUsers").Return(users, nil)

	handler := userDelivery.Delivery{
		UserUC: mockUCase,
	}

	e := echo.New()

	cases := map[string]TestCaseGetProfile {
		"success": {
			Error: nil,
			StatusCode: http.StatusOK,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(echo.POST, "/users", strings.NewReader(""))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/users")

			err := handler.GetUsers(c)
			require.Equal(t, test.Error, err)
			assert.Equal(t, test.StatusCode, rec.Code)
		})
	}

	mockUCase.AssertExpectations(t)
}


