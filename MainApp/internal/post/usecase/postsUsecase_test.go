package usecase_test

import (
	"testing"

	"github.com/bxcodec/faker"
	postMocks "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/post/repository/mocks"
	imageMocks "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/image/repository/mocks"
	userMocks "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/user/repository/mocks"
	postUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/post/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestCaseGetPostById struct {
	ArgData uint64
	ExpectedRes *models.Post
	Error error
}
func TestUsecaseGetPostById(t *testing.T) {
	var mockPost models.Post
	err := faker.FakeData(&mockPost)
	assert.NoError(t, err)
	mockPost.Images = nil

	var mockUser models.User
	err = faker.FakeData(&mockUser)
	assert.NoError(t, err)

	mockImages := make([]*models.Image, 0, 10)
	err = faker.FakeData(&mockImages)
	assert.NoError(t, err)
	for _, image := range mockImages {
		mockPost.Images = append(mockPost.Images, *image)
	}

	mockPost.UserLastName = mockUser.LastName
	mockPost.AvatarID = mockUser.Avatar
	mockPost.UserFirstName = mockUser.FirstName

	mockPostRepo := postMocks.NewRepositoryI(t)
	mockImageRepo := imageMocks.NewRepositoryI(t)
	mockUserRepo := userMocks.NewRepositoryI(t)

	var userId uint64 = mockPost.UserID

	mockPostRepo.On("GetPostById", mockPost.ID).Return(&mockPost, nil)
	mockImageRepo.On("GetPostImages", mockPost.ID).Return(mockImages, nil)
	mockUserRepo.On("SelectUserById", mockPost.UserID).Return(&mockUser, nil)
	mockPostRepo.On("GetCountLikesPost", mockPost.ID).Return(uint64(50), nil)
	mockPostRepo.On("CheckLikePost", mockPost.ID, mockPost.UserID).Return(true, nil)

	useCase := postUsecase.NewPostUsecase(mockPostRepo, mockImageRepo, mockUserRepo)

	cases := map[string]TestCaseGetPostById {
		"success": {
			ArgData:   mockPost.ID,
			ExpectedRes: &mockPost,
			Error: nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			post, err := useCase.GetPostById(test.ArgData, userId)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.ExpectedRes, post)
			}
		})
	}
	mockPostRepo.AssertExpectations(t)
}

func TestUsecaseGetUserPosts(t *testing.T) {
	var mockPost models.Post
	err := faker.FakeData(&mockPost)
	assert.NoError(t, err)
	mockPost.Images = nil

	var mockUser models.User
	err = faker.FakeData(&mockUser)
	assert.NoError(t, err)

	mockImages := make([]*models.Image, 0, 10)
	err = faker.FakeData(&mockImages)
	assert.NoError(t, err)
	for _, image := range mockImages {
		mockPost.Images = append(mockPost.Images, *image)
	}

	mockPost.UserLastName = mockUser.LastName
	mockPost.AvatarID = mockUser.Avatar
	mockPost.UserFirstName = mockUser.FirstName

	var mockPosts []*models.Post
	mockPosts = append(mockPosts, &mockPost)

	mockPostRepo := postMocks.NewRepositoryI(t)
	mockImageRepo := imageMocks.NewRepositoryI(t)
	mockUserRepo := userMocks.NewRepositoryI(t)

	var userId uint64 = mockPost.UserID

	mockPostRepo.On("GetUserPosts", mockPost.UserID).Return(mockPosts, nil)
	mockImageRepo.On("GetPostImages", mockPost.ID).Return(mockImages, nil)
	mockUserRepo.On("SelectUserById", mockPost.UserID).Return(&mockUser, nil)
	mockPostRepo.On("GetCountLikesPost", mockPost.ID).Return(uint64(50), nil)
	mockPostRepo.On("CheckLikePost", mockPost.ID, mockPost.UserID).Return(true, nil)

	useCase := postUsecase.NewPostUsecase(mockPostRepo, mockImageRepo, mockUserRepo)

	cases := map[string]TestCaseGetPostById {
		"success": {
			ArgData:   mockPost.UserID,
			ExpectedRes: &mockPost,
			Error: nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := useCase.GetUserPosts(test.ArgData, userId)
			require.Equal(t, test.Error, err)
		})
	}
	mockPostRepo.AssertExpectations(t)
}

