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
	communityDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/communities/delivery"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/communities/usecase/mocks"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestCase struct {
	ArgData string
	ArgDataContext uint64
	Error error
	StatusCode int
}

func TestDeliveryCreateCommunity(t *testing.T) {
	mockCommunity := models.ReqCommunityCreate {
		AvatarID: 1,
		Name: "name",
		Description: "descr",
	}

	var mockCommunityInternalErr models.ReqCommunityCreate
	err := faker.FakeData(&mockCommunityInternalErr)
	assert.NoError(t, err)
	mockCommunityInternalErr.AvatarID = 2

	mockCommunityInvalid := models.ReqCommunityCreate{}

	comm := models.ReqCreateToComm(mockCommunity)
	comm.OwnerID = 1

	commInternalErr := models.ReqCreateToComm(mockCommunityInternalErr)
	commInternalErr.OwnerID = 1

	jsonCommunity, err := json.Marshal(mockCommunity)
	assert.NoError(t, err)

	jsonCommunityInvalid, err := json.Marshal(mockCommunityInvalid)
	assert.NoError(t, err)

	jsonCommunityInternalErr, err := json.Marshal(mockCommunityInternalErr)
	assert.NoError(t, err)

	mockUCase := mocks.NewUseCaseI(t)

	mockUCase.On("CreateCommunity", &comm).Return(nil)
	mockUCase.On("CreateCommunity", &commInternalErr).Return(models.ErrInternalServerError)

	handler := communityDelivery.Delivery {
		CommUC: mockUCase,
	}

	e := echo.New()
	communityDelivery.NewDelivery(e, mockUCase)

	cases := map[string]TestCase {
		"success": {
			ArgData:   string(jsonCommunity),
			ArgDataContext: 1,
			Error: nil,
			StatusCode: http.StatusOK,
		},
		"bad_request": {
			ArgData:   "aaa",
			ArgDataContext: 1,
			Error: &echo.HTTPError{
				Code: http.StatusUnprocessableEntity,
				Message: models.ErrBadRequest.Error(),
			},
		},
		"invalid_request": {
			ArgData:   string(jsonCommunityInvalid),
			ArgDataContext: 1,
			Error: &echo.HTTPError{
				Code: http.StatusBadRequest,
				Message: models.ErrBadRequest.Error(),
			},
		},
		"internal_error": {
			ArgData:   string(jsonCommunityInternalErr),
			ArgDataContext: 1,
			Error: &echo.HTTPError{
				Code: http.StatusInternalServerError,
				Message: models.ErrInternalServerError.Error(),
			},
		},
		"invalid_context": {
			ArgData:   string(jsonCommunity),
			Error: &echo.HTTPError{
				Code: http.StatusInternalServerError,
				Message: models.ErrInternalServerError.Error(),
			},
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(echo.POST, "/communities/create", strings.NewReader(test.ArgData))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/communities/create")
			if name != "invalid_context" {
				c.Set("user_id", test.ArgDataContext)
			}

			err = handler.CreateCommunity(c)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.StatusCode, rec.Code)
			}
		})
	}

	mockUCase.AssertExpectations(t)
}

func TestDeliveryUpdateCommunity(t *testing.T) {
	mockCommunity := models.Community {
		ID: 1,
		OwnerID: 1,
		AvatarID: 1,
		Name: "name",
		Description: "descr",
		CreateDate: time.Date(2022, time.September, 5, 1, 12, 12, 12, time.Local),
	}

	mockCommunityInternalErr := models.Community {
		ID: 2,
		OwnerID: 1,
		AvatarID: 1,
		Name: "name",
		Description: "descr",
		CreateDate: time.Date(2022, time.September, 5, 1, 12, 12, 12, time.Local),
	}

	jsonCommunity, err := json.Marshal(mockCommunity)
	assert.NoError(t, err)

	jsonCommunityInternalErr, err := json.Marshal(mockCommunityInternalErr)
	assert.NoError(t, err)

	mockUCase := mocks.NewUseCaseI(t)

	mockUCase.On("UpdateCommunity", &mockCommunity).Return(nil)
	mockUCase.On("UpdateCommunity", &mockCommunityInternalErr).Return(models.ErrInternalServerError)

	handler := communityDelivery.Delivery {
		CommUC: mockUCase,
	}

	e := echo.New()
	communityDelivery.NewDelivery(e, mockUCase)

	cases := map[string]TestCase {
		"success": {
			ArgData:   string(jsonCommunity),
			ArgDataContext: 1,
			Error: nil,
			StatusCode: http.StatusOK,
		},
		"bad_request": {
			ArgData:   "aaa",
			ArgDataContext: 1,
			Error: &echo.HTTPError{
				Code: http.StatusUnprocessableEntity,
				Message: models.ErrBadRequest.Error(),
			},
		},
		"internal_error": {
			ArgData:   string(jsonCommunityInternalErr),
			ArgDataContext: 1,
			Error: &echo.HTTPError{
				Code: http.StatusInternalServerError,
				Message: models.ErrInternalServerError.Error(),
			},
		},
		"invalid_context": {
			ArgData:   string(jsonCommunity),
			Error: &echo.HTTPError{
				Code: http.StatusInternalServerError,
				Message: models.ErrInternalServerError.Error(),
			},
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(echo.POST, "/communities/edit", strings.NewReader(test.ArgData))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/communities/edit")
			if name != "invalid_context" {
				c.Set("user_id", test.ArgDataContext)
			}

			err = handler.UpdateCommunity(c)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.StatusCode, rec.Code)
			}
		})
	}

	mockUCase.AssertExpectations(t)
}

func TestDeliveryDeleteCommunity(t *testing.T) {
	var userId uint64 = 1

	var communityId uint64 = 1
	var communityIdErr uint64 = 2

	mockUCase := mocks.NewUseCaseI(t)

	mockUCase.On("DeleteCommunity", communityId, userId).Return(nil)
	mockUCase.On("DeleteCommunity", communityIdErr, userId).Return(models.ErrInternalServerError)

	handler := communityDelivery.Delivery {
		CommUC: mockUCase,
	}

	e := echo.New()
	communityDelivery.NewDelivery(e, mockUCase)

	cases := map[string]TestCase {
		"success": {
			ArgData:   strconv.Itoa(int(communityId)),
			ArgDataContext: 1,
			Error: nil,
			StatusCode: http.StatusNoContent,
		},
		"bad_request": {
			ArgData:   "aaa",
			ArgDataContext: 1,
			Error: &echo.HTTPError{
				Code: http.StatusNotFound,
				Message: models.ErrNotFound.Error(),
			},
		},
		"internal_error": {
			ArgData:   strconv.Itoa(int(communityIdErr)),
			ArgDataContext: 1,
			Error: &echo.HTTPError{
				Code: http.StatusInternalServerError,
				Message: models.ErrInternalServerError.Error(),
			},
		},
		"invalid_context": {
			ArgData:   strconv.Itoa(int(communityId)),
			Error: &echo.HTTPError{
				Code: http.StatusInternalServerError,
				Message: models.ErrInternalServerError.Error(),
			},
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(echo.DELETE, "/communities/:id", strings.NewReader(test.ArgData))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/communities/:id")
			c.SetParamNames("id")
			c.SetParamValues(test.ArgData)
			if name != "invalid_context" {
				c.Set("user_id", test.ArgDataContext)
			}

			err := handler.DeleteCommunity(c)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.StatusCode, rec.Code)
			}
		})
	}

	mockUCase.AssertExpectations(t)
}

