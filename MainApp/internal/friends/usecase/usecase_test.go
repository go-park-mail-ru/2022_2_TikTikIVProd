package usecase_test

import (
	"testing"

	"github.com/bxcodec/faker"
	friendsMocks "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/friends/repository/mocks"
	friendsUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/friends/usecase"
	userMocks "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/user/repository/mocks"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestCaseAddFriend struct {
	ArgData models.Friends
	Error error
}

type TestCaseCheckIsFriend struct {
	ArgData models.Friends
	ExpectedRes bool
	Error error
}

type TestCaseDeleteFriend struct {
	ArgData models.Friends
	Error error
}

type TestCaseSelectFriends struct {
	ArgData uint64
	ExpectedRes []models.User
	Error error
}

func TestUsecaseAddFriend(t *testing.T) {
	mockFriendsSuccess := models.Friends {
		Id1: 1,
		Id2: 2,
	}

	mockFriendsConflict := models.Friends {
		Id1: 3,
		Id2: 4,
	}

	mockFriendsBadRequest := models.Friends {
		Id1: 5,
		Id2: 5,
	}

	mockFriendsRepo := friendsMocks.NewRepositoryI(t)
	mockUserRepo := userMocks.NewRepositoryI(t)

	mockUserRepo.On("SelectUserById", mockFriendsSuccess.Id2).Return(nil, nil)
	mockFriendsRepo.On("CheckFriends", mockFriendsSuccess).Return(false, nil)
	mockFriendsRepo.On("AddFriend", mockFriendsSuccess).Return(nil)

	mockUserRepo.On("SelectUserById", mockFriendsConflict.Id2).Return(nil, nil)
	mockFriendsRepo.On("CheckFriends", mockFriendsConflict).Return(true, nil)

	useCase := friendsUsecase.New(mockFriendsRepo, mockUserRepo)

	cases := map[string]TestCaseAddFriend {
		"success": {
			ArgData:   mockFriendsSuccess,
			Error: nil,
		},
		"conflict": {
			ArgData:   mockFriendsConflict,
			Error: models.ErrConflictFriend,
		},
		"bad_request": {
			ArgData:   mockFriendsBadRequest,
			Error: models.ErrBadRequest,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := useCase.AddFriend(test.ArgData)
			require.Equal(t, test.Error, err)
		})
	}
	mockUserRepo.AssertExpectations(t)
	mockFriendsRepo.AssertExpectations(t)
}

func TestUsecaseDeleteFriend(t *testing.T) {
	mockFriendsSuccess := models.Friends {
		Id1: 1,
		Id2: 2,
	}

	mockFriendsNotFound := models.Friends {
		Id1: 3,
		Id2: 4,
	}

	mockFriendsBadRequest := models.Friends {
		Id1: 5,
		Id2: 5,
	}

	mockFriendsRepo := friendsMocks.NewRepositoryI(t)
	mockUserRepo := userMocks.NewRepositoryI(t)

	mockUserRepo.On("SelectUserById", mockFriendsSuccess.Id2).Return(nil, nil)
	mockFriendsRepo.On("CheckFriends", mockFriendsSuccess).Return(true, nil)
	mockFriendsRepo.On("DeleteFriend", mockFriendsSuccess).Return(nil)

	mockUserRepo.On("SelectUserById", mockFriendsNotFound.Id2).Return(nil, nil)
	mockFriendsRepo.On("CheckFriends", mockFriendsNotFound).Return(false, nil)

	useCase := friendsUsecase.New(mockFriendsRepo, mockUserRepo)

	cases := map[string]TestCaseDeleteFriend {
		"success": {
			ArgData:   mockFriendsSuccess,
			Error: nil,
		},
		"not_found": {
			ArgData:   mockFriendsNotFound,
			Error: models.ErrNotFound,
		},
		"bad_request": {
			ArgData:   mockFriendsBadRequest,
			Error: models.ErrBadRequest,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := useCase.DeleteFriend(test.ArgData)
			require.Equal(t, test.Error, err)
		})
	}
	mockUserRepo.AssertExpectations(t)
	mockFriendsRepo.AssertExpectations(t)
}

func TestUsecaseSelectFriends(t *testing.T) {
	mockFriends := make([]models.User, 0, 10)
	err := faker.FakeData(&mockFriends)
	assert.NoError(t, err)

	var userId uint64 = 1

	mockFriendsRepo := friendsMocks.NewRepositoryI(t)
	mockUserRepo := userMocks.NewRepositoryI(t)

	mockUserRepo.On("SelectUserById", userId).Return(nil, nil)
	mockFriendsRepo.On("SelectFriends", userId).Return(mockFriends, nil)

	useCase := friendsUsecase.New(mockFriendsRepo, mockUserRepo)

	cases := map[string]TestCaseSelectFriends {
		"success": {
			ArgData:   userId,
			ExpectedRes: mockFriends,
			Error: nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			friends, err := useCase.SelectFriends(test.ArgData)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.ExpectedRes, friends)
			}
		})
	}
	mockUserRepo.AssertExpectations(t)
	mockFriendsRepo.AssertExpectations(t)
}

func TestUsecaseCheckIsFriend(t *testing.T) {
	mockFriendsSuccess := models.Friends {
		Id1: 1,
		Id2: 2,
	}

	mockFriendsBadRequest := models.Friends {
		Id1: 5,
		Id2: 5,
	}

	mockFriendsRepo := friendsMocks.NewRepositoryI(t)
	mockUserRepo := userMocks.NewRepositoryI(t)

	mockUserRepo.On("SelectUserById", mockFriendsSuccess.Id2).Return(nil, nil)
	mockFriendsRepo.On("CheckFriends", mockFriendsSuccess).Return(true, nil)

	useCase := friendsUsecase.New(mockFriendsRepo, mockUserRepo)

	cases := map[string]TestCaseCheckIsFriend {
		"success": {
			ArgData:   mockFriendsSuccess,
			ExpectedRes: true,
			Error: nil,
		},
		"bad_request": {
			ArgData:   mockFriendsBadRequest,
			ExpectedRes: false,
			Error: models.ErrBadRequest,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			friendExists, err := useCase.CheckIsFriend(test.ArgData)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.ExpectedRes, friendExists)
			}
		})
	}
	mockUserRepo.AssertExpectations(t)
	mockFriendsRepo.AssertExpectations(t)
}

