package delivery_test

// import (
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"strconv"
// 	"strings"
// 	"testing"

// 	"github.com/bxcodec/faker"
// 	postDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/post/delivery"
// 	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/post/usecase/mocks"
// 	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
// 	"github.com/labstack/echo/v4"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// )

// type TestCaseGetPost struct {
// 	ArgData        string
// 	ArgDataContext uint64
// 	Error          error
// 	StatusCode     int
// }

// type TestCaseCreatePost struct {
// 	ArgDataBody    string
// 	ArgDataContext uint64
// 	Error          error
// 	StatusCode     int
// }

// type TestCaseDeletePost struct {
// 	ArgDataContext uint64
// 	Error          error
// 	StatusCode     int
// 	ID             uint64
// }

// type TestCaseFeed struct {
// 	ArgDataContext uint64
// 	Error          error
// 	StatusCode     int
// }

// func TestDeliveryGetPost(t *testing.T) {
// 	var post models.Post
// 	err := faker.FakeData(&post)
// 	assert.NoError(t, err)

// 	postIdBadRequest := "hgcv"

// 	mockUCase := mocks.NewPostUseCaseI(t)

// 	var userId uint64 = 1

// 	mockUCase.On("GetPostById", post.ID, userId).Return(&post, nil)

// 	handler := postDelivery.Delivery{
// 		PUsecase: mockUCase,
// 	}

// 	e := echo.New()
// 	postDelivery.NewDelivery(e, mockUCase)

// 	cases := map[string]TestCaseGetPost{
// 		"success": {
// 			ArgData:        strconv.Itoa(int(post.ID)),
// 			ArgDataContext: userId,
// 			Error:          nil,
// 			StatusCode:     http.StatusOK,
// 		},
// 		"bad_request": {
// 			ArgData:        postIdBadRequest,
// 			ArgDataContext: userId,
// 			Error: &echo.HTTPError{
// 				Code:    http.StatusBadRequest,
// 				Message: "bad request",
// 			},
// 		},
// 		"invalid_context": {
// 			ArgData: strconv.Itoa(int(post.ID)),
// 			Error: &echo.HTTPError{
// 				Code:    http.StatusInternalServerError,
// 				Message: models.ErrInternalServerError.Error(),
// 			},
// 		},
// 	}

// 	for name, test := range cases {
// 		t.Run(name, func(t *testing.T) {
// 			req := httptest.NewRequest(echo.GET, "/post/:id", strings.NewReader(""))
// 			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 			rec := httptest.NewRecorder()
// 			c := e.NewContext(req, rec)
// 			c.SetPath("/post/:id")
// 			c.SetParamNames("id")
// 			c.SetParamValues(test.ArgData)
// 			if name != "invalid_context" {
// 				c.Set("user_id", test.ArgDataContext)
// 			}

// 			err := handler.GetPost(c)
// 			require.Equal(t, test.Error, err)

// 			if err == nil {
// 				assert.Equal(t, test.StatusCode, rec.Code)
// 			}
// 		})
// 	}

// 	mockUCase.AssertExpectations(t)
// }

// func TestDeliveryCreatePost(t *testing.T) {
// 	mockPostValid := models.Post{Message: "123", Attachments: []models.Attachment{}}
// 	mockPostInValid := models.Post{Attachments: []models.Attachment{}}

// 	jsonPostValid, err := json.Marshal(mockPostValid)
// 	assert.NoError(t, err)
// 	jsonPostInValid, err := json.Marshal(mockPostInValid)
// 	assert.NoError(t, err)

// 	mockUCase := mocks.NewPostUseCaseI(t)

// 	mockUCase.On("CreatePost", &mockPostValid).Return(nil)

// 	handler := postDelivery.Delivery{
// 		PUsecase: mockUCase,
// 	}

// 	e := echo.New()
// 	postDelivery.NewDelivery(e, mockUCase)

// 	cases := map[string]TestCaseCreatePost{
// 		"success": {
// 			ArgDataBody:    string(jsonPostValid),
// 			ArgDataContext: mockPostValid.UserID,
// 			Error:          nil,
// 			StatusCode:     http.StatusOK,
// 		},
// 		"bad_request": {
// 			ArgDataBody:    string(jsonPostInValid),
// 			ArgDataContext: mockPostValid.UserID,
// 			Error: &echo.HTTPError{
// 				Code:    http.StatusInternalServerError,
// 				Message: models.ErrInternalServerError.Error(),
// 			},
// 		},
// 	}

// 	for name, test := range cases {
// 		t.Run(name, func(t *testing.T) {
// 			req := httptest.NewRequest(echo.POST, "/post/create", strings.NewReader(test.ArgDataBody))
// 			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 			rec := httptest.NewRecorder()
// 			c := e.NewContext(req, rec)
// 			c.SetPath("/post/create")
// 			c.Set("user_id", test.ArgDataContext)

// 			err := handler.CreatePost(c)
// 			require.Equal(t, test.Error, err)

// 			if err == nil {
// 				assert.Equal(t, test.StatusCode, rec.Code)
// 			}
// 		})
// 	}

// 	mockUCase.AssertExpectations(t)
// }

