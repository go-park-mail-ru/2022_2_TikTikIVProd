package usecase_test

// import (
// 	"testing"

// 	"github.com/bxcodec/faker"
// 	attachmentMocks "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/attachment/repository/mocks"
// 	postMocks "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/post/repository/mocks"
// 	postUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/post/usecase"
// 	userMocks "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/user/repository/mocks"
// 	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// )

// type TestCaseGetPostById struct {
// 	ArgData     uint64
// 	ExpectedRes *models.Post
// 	Error       error
// }

// func TestUsecaseGetPostById(t *testing.T) {
// 	var mockPost models.Post
// 	err := faker.FakeData(&mockPost)
// 	assert.NoError(t, err)
// 	mockPost.Attachments = nil

// 	var mockUser models.User
// 	err = faker.FakeData(&mockUser)
// 	assert.NoError(t, err)

// 	mockAttachments := make([]*models.Attachment, 0, 10)
// 	err = faker.FakeData(&mockAttachments)
// 	assert.NoError(t, err)
// 	for _, Attachment := range mockAttachments {
// 		mockPost.Attachments = append(mockPost.Attachments, *Attachment)
// 	}

// 	mockPost.UserLastName = mockUser.LastName
// 	mockPost.AvatarID = mockUser.Avatar
// 	mockPost.UserFirstName = mockUser.FirstName

// 	mockPostRepo := postMocks.NewRepositoryI(t)
// 	mockAttachmentRepo := attachmentMocks.NewRepositoryI(t)
// 	mockUserRepo := userMocks.NewRepositoryI(t)

// 	var userId uint64 = mockPost.UserID

// 	mockPostRepo.On("GetPostById", mockPost.ID).Return(&mockPost, nil)
// 	mockAttachmentRepo.On("GetPostAttachments", mockPost.ID).Return(mockAttachments, nil)
// 	mockUserRepo.On("SelectUserById", mockPost.UserID).Return(&mockUser, nil)
// 	mockPostRepo.On("GetCountLikesPost", mockPost.ID).Return(uint64(50), nil)
// 	mockPostRepo.On("CheckLikePost", mockPost.ID, mockPost.UserID).Return(true, nil)

// 	useCase := postUsecase.NewPostUsecase(mockPostRepo, mockAttachmentRepo, mockUserRepo)

// 	cases := map[string]TestCaseGetPostById{
// 		"success": {
// 			ArgData:     mockPost.ID,
// 			ExpectedRes: &mockPost,
// 			Error:       nil,
// 		},
// 	}

// 	for name, test := range cases {
// 		t.Run(name, func(t *testing.T) {
// 			post, err := useCase.GetPostById(test.ArgData, userId)
// 			require.Equal(t, test.Error, err)

// 			if err == nil {
// 				assert.Equal(t, test.ExpectedRes, post)
// 			}
// 		})
// 	}
// 	mockPostRepo.AssertExpectations(t)
// }

// func TestUsecaseGetUserPosts(t *testing.T) {
// 	var mockPost models.Post
// 	err := faker.FakeData(&mockPost)
// 	assert.NoError(t, err)
// 	mockPost.Attachments = nil

// 	var mockUser models.User
// 	err = faker.FakeData(&mockUser)
// 	assert.NoError(t, err)

// 	mockAttachments := make([]*models.Attachment, 0, 10)
// 	err = faker.FakeData(&mockAttachments)
// 	assert.NoError(t, err)
// 	for _, Attachment := range mockAttachments {
// 		mockPost.Attachments = append(mockPost.Attachments, *Attachment)
// 	}

// 	mockPost.UserLastName = mockUser.LastName
// 	mockPost.AvatarID = mockUser.Avatar
// 	mockPost.UserFirstName = mockUser.FirstName

// 	var mockPosts []*models.Post
// 	mockPosts = append(mockPosts, &mockPost)

// 	mockPostRepo := postMocks.NewRepositoryI(t)
// 	mockAttachmentRepo := attachmentMocks.NewRepositoryI(t)
// 	mockUserRepo := userMocks.NewRepositoryI(t)

// 	var userId uint64 = mockPost.UserID

// 	mockPostRepo.On("GetUserPosts", mockPost.UserID).Return(mockPosts, nil)
// 	mockAttachmentRepo.On("GetPostAttachments", mockPost.ID).Return(mockAttachments, nil)
// 	mockUserRepo.On("SelectUserById", mockPost.UserID).Return(&mockUser, nil)
// 	mockPostRepo.On("GetCountLikesPost", mockPost.ID).Return(uint64(50), nil)
// 	mockPostRepo.On("CheckLikePost", mockPost.ID, mockPost.UserID).Return(true, nil)

// 	useCase := postUsecase.NewPostUsecase(mockPostRepo, mockAttachmentRepo, mockUserRepo)

// 	cases := map[string]TestCaseGetPostById{
// 		"success": {
// 			ArgData:     mockPost.UserID,
// 			ExpectedRes: &mockPost,
// 			Error:       nil,
// 		},
// 	}

// 	for name, test := range cases {
// 		t.Run(name, func(t *testing.T) {
// 			_, err := useCase.GetUserPosts(test.ArgData, userId)
// 			require.Equal(t, test.Error, err)
// 		})
// 	}
// 	mockPostRepo.AssertExpectations(t)
// }
