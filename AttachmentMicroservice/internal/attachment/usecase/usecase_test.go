package usecase_test

import (
	"testing"

	"github.com/bxcodec/faker"
	attachmentMocks "github.com/go-park-mail-ru/2022_2_TikTikIVProd/AttachmentMicroservice/internal/attachment/repository/mocks"
	attachmentUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/AttachmentMicroservice/internal/attachment/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/AttachmentMicroservice/models"
	Attachment "github.com/go-park-mail-ru/2022_2_TikTikIVProd/AttachmentMicroservice/proto"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestCaseGetPostAttachments struct {
	ArgData     *Attachment.GetPostAttachmentsRequest
	ExpectedRes *Attachment.GetPostAttachmentsResponse
	Error       error
}

type TestCaseGetMessageAttachments struct {
	ArgData     *Attachment.GetMessageAttachmentsRequest
	ExpectedRes *Attachment.GetMessageAttachmentsResponse
	Error       error
}

type TestCaseGetAttachment struct {
	ArgData     *Attachment.AttachmentId
	ExpectedRes *Attachment.Attachment
	Error       error
}

type TestCaseCreateAttachment struct {
	ArgData *Attachment.Attachment
	Error   error
}

func TestUsecaseCreateAttachment(t *testing.T) {
	var mockPbAttachmentSuccess Attachment.Attachment
	err := faker.FakeData(&mockPbAttachmentSuccess)
	assert.NoError(t, err)

	modelAttachmentSuccess := models.Attachment{
		ID:      mockPbAttachmentSuccess.Id,
		AttLink: mockPbAttachmentSuccess.AttLink,
		Type: mockPbAttachmentSuccess.Type,
	}

	var mockPbAttachmentError Attachment.Attachment
	err = faker.FakeData(&mockPbAttachmentError)
	assert.NoError(t, err)

	modelAttachmentError := models.Attachment{
		ID:      mockPbAttachmentError.Id,
		AttLink: mockPbAttachmentError.AttLink,
		Type: mockPbAttachmentError.Type,
	}

	createErr := errors.New("error")

	mockAttachmentRepo := attachmentMocks.NewRepositoryI(t)

	mockAttachmentRepo.On("CreateAttachment", &modelAttachmentSuccess).Return(nil)
	mockAttachmentRepo.On("CreateAttachment", &modelAttachmentError).Return(createErr)

	useCase := attachmentUsecase.New(mockAttachmentRepo)

	cases := map[string]TestCaseCreateAttachment{
		"success": {
			ArgData: &mockPbAttachmentSuccess,
			Error:   nil,
		},
		"error": {
			ArgData: &mockPbAttachmentError,
			Error:   createErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := useCase.CreateAttachment(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
		})
	}
	mockAttachmentRepo.AssertExpectations(t)
}

func TestUsecaseGetAttachment(t *testing.T) {
	mockPbAttachmentIdSuccess := Attachment.AttachmentId{
		AttachmentId: 1,
	}

	mockPbAttachmentSuccess := Attachment.Attachment{
		Id:      mockPbAttachmentIdSuccess.AttachmentId,
		AttLink: "link",
	}

	modelAttachmentSuccess := models.Attachment{
		ID:      mockPbAttachmentSuccess.Id,
		AttLink: mockPbAttachmentSuccess.AttLink,
	}

	mockPbAttachmentIdError := Attachment.AttachmentId{
		AttachmentId: 2,
	}

	getErr := errors.New("error")

	mockAttachmentRepo := attachmentMocks.NewRepositoryI(t)

	mockAttachmentRepo.On("GetAttachment", mockPbAttachmentIdSuccess.AttachmentId).Return(&modelAttachmentSuccess, nil)
	mockAttachmentRepo.On("GetAttachment", mockPbAttachmentIdError.AttachmentId).Return(nil, getErr)

	useCase := attachmentUsecase.New(mockAttachmentRepo)

	cases := map[string]TestCaseGetAttachment{
		"success": {
			ArgData:     &mockPbAttachmentIdSuccess,
			ExpectedRes: &mockPbAttachmentSuccess,
			Error:       nil,
		},
		"error": {
			ArgData: &mockPbAttachmentIdError,
			Error:   getErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			att, err := useCase.GetAttachment(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))

			if err == nil {
				assert.Equal(t, test.ExpectedRes, att)
			}
		})
	}
	mockAttachmentRepo.AssertExpectations(t)
}