// func TestDeliveryUpdatePost(t *testing.T) {
// 	mockPostValid := models.Post{ID: 2, Message: "123", Attachments: []models.Attachment{}}
// 	mockPostInValid := models.Post{Attachments: []models.Attachment{}}

// 	jsonPostValid, err := json.Marshal(mockPostValid)
// 	assert.NoError(t, err)
// 	jsonPostInValid, err := json.Marshal(mockPostInValid)
// 	assert.NoError(t, err)

// 	mockUCase := mocks.NewPostUseCaseI(t)

// 	mockUCase.On("UpdatePost", &mockPostValid).Return(nil)

// 	handler := postDelivery.Delivery{
// 		PUsecase: mockUCase,
// 	}

// 	e := echo.New()
// 	postDelivery.NewDelivery(e, mockUCase)

// 	cases := map[string]TestCaseCreatePost{
// 		"success": {
// 			ArgDataBody:    string(jsonPostValid),
// 			ArgDataContext: mockPostValid.UserID,
// 			Error:          nil,
// 			StatusCode:     http.StatusOK,
// 		},
// 		"bad_request": {
// 			ArgDataBody:    string(jsonPostInValid),
// 			ArgDataContext: mockPostValid.UserID,
// 			Error: &echo.HTTPError{
// 				Code:    http.StatusBadRequest,
// 				Message: models.ErrBadRequest.Error(),
// 			},
// 		},
// 	}

// 	for name, test := range cases {
// 		t.Run(name, func(t *testing.T) {
// 			req := httptest.NewRequest(echo.POST, "/post/edit", strings.NewReader(test.ArgDataBody))
// 			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 			rec := httptest.NewRecorder()
// 			c := e.NewContext(req, rec)
// 			c.SetPath("/post/edit")
// 			c.Set("user_id", test.ArgDataContext)

// 			err := handler.UpdatePost(c)
// 			require.Equal(t, test.Error, err)

// 			if err == nil {
// 				assert.Equal(t, test.StatusCode, rec.Code)
// 			}
// 		})
// 	}

// 	mockUCase.AssertExpectations(t)
// }

// func TestDeliveryDeletePost(t *testing.T) {
// 	var validPostID uint64 = 1
// 	var validUserID uint64 = 1

// 	mockUCase := mocks.NewPostUseCaseI(t)
// 	mockUCase.On("DeletePost", validPostID, validUserID).Return(nil)

// 	handler := postDelivery.Delivery{
// 		PUsecase: mockUCase,
// 	}

// 	e := echo.New()
// 	postDelivery.NewDelivery(e, mockUCase)

// 	cases := map[string]TestCaseDeletePost{
// 		"success": {
// 			ArgDataContext: validUserID,
// 			Error:          nil,
// 			StatusCode:     http.StatusNoContent,
// 			ID:             validPostID,
// 		},
// 		"invalid_context": {
// 			ID: validPostID,
// 			Error: &echo.HTTPError{
// 				Code:    http.StatusInternalServerError,
// 				Message: models.ErrInternalServerError.Error(),
// 			},
// 		},
// 	}

// 	for name, test := range cases {
// 		t.Run(name, func(t *testing.T) {
// 			req := httptest.NewRequest(echo.DELETE, "/post/:id", strings.NewReader(""))
// 			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 			rec := httptest.NewRecorder()
// 			c := e.NewContext(req, rec)
// 			c.SetPath("/post/:id")
// 			c.SetParamNames("id")
// 			c.SetParamValues(strconv.Itoa(int(test.ID)))
// 			if name != "invalid_context" {
// 				c.Set("user_id", test.ArgDataContext)
// 			}

// 			err := handler.DeletePost(c)
// 			require.Equal(t, test.Error, err)

// 			if err == nil {
// 				assert.Equal(t, test.StatusCode, rec.Code)
// 			}
// 		})
// 	}

// 	mockUCase.AssertExpectations(t)
// }

// func TestDeliveryFeed(t *testing.T) {
// 	var validUserID uint64 = 1

// 	mockUCase := mocks.NewPostUseCaseI(t)
// 	mockUCase.On("GetAllPosts", validUserID).Return([]*models.Post{}, nil)

// 	handler := postDelivery.Delivery{
// 		PUsecase: mockUCase,
// 	}

// 	e := echo.New()
// 	postDelivery.NewDelivery(e, mockUCase)

// 	cases := map[string]TestCaseFeed{
// 		"success": {
// 			ArgDataContext: validUserID,
// 			Error:          nil,
// 			StatusCode:     http.StatusOK,
// 		},
// 	}

// 	for name, test := range cases {
// 		t.Run(name, func(t *testing.T) {
// 			req := httptest.NewRequest(echo.GET, "/feed", strings.NewReader(""))
// 			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 			rec := httptest.NewRecorder()
// 			c := e.NewContext(req, rec)
// 			c.SetPath("/feed")
// 			c.Set("user_id", test.ArgDataContext)

// 			err := handler.Feed(c)
// 			require.Equal(t, test.Error, err)

// 			if err == nil {
// 				assert.Equal(t, test.StatusCode, rec.Code)
// 			}
// 		})
// 	}

// 	mockUCase.AssertExpectations(t)
// }
