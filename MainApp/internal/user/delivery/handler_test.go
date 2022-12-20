package delivery_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/bxcodec/faker"
	userDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/user/delivery"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/user/usecase/mocks"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
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
	ArgDataContext uint64
	Error error
	StatusCode int
}

func TestDeliveryGetProfile(t *testing.T) {
	var user models.User
	err := faker.FakeData(&user)
	assert.NoError(t, err)
	user.Id = 1

	userIdBadRequest := "hgcv"
	var userIdNotFound uint64 = 2
	var userIdInternalErr uint64 = 3

	mockUCase := mocks.NewUseCaseI(t)

	mockUCase.On("SelectUserById", user.Id).Return(&user, nil)
	mockUCase.On("SelectUserById", userIdNotFound).Return(nil, models.ErrNotFound)
	mockUCase.On("SelectUserById", userIdInternalErr).Return(nil, models.ErrInternalServerError)

	handler := userDelivery.Delivery{
		UserUC: mockUCase,
	}

	e := echo.New()
	userDelivery.NewDelivery(e, mockUCase)

	cases := map[string]TestCaseGetProfile {
		"success": {
			ArgData: strconv.Itoa(int(user.Id)),
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
		"not_found": {
			ArgData: strconv.Itoa(int(userIdNotFound)),
			Error: &echo.HTTPError{
				Code: http.StatusNotFound,
				Message: models.ErrNotFound.Error(),
			},
		},
		"internal_error": {
			ArgData: strconv.Itoa(int(userIdInternalErr)),
			Error: &echo.HTTPError{
				Code: http.StatusInternalServerError,
				Message: models.ErrInternalServerError.Error(),
			},
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(echo.GET, "/users/:id", strings.NewReader(""))
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
	mockUser.Id = 1
	mockUser.CreatedAt = time.Date(2022, time.September, 5, 1, 12, 12, 12, time.Local)
	

	jsonUser, err := json.Marshal(mockUser)
	assert.NoError(t, err)

	var mockUserNotFound models.User
	err = faker.FakeData(&mockUserNotFound)
	assert.NoError(t, err)
	mockUserNotFound.Id = 2
	mockUserNotFound.CreatedAt = time.Date(2022, time.September, 5, 1, 12, 12, 12, time.Local)

	jsonUserNotFound, err := json.Marshal(mockUserNotFound)
	assert.NoError(t, err)

	var mockUserInternalErr models.User
	err = faker.FakeData(&mockUserInternalErr)
	assert.NoError(t, err)
	mockUserInternalErr.Id = 3
	mockUserInternalErr.CreatedAt = time.Date(2022, time.September, 5, 1, 12, 12, 12, time.Local)

	jsonUserInternalErr, err := json.Marshal(mockUserInternalErr)
	assert.NoError(t, err)

	mockUCase := mocks.NewUseCaseI(t)

	mockUCase.On("UpdateUser", mockUser).Return(nil)
	mockUCase.On("UpdateUser", mockUserNotFound).Return(models.ErrNotFound)
	mockUCase.On("UpdateUser", mockUserInternalErr).Return(models.ErrInternalServerError)

	handler := userDelivery.Delivery{
		UserUC: mockUCase,
	}

	e := echo.New()
	userDelivery.NewDelivery(e, mockUCase)

	cases := map[string]TestCaseUpdateUser {
		"success": {
			ArgDataBody: string(jsonUser),
			ArgDataContext: mockUser.Id,
			Error: nil,
			StatusCode: http.StatusNoContent,
		},
		"not_found": {
			ArgDataBody: string(jsonUserNotFound),
			ArgDataContext: mockUserNotFound.Id,
			Error: &echo.HTTPError{
				Code: http.StatusNotFound,
				Message: models.ErrNotFound.Error(),
			},
		},
		"internal_error": {
			ArgDataBody: string(jsonUserInternalErr),
			ArgDataContext: mockUserInternalErr.Id,
			Error: &echo.HTTPError{
				Code: http.StatusInternalServerError,
				Message: models.ErrInternalServerError.Error(),
			},
		},
		"bad_request": {
			ArgDataBody: "sffvfb",
			Error: &echo.HTTPError{
				Code: http.StatusBadRequest,
				Message: models.ErrBadRequest.Error(),
			},
		},
		"invalid_context": {
			ArgDataBody:   string(jsonUser),
			Error: &echo.HTTPError{
				Code: http.StatusInternalServerError,
				Message: models.ErrInternalServerError.Error(),
			},
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(echo.PUT, "/users/update", strings.NewReader(test.ArgDataBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/users/update")
			if name != "invalid_context" {
				c.Set("user_id", test.ArgDataContext)
			}

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
	userDelivery.NewDelivery(e, mockUCase)

	cases := map[string]TestCaseGetProfile {
		"success": {
			Error: nil,
			StatusCode: http.StatusOK,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(echo.GET, "/users", strings.NewReader(""))
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