func TestUsecaseGetPostAttachments(t *testing.T) {
	mockPbPostIdSuccess := Attachment.GetPostAttachmentsRequest{
		PostId: 1,
	}

	modelAttachments := make([]*models.Attachment, 0)
	err := faker.FakeData(&modelAttachments)
	assert.NoError(t, err)

	mockPbAttachments := &Attachment.GetPostAttachmentsResponse{}

	for idx := range modelAttachments {
		att := &Attachment.Attachment{
			Id:      modelAttachments[idx].ID,
			AttLink: modelAttachments[idx].AttLink,
			Type: modelAttachments[idx].Type,
		}
		mockPbAttachments.Attachments = append(mockPbAttachments.Attachments, att)
	}

	mockPbPostIdError := Attachment.GetPostAttachmentsRequest{
		PostId: 2,
	}

	getErr := errors.New("error")

	mockAttachmentRepo := attachmentMocks.NewRepositoryI(t)

	mockAttachmentRepo.On("GetPostAttachments", mockPbPostIdSuccess.PostId).Return(modelAttachments, nil)
	mockAttachmentRepo.On("GetPostAttachments", mockPbPostIdError.PostId).Return(nil, getErr)

	useCase := attachmentUsecase.New(mockAttachmentRepo)

	cases := map[string]TestCaseGetPostAttachments{
		"success": {
			ArgData:     &mockPbPostIdSuccess,
			ExpectedRes: mockPbAttachments,
			Error:       nil,
		},
		"error": {
			ArgData: &mockPbPostIdError,
			Error:   getErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			atts, err := useCase.GetPostAttachments(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))

			if err == nil {
				assert.Equal(t, test.ExpectedRes, atts)
			}
		})
	}
	mockAttachmentRepo.AssertExpectations(t)
}

func TestUsecaseGetMessageAttachments(t *testing.T) {
	mockPbMessageIdSuccess := Attachment.GetMessageAttachmentsRequest{
		MessageId: 1,
	}

	modelAttachments := make([]*models.Attachment, 0)
	err := faker.FakeData(&modelAttachments)
	assert.NoError(t, err)

	mockPbAttachments := &Attachment.GetMessageAttachmentsResponse{}

	for idx := range modelAttachments {
		att := &Attachment.Attachment{
			Id:      modelAttachments[idx].ID,
			AttLink: modelAttachments[idx].AttLink,
			Type: modelAttachments[idx].Type,
		}
		mockPbAttachments.Attachments = append(mockPbAttachments.Attachments, att)
	}

	mockPbMessageIdError := Attachment.GetMessageAttachmentsRequest{
		MessageId: 2,
	}

	getErr := errors.New("error")

	mockAttachmentRepo := attachmentMocks.NewRepositoryI(t)

	mockAttachmentRepo.On("GetMessageAttachments", mockPbMessageIdSuccess.MessageId).Return(modelAttachments, nil)
	mockAttachmentRepo.On("GetMessageAttachments", mockPbMessageIdError.MessageId).Return(nil, getErr)

	useCase := attachmentUsecase.New(mockAttachmentRepo)

	cases := map[string]TestCaseGetMessageAttachments{
		"success": {
			ArgData:     &mockPbMessageIdSuccess,
			ExpectedRes: mockPbAttachments,
			Error:       nil,
		},
		"error": {
			ArgData: &mockPbMessageIdError,
			Error:   getErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			atts, err := useCase.GetMessageAttachments(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))

			if err == nil {
				assert.Equal(t, test.ExpectedRes, atts)
			}
		})
	}
	mockAttachmentRepo.AssertExpectations(t)
}

