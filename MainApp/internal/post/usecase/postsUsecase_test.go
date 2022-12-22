package usecase_test

import (
	"testing"

	"github.com/bxcodec/faker"
	attachmentMocks "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/attachment/repository/mocks"
	postMocks "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/post/repository/mocks"
	postUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/post/usecase"
	userMocks "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/user/repository/mocks"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestCaseGetPostById struct {
	ArgData     uint64
	ExpectedRes *models.Post
	Error       error
}

type TestCaseCreatePost struct {
	ArgData     *models.Post
	Error       error
}

type TestCaseAddComment struct {
	ArgData     *models.Comment
	Error       error
}

func TestUsecaseGetPostById(t *testing.T) {
	var mockPost models.Post
	err := faker.FakeData(&mockPost)
	assert.NoError(t, err)
	mockPost.Attachments = nil

	var mockUser models.User
	err = faker.FakeData(&mockUser)
	assert.NoError(t, err)

	mockAttachments := make([]*models.Attachment, 0, 10)
	err = faker.FakeData(&mockAttachments)
	assert.NoError(t, err)
	for _, Attachment := range mockAttachments {
		mockPost.Attachments = append(mockPost.Attachments, *Attachment)
	}

	mockPost.UserLastName = mockUser.LastName
	mockPost.AvatarID = mockUser.Avatar
	mockPost.UserFirstName = mockUser.FirstName

	mockPostRepo := postMocks.NewRepositoryI(t)
	mockAttachmentRepo := attachmentMocks.NewRepositoryI(t)
	mockUserRepo := userMocks.NewRepositoryI(t)

	var userId uint64 = mockPost.UserID

	mockPostRepo.On("GetPostById", mockPost.ID).Return(&mockPost, nil)
	mockAttachmentRepo.On("GetPostAttachments", mockPost.ID).Return(mockAttachments, nil)
	mockUserRepo.On("SelectUserById", mockPost.UserID).Return(&mockUser, nil)
	mockPostRepo.On("GetCountLikesPost", mockPost.ID).Return(uint64(50), nil)
	mockPostRepo.On("CheckLikePost", mockPost.ID, mockPost.UserID).Return(true, nil)

	useCase := postUsecase.NewPostUsecase(mockPostRepo, mockAttachmentRepo, mockUserRepo)

	cases := map[string]TestCaseGetPostById{
		"success": {
			ArgData:     mockPost.ID,
			ExpectedRes: &mockPost,
			Error:       nil,
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
	mockUserRepo.AssertExpectations(t)
	mockAttachmentRepo.AssertExpectations(t)
}

func TestUsecaseGetUserPosts(t *testing.T) {
	var mockPost models.Post
	err := faker.FakeData(&mockPost)
	assert.NoError(t, err)
	mockPost.Attachments = nil

	var mockUser models.User
	err = faker.FakeData(&mockUser)
	assert.NoError(t, err)

	mockAttachments := make([]*models.Attachment, 0, 10)
	err = faker.FakeData(&mockAttachments)
	assert.NoError(t, err)
	for _, Attachment := range mockAttachments {
		mockPost.Attachments = append(mockPost.Attachments, *Attachment)
	}

	mockPost.UserLastName = mockUser.LastName
	mockPost.AvatarID = mockUser.Avatar
	mockPost.UserFirstName = mockUser.FirstName

	var mockPosts []*models.Post
	mockPosts = append(mockPosts, &mockPost)

	mockPostRepo := postMocks.NewRepositoryI(t)
	mockAttachmentRepo := attachmentMocks.NewRepositoryI(t)
	mockUserRepo := userMocks.NewRepositoryI(t)

	var userId uint64 = mockPost.UserID

	mockPostRepo.On("GetUserPosts", mockPost.UserID).Return(mockPosts, nil)
	mockAttachmentRepo.On("GetPostAttachments", mockPost.ID).Return(mockAttachments, nil)
	mockUserRepo.On("SelectUserById", mockPost.UserID).Return(&mockUser, nil)
	mockPostRepo.On("GetCountLikesPost", mockPost.ID).Return(uint64(50), nil)
	mockPostRepo.On("CheckLikePost", mockPost.ID, mockPost.UserID).Return(true, nil)

	useCase := postUsecase.NewPostUsecase(mockPostRepo, mockAttachmentRepo, mockUserRepo)

	cases := map[string]TestCaseGetPostById{
		"success": {
			ArgData:     mockPost.UserID,
			ExpectedRes: &mockPost,
			Error:       nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := useCase.GetUserPosts(test.ArgData, userId)
			require.Equal(t, test.Error, err)
		})
	}
	mockPostRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
	mockAttachmentRepo.AssertExpectations(t)
}

func TestUsecaseGetCommunityPosts(t *testing.T) {
	var mockPost models.Post
	err := faker.FakeData(&mockPost)
	assert.NoError(t, err)
	mockPost.Attachments = nil

	var mockUser models.User
	err = faker.FakeData(&mockUser)
	assert.NoError(t, err)

	mockAttachments := make([]*models.Attachment, 0, 10)
	err = faker.FakeData(&mockAttachments)
	assert.NoError(t, err)
	for _, Attachment := range mockAttachments {
		mockPost.Attachments = append(mockPost.Attachments, *Attachment)
	}

	mockPost.UserLastName = mockUser.LastName
	mockPost.AvatarID = mockUser.Avatar
	mockPost.UserFirstName = mockUser.FirstName

	var mockPosts []*models.Post
	mockPosts = append(mockPosts, &mockPost)

	mockPostRepo := postMocks.NewRepositoryI(t)
	mockAttachmentRepo := attachmentMocks.NewRepositoryI(t)
	mockUserRepo := userMocks.NewRepositoryI(t)

	var userId uint64 = mockPost.UserID

	mockPostRepo.On("GetCommunityPosts", mockPost.UserID).Return(mockPosts, nil)
	mockAttachmentRepo.On("GetPostAttachments", mockPost.ID).Return(mockAttachments, nil)
	mockUserRepo.On("SelectUserById", mockPost.UserID).Return(&mockUser, nil)
	mockPostRepo.On("GetCountLikesPost", mockPost.ID).Return(uint64(50), nil)
	mockPostRepo.On("CheckLikePost", mockPost.ID, mockPost.UserID).Return(true, nil)

	useCase := postUsecase.NewPostUsecase(mockPostRepo, mockAttachmentRepo, mockUserRepo)

	cases := map[string]TestCaseGetPostById{
		"success": {
			ArgData:     mockPost.UserID,
			ExpectedRes: &mockPost,
			Error:       nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := useCase.GetCommunityPosts(test.ArgData, userId)
			require.Equal(t, test.Error, err)
		})
	}
	mockPostRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
	mockAttachmentRepo.AssertExpectations(t)
}

func TestUsecaseGetAllPosts(t *testing.T) {
	var mockPost models.Post
	err := faker.FakeData(&mockPost)
	assert.NoError(t, err)
	mockPost.Attachments = nil

	var mockUser models.User
	err = faker.FakeData(&mockUser)
	assert.NoError(t, err)

	mockAttachments := make([]*models.Attachment, 0, 10)
	err = faker.FakeData(&mockAttachments)
	assert.NoError(t, err)
	for _, Attachment := range mockAttachments {
		mockPost.Attachments = append(mockPost.Attachments, *Attachment)
	}

	mockPost.UserLastName = mockUser.LastName
	mockPost.AvatarID = mockUser.Avatar
	mockPost.UserFirstName = mockUser.FirstName

	var mockPosts []*models.Post
	mockPosts = append(mockPosts, &mockPost)

	mockPostRepo := postMocks.NewRepositoryI(t)
	mockAttachmentRepo := attachmentMocks.NewRepositoryI(t)
	mockUserRepo := userMocks.NewRepositoryI(t)

	mockPostRepo.On("GetAllPosts").Return(mockPosts, nil)
	mockAttachmentRepo.On("GetPostAttachments", mockPost.ID).Return(mockAttachments, nil)
	mockUserRepo.On("SelectUserById", mockPost.UserID).Return(&mockUser, nil)
	mockPostRepo.On("GetCountLikesPost", mockPost.ID).Return(uint64(50), nil)
	mockPostRepo.On("CheckLikePost", mockPost.ID, mockPost.UserID).Return(true, nil)

	useCase := postUsecase.NewPostUsecase(mockPostRepo, mockAttachmentRepo, mockUserRepo)

	cases := map[string]TestCaseGetPostById{
		"success": {
			ArgData:     mockPost.UserID,
			ExpectedRes: &mockPost,
			Error:       nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := useCase.GetAllPosts(test.ArgData)
			require.Equal(t, test.Error, err)
		})
	}
	mockPostRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
	mockAttachmentRepo.AssertExpectations(t)
}

