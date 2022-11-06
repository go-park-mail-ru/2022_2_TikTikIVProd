package delivery_test

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	friendsDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/friends/delivery"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/friends/usecase/mocks"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestCase struct {
	ArgData models.Friends
	Error error
	StatusCode int
}

type TestCaseSelectFriends struct {
	ArgData string
	Error error
	StatusCode int
}

func TestDeliveryAddFriend(t *testing.T) {
	mockFriendsSuccess := models.Friends {
		Id1: 1,
		Id2: 2,
	}
	mockFriendsBadRequest := models.Friends {
		Id1: 1,
		Id2: 0,
	}

	mockUCase := new(mocks.UseCaseI)

	mockUCase.On("AddFriend", mockFriendsSuccess).Return(nil)

	handler := friendsDelivery.Delivery{
		FriendsUC: mockUCase,
	}

	e := echo.New()

	cases := map[string]TestCase {
		"success": {
			ArgData:   mockFriendsSuccess,
			Error: nil,
			StatusCode: http.StatusCreated,
		},
		"bad_request": {
			ArgData:   mockFriendsBadRequest,
			Error: &echo.HTTPError{
				Code: http.StatusBadRequest,
				Message: "bad request",
			},
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(echo.POST, "/friends/add", strings.NewReader(""))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/friends/add/:friend_id")
			c.SetParamNames("friend_id")
			c.SetParamValues(strconv.Itoa(test.ArgData.Id2))
			c.Set("user_id", test.ArgData.Id1)

			err := handler.AddFriend(c)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.StatusCode, rec.Code)
			}
		})
	}
	mockUCase.AssertExpectations(t)
}

func TestDeliveryDeleteFriend(t *testing.T) {
	mockFriendsSuccess := models.Friends {
		Id1: 1,
		Id2: 2,
	}

	mockFriendsBadRequest := models.Friends {
		Id1: 1,
		Id2: 0,
	}

	mockFriendsNotFound := models.Friends {
		Id1: 100,
		Id2: 101,
	}

	mockUCase := new(mocks.UseCaseI)

	mockUCase.On("DeleteFriend", mockFriendsSuccess).Return(nil)
	mockUCase.On("DeleteFriend", mockFriendsNotFound).Return(models.ErrNotFound)

	handler := friendsDelivery.Delivery{
		FriendsUC: mockUCase,
	}

	e := echo.New()

	cases := map[string]TestCase {
		"success": {
			ArgData: mockFriendsSuccess,
			Error: nil,
			StatusCode: http.StatusNoContent,
		},
		"bad_request": {
			ArgData: mockFriendsBadRequest,
			Error: &echo.HTTPError{
				Code: http.StatusBadRequest,
				Message: models.ErrBadRequest.Error(),
			},
		},
		"not_found": {
			ArgData: mockFriendsNotFound,
			Error: &echo.HTTPError{
				Code: http.StatusNotFound,
				Message: models.ErrNotFound.Error(),
			},
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(echo.POST, "/friends/delete/:friend_id", strings.NewReader(""))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/friends/delete/:friend_id")
			c.SetParamNames("friend_id")
			c.SetParamValues(strconv.Itoa(test.ArgData.Id2))
			c.Set("user_id", test.ArgData.Id1)

			err := handler.DeleteFriend(c)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.StatusCode, rec.Code)
			}
		})
	}

	mockUCase.AssertExpectations(t)
}


func TestDeliverySelectFriends(t *testing.T) {
	var friends []models.User
	err := faker.FakeData(&friends)
	assert.NoError(t, err)

	userIdSuccess := 1
	userIdBadRequest := "hgcv"

	mockUCase := new(mocks.UseCaseI)

	mockUCase.On("SelectFriends", userIdSuccess).Return(friends, nil)

	handler := friendsDelivery.Delivery{
		FriendsUC: mockUCase,
	}

	e := echo.New()

	cases := map[string]TestCaseSelectFriends {
		"success": {
			ArgData: strconv.Itoa(userIdSuccess),
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
			req := httptest.NewRequest(echo.POST, "/friends/:user_id", strings.NewReader(""))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/friends/:user_id")
			c.SetParamNames("user_id")
			c.SetParamValues(test.ArgData)

			err := handler.SelectFriends(c)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.StatusCode, rec.Code)
			}
		})
	}

	mockUCase.AssertExpectations(t)
}

