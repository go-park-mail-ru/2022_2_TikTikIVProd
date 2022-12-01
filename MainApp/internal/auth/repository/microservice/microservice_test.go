package microservice_test

import (
	"context"
	"testing"

	"github.com/bxcodec/faker"
	authRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/auth/repository/microservice"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	authMocks "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/proto/auth/mocks"
	auth "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/proto/auth"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestCaseCreateCookie struct {
	ArgData *models.Cookie
	Error error
}

type TestCaseGetCookie struct {
	ArgData string
	ExpectedRes string
	Error error
}

type TestCaseDeleteCookie struct {
	ArgData string
	Error error
}

func TestMicroserviceCreateCookie(t *testing.T) {
	var mockPbCookie auth.Cookie
	err := faker.FakeData(&mockPbCookie)
	assert.NoError(t, err)
	mockPbCookie.UserId = 1

	cookie := &models.Cookie {
		SessionToken: mockPbCookie.SessionToken,
		UserId: mockPbCookie.UserId,
		MaxAge: mockPbCookie.MaxAge,
	}

	var mockPbCookieError auth.Cookie
	err = faker.FakeData(&mockPbCookieError)
	assert.NoError(t, err)
	mockPbCookieError.UserId = 2

	cookieError := &models.Cookie {
		SessionToken: mockPbCookieError.SessionToken,
		UserId: mockPbCookieError.UserId,
		MaxAge: mockPbCookieError.MaxAge,
	}

	pbNothing := auth.Nothing{Dummy: true}

	mockAuthClient := authMocks.NewAuthClient(t)

	ctx := context.Background()

	createErr := errors.New("error")

	mockAuthClient.On("CreateCookie", ctx, &mockPbCookie).Return(&pbNothing, nil)
	mockAuthClient.On("CreateCookie", ctx, &mockPbCookieError).Return(nil, createErr)

	repository := authRep.New(mockAuthClient)

	cases := map[string]TestCaseCreateCookie {
		"success": {
			ArgData:   cookie,
			Error: nil,
		},
		"error": {
			ArgData:   cookieError,
			Error: createErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := repository.CreateCookie(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
		})
	}
	mockAuthClient.AssertExpectations(t)
}

func TestMicroserviceGetCookie(t *testing.T) {
	pbValueCookie := auth.ValueCookieRequest {
		ValueCookie: "cookie1",
	}
	pbResp := auth.GetCookieResponse {
		UserId: "1",
	}

	pbValueCookieError := auth.ValueCookieRequest {
		ValueCookie: "cookie2",
	}
	userId := ""
	
	mockAuthClient := authMocks.NewAuthClient(t)

	ctx := context.Background()

	getErr := errors.New("error")

	mockAuthClient.On("GetCookie", ctx, &pbValueCookie).Return(&pbResp, nil)
	mockAuthClient.On("GetCookie", ctx, &pbValueCookieError).Return(nil, getErr)

	repository := authRep.New(mockAuthClient)

	cases := map[string]TestCaseGetCookie {
		"success": {
			ArgData:   pbValueCookie.ValueCookie,
			ExpectedRes: pbResp.UserId,
			Error: nil,
		},
		"error": {
			ArgData:   pbValueCookieError.ValueCookie,
			ExpectedRes: userId,
			Error: getErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			id, err := repository.GetCookie(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))

			if err == nil {
				assert.Equal(t, test.ExpectedRes, id)
			}
		})
	}
	mockAuthClient.AssertExpectations(t)
}

func TestMicroserviceDeleteCookie(t *testing.T) {
	pbValueCookie := auth.ValueCookieRequest {
		ValueCookie: "cookie1",
	}

	pbValueCookieError := auth.ValueCookieRequest {
		ValueCookie: "cookie2",
	}

	pbNothing := auth.Nothing{Dummy: true}
	
	mockAuthClient := authMocks.NewAuthClient(t)

	ctx := context.Background()

	deleteErr := errors.New("error")

	mockAuthClient.On("DeleteCookie", ctx, &pbValueCookie).Return(&pbNothing, nil)
	mockAuthClient.On("DeleteCookie", ctx, &pbValueCookieError).Return(&pbNothing, deleteErr)

	repository := authRep.New(mockAuthClient)

	cases := map[string]TestCaseDeleteCookie {
		"success": {
			ArgData:   pbValueCookie.ValueCookie,
			Error: nil,
		},
		"error": {
			ArgData:   pbValueCookieError.ValueCookie,
			Error: deleteErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := repository.DeleteCookie(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
		})
	}
	mockAuthClient.AssertExpectations(t)
}

