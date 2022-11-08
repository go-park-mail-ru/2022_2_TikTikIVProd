package usecase_test

import (
	"testing"

	"github.com/bxcodec/faker"
	userMocks "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/repository/mocks"
	userUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestCaseSelectUser struct {
	ArgData int
	ExpectedRes *models.User
	Error error
}

type TestCaseUpdateUser struct {
	ArgData models.User
	Error error
}

type TestCaseSelectUsers struct {
	ExpectedRes []models.User
	Error error
}

func TestUsecaseSelectUserById(t *testing.T) {
	var mockUserRes models.User
	err := faker.FakeData(&mockUserRes)
	assert.NoError(t, err)

	mockExpectedUser := mockUserRes
	mockExpectedUser.Password = ""

	mockUserRepo := userMocks.NewRepositoryI(t)

	mockUserRepo.On("SelectUserById", mockUserRes.Id).Return(&mockUserRes, nil)

	useCase := userUsecase.New(mockUserRepo)

	cases := map[string]TestCaseSelectUser {
		"success": {
			ArgData:   mockUserRes.Id,
			ExpectedRes: &mockExpectedUser,
			Error: nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			user, err := useCase.SelectUserById(test.ArgData)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.ExpectedRes, user)
			}
		})
	}
	mockUserRepo.AssertExpectations(t)
}

func TestUsecaseUpdateUser(t *testing.T) {
	var mockUser models.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)

	mockUserRepo := userMocks.NewRepositoryI(t)

	mockUserRepo.On("SelectUserById", mockUser.Id).Return(&mockUser, nil)
	mockUserRepo.On("UpdateUser", mock.AnythingOfType("models.User")).Return(nil)

	useCase := userUsecase.New(mockUserRepo)

	cases := map[string]TestCaseUpdateUser {
		"success": {
			ArgData:   mockUser,
			Error: nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := useCase.UpdateUser(test.ArgData)
			require.Equal(t, test.Error, err)
		})
	}
	mockUserRepo.AssertExpectations(t)
}

func TestUsecaseSelectUsers(t *testing.T) {
	users := make([]models.User, 0, 10)
	err := faker.FakeData(&users)
	assert.NoError(t, err)

	mockUserRepo := userMocks.NewRepositoryI(t)

	mockUserRepo.On("SelectAllUsers").Return(users, nil)

	useCase := userUsecase.New(mockUserRepo)

	cases := map[string]TestCaseSelectUsers {
		"success": {
			ExpectedRes: users,
			Error: nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			users, err := useCase.SelectUsers()
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.ExpectedRes, users)
			}
		})
	}
	mockUserRepo.AssertExpectations(t)
}

