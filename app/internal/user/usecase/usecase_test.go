package usecase_test

// import (
// 	"testing"

// 	"github.com/bxcodec/faker"
// 	userUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/usecase"
// 	userMocks "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/repository/mocks"
// 	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// 	"golang.org/x/crypto/bcrypt"
// )

// type TestCaseSignUp struct {
// 	ArgData *models.User
// 	ExpectedRes int
// 	Error error
// }

// type TestCaseSignIn struct {
// 	ArgData models.UserSignIn
// 	ExpectedResUser *models.User
// 	ExpectedResCookie int
// 	Error error
// }

// type TestCaseDeleteCookie struct {
// 	Data string
// 	Error error
// }

// type TestCaseAuth struct {
// 	Data string
// 	Expected models.User
// 	Error error
// }

// func TestUsecaseSignUp(t *testing.T) {
// 	var mockUserSuccess models.User
// 	err := faker.FakeData(&mockUserSuccess)
// 	assert.NoError(t, err)

// 	var mockUserBadRequest models.User
// 	err = faker.FakeData(&mockUserBadRequest)
// 	assert.NoError(t, err)
// 	mockUserBadRequest.Email = ""

// 	var mockUserConflictNickName models.User
// 	err = faker.FakeData(&mockUserConflictNickName)
// 	assert.NoError(t, err)

// 	var mockUserConflictEmail models.User
// 	err = faker.FakeData(&mockUserConflictEmail)
// 	assert.NoError(t, err)

// 	mockUserRepo := new(userMocks.RepositoryI)

// 	mockUserRepo.On("SelectUserByNickName", mockUserSuccess.NickName).Return(nil, models.ErrNotFound)
// 	mockUserRepo.On("SelectUserByEmail", mockUserSuccess.Email).Return(nil, models.ErrNotFound)
// 	mockUserRepo.On("CreateUser", &mockUserSuccess).Return(nil)

// 	mockUserRepo.On("SelectUserByNickName", mockUserConflictNickName.NickName).Return(&mockUserConflictNickName, nil)
	
// 	mockUserRepo.On("SelectUserByNickName", mockUserConflictEmail.NickName).Return(nil, models.ErrNotFound)
// 	mockUserRepo.On("SelectUserByEmail", mockUserConflictEmail.Email).Return(&mockUserConflictEmail, nil)

// 	useCase := userUsecase.New(mockUserRepo)

// 	cases := map[string]TestCaseSignUp {
// 		"success": {
// 			ArgData:   &mockUserSuccess,
// 			ExpectedRes: mockUserSuccess.Id,
// 			Error: nil,
// 		},
// 		"bad_request": {
// 			ArgData:   &mockUserBadRequest,
// 			Error: models.ErrBadRequest,
// 		},
// 		"conflict_nickname": {
// 			ArgData:   &mockUserConflictNickName,
// 			Error: models.ErrConflictNickname,
// 		},
// 		"conflict_email": {
// 			ArgData:   &mockUserConflictEmail,
// 			Error: models.ErrConflictEmail,
// 		},
// 	}

// 	for name, test := range cases {
// 		t.Run(name, func(t *testing.T) {
// 			err := useCase.SignUp(test.ArgData)
// 			require.Equal(t, test.Error, err)

// 			// if err == nil {
// 			// 	assert.Equal(t, test.ExpectedRes, cookie.UserId)
// 			// }
// 		})
// 	}
// 	mockUserRepo.AssertExpectations(t)
// }

// func TestUsecaseSignIn(t *testing.T) {
// 	var mockUser models.User

// 	err := faker.FakeData(&mockUser)
// 	assert.NoError(t, err)

// 	var mockUserSignIn models.UserSignIn
// 	mockUserSignIn.Email = mockUser.Email
// 	mockUserSignIn.Password = mockUser.Password

// 	var mockUserSignInBadRequest models.UserSignIn
// 	mockUserSignInBadRequest.Email = mockUser.Email
// 	mockUserSignInBadRequest.Password = ""

