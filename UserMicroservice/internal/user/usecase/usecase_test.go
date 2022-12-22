package usecase_test

import (
	"testing"

	"github.com/bxcodec/faker"
	userMocks "github.com/go-park-mail-ru/2022_2_TikTikIVProd/UserMicroservice/internal/user/repository/mocks"
	userUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/UserMicroservice/internal/user/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/UserMicroservice/models"
	user "github.com/go-park-mail-ru/2022_2_TikTikIVProd/UserMicroservice/proto"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestCaseSelectUserByNickName struct {
	ArgData *user.SelectUserByNickNameRequest
	ExpectedRes *user.User
	Error error
}

type TestCaseSelectUserByEmail struct {
	ArgData *user.SelectUserByEmailRequest
	ExpectedRes *user.User
	Error error
}

type TestCaseSelectUserById struct {
	ArgData *user.UserId
	ExpectedRes *user.User
	Error error
}

type TestCaseSelectAllUsers struct {
	ExpectedRes *user.UsersList
	Error error
}

type TestCaseSearchUsers struct {
	ArgData *user.SearchUsersRequest
	ExpectedRes *user.UsersList
	Error error
}

type TestCaseCreateUser struct {
	ArgData *user.User
	Error error
}

type TestCaseUpdateUser struct {
	ArgData *user.User
	Error error
}

type TestCaseAddFriend struct {
	ArgData *user.Friends
	Error error
}

type TestCaseDeleteFriend struct {
	ArgData *user.Friends
	Error error
}

type TestCaseCheckFriends struct {
	ArgData *user.Friends
	ExpectedRes *user.CheckFriendsResponse
	Error error
}

type TestCaseSelectFriends struct {
	ArgData *user.UserId
	ExpectedRes *user.UsersList
	Error error
}

func TestUsecaseSelectUserByNickName(t *testing.T) {
	var mockPbNickname user.SelectUserByNickNameRequest
	err := faker.FakeData(&mockPbNickname)
	assert.NoError(t, err)

	var mockPbUser user.User
	err = faker.FakeData(&mockPbUser)
	assert.NoError(t, err)

	usr := models.User {
		Id: mockPbUser.Id,
		FirstName: mockPbUser.FirstName,
		LastName: mockPbUser.LastName,
		NickName: mockPbUser.NickName,
		Avatar: mockPbUser.Avatar,
		Email: mockPbUser.Email,
		Password: mockPbUser.Password,
		CreatedAt: mockPbUser.CreatedAt.AsTime(),
	}

	var mockPbNicknameError user.SelectUserByNickNameRequest
	err = faker.FakeData(&mockPbNicknameError)
	assert.NoError(t, err)

	selectErr := errors.New("error")

	mockUserRepo := userMocks.NewRepositoryI(t)

	mockUserRepo.On("SelectUserByNickName", mockPbNickname.NickName).Return(&usr, nil)
	mockUserRepo.On("SelectUserByNickName", mockPbNicknameError.NickName).Return(nil, selectErr)

	useCase := userUsecase.New(mockUserRepo)

	cases := map[string]TestCaseSelectUserByNickName {
		"success": {
			ArgData:   &mockPbNickname,
			ExpectedRes: &mockPbUser,
			Error: nil,
		},
		"error": {
			ArgData:   &mockPbNicknameError,
			Error: selectErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			userId, err := useCase.SelectUserByNickName(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))

			if err == nil {
				assert.Equal(t, test.ExpectedRes, userId)
			}
		})
	}
	mockUserRepo.AssertExpectations(t)
}

