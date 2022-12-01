package usecase_test

import (
	"testing"

	"github.com/bxcodec/faker"
	authMocks "github.com/go-park-mail-ru/2022_2_TikTikIVProd/AuthMicroservice/internal/auth/repository/mocks"
	authUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/AuthMicroservice/internal/auth/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/AuthMicroservice/models"
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

	modelCookieSuccess := models.Cookie{
		SessionToken: mockPbCookieSuccess.SessionToken,
		UserId:       mockPbCookieSuccess.UserId,
		MaxAge:       mockPbCookieSuccess.MaxAge,
	}

	var mockPbCookieError auth.Cookie
	err = faker.FakeData(&mockPbCookieError)
	assert.NoError(t, err)

	modelCookieError := models.Cookie{
		SessionToken: mockPbCookieError.SessionToken,
		UserId:       mockPbCookieError.UserId,
		MaxAge:       mockPbCookieError.MaxAge,
	}

	createErr := errors.New("error")

	mockAuthRepo := authMocks.NewRepositoryI(t)

	mockAuthRepo.On("CreateCookie", &modelCookieSuccess).Return(nil)
	mockAuthRepo.On("CreateCookie", &modelCookieError).Return(createErr)

	useCase := authUsecase.New(mockAuthRepo)

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
			_, err := useCase.CreateCookie(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
		})
	}
	mockAuthRepo.AssertExpectations(t)
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

	mockAuthRepo := authMocks.NewRepositoryI(t)

	mockAuthRepo.On("GetCookie", mockPbValueCookieSuccess.ValueCookie).Return(mockPbUserId.UserId, nil)
	mockAuthRepo.On("GetCookie", mockPbValueCookieError.ValueCookie).Return("", getErr)

	useCase := authUsecase.New(mockAuthRepo)

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
			userId, err := useCase.GetCookie(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))

			if err == nil {
				assert.Equal(t, test.ExpectedRes, userId)
			}
		})
	}
	mockAuthRepo.AssertExpectations(t)
}

func TestUsecaseDeleteCookie(t *testing.T) {
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

	mockAuthRepo := authMocks.NewRepositoryI(t)

	mockAuthRepo.On("DeleteCookie", mockPbValueCookieSuccess.ValueCookie).Return(nil)
	mockAuthRepo.On("DeleteCookie", mockPbValueCookieError.ValueCookie).Return(getErr)

	useCase := authUsecase.New(mockAuthRepo)

	cases := map[string]TestCaseDeleteCookie {
		"success": {
			ArgData:   &mockPbValueCookieSuccess,
			Error: nil,
		},
		"error": {
			ArgData:   &mockPbValueCookieError,
			Error: getErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := useCase.DeleteCookie(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
		})
	}
	mockAuthRepo.AssertExpectations(t)
}