// 	var mockUserSignInInvalidPassword models.UserSignIn
// 	mockUserSignInInvalidPassword.Email = mockUser.Email
// 	mockUserSignInInvalidPassword.Password = "dfvdf"

// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(mockUser.Password), 8)
// 	assert.NoError(t, err)

// 	mockUser.Password = string(hashedPassword)

// 	mockUserRepo := new(userMocks.RepositoryI)

// 	mockUserRepo.On("SelectUserByEmail", mockUserSignIn.Email).Return(&mockUser, nil)
// 	mockUserRepo.On("SelectUserByEmail", mockUserSignInInvalidPassword.Email).Return(&mockUser, nil)

// 	useCase := userUsecase.New(mockUserRepo)

// 	expectedUser := mockUser
// 	expectedUser.Password = ""

// 	cases := map[string]TestCaseSignIn {
// 		"invalid_password": {
// 			ArgData:   mockUserSignInInvalidPassword,
// 			Error: models.ErrInvalidPassword,
// 		},
// 		"bad_request": {
// 			ArgData:   mockUserSignInBadRequest,
// 			Error: models.ErrBadRequest,
// 		},
// 		"success": {
// 			ArgData:   mockUserSignIn,
// 			ExpectedResUser:   &expectedUser,
// 			ExpectedResCookie: mockUser.Id,
// 			Error: nil,
// 		},
// 	}

// 	for name, test := range cases {
// 		t.Run(name, func(t *testing.T) {
// 			user, err := useCase.SignIn(test.ArgData)
// 			require.Equal(t, test.Error, err)

// 			if err == nil {
// 				assert.Equal(t, test.ExpectedResUser, user)
// 			}
// 		})
// 	}
// 	mockUserRepo.AssertExpectations(t)
// }

// // func TestUsecaseSelectUserById(t *testing.T) {
// // 	var cookie models.Cookie

// // 	err := faker.FakeData(&cookie)
// // 	assert.NoError(t, err)

// // 	mockUserRepo := new(userMocks.RepositoryI)

// // 	user := models.User {
// // 		Id: cookie.UserId,
// // 	}

// // 	mockUserRepo.On("SelectUserById", cookie.UserId).Return(&user, nil)

// // 	useCase := userUsecase.New(mockUserRepo)

// // 	cases := map[string]TestCaseDeleteCookie {
// // 		"success": {
// // 			Data:   cookie.SessionToken,
// // 			Error: nil,
// // 		},
// // 	}

// // 	for name, test := range cases {
// // 		t.Run(name, func(t *testing.T) {
// // 			user, err := useCase.SelectUserById(1)
// // 			require.Equal(t, test.Error, err)
// // 		})
// // 	}
// // 	mockAuthRepo.AssertExpectations(t)
// // }

// // func TestUsecaseAuth(t *testing.T) {
// // 	var cookie models.Cookie

// // 	err := faker.FakeData(&cookie)
// // 	assert.NoError(t, err)

// // 	mockAuthRepo := new(authMocks.RepositoryI)
// // 	mockUserRepo := new(userMocks.RepositoryI)

// // 	var user models.User
// // 	err = faker.FakeData(&user)
// // 	assert.NoError(t, err)

// // 	user.Id = cookie.UserId

// // 	mockUserRepo.On("SelectUserById", cookie.UserId).Return(&user, nil)
// // 	mockAuthRepo.On("SelectCookie", cookie.SessionToken).Return(&cookie, nil)

// // 	user.Password = ""

// // 	useCase := authUsecase.New(mockAuthRepo, mockUserRepo)

// // 	cases := map[string]TestCaseAuth {
// // 		"success": {
// // 			Data:   cookie.SessionToken,
// // 			Expected: user,
// // 			Error: nil,
// // 		},
// // 	}

// // 	for name, test := range cases {
// // 		t.Run(name, func(t *testing.T) {
// // 			gotUser, err := useCase.Auth(test.Data)
// // 			require.Equal(t, test.Error, err)
// // 			if err == nil {
// // 				assert.Equal(t, &test.Expected, gotUser)
// // 			}
// // 		})
// // 	}
// // 	mockAuthRepo.AssertExpectations(t)
// // }