func TestUsecaseSelectUserByEmail(t *testing.T) {
	var mockPbEmail user.SelectUserByEmailRequest
	err := faker.FakeData(&mockPbEmail)
	assert.NoError(t, err)

	var mockPbUser user.User
	err = faker.FakeData(&mockPbUser)
	assert.NoError(t, err)

	usr := models.User {
		Id: mockPbUser.Id,
		FirstName: mockPbUser.FirstName,
		LastName: mockPbUser.LastName,
		NickName: mockPbUser.NickName,
		Avatar: mockPbUser.Avatar,
		Email: mockPbUser.Email,
		Password: mockPbUser.Password,
		CreatedAt: mockPbUser.CreatedAt.AsTime(),
	}

	var mockPbEmailError user.SelectUserByEmailRequest
	err = faker.FakeData(&mockPbEmailError)
	assert.NoError(t, err)

	selectErr := errors.New("error")

	mockUserRepo := userMocks.NewRepositoryI(t)

	mockUserRepo.On("SelectUserByEmail", mockPbEmail.Email).Return(&usr, nil)
	mockUserRepo.On("SelectUserByEmail", mockPbEmailError.Email).Return(nil, selectErr)

	useCase := userUsecase.New(mockUserRepo)

	cases := map[string]TestCaseSelectUserByEmail {
		"success": {
			ArgData:   &mockPbEmail,
			ExpectedRes: &mockPbUser,
			Error: nil,
		},
		"error": {
			ArgData:   &mockPbEmailError,
			Error: selectErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			userId, err := useCase.SelectUserByEmail(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))

			if err == nil {
				assert.Equal(t, test.ExpectedRes, userId)
			}
		})
	}
	mockUserRepo.AssertExpectations(t)
}

func TestUsecaseSelectUserById(t *testing.T) {
	mockPbUserId := user.UserId {
		Id: 1,
	}

	var mockPbUser user.User
	err := faker.FakeData(&mockPbUser)
	assert.NoError(t, err)
	mockPbUser.Id = mockPbUserId.Id

	usr := models.User {
		Id: mockPbUser.Id,
		FirstName: mockPbUser.FirstName,
		LastName: mockPbUser.LastName,
		NickName: mockPbUser.NickName,
		Avatar: mockPbUser.Avatar,
		Email: mockPbUser.Email,
		Password: mockPbUser.Password,
		CreatedAt: mockPbUser.CreatedAt.AsTime(),
	}

	mockPbUserIdError := user.UserId {
		Id: 2,
	}

	selectErr := errors.New("error")

	mockUserRepo := userMocks.NewRepositoryI(t)

	mockUserRepo.On("SelectUserById", mockPbUserId.Id).Return(&usr, nil)
	mockUserRepo.On("SelectUserById", mockPbUserIdError.Id).Return(nil, selectErr)

	useCase := userUsecase.New(mockUserRepo)

	cases := map[string]TestCaseSelectUserById {
		"success": {
			ArgData:   &mockPbUserId,
			ExpectedRes: &mockPbUser,
			Error: nil,
		},
		"error": {
			ArgData:   &mockPbUserIdError,
			Error: selectErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			userId, err := useCase.SelectUserById(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))

			if err == nil {
				assert.Equal(t, test.ExpectedRes, userId)
			}
		})
	}
	mockUserRepo.AssertExpectations(t)
}

func TestUsecaseSelectAllUsers(t *testing.T) {
	var mockPbUsers user.UsersList
	err := faker.FakeData(&mockPbUsers)
	assert.NoError(t, err)

	users := make([]models.User, len(mockPbUsers.Users))

	for idx := range mockPbUsers.Users {
		usr := models.User {
			Id: mockPbUsers.Users[idx].Id,
			FirstName: mockPbUsers.Users[idx].FirstName,
			LastName: mockPbUsers.Users[idx].LastName,
			NickName: mockPbUsers.Users[idx].NickName,
			Avatar: mockPbUsers.Users[idx].Avatar,
			Email: mockPbUsers.Users[idx].Email,
			Password: mockPbUsers.Users[idx].Password,
			CreatedAt: mockPbUsers.Users[idx].CreatedAt.AsTime(),
		}

		users[idx] = usr
	}

	pbNothing := user.Nothing{Dummy: true}

	mockUserRepo := userMocks.NewRepositoryI(t)

	mockUserRepo.On("SelectAllUsers").Return(users, nil)

	useCase := userUsecase.New(mockUserRepo)

	cases := map[string]TestCaseSelectAllUsers {
		"success": {
			ExpectedRes: &mockPbUsers,
			Error: nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			usersList, err := useCase.SelectAllUsers(&pbNothing)
			require.Equal(t, test.Error, errors.Cause(err))

			if err == nil {
				assert.Equal(t, test.ExpectedRes, usersList)
			}
		})
	}
	mockUserRepo.AssertExpectations(t)
}

