package delivery_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	postDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/delivery"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/usecase/mocks"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestCaseGetPost struct {
	ArgData    string
	Error      error
	StatusCode int
}

type TestCaseCreatePost struct {
	ArgDataBody    string
	ArgDataContext int
	Error          error
	StatusCode     int
}

func TestDeliveryGetPost(t *testing.T) {
	var post models.Post
	err := faker.FakeData(&post)
	assert.NoError(t, err)

	postIdBadRequest := "hgcv"

	mockUCase := mocks.NewPostUseCaseI(t)

	mockUCase.On("GetPostById", post.ID).Return(&post, nil)

	handler := postDelivery.Delivery{
		PUsecase: mockUCase,
	}

	e := echo.New()
	postDelivery.NewDelivery(e, mockUCase)

	cases := map[string]TestCaseGetPost{
		"success": {
			ArgData:    strconv.Itoa(post.ID),
			Error:      nil,
			StatusCode: http.StatusOK,
		},
		"bad_request": {
			ArgData: postIdBadRequest,
			Error: &echo.HTTPError{
				Code:    http.StatusNotFound,
				Message: "Post not found",
			},
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(echo.GET, "/post/:id", strings.NewReader(""))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/post/:id")
			c.SetParamNames("id")
			c.SetParamValues(test.ArgData)

			err := handler.GetPost(c)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.StatusCode, rec.Code)
			}
		})
	}

	mockUCase.AssertExpectations(t)
}

func TestDeliveryCreatePost(t *testing.T) {
	mockPost := models.Post{Message: "123", Images: []models.Image{}}

	jsonPost, err := json.Marshal(mockPost)
	fmt.Println(string(jsonPost))
	assert.NoError(t, err)

	mockUCase := mocks.NewPostUseCaseI(t)

	mockUCase.On("CreatePost", &mockPost).Return(nil)

	handler := postDelivery.Delivery{
		PUsecase: mockUCase,
	}

	e := echo.New()
	postDelivery.NewDelivery(e, mockUCase)

	cases := map[string]TestCaseCreatePost{
		"success": {
			ArgDataBody:    string(jsonPost),
			ArgDataContext: mockPost.UserID,
			Error:          nil,
			StatusCode:     http.StatusOK,
		},
		//"bad_request": {
		//	ArgDataBody: "sffvfb",
		//	Error: &echo.HTTPError{
		//		Code:    http.StatusBadRequest,
		//		Message: models.ErrBadRequest.Error(),
		//	},
		//},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(echo.POST, "/post/create", strings.NewReader(test.ArgDataBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/post/create")
			c.Set("user_id", test.ArgDataContext)

			err := handler.CreatePost(c)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.StatusCode, rec.Code)
			}
		})
	}

	mockUCase.AssertExpectations(t)
}
