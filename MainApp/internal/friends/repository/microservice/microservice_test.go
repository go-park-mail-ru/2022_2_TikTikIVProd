package microservice_test

import (
	"context"
	"testing"

	"github.com/bxcodec/faker"
	friendsRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/friends/repository/microservice"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	user "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/proto/user"
	userMocks "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/proto/user/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestCaseAddFriend struct {
	ArgData models.Friends
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

type TestCaseCheckFriends struct {
	ArgData models.Friends
	ExpectedRes bool
	Error error
}

func TestMicroserviceDeleteFriend(t *testing.T) {
	mockPbFriends := user.Friends {
		Id1: 1,
		Id2: 2,
	}

	friends := models.Friends {
		Id1: mockPbFriends.Id1,
		Id2: mockPbFriends.Id2,
	}

	mockPbFriendsError := user.Friends {
		Id1: 3,
		Id2: 4,
	}

	friendsError := models.Friends {
		Id1: mockPbFriendsError.Id1,
		Id2: mockPbFriendsError.Id2,
	}

	pbNothing := user.Nothing{Dummy: true}

	mockUserClient := userMocks.NewUsersClient(t)

	ctx := context.Background()

	deleteErr := errors.New("error")

	mockUserClient.On("DeleteFriend", ctx, &mockPbFriends).Return(&pbNothing, nil)
	mockUserClient.On("DeleteFriend", ctx, &mockPbFriendsError).Return(nil, deleteErr)

	repository := friendsRep.New(mockUserClient)

	cases := map[string]TestCaseDeleteFriend {
		"success": {
			ArgData:   friends,
			Error: nil,
		},
		"error": {
			ArgData:   friendsError,
			Error: deleteErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := repository.DeleteFriend(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
		})
	}
	mockUserClient.AssertExpectations(t)
}

func TestMicroserviceCheckFriends(t *testing.T) {
	mockPbFriends := user.Friends {
		Id1: 1,
		Id2: 2,
	}

	friends := models.Friends {
		Id1: mockPbFriends.Id1,
		Id2: mockPbFriends.Id2,
	}

	mockPbIsExists := user.CheckFriendsResponse {
		IsExists: true,
	}

	mockPbFriendsError := user.Friends {
		Id1: 3,
		Id2: 4,
	}

	friendsError := models.Friends {
		Id1: mockPbFriendsError.Id1,
		Id2: mockPbFriendsError.Id2,
	}

	mockPbIsExistsError := user.CheckFriendsResponse {
		IsExists: false,
	}

	mockUserClient := userMocks.NewUsersClient(t)

	ctx := context.Background()

	addErr := errors.New("error")

	mockUserClient.On("CheckFriends", ctx, &mockPbFriends).Return(&mockPbIsExists, nil)
	mockUserClient.On("CheckFriends", ctx, &mockPbFriendsError).Return(&mockPbIsExistsError, addErr)

	repository := friendsRep.New(mockUserClient)

	cases := map[string]TestCaseCheckFriends {
		"success": {
			ArgData:   friends,
			ExpectedRes: mockPbIsExists.IsExists,
			Error: nil,
		},
		"error": {
			ArgData:   friendsError,
			ExpectedRes: mockPbIsExistsError.IsExists,
			Error: addErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			isExists, err := repository.CheckFriends(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
			require.Equal(t, test.ExpectedRes, isExists)
		})
	}
	mockUserClient.AssertExpectations(t)
}

func TestMicroserviceAddFriend(t *testing.T) {
	mockPbFriends := user.Friends {
		Id1: 1,
		Id2: 2,
	}

	friends := models.Friends {
		Id1: mockPbFriends.Id1,
		Id2: mockPbFriends.Id2,
	}

	mockPbFriendsError := user.Friends {
		Id1: 3,
		Id2: 4,
	}

	friendsError := models.Friends {
		Id1: mockPbFriendsError.Id1,
		Id2: mockPbFriendsError.Id2,
	}

	pbNothing := user.Nothing{Dummy: true}

	mockUserClient := userMocks.NewUsersClient(t)

	ctx := context.Background()

	addErr := errors.New("error")

	mockUserClient.On("AddFriend", ctx, &mockPbFriends).Return(&pbNothing, nil)
	mockUserClient.On("AddFriend", ctx, &mockPbFriendsError).Return(nil, addErr)

	repository := friendsRep.New(mockUserClient)

	cases := map[string]TestCaseAddFriend {
		"success": {
			ArgData:   friends,
			Error: nil,
		},
		"error": {
			ArgData:   friendsError,
			Error: addErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := repository.AddFriend(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
		})
	}
	mockUserClient.AssertExpectations(t)
}

func TestMicroserviceSelectFriends(t *testing.T) {
	pbUserId := user.UserId {
		Id: 1,
	}

	var mockPbUsers user.UsersList
	err := faker.FakeData(&mockPbUsers)
	assert.NoError(t, err)

	friends := make([]models.User, 0)

	for idx := range mockPbUsers.Users {
		friend := models.User {
			Id: mockPbUsers.Users[idx].Id,
			FirstName: mockPbUsers.Users[idx].FirstName,
			LastName: mockPbUsers.Users[idx].LastName,
			NickName: mockPbUsers.Users[idx].NickName,
			Avatar: mockPbUsers.Users[idx].Avatar,
			Email: mockPbUsers.Users[idx].Email,
			Password: mockPbUsers.Users[idx].Password,
			CreatedAt: mockPbUsers.Users[idx].CreatedAt.AsTime(),
		}

		friends = append(friends, friend)
	}

	pbUserIdError := user.UserId {
		Id: 2,
	}
	
	mockUserClient := userMocks.NewUsersClient(t)

	ctx := context.Background()

	selectErr := errors.New("error")

	mockUserClient.On("SelectFriends", ctx, &pbUserId).Return(&mockPbUsers, nil)
	mockUserClient.On("SelectFriends", ctx, &pbUserIdError).Return(nil, selectErr)

	repository := friendsRep.New(mockUserClient)

	cases := map[string]TestCaseSelectFriends {
		"success": {
			ArgData:   pbUserId.Id,
			ExpectedRes: friends,
			Error: nil,
		},
		"error": {
			ArgData:   pbUserIdError.Id,
			ExpectedRes: nil,
			Error: selectErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			selectedFriends, err := repository.SelectFriends(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))

			if err == nil {
				assert.Equal(t, test.ExpectedRes, selectedFriends)
			}
		})
	}
	mockUserClient.AssertExpectations(t)
}

