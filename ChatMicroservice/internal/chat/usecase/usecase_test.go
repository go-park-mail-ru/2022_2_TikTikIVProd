package usecase_test

import (
	"testing"

	"github.com/bxcodec/faker"
	chatMocks "github.com/go-park-mail-ru/2022_2_TikTikIVProd/ChatMicroservice/internal/chat/repository/mocks"
	chatUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/ChatMicroservice/internal/chat/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/ChatMicroservice/models"
	chat "github.com/go-park-mail-ru/2022_2_TikTikIVProd/ChatMicroservice/proto"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestCaseSelectDialog struct {
	ArgData *chat.DialogId
	ExpectedRes *chat.Dialog
	Error error
}

type TestCaseSelectDialogByUsers struct {
	ArgData *chat.SelectDialogByUsersRequest
	ExpectedRes *chat.Dialog
	Error error
}

type TestCaseSelectAllDialogs struct {
	ArgData *chat.SelectAllDialogsRequest
	ExpectedRes *chat.SelectAllDialogsResponse
	Error error
}

type TestCaseSelectMessages struct {
	ArgData *chat.DialogId
	ExpectedRes *chat.SelectMessagesResponse
	Error error
}

type TestCaseCreateDialog struct {
	ArgData *chat.Dialog
	Error error
}

type TestCaseCreateMessage struct {
	ArgData *chat.Message
	Error error
}

func TestUsecaseSelectDialog(t *testing.T) {
	mockPbDialogId := chat.DialogId {
		Id: 1,
	}

	var mockPbDialog chat.Dialog
	err := faker.FakeData(&mockPbDialog)
	assert.NoError(t, err)

	dialog := &models.Dialog {
		Id: mockPbDialog.Id,
		UserId1: mockPbDialog.UserId1,
		UserId2: mockPbDialog.UserId2,
	}

	for idx := range mockPbDialog.Messages {
		msg := models.Message {
			ID: mockPbDialog.Messages[idx].Id,
			DialogID: mockPbDialog.Messages[idx].DialogId,
			SenderID: mockPbDialog.Messages[idx].SenderId,
			ReceiverID: mockPbDialog.Messages[idx].ReceiverId,
			Body: mockPbDialog.Messages[idx].Body,
			CreatedAt: mockPbDialog.Messages[idx].CreatedAt.AsTime(),
		}
		dialog.Messages = append(dialog.Messages, msg)
	}

	mockPbDialogIdError := chat.DialogId {
		Id: 2,
	}

	selectErr := errors.New("error")

	mockChatRepo := chatMocks.NewRepositoryI(t)

	mockChatRepo.On("SelectDialog", mockPbDialogId.Id).Return(dialog, nil)
	mockChatRepo.On("SelectDialog", mockPbDialogIdError.Id).Return(nil, selectErr)

	useCase := chatUsecase.New(mockChatRepo)

	cases := map[string]TestCaseSelectDialog {
		"success": {
			ArgData:   &mockPbDialogId,
			ExpectedRes: &mockPbDialog,
			Error: nil,
		},
		"error": {
			ArgData:   &mockPbDialogIdError,
			Error: selectErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			dlg, err := useCase.SelectDialog(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))

			if err == nil {
				assert.Equal(t, test.ExpectedRes, dlg)
			}
		})
	}
	mockChatRepo.AssertExpectations(t)
}

