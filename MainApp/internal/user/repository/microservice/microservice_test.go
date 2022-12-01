package microservice_test

import (
	"context"
	"testing"

	"github.com/bxcodec/faker"
	usersRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/user/repository/microservice"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	user "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/proto/user"
	userMocks "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/proto/user/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestCaseSelectUserByNickName struct {
	ArgData string
	ExpectedRes *models.User
	Error error
}

type TestCaseSelectUserByEmail struct {
	ArgData string
	ExpectedRes *models.User
	Error error
}

type TestCaseSelectUserById struct {
	ArgData uint64
	ExpectedRes *models.User
	Error error
}

type TestCaseSelectAllUsers struct {
	ExpectedRes []models.User
	Error error
}

type TestCaseSearchUsers struct {
	ArgData string
	ExpectedRes []models.User
	Error error
}

type TestCaseCreateUser struct {
	ArgData *models.User
	Error error
}

type TestCaseUpdateUser struct {
	ArgData models.User
	Error error
}

func TestMicroserviceCreateUser(t *testing.T) {
	var mockPbUser user.User
	err := faker.FakeData(&mockPbUser)
	assert.NoError(t, err)
	mockPbUser.Id = 1

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

	mockPbUserId := user.UserId {
		Id: mockPbUser.Id,
	}

	var mockPbUserError user.User
	err = faker.FakeData(&mockPbUserError)
	assert.NoError(t, err)
	mockPbUserError.Id = 2

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

	mockUserClient := userMocks.NewUsersClient(t)

	ctx := context.Background()

	createErr := errors.New("error")

	mockUserClient.On("CreateUser", ctx, &mockPbUser).Return(&mockPbUserId, nil)
	mockUserClient.On("CreateUser", ctx, &mockPbUserError).Return(nil, createErr)

	repository := usersRep.New(mockUserClient)

	cases := map[string]TestCaseCreateUser {
		"success": {
			ArgData:   &usr,
			Error: nil,
		},
		"error": {
			ArgData:   &usrError,
			Error: createErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := repository.CreateUser(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
		})
	}
	mockUserClient.AssertExpectations(t)
}

func TestMicroserviceUpdateUser(t *testing.T) {
	var mockPbUser user.User
	err := faker.FakeData(&mockPbUser)
	assert.NoError(t, err)
	mockPbUser.Id = 1

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

	pbNothing := user.Nothing{Dummy: true}

	var mockPbUserError user.User
	err = faker.FakeData(&mockPbUserError)
	assert.NoError(t, err)
	mockPbUserError.Id = 2

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

	mockUserClient := userMocks.NewUsersClient(t)

	ctx := context.Background()

	createErr := errors.New("error")

	mockUserClient.On("UpdateUser", ctx, &mockPbUser).Return(&pbNothing, nil)
	mockUserClient.On("UpdateUser", ctx, &mockPbUserError).Return(nil, createErr)

	repository := usersRep.New(mockUserClient)

	cases := map[string]TestCaseUpdateUser {
		"success": {
			ArgData:   usr,
			Error: nil,
		},
		"error": {
			ArgData:   usrError,
			Error: createErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := repository.UpdateUser(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
		})
	}
	mockUserClient.AssertExpectations(t)
}

func TestMicroserviceSelectUserByNickName(t *testing.T) {
	pbUserNickName := user.SelectUserByNickNameRequest {
		NickName: "nick1",
	}

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

	pbUserNickNameError := user.SelectUserByNickNameRequest {
		NickName: "nick2",
	}
	
	mockUserClient := userMocks.NewUsersClient(t)

	ctx := context.Background()

	selectErr := errors.New("error")

	mockUserClient.On("SelectUserByNickName", ctx, &pbUserNickName).Return(&mockPbUser, nil)
	mockUserClient.On("SelectUserByNickName", ctx, &pbUserNickNameError).Return(nil, selectErr)

	repository := usersRep.New(mockUserClient)

	cases := map[string]TestCaseSelectUserByNickName {
		"success": {
			ArgData:   pbUserNickName.NickName,
			ExpectedRes: &usr,
			Error: nil,
		},
		"error": {
			ArgData:   pbUserNickNameError.NickName,
			ExpectedRes: nil,
			Error: selectErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			selectedUser, err := repository.SelectUserByNickName(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
			assert.Equal(t, test.ExpectedRes, selectedUser)
		})
	}
	mockUserClient.AssertExpectations(t)
}