func TestUsecaseSearchUsers(t *testing.T) {
	var mockPbName user.SearchUsersRequest
	err := faker.FakeData(&mockPbName)
	assert.NoError(t, err)

	var mockPbUsers user.UsersList
	err = faker.FakeData(&mockPbUsers)
	assert.NoError(t, err)

	users := make([]models.User, len(mockPbUsers.Users))

	for idx := range mockPbUsers.Users {
		usr := models.User {
			Id: mockPbUsers.Users[idx].Id,
			FirstName: mockPbUsers.Users[idx].FirstName,
			LastName: mockPbUsers.Users[idx].LastName,
			NickName: mockPbUsers.Users[idx].NickName,
			Avatar: mockPbUsers.Users[idx].Avatar,
			Email: mockPbUsers.Users[idx].Email,
			Password: mockPbUsers.Users[idx].Password,
			CreatedAt: mockPbUsers.Users[idx].CreatedAt.AsTime(),
		}

		users[idx] = usr
	}

	var mockPbNameError user.SearchUsersRequest
	err = faker.FakeData(&mockPbNameError)
	assert.NoError(t, err)

	searchErr := errors.New("error")

	mockUserRepo := userMocks.NewRepositoryI(t)

	mockUserRepo.On("SearchUsers", mockPbName.Name).Return(users, nil)
	mockUserRepo.On("SearchUsers", mockPbNameError.Name).Return(nil, searchErr)

	useCase := userUsecase.New(mockUserRepo)

	cases := map[string]TestCaseSearchUsers {
		"success": {
			ArgData: &mockPbName,
			ExpectedRes: &mockPbUsers,
			Error: nil,
		},
		"error": {
			ArgData: &mockPbNameError,
			Error: searchErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			usersList, err := useCase.SearchUsers(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))

			if err == nil {
				assert.Equal(t, test.ExpectedRes, usersList)
			}
		})
	}
	mockUserRepo.AssertExpectations(t)
}

func TestUsecaseCreateUser(t *testing.T) {
	var mockPbUser user.User
	err := faker.FakeData(&mockPbUser)
	assert.NoError(t, err)

	usr := models.User {
		Id: mockPbUser.Id,
		FirstName: mockPbUser.FirstName,
		LastName: mockPbUser.LastName,
		NickName: mockPbUser.NickName,
		Avatar: mockPbUser.Avatar,
		Email: mockPbUser.Email,
		Password: mockPbUser.Password,
		CreatedAt: mockPbUser.CreatedAt.AsTime(),
	}

	var mockPbUserError user.User
	err = faker.FakeData(&mockPbUserError)
	assert.NoError(t, err)

	usrError := models.User {
		Id: mockPbUserError.Id,
		FirstName: mockPbUserError.FirstName,
		LastName: mockPbUserError.LastName,
		NickName: mockPbUserError.NickName,
		Avatar: mockPbUserError.Avatar,
		Email: mockPbUserError.Email,
		Password: mockPbUserError.Password,
		CreatedAt: mockPbUserError.CreatedAt.AsTime(),
	}

	createErr := errors.New("error")

	mockUserRepo := userMocks.NewRepositoryI(t)

	mockUserRepo.On("CreateUser", &usr).Return(nil)
	mockUserRepo.On("CreateUser", &usrError).Return(createErr)

	useCase := userUsecase.New(mockUserRepo)

	cases := map[string]TestCaseCreateUser {
		"success": {
			ArgData:   &mockPbUser,
			Error: nil,
		},
		"error": {
			ArgData:   &mockPbUserError,
			Error: createErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := useCase.CreateUser(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
		})
	}
	mockUserRepo.AssertExpectations(t)
}

