package delivery_test

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	// friendsDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/friends/delivery"
	// "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/friends/usecase/mocks"
	// "github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestCase struct {
	ArgDataUserId int
	ArgDataFriendId string
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
	mockFriendsInvalid := models.Friends {
		Id1: 1,
		Id2: 0,
	}
	mockFriendsNotFound := models.Friends {
		Id1: 3,
		Id2: 4,
	}
	mockFriendsConflict := models.Friends {
		Id1: 5,
		Id2: 6,
	}
	mockFriendsEqual := models.Friends {
		Id1: 5,
		Id2: 5,
	}
	mockFriendsInternalErr := models.Friends {
		Id1: 6,
		Id2: 7,
	}

	mockUCase := mocks.NewUseCaseI(t)

	mockUCase.On("AddFriend", mockFriendsSuccess).Return(nil)
	mockUCase.On("AddFriend", mockFriendsNotFound).Return(models.ErrNotFound)
	mockUCase.On("AddFriend", mockFriendsConflict).Return(models.ErrConflictFriend)
	mockUCase.On("AddFriend", mockFriendsEqual).Return(models.ErrBadRequest)
	mockUCase.On("AddFriend", mockFriendsInternalErr).Return(models.ErrInternalServerError)

	handler := friendsDelivery.Delivery{
		FriendsUC: mockUCase,
	}

	e := echo.New()
	friendsDelivery.NewDelivery(e, mockUCase)

	cases := map[string]TestCase {
		"success": {
			ArgDataUserId: mockFriendsSuccess.Id1,
			ArgDataFriendId: strconv.Itoa(mockFriendsSuccess.Id2),
			Error: nil,
			StatusCode: http.StatusCreated,
		},
		"ivalid_friends": {
			ArgDataUserId: mockFriendsInvalid.Id1,
			ArgDataFriendId:   strconv.Itoa(mockFriendsInvalid.Id2),
			Error: &echo.HTTPError{
				Code: http.StatusBadRequest,
				Message: models.ErrBadRequest.Error(),
			},
		},
		"bad_request": {
			ArgDataUserId: 1,
			ArgDataFriendId:   "aaa",
			Error: &echo.HTTPError{
				Code: http.StatusBadRequest,
				Message: models.ErrBadRequest.Error(),
			},
		},
		"invalid_context": {
			ArgDataFriendId:   "2",
			Error: &echo.HTTPError{
				Code: http.StatusInternalServerError,
				Message: models.ErrInternalServerError.Error(),
			},
		},
		"not_found": {
			ArgDataUserId: mockFriendsNotFound.Id1,
			ArgDataFriendId:   strconv.Itoa(mockFriendsNotFound.Id2),
			Error: &echo.HTTPError{
				Code: http.StatusNotFound,
				Message: models.ErrNotFound.Error(),
			},
		},
		"conflict": {
			ArgDataUserId: mockFriendsConflict.Id1,
			ArgDataFriendId:   strconv.Itoa(mockFriendsConflict.Id2),
			Error: &echo.HTTPError{
				Code: http.StatusConflict,
				Message: models.ErrConflictFriend.Error(),
			},
		},
		"equal_id": {
			ArgDataUserId: mockFriendsEqual.Id1,
			ArgDataFriendId:   strconv.Itoa(mockFriendsEqual.Id2),
			Error: &echo.HTTPError{
				Code: http.StatusBadRequest,
				Message: models.ErrBadRequest.Error(),
			},
		},
		"internal_error": {
			ArgDataUserId: mockFriendsInternalErr.Id1,
			ArgDataFriendId:   strconv.Itoa(mockFriendsInternalErr.Id2),
			Error: &echo.HTTPError{
				Code: http.StatusInternalServerError,
				Message: models.ErrInternalServerError.Error(),
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
			c.SetParamValues(test.ArgDataFriendId)
			if name != "invalid_context" {
				c.Set("user_id", test.ArgDataUserId)
			}

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
	mockFriendsInvalid := models.Friends {
		Id1: 1,
		Id2: 0,
	}
	mockFriendsNotFound := models.Friends {
		Id1: 100,
		Id2: 101,
	}
	mockFriendsEqual := models.Friends {
		Id1: 5,
		Id2: 5,
	}
	mockFriendsInternalErr := models.Friends {
		Id1: 6,
		Id2: 7,
	}

	mockUCase := mocks.NewUseCaseI(t)

	mockUCase.On("DeleteFriend", mockFriendsSuccess).Return(nil)
	mockUCase.On("DeleteFriend", mockFriendsNotFound).Return(models.ErrNotFound)
	mockUCase.On("DeleteFriend", mockFriendsEqual).Return(models.ErrBadRequest)
	mockUCase.On("DeleteFriend", mockFriendsInternalErr).Return(models.ErrInternalServerError)

	handler := friendsDelivery.Delivery{
		FriendsUC: mockUCase,
	}

	e := echo.New()
	friendsDelivery.NewDelivery(e, mockUCase)

	cases := map[string]TestCase {
		"success": {
			ArgDataUserId: mockFriendsSuccess.Id1,
			ArgDataFriendId: strconv.Itoa(mockFriendsSuccess.Id2),
			Error: nil,
			StatusCode: http.StatusNoContent,
		},
		"invalid_friends": {
			ArgDataUserId: mockFriendsInvalid.Id1,
			ArgDataFriendId: strconv.Itoa(mockFriendsInvalid.Id2),
			Error: &echo.HTTPError{
				Code: http.StatusBadRequest,
				Message: models.ErrBadRequest.Error(),
			},
		},
		"not_found": {
			ArgDataUserId: mockFriendsNotFound.Id1,
			ArgDataFriendId: strconv.Itoa(mockFriendsNotFound.Id2),
			Error: &echo.HTTPError{
				Code: http.StatusNotFound,
				Message: models.ErrNotFound.Error(),
			},
		},
		"bad_request": {
			ArgDataUserId: 1,
			ArgDataFriendId: "sjhdb",
			Error: &echo.HTTPError{
				Code: http.StatusBadRequest,
				Message: models.ErrBadRequest.Error(),
			},
		},
		"invalid_context": {
			ArgDataFriendId:   "2",
			Error: &echo.HTTPError{
				Code: http.StatusInternalServerError,
				Message: models.ErrInternalServerError.Error(),
			},
		},
		"equal_id": {
			ArgDataUserId: mockFriendsEqual.Id1,
			ArgDataFriendId:   strconv.Itoa(mockFriendsEqual.Id2),
			Error: &echo.HTTPError{
				Code: http.StatusBadRequest,
				Message: models.ErrBadRequest.Error(),
			},
		},
		"internal_error": {
			ArgDataUserId: mockFriendsInternalErr.Id1,
			ArgDataFriendId:   strconv.Itoa(mockFriendsInternalErr.Id2),
			Error: &echo.HTTPError{
				Code: http.StatusInternalServerError,
				Message: models.ErrInternalServerError.Error(),
			},
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(echo.DELETE, "/friends/delete/:friend_id", strings.NewReader(""))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/friends/delete/:friend_id")
			c.SetParamNames("friend_id")
			c.SetParamValues(test.ArgDataFriendId)
			if name != "invalid_context" {
				c.Set("user_id", test.ArgDataUserId)
			}

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
	friends := make([]models.User, 0, 10)
	err := faker.FakeData(&friends)
	assert.NoError(t, err)

	userIdSuccess := 1
	userIdInternalErr := 2
	userIdNotFound := 3
	userIdBadRequest := "hgcv"

	mockUCase := mocks.NewUseCaseI(t)

	mockUCase.On("SelectFriends", userIdSuccess).Return(friends, nil)
	mockUCase.On("SelectFriends", userIdNotFound).Return(nil, models.ErrNotFound)
	mockUCase.On("SelectFriends", userIdInternalErr).Return(nil, models.ErrInternalServerError)

	handler := friendsDelivery.Delivery{
		FriendsUC: mockUCase,
	}

	e := echo.New()
	friendsDelivery.NewDelivery(e, mockUCase)

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
		"internal_error": {
			ArgData: strconv.Itoa(userIdInternalErr),
			Error: &echo.HTTPError{
				Code: http.StatusInternalServerError,
				Message: models.ErrInternalServerError.Error(),
			},
		},
		"not_found": {
			ArgData: strconv.Itoa(userIdNotFound),
			Error: &echo.HTTPError{
				Code: http.StatusNotFound,
				Message: models.ErrNotFound.Error(),
			},
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(echo.GET, "/friends/:user_id", strings.NewReader(""))
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

func TestDeliveryCheckIsFriend(t *testing.T) {
	friendIdSuccess := 2
	friendIdBadRequest := "hgcv"

	userId := 1

	mockFriendsSuccess := models.Friends {
		Id1: userId,
		Id2: friendIdSuccess,
	}
	mockFriendsNotFound := models.Friends {
		Id1: 100,
		Id2: 101,
	}
	mockFriendsEqual := models.Friends {
		Id1: 5,
		Id2: 5,
	}
	mockFriendsInternalErr := models.Friends {
		Id1: 6,
		Id2: 7,
	}

	mockUCase := mocks.NewUseCaseI(t)

	mockUCase.On("CheckIsFriend", mockFriendsSuccess).Return(true, nil)
	mockUCase.On("CheckIsFriend", mockFriendsNotFound).Return(false, models.ErrNotFound)
	mockUCase.On("CheckIsFriend", mockFriendsEqual).Return(false, models.ErrBadRequest)
	mockUCase.On("CheckIsFriend", mockFriendsInternalErr).Return(false, models.ErrInternalServerError)

	handler := friendsDelivery.Delivery{
		FriendsUC: mockUCase,
	}

	e := echo.New()
	friendsDelivery.NewDelivery(e, mockUCase)

	cases := map[string]TestCase {
		"success": {
			ArgDataFriendId: strconv.Itoa(friendIdSuccess),
			ArgDataUserId: userId,
			Error: nil,
			StatusCode: http.StatusOK,
		},
		"bad_request": {
			ArgDataFriendId: friendIdBadRequest,
			ArgDataUserId: userId,
			Error: &echo.HTTPError{
				Code: http.StatusBadRequest,
				Message: models.ErrBadRequest.Error(),
			},
		},
		"invalid_context": {
			ArgDataFriendId:   "2",
			Error: &echo.HTTPError{
				Code: http.StatusInternalServerError,
				Message: models.ErrInternalServerError.Error(),
			},
		},
		"equal_id": {
			ArgDataUserId: mockFriendsEqual.Id1,
			ArgDataFriendId:   strconv.Itoa(mockFriendsEqual.Id2),
			Error: &echo.HTTPError{
				Code: http.StatusBadRequest,
				Message: models.ErrBadRequest.Error(),
			},
		},
		"internal_error": {
			ArgDataUserId: mockFriendsInternalErr.Id1,
			ArgDataFriendId:   strconv.Itoa(mockFriendsInternalErr.Id2),
			Error: &echo.HTTPError{
				Code: http.StatusInternalServerError,
				Message: models.ErrInternalServerError.Error(),
			},
		},
		"not_found": {
			ArgDataUserId: mockFriendsNotFound.Id1,
			ArgDataFriendId: strconv.Itoa(mockFriendsNotFound.Id2),
			Error: &echo.HTTPError{
				Code: http.StatusNotFound,
				Message: models.ErrNotFound.Error(),
			},
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(echo.GET, "/friends/check/:friend_id", strings.NewReader(""))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/friends/check/:friend_id")
			c.SetParamNames("friend_id")
			c.SetParamValues(test.ArgDataFriendId)
			if name != "invalid_context" {
				c.Set("user_id", test.ArgDataUserId)
			}

			err := handler.CheckIsFriend(c)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.StatusCode, rec.Code)
			}
		})
	}

	mockUCase.AssertExpectations(t)
}

