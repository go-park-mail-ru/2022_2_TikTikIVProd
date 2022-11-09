package delivery_test

import (
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
	ArgData string
	Error error
	StatusCode int
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

	cases := map[string]TestCaseGetPost {
		"success": {
			ArgData: strconv.Itoa(post.ID),
			Error: nil,
			StatusCode: http.StatusOK,
		},
		"bad_request": {
			ArgData: postIdBadRequest,
			Error: &echo.HTTPError{
				Code: http.StatusNotFound,
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