func TestUsecaseUpdateUser(t *testing.T) {
	var mockPbUser user.User
	err := faker.FakeData(&mockPbUser)
	assert.NoError(t, err)

	usr := models.User {
		Id: mockPbUser.Id,
		FirstName: mockPbUser.FirstName,
		LastName: mockPbUser.LastName,
		NickName: mockPbUser.NickName,
		Avatar: mockPbUser.Avatar,
		Email: mockPbUser.Email,
		Password: mockPbUser.Password,
		CreatedAt: mockPbUser.CreatedAt.AsTime(),
	}

	var mockPbUserError user.User
	err = faker.FakeData(&mockPbUserError)
	assert.NoError(t, err)

	usrError := models.User {
		Id: mockPbUserError.Id,
		FirstName: mockPbUserError.FirstName,
		LastName: mockPbUserError.LastName,
		NickName: mockPbUserError.NickName,
		Avatar: mockPbUserError.Avatar,
		Email: mockPbUserError.Email,
		Password: mockPbUserError.Password,
		CreatedAt: mockPbUserError.CreatedAt.AsTime(),
	}

	createErr := errors.New("error")

	mockUserRepo := userMocks.NewRepositoryI(t)

	mockUserRepo.On("UpdateUser", usr).Return(nil)
	mockUserRepo.On("UpdateUser", usrError).Return(createErr)

	useCase := userUsecase.New(mockUserRepo)

	cases := map[string]TestCaseUpdateUser {
		"success": {
			ArgData:   &mockPbUser,
			Error: nil,
		},
		"error": {
			ArgData:   &mockPbUserError,
			Error: createErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := useCase.UpdateUser(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
		})
	}
	mockUserRepo.AssertExpectations(t)
}

func TestUsecaseAddFriend(t *testing.T) {
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

	addErr := errors.New("error")

	mockUserRepo := userMocks.NewRepositoryI(t)

	mockUserRepo.On("AddFriend", friends).Return(nil)
	mockUserRepo.On("AddFriend", friendsError).Return(addErr)

	useCase := userUsecase.New(mockUserRepo)

	cases := map[string]TestCaseAddFriend {
		"success": {
			ArgData:   &mockPbFriends,
			Error: nil,
		},
		"error": {
			ArgData:   &mockPbFriendsError,
			Error: addErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := useCase.AddFriend(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
		})
	}
	mockUserRepo.AssertExpectations(t)
}

func TestUsecaseDeleteFriend(t *testing.T) {
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

	deleteErr := errors.New("error")

	mockUserRepo := userMocks.NewRepositoryI(t)

	mockUserRepo.On("DeleteFriend", friends).Return(nil)
	mockUserRepo.On("DeleteFriend", friendsError).Return(deleteErr)

	useCase := userUsecase.New(mockUserRepo)

	cases := map[string]TestCaseAddFriend {
		"success": {
			ArgData:   &mockPbFriends,
			Error: nil,
		},
		"error": {
			ArgData:   &mockPbFriendsError,
			Error: deleteErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := useCase.DeleteFriend(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
		})
	}
	mockUserRepo.AssertExpectations(t)
}

func TestUsecaseCheckFriends(t *testing.T) {
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

	var mockPbCheckResp user.CheckFriendsResponse
	mockPbCheckResp.IsExists = true

	var mockPbCheckRespError user.CheckFriendsResponse
	mockPbCheckRespError.IsExists = false

	deleteErr := errors.New("error")

	mockUserRepo := userMocks.NewRepositoryI(t)

	mockUserRepo.On("CheckFriends", friends).Return(mockPbCheckResp.IsExists, nil)
	mockUserRepo.On("CheckFriends", friendsError).Return(mockPbCheckRespError.IsExists, deleteErr)

	useCase := userUsecase.New(mockUserRepo)

	cases := map[string]TestCaseCheckFriends {
		"success": {
			ArgData:   &mockPbFriends,
			ExpectedRes: &mockPbCheckResp,
			Error: nil,
		},
		"error": {
			ArgData:   &mockPbFriendsError,
			ExpectedRes: &mockPbCheckRespError,
			Error: deleteErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := useCase.CheckFriends(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
		})
	}
	mockUserRepo.AssertExpectations(t)
}