func TestUsecaseSelectDialogByUsers(t *testing.T) {
	var mockPbUsersId chat.SelectDialogByUsersRequest
	err := faker.FakeData(&mockPbUsersId)
	assert.NoError(t, err)

	var mockPbDialog chat.Dialog
	err = faker.FakeData(&mockPbDialog)
	assert.NoError(t, err)

	dialog := &models.Dialog {
		Id: mockPbDialog.Id,
		UserId1: mockPbDialog.UserId1,
		UserId2: mockPbDialog.UserId2,
	}

	for idx := range mockPbDialog.Messages {
		msg := models.Message {
			ID: mockPbDialog.Messages[idx].Id,
			DialogID: mockPbDialog.Messages[idx].DialogId,
			SenderID: mockPbDialog.Messages[idx].SenderId,
			ReceiverID: mockPbDialog.Messages[idx].ReceiverId,
			Body: mockPbDialog.Messages[idx].Body,
			CreatedAt: mockPbDialog.Messages[idx].CreatedAt.AsTime(),
		}
		dialog.Messages = append(dialog.Messages, msg)
	}

	var mockPbUsersIdError chat.SelectDialogByUsersRequest
	err = faker.FakeData(&mockPbUsersIdError)
	assert.NoError(t, err)

	selectErr := errors.New("error")

	mockChatRepo := chatMocks.NewRepositoryI(t)

	mockChatRepo.On("SelectDialogByUsers", mockPbUsersId.UserId, mockPbUsersId.FriendId).Return(dialog, nil)
	mockChatRepo.On("SelectDialogByUsers", mockPbUsersIdError.UserId, mockPbUsersIdError.FriendId).Return(nil, selectErr)

	useCase := chatUsecase.New(mockChatRepo)

	cases := map[string]TestCaseSelectDialogByUsers {
		"success": {
			ArgData:   &mockPbUsersId,
			ExpectedRes: &mockPbDialog,
			Error: nil,
		},
		"error": {
			ArgData:   &mockPbUsersIdError,
			Error: selectErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			dlg, err := useCase.SelectDialogByUsers(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))

			if err == nil {
				assert.Equal(t, test.ExpectedRes, dlg)
			}
		})
	}
	mockChatRepo.AssertExpectations(t)
}

func TestUsecaseSelectAllDialogs(t *testing.T) {
	var mockPbUserId chat.SelectAllDialogsRequest
	err := faker.FakeData(&mockPbUserId)
	assert.NoError(t, err)

	var mockPbDialogs chat.SelectAllDialogsResponse
	err = faker.FakeData(&mockPbDialogs)
	assert.NoError(t, err)

	dialogs := make([]models.Dialog, 0)

	for idx := range mockPbDialogs.Dialogs {
		dialog := models.Dialog {
			Id: mockPbDialogs.Dialogs[idx].Id,
			UserId1: mockPbDialogs.Dialogs[idx].UserId1,
			UserId2: mockPbDialogs.Dialogs[idx].UserId2,
		}
	
		for idx2 := range mockPbDialogs.Dialogs[idx].Messages {
			msg := models.Message {
				ID: mockPbDialogs.Dialogs[idx].Messages[idx2].Id,
				DialogID: mockPbDialogs.Dialogs[idx].Messages[idx2].DialogId,
				SenderID: mockPbDialogs.Dialogs[idx].Messages[idx2].SenderId,
				ReceiverID: mockPbDialogs.Dialogs[idx].Messages[idx2].ReceiverId,
				Body: mockPbDialogs.Dialogs[idx].Messages[idx2].Body,
				CreatedAt: mockPbDialogs.Dialogs[idx].Messages[idx2].CreatedAt.AsTime(),
			}
			dialog.Messages = append(dialog.Messages, msg)
		}

		dialogs = append(dialogs, dialog)
	}

	var mockPbUserIdError chat.SelectAllDialogsRequest
	err = faker.FakeData(&mockPbUserIdError)
	assert.NoError(t, err)

	selectErr := errors.New("error")

	mockChatRepo := chatMocks.NewRepositoryI(t)

	mockChatRepo.On("SelectAllDialogs", mockPbUserId.UserId).Return(dialogs, nil)
	mockChatRepo.On("SelectAllDialogs", mockPbUserIdError.UserId).Return(nil, selectErr)

	useCase := chatUsecase.New(mockChatRepo)

	cases := map[string]TestCaseSelectAllDialogs {
		"success": {
			ArgData:   &mockPbUserId,
			ExpectedRes: &mockPbDialogs,
			Error: nil,
		},
		"error": {
			ArgData:   &mockPbUserIdError,
			Error: selectErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			dlgs, err := useCase.SelectAllDialogs(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))

			if err == nil {
				assert.Equal(t, test.ExpectedRes, dlgs)
			}
		})
	}
	mockChatRepo.AssertExpectations(t)
}