func TestMicroserviceSelectUserByEmail(t *testing.T) {
	pbUserEmail := user.SelectUserByEmailRequest {
		Email: "email1",
	}

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

	pbUserEmailError := user.SelectUserByEmailRequest {
		Email: "email2",
	}
	
	mockUserClient := userMocks.NewUsersClient(t)

	ctx := context.Background()

	selectErr := errors.New("error")

	mockUserClient.On("SelectUserByEmail", ctx, &pbUserEmail).Return(&mockPbUser, nil)
	mockUserClient.On("SelectUserByEmail", ctx, &pbUserEmailError).Return(nil, selectErr)

	repository := usersRep.New(mockUserClient)

	cases := map[string]TestCaseSelectUserByEmail {
		"success": {
			ArgData:   pbUserEmail.Email,
			ExpectedRes: &usr,
			Error: nil,
		},
		"error": {
			ArgData:   pbUserEmailError.Email,
			ExpectedRes: nil,
			Error: selectErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			selectedUser, err := repository.SelectUserByEmail(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
			assert.Equal(t, test.ExpectedRes, selectedUser)
		})
	}
	mockUserClient.AssertExpectations(t)
}

func TestMicroserviceSelectUserById(t *testing.T) {
	pbUserId := user.UserId {
		Id: 1,
	}

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

	pbUserIdError := user.UserId {
		Id: 2,
	}
	
	mockUserClient := userMocks.NewUsersClient(t)

	ctx := context.Background()

	selectErr := errors.New("error")

	mockUserClient.On("SelectUserById", ctx, &pbUserId).Return(&mockPbUser, nil)
	mockUserClient.On("SelectUserById", ctx, &pbUserIdError).Return(nil, selectErr)

	repository := usersRep.New(mockUserClient)

	cases := map[string]TestCaseSelectUserById {
		"success": {
			ArgData:   pbUserId.Id,
			ExpectedRes: &usr,
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
			selectedUser, err := repository.SelectUserById(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
			assert.Equal(t, test.ExpectedRes, selectedUser)
		})
	}
	mockUserClient.AssertExpectations(t)
}

func TestMicroserviceSelectAllUsers(t *testing.T) {
	var mockPbUsers user.UsersList
	err := faker.FakeData(&mockPbUsers)
	assert.NoError(t, err)

	users := make([]models.User, 0)

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

		users = append(users, usr)
	}

	pbNothing := user.Nothing{Dummy: true}
	
	mockUserClient := userMocks.NewUsersClient(t)

	ctx := context.Background()

	mockUserClient.On("SelectAllUsers", ctx, &pbNothing).Return(&mockPbUsers, nil)

	repository := usersRep.New(mockUserClient)

	cases := map[string]TestCaseSelectAllUsers {
		"success": {
			ExpectedRes: users,
			Error: nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			selectedUsers, err := repository.SelectAllUsers()
			require.Equal(t, test.Error, errors.Cause(err))
			assert.Equal(t, test.ExpectedRes, selectedUsers)
		})
	}
	mockUserClient.AssertExpectations(t)
}

func TestMicroserviceSearchUsers(t *testing.T) {
	pbName := user.SearchUsersRequest {
		Name: "name1",
	}

	var mockPbUsers user.UsersList
	err := faker.FakeData(&mockPbUsers)
	assert.NoError(t, err)

	users := make([]models.User, 0)

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

		users = append(users, usr)
	}

	pbNameError := user.SearchUsersRequest {
		Name: "name2",
	}

	searchErr := errors.New("error")
	
	mockUserClient := userMocks.NewUsersClient(t)

	ctx := context.Background()

	mockUserClient.On("SearchUsers", ctx, &pbName).Return(&mockPbUsers, nil)
	mockUserClient.On("SearchUsers", ctx, &pbNameError).Return(nil, searchErr)

	repository := usersRep.New(mockUserClient)

	cases := map[string]TestCaseSearchUsers {
		"success": {
			ArgData: pbName.Name,
			ExpectedRes: users,
			Error: nil,
		},
		"error": {
			ArgData: pbNameError.Name,
			ExpectedRes: nil,
			Error: searchErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			selectedUsers, err := repository.SearchUsers(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
			assert.Equal(t, test.ExpectedRes, selectedUsers)
		})
	}
	mockUserClient.AssertExpectations(t)
}

