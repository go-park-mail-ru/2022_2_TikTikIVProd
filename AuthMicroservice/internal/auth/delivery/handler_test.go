package delivery_test

import (
	"context"
	"testing"

	"github.com/bxcodec/faker"
	authDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/AuthMicroservice/internal/auth/delivery"
	authMocks "github.com/go-park-mail-ru/2022_2_TikTikIVProd/AuthMicroservice/internal/auth/usecase/mocks"
	auth "github.com/go-park-mail-ru/2022_2_TikTikIVProd/AuthMicroservice/proto"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestCaseGetCookie struct {
	ArgData *auth.ValueCookieRequest
	ExpectedRes *auth.GetCookieResponse
	Error error
}

type TestCaseDeleteCookie struct {
	ArgData *auth.ValueCookieRequest
	Error error
}

type TestCaseCreateCookie struct {
	ArgData *auth.Cookie
	Error error
}

func TestUsecaseCreateCookie(t *testing.T) {
	var mockPbCookieSuccess auth.Cookie
	err := faker.FakeData(&mockPbCookieSuccess)
	assert.NoError(t, err)

	var mockPbCookieError auth.Cookie
	err = faker.FakeData(&mockPbCookieError)
	assert.NoError(t, err)

	pbNothing := auth.Nothing{Dummy: true}

	createErr := errors.New("error")
	ctx := context.Background()

	mockAuthUsecase := authMocks.NewUseCaseI(t)

	mockAuthUsecase.On("CreateCookie", &mockPbCookieSuccess).Return(&pbNothing, nil)
	mockAuthUsecase.On("CreateCookie", &mockPbCookieError).Return(nil, createErr)

	delivery := authDelivery.New(mockAuthUsecase)

	cases := map[string]TestCaseCreateCookie {
		"success": {
			ArgData:   &mockPbCookieSuccess,
			Error: nil,
		},
		"error": {
			ArgData:   &mockPbCookieError,
			Error: createErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := delivery.CreateCookie(ctx, test.ArgData)
			require.Equal(t, test.Error, err)
		})
	}
	mockAuthUsecase.AssertExpectations(t)
}

func TestUsecaseGetCookie(t *testing.T) {
	var mockPbValueCookieSuccess auth.ValueCookieRequest
	err := faker.FakeData(&mockPbValueCookieSuccess)
	assert.NoError(t, err)

	var mockPbUserId auth.GetCookieResponse
	err = faker.FakeData(&mockPbUserId)
	assert.NoError(t, err)

	var mockPbValueCookieError auth.ValueCookieRequest
	err = faker.FakeData(&mockPbValueCookieError)
	assert.NoError(t, err)

	getErr := errors.New("error")

	ctx := context.Background()

	mockAuthUsecase := authMocks.NewUseCaseI(t)

	mockAuthUsecase.On("GetCookie", &mockPbValueCookieSuccess).Return(&mockPbUserId, nil)
	mockAuthUsecase.On("GetCookie", &mockPbValueCookieError).Return(nil, getErr)

	delivery := authDelivery.New(mockAuthUsecase)

	cases := map[string]TestCaseGetCookie {
		"success": {
			ArgData:   &mockPbValueCookieSuccess,
			ExpectedRes: &mockPbUserId,
			Error: nil,
		},
		"error": {
			ArgData:   &mockPbValueCookieError,
			Error: getErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			userId, err := delivery.GetCookie(ctx, test.ArgData)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.ExpectedRes, userId)
			}
		})
	}
	mockAuthUsecase.AssertExpectations(t)
}

func TestUsecaseDeleteCookie(t *testing.T) {
	var mockPbValueCookieSuccess auth.ValueCookieRequest
	err := faker.FakeData(&mockPbValueCookieSuccess)
	assert.NoError(t, err)

	var mockPbValueCookieError auth.ValueCookieRequest
	err = faker.FakeData(&mockPbValueCookieError)
	assert.NoError(t, err)

	deleteErr := errors.New("error")
	pbNothing := auth.Nothing{Dummy: true}

	ctx := context.Background()

	mockAuthUsecase := authMocks.NewUseCaseI(t)

	mockAuthUsecase.On("DeleteCookie", &mockPbValueCookieSuccess).Return(&pbNothing, nil)
	mockAuthUsecase.On("DeleteCookie", &mockPbValueCookieError).Return(nil, deleteErr)

	delivery := authDelivery.New(mockAuthUsecase)

	cases := map[string]TestCaseDeleteCookie {
		"success": {
			ArgData:   &mockPbValueCookieSuccess,
			Error: nil,
		},
		"error": {
			ArgData:   &mockPbValueCookieError,
			Error: deleteErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := delivery.DeleteCookie(ctx, test.ArgData)
			require.Equal(t, test.Error, err)
		})
	}
	mockAuthUsecase.AssertExpectations(t)
}