func TestUsecaseSelectMessages(t *testing.T) {
	mockPbDialogId := chat.DialogId {
		Id: 1,
	}

	var mockPbMessages chat.SelectMessagesResponse
	err := faker.FakeData(&mockPbMessages)
	assert.NoError(t, err)

	messages := make([]models.Message, 0)
	
	for idx := range mockPbMessages.Messages {
		msg := models.Message {
			ID: mockPbMessages.Messages[idx].Id,
			DialogID: mockPbMessages.Messages[idx].DialogId,
			SenderID: mockPbMessages.Messages[idx].SenderId,
			ReceiverID: mockPbMessages.Messages[idx].ReceiverId,
			Body: mockPbMessages.Messages[idx].Body,
			CreatedAt: mockPbMessages.Messages[idx].CreatedAt.AsTime(),
		}
		messages = append(messages, msg)
	}

	mockPbDialogIdError := chat.DialogId {
		Id: 2,
	}

	selectErr := errors.New("error")

	mockChatRepo := chatMocks.NewRepositoryI(t)

	mockChatRepo.On("SelectMessages", mockPbDialogId.Id).Return(messages, nil)
	mockChatRepo.On("SelectMessages", mockPbDialogIdError.Id).Return(nil, selectErr)

	useCase := chatUsecase.New(mockChatRepo)

	cases := map[string]TestCaseSelectMessages {
		"success": {
			ArgData:   &mockPbDialogId,
			ExpectedRes: &mockPbMessages,
			Error: nil,
		},
		"error": {
			ArgData:   &mockPbDialogIdError,
			Error: selectErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			msgs, err := useCase.SelectMessages(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))

			if err == nil {
				assert.Equal(t, test.ExpectedRes, msgs)
			}
		})
	}
	mockChatRepo.AssertExpectations(t)
}

func TestUsecaseCreateDialog(t *testing.T) {
	var mockPbDialog chat.Dialog
	err := faker.FakeData(&mockPbDialog)
	assert.NoError(t, err)
	mockPbDialog.Id = 1

	dialog := &models.Dialog {
		Id: mockPbDialog.Id,
		UserId1: mockPbDialog.UserId1,
		UserId2: mockPbDialog.UserId2,
	}

	for idx := range mockPbDialog.Messages {
		msg := models.Message {
			ID: mockPbDialog.Messages[idx].Id,
			DialogID: mockPbDialog.Messages[idx].DialogId,
			SenderID: mockPbDialog.Messages[idx].SenderId,
			ReceiverID: mockPbDialog.Messages[idx].ReceiverId,
			Body: mockPbDialog.Messages[idx].Body,
			CreatedAt: mockPbDialog.Messages[idx].CreatedAt.AsTime(),
		}
		dialog.Messages = append(dialog.Messages, msg)
	}

	var mockPbDialogError chat.Dialog
	err = faker.FakeData(&mockPbDialogError)
	assert.NoError(t, err)
	mockPbDialogError.Id = 2

	dialogError := &models.Dialog {
		Id: mockPbDialogError.Id,
		UserId1: mockPbDialogError.UserId1,
		UserId2: mockPbDialogError.UserId2,
	}

	for idx := range mockPbDialogError.Messages {
		msg := models.Message {
			ID: mockPbDialogError.Messages[idx].Id,
			DialogID: mockPbDialogError.Messages[idx].DialogId,
			SenderID: mockPbDialogError.Messages[idx].SenderId,
			ReceiverID: mockPbDialogError.Messages[idx].ReceiverId,
			Body: mockPbDialogError.Messages[idx].Body,
			CreatedAt: mockPbDialogError.Messages[idx].CreatedAt.AsTime(),
		}
		dialogError.Messages = append(dialogError.Messages, msg)
	}

	createErr := errors.New("error")

	mockChatRepo := chatMocks.NewRepositoryI(t)

	mockChatRepo.On("CreateDialog", dialog).Return(nil)
	mockChatRepo.On("CreateDialog", dialogError).Return(createErr)

	useCase := chatUsecase.New(mockChatRepo)

	cases := map[string]TestCaseCreateDialog {
		"success": {
			ArgData:   &mockPbDialog,
			Error: nil,
		},
		"error": {
			ArgData:   &mockPbDialogError,
			Error: createErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := useCase.CreateDialog(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
		})
	}
	mockChatRepo.AssertExpectations(t)
}