func TestUsecaseGetComments(t *testing.T) {
	var mockComment models.Comment
	err := faker.FakeData(&mockComment)
	assert.NoError(t, err)

	var mockUser models.User
	err = faker.FakeData(&mockUser)
	assert.NoError(t, err)

	var mockPost models.Post
	err = faker.FakeData(&mockPost)
	assert.NoError(t, err)

	mockComment.PostID = mockPost.ID

	var mockComments []*models.Comment
	mockComments = append(mockComments, &mockComment)

	mockPostRepo := postMocks.NewRepositoryI(t)
	mockAttachmentRepo := attachmentMocks.NewRepositoryI(t)
	mockUserRepo := userMocks.NewRepositoryI(t)

	mockPostRepo.On("GetPostById", mockComment.PostID).Return(&mockPost, nil)
	mockPostRepo.On("GetComments", mockComment.PostID).Return(mockComments, nil)
	mockUserRepo.On("SelectUserById", mockComment.UserID).Return(&mockUser, nil)

	useCase := postUsecase.NewPostUsecase(mockPostRepo, mockAttachmentRepo, mockUserRepo)

	cases := map[string]TestCaseGetPostById{
		"success": {
			ArgData:     mockComment.PostID,
			Error:       nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := useCase.GetComments(test.ArgData)
			require.Equal(t, test.Error, err)
		})
	}
	mockPostRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestUsecaseCreatePost(t *testing.T) {
	var mockPost models.Post
	err := faker.FakeData(&mockPost)
	assert.NoError(t, err)

	var mockUser models.User
	err = faker.FakeData(&mockUser)
	assert.NoError(t, err)

	mockPostRepo := postMocks.NewRepositoryI(t)
	mockAttachmentRepo := attachmentMocks.NewRepositoryI(t)
	mockUserRepo := userMocks.NewRepositoryI(t)

	mockPostRepo.On("CreatePost", &mockPost).Return(nil)
	mockUserRepo.On("SelectUserById", mockPost.UserID).Return(&mockUser, nil)

	useCase := postUsecase.NewPostUsecase(mockPostRepo, mockAttachmentRepo, mockUserRepo)

	cases := map[string]TestCaseCreatePost{
		"success": {
			ArgData:     &mockPost,
			Error:       nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := useCase.CreatePost(test.ArgData)
			require.Equal(t, test.Error, err)
		})
	}
	mockPostRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestUsecaseUpdatePost(t *testing.T) {
	var mockPost models.Post
	err := faker.FakeData(&mockPost)
	assert.NoError(t, err)

	var mockUser models.User
	err = faker.FakeData(&mockUser)
	assert.NoError(t, err)

	mockPostRepo := postMocks.NewRepositoryI(t)
	mockAttachmentRepo := attachmentMocks.NewRepositoryI(t)
	mockUserRepo := userMocks.NewRepositoryI(t)

	mockPostRepo.On("GetPostById", mockPost.ID).Return(&mockPost, nil)
	mockPostRepo.On("UpdatePost", &mockPost).Return(nil)
	mockUserRepo.On("SelectUserById", mockPost.UserID).Return(&mockUser, nil)

	useCase := postUsecase.NewPostUsecase(mockPostRepo, mockAttachmentRepo, mockUserRepo)

	cases := map[string]TestCaseCreatePost{
		"success": {
			ArgData:     &mockPost,
			Error:       nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := useCase.UpdatePost(test.ArgData)
			require.Equal(t, test.Error, err)
		})
	}
	mockPostRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestUsecaseAddComment(t *testing.T) {
	var mockComment models.Comment
	err := faker.FakeData(&mockComment)
	assert.NoError(t, err)

	var mockUser models.User
	err = faker.FakeData(&mockUser)
	assert.NoError(t, err)

	mockPostRepo := postMocks.NewRepositoryI(t)
	mockAttachmentRepo := attachmentMocks.NewRepositoryI(t)
	mockUserRepo := userMocks.NewRepositoryI(t)

	mockPostRepo.On("AddComment", &mockComment).Return(nil)
	mockUserRepo.On("SelectUserById", mockComment.UserID).Return(&mockUser, nil)

	useCase := postUsecase.NewPostUsecase(mockPostRepo, mockAttachmentRepo, mockUserRepo)

	cases := map[string]TestCaseAddComment{
		"success": {
			ArgData:     &mockComment,
			Error:       nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := useCase.AddComment(test.ArgData)
			require.Equal(t, test.Error, err)
		})
	}
	mockPostRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestUsecaseDeletePost(t *testing.T) {
	var mockPost models.Post
	err := faker.FakeData(&mockPost)
	assert.NoError(t, err)
	mockPost.UserID = 1

	mockPostRepo := postMocks.NewRepositoryI(t)
	mockAttachmentRepo := attachmentMocks.NewRepositoryI(t)
	mockUserRepo := userMocks.NewRepositoryI(t)

	mockPostRepo.On("GetPostById", mockPost.ID).Return(&mockPost, nil)
	mockPostRepo.On("DeletePostById", mockPost.ID).Return(nil)

	useCase := postUsecase.NewPostUsecase(mockPostRepo, mockAttachmentRepo, mockUserRepo)

	cases := map[string]TestCaseGetPostById{
		"success": {
			ArgData:     mockPost.ID,
			Error:       nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := useCase.DeletePost(test.ArgData, 1)
			require.Equal(t, test.Error, err)
		})
	}
	mockPostRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestUsecaseDeleteComment(t *testing.T) {
	var mockComment models.Comment
	err := faker.FakeData(&mockComment)
	assert.NoError(t, err)
	mockComment.UserID = 1

	mockPostRepo := postMocks.NewRepositoryI(t)
	mockAttachmentRepo := attachmentMocks.NewRepositoryI(t)
	mockUserRepo := userMocks.NewRepositoryI(t)

	mockPostRepo.On("GetCommentById", mockComment.ID).Return(&mockComment, nil)
	mockPostRepo.On("DeleteComment", mockComment.ID).Return(nil)

	useCase := postUsecase.NewPostUsecase(mockPostRepo, mockAttachmentRepo, mockUserRepo)

	cases := map[string]TestCaseGetPostById{
		"success": {
			ArgData:     mockComment.ID,
			Error:       nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := useCase.DeleteComment(test.ArgData, 1)
			require.Equal(t, test.Error, err)
		})
	}
	mockPostRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestUsecaseLikePost(t *testing.T) {
	var mockPost models.Post
	err := faker.FakeData(&mockPost)
	assert.NoError(t, err)
	mockPost.UserID = 1

	mockPostRepo := postMocks.NewRepositoryI(t)
	mockAttachmentRepo := attachmentMocks.NewRepositoryI(t)
	mockUserRepo := userMocks.NewRepositoryI(t)

	var userId uint64 = 1

	mockPostRepo.On("GetPostById", mockPost.ID).Return(&mockPost, nil)
	mockPostRepo.On("LikePost", mockPost.ID, userId).Return(nil)

	useCase := postUsecase.NewPostUsecase(mockPostRepo, mockAttachmentRepo, mockUserRepo)

	cases := map[string]TestCaseGetPostById{
		"success": {
			ArgData:     mockPost.ID,
			Error:       nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := useCase.LikePost(test.ArgData, userId)
			require.Equal(t, test.Error, err)
		})
	}
	mockPostRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestUsecaseUnLikePost(t *testing.T) {
	var mockPost models.Post
	err := faker.FakeData(&mockPost)
	assert.NoError(t, err)
	mockPost.UserID = 1

	mockPostRepo := postMocks.NewRepositoryI(t)
	mockAttachmentRepo := attachmentMocks.NewRepositoryI(t)
	mockUserRepo := userMocks.NewRepositoryI(t)

	var userId uint64 = 1

	mockPostRepo.On("GetPostById", mockPost.ID).Return(&mockPost, nil)
	mockPostRepo.On("UnLikePost", mockPost.ID, userId).Return(nil)

	useCase := postUsecase.NewPostUsecase(mockPostRepo, mockAttachmentRepo, mockUserRepo)

	cases := map[string]TestCaseGetPostById{
		"success": {
			ArgData:     mockPost.ID,
			Error:       nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := useCase.UnLikePost(test.ArgData, userId)
			require.Equal(t, test.Error, err)
		})
	}
	mockPostRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

