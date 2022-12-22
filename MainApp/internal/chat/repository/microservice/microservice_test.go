package microservice_test

import (
	"context"
	"testing"

	"github.com/bxcodec/faker"
	chatRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/chat/repository/microservice"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	chat "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/proto/chat"
	chatMocks "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/proto/chat/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestCaseCreateMessage struct {
	ArgData *models.Message
	Error error
}

type TestCaseCreateDialog struct {
	ArgData *models.Dialog
	Error error
}

type TestCaseSelectDialog struct {
	ArgData uint64
	ExpectedRes *models.Dialog
	Error error
}

type TestCaseSelectDialogByUsers struct {
	ArgDataUserId uint64
	ArgDataFriendId uint64
	ExpectedRes *models.Dialog
	Error error
}

type TestCaseSelectAllDialogs struct {
	ArgData uint64
	ExpectedRes []models.Dialog
	Error error
models.Dialog}

type TestCaseSelectMessages struct {
	ArgData uint64
	ExpectedRes []models.Message 
	Error error
}

func TestMicroserviceCreateMessage(t *testing.T) {
	var mockPbMessage chat.Message
	err := faker.FakeData(&mockPbMessage)
	assert.NoError(t, err)
	mockPbMessage.Id = 1
	mockPbMessage.AttachmentsIds = nil

	message := &models.Message {
		ID: mockPbMessage.Id,
		DialogID: mockPbMessage.DialogId,
		SenderID: mockPbMessage.SenderId,
		ReceiverID: mockPbMessage.ReceiverId,
		Body: mockPbMessage.Body,
		CreatedAt: mockPbMessage.CreatedAt.AsTime(),
		StickerID: mockPbMessage.StickerId,
	}

	var mockPbMessageError chat.Message
	err = faker.FakeData(&mockPbMessageError)
	assert.NoError(t, err)
	mockPbMessageError.Id = 2
	mockPbMessageError.AttachmentsIds = nil

	messageError := &models.Message {
		ID: mockPbMessageError.Id,
		DialogID: mockPbMessageError.DialogId,
		SenderID: mockPbMessageError.SenderId,
		ReceiverID: mockPbMessageError.ReceiverId,
		Body: mockPbMessageError.Body,
		CreatedAt: mockPbMessageError.CreatedAt.AsTime(),
		StickerID: mockPbMessageError.StickerId,
	}

	pbNothing := chat.Nothing{Dummy: true}

	mockChatClient := chatMocks.NewChatClient(t)

	ctx := context.Background()

	createErr := errors.New("error")

	mockChatClient.On("CreateMessage", ctx, &mockPbMessage).Return(&pbNothing, nil)
	mockChatClient.On("CreateMessage", ctx, &mockPbMessageError).Return(nil, createErr)

	repository := chatRep.New(mockChatClient)

	cases := map[string]TestCaseCreateMessage {
		"success": {
			ArgData:   message,
			Error: nil,
		},
		"error": {
			ArgData:   messageError,
			Error: createErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := repository.CreateMessage(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
		})
	}
	mockChatClient.AssertExpectations(t)
}

func TestMicroserviceCreateDialog(t *testing.T) {
	var mockPbDialog chat.Dialog
	err := faker.FakeData(&mockPbDialog)
	assert.NoError(t, err)
	mockPbDialog.Id = 1
	mockPbDialog.Messages = nil

	dialogId := chat.DialogId {
		Id: mockPbDialog.Id,
	}

	dialog := &models.Dialog {
		Id: mockPbDialog.Id,
		UserId1: mockPbDialog.UserId1,
		UserId2: mockPbDialog.UserId2,
	}

	// for idx := range mockPbDialog.Messages {
	// 	mockPbDialog.Messages[idx].AttachmentsIds = nil
	// 	msg := models.Message {
	// 		ID: mockPbDialog.Messages[idx].Id,
	// 		DialogID: mockPbDialog.Messages[idx].DialogId,
	// 		SenderID: mockPbDialog.Messages[idx].SenderId,
	// 		ReceiverID: mockPbDialog.Messages[idx].ReceiverId,
	// 		Body: mockPbDialog.Messages[idx].Body,
	// 		CreatedAt: mockPbDialog.Messages[idx].CreatedAt.AsTime(),
	// 		StickerID: mockPbDialog.Messages[idx].StickerId,
	// 	}
	// 	dialog.Messages = append(dialog.Messages, msg)
	// }

	var mockPbDialogError chat.Dialog
	err = faker.FakeData(&mockPbDialogError)
	assert.NoError(t, err)
	mockPbDialogError.Id = 2
	mockPbDialogError.Messages = nil

	dialogError := &models.Dialog {
		Id: mockPbDialogError.Id,
		UserId1: mockPbDialogError.UserId1,
		UserId2: mockPbDialogError.UserId2,
	}

	// for idx := range mockPbDialogError.Messages {
	// 	mockPbDialogError.Messages[idx].AttachmentsIds = nil
	// 	msg := models.Message {
	// 		ID: mockPbDialogError.Messages[idx].Id,
	// 		DialogID: mockPbDialogError.Messages[idx].DialogId,
	// 		SenderID: mockPbDialogError.Messages[idx].SenderId,
	// 		ReceiverID: mockPbDialogError.Messages[idx].ReceiverId,
	// 		Body: mockPbDialogError.Messages[idx].Body,
	// 		CreatedAt: mockPbDialogError.Messages[idx].CreatedAt.AsTime(),
	// 		StickerID: mockPbDialogError.Messages[idx].StickerId,
	// 	}
	// 	dialogError.Messages = append(dialogError.Messages, msg)
	// }

	mockChatClient := chatMocks.NewChatClient(t)

	ctx := context.Background()

	createErr := errors.New("error")

	mockChatClient.On("CreateDialog", ctx, &mockPbDialog).Return(&dialogId, nil)
	mockChatClient.On("CreateDialog", ctx, &mockPbDialogError).Return(nil, createErr)

	repository := chatRep.New(mockChatClient)

	cases := map[string]TestCaseCreateDialog {
		"success": {
			ArgData:   dialog,
			Error: nil,
		},
		"error": {
			ArgData:   dialogError,
			Error: createErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := repository.CreateDialog(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
		})
	}
	mockChatClient.AssertExpectations(t)
}


func TestMicroserviceSelectDialog(t *testing.T) {
	pbDialogId := chat.DialogId {
		Id: 1,
	}

	var mockPbDialog chat.Dialog
	err := faker.FakeData(&mockPbDialog)
	assert.NoError(t, err)
	mockPbDialog.Id = pbDialogId.Id

	dialog := &models.Dialog {
		Id: mockPbDialog.Id,
		UserId1: mockPbDialog.UserId1,
		UserId2: mockPbDialog.UserId2,
	}

	for idx := range mockPbDialog.Messages {
		mockPbDialog.Messages[idx].AttachmentsIds = nil
		msg := models.Message {
			ID: mockPbDialog.Messages[idx].Id,
			DialogID: mockPbDialog.Messages[idx].DialogId,
			SenderID: mockPbDialog.Messages[idx].SenderId,
			ReceiverID: mockPbDialog.Messages[idx].ReceiverId,
			Body: mockPbDialog.Messages[idx].Body,
			CreatedAt: mockPbDialog.Messages[idx].CreatedAt.AsTime(),
			StickerID: mockPbDialog.Messages[idx].StickerId,
		}
		dialog.Messages = append(dialog.Messages, msg)
	}

	pbDialogIdError := chat.DialogId {
		Id: 2,
	}
	
	mockChatClient := chatMocks.NewChatClient(t)

	ctx := context.Background()

	selectErr := errors.New("error")

	mockChatClient.On("SelectDialog", ctx, &pbDialogId).Return(&mockPbDialog, nil)
	mockChatClient.On("SelectDialog", ctx, &pbDialogIdError).Return(nil, selectErr)

	repository := chatRep.New(mockChatClient)

	cases := map[string]TestCaseSelectDialog {
		"success": {
			ArgData:   pbDialogId.Id,
			ExpectedRes: dialog,
			Error: nil,
		},
		"error": {
			ArgData:   pbDialogIdError.Id,
			ExpectedRes: nil,
			Error: selectErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			selectedDialog, err := repository.SelectDialog(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))

			if err == nil {
				assert.Equal(t, test.ExpectedRes, selectedDialog)
			}
		})
	}
	mockChatClient.AssertExpectations(t)
}

func TestMicroserviceSelectDialogByUsers(t *testing.T) {
	pbUsersId := chat.SelectDialogByUsersRequest {
		UserId: 1,
		FriendId: 2,
	}

	var mockPbDialog chat.Dialog
	err := faker.FakeData(&mockPbDialog)
	assert.NoError(t, err)
	mockPbDialog.UserId1 = pbUsersId.UserId
	mockPbDialog.UserId2 = pbUsersId.FriendId

	dialog := &models.Dialog {
		Id: mockPbDialog.Id,
		UserId1: mockPbDialog.UserId1,
		UserId2: mockPbDialog.UserId2,
	}

	for idx := range mockPbDialog.Messages {
		mockPbDialog.Messages[idx].AttachmentsIds = nil
		msg := models.Message {
			ID: mockPbDialog.Messages[idx].Id,
			DialogID: mockPbDialog.Messages[idx].DialogId,
			SenderID: mockPbDialog.Messages[idx].SenderId,
			ReceiverID: mockPbDialog.Messages[idx].ReceiverId,
			Body: mockPbDialog.Messages[idx].Body,
			CreatedAt: mockPbDialog.Messages[idx].CreatedAt.AsTime(),
			StickerID: mockPbDialog.Messages[idx].StickerId,
		}
		dialog.Messages = append(dialog.Messages, msg)
	}

	pbUsersIdError := chat.SelectDialogByUsersRequest {
		UserId: 3,
		FriendId: 4,
	}
	
	mockChatClient := chatMocks.NewChatClient(t)

	ctx := context.Background()

	selectErr := errors.New("error")

	mockChatClient.On("SelectDialogByUsers", ctx, &pbUsersId).Return(&mockPbDialog, nil)
	mockChatClient.On("SelectDialogByUsers", ctx, &pbUsersIdError).Return(nil, selectErr)

	repository := chatRep.New(mockChatClient)

	cases := map[string]TestCaseSelectDialogByUsers {
		"success": {
			ArgDataUserId:   pbUsersId.UserId,
			ArgDataFriendId:   pbUsersId.FriendId,
			ExpectedRes: dialog,
			Error: nil,
		},
		"error": {
			ArgDataUserId:   pbUsersIdError.UserId,
			ArgDataFriendId:   pbUsersIdError.FriendId,
			ExpectedRes: nil,
			Error: selectErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			selectedDialog, err := repository.SelectDialogByUsers(test.ArgDataUserId, test.ArgDataFriendId)
			require.Equal(t, test.Error, errors.Cause(err))

			if err == nil {
				assert.Equal(t, test.ExpectedRes, selectedDialog)
			}
		})
	}
	mockChatClient.AssertExpectations(t)
}

func TestMicroserviceSelectAllDialogs(t *testing.T) {
	pbUserId := chat.SelectAllDialogsRequest {
		UserId: 1,
	}

	var mockPbDialogs chat.SelectAllDialogsResponse
	err := faker.FakeData(&mockPbDialogs)
	assert.NoError(t, err)

	dialogs := make([]models.Dialog, 0)

	for idx := range mockPbDialogs.Dialogs {
		dialog := models.Dialog {
			Id: mockPbDialogs.Dialogs[idx].Id,
			UserId1: mockPbDialogs.Dialogs[idx].UserId1,
			UserId2: mockPbDialogs.Dialogs[idx].UserId2,
		}
	
		for idx2 := range mockPbDialogs.Dialogs[idx].Messages {
			mockPbDialogs.Dialogs[idx].Messages[idx2].AttachmentsIds = nil
			msg := models.Message {
				ID: mockPbDialogs.Dialogs[idx].Messages[idx2].Id,
				DialogID: mockPbDialogs.Dialogs[idx].Messages[idx2].DialogId,
				SenderID: mockPbDialogs.Dialogs[idx].Messages[idx2].SenderId,
				ReceiverID: mockPbDialogs.Dialogs[idx].Messages[idx2].ReceiverId,
				Body: mockPbDialogs.Dialogs[idx].Messages[idx2].Body,
				CreatedAt: mockPbDialogs.Dialogs[idx].Messages[idx2].CreatedAt.AsTime(),
				StickerID: mockPbDialogs.Dialogs[idx].Messages[idx2].StickerId,
			}
			dialog.Messages = append(dialog.Messages, msg)
		}

		dialogs = append(dialogs, dialog)
	}

	pbUserIdError := chat.SelectAllDialogsRequest {
		UserId: 2,
	}
	
	mockChatClient := chatMocks.NewChatClient(t)

	ctx := context.Background()

	selectErr := errors.New("error")

	mockChatClient.On("SelectAllDialogs", ctx, &pbUserId).Return(&mockPbDialogs, nil)
	mockChatClient.On("SelectAllDialogs", ctx, &pbUserIdError).Return(nil, selectErr)

	repository := chatRep.New(mockChatClient)

	cases := map[string]TestCaseSelectAllDialogs {
		"success": {
			ArgData:   pbUserId.UserId,
			ExpectedRes: dialogs,
			Error: nil,
		},
		"error": {
			ArgData:   pbUserIdError.UserId,
			ExpectedRes: nil,
			Error: selectErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			selectedDialogs, err := repository.SelectAllDialogs(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))

			if err == nil {
				assert.Equal(t, test.ExpectedRes, selectedDialogs)
			}
		})
	}
	mockChatClient.AssertExpectations(t)
}

func TestMicroserviceSelectMessages(t *testing.T) {
	pbDialogId := chat.DialogId {
		Id: 1,
	}

	var mockPbMessages chat.SelectMessagesResponse
	err := faker.FakeData(&mockPbMessages)
	assert.NoError(t, err)

	messages := make([]models.Message, 0)

	for idx := range mockPbMessages.Messages {
		mockPbMessages.Messages[idx].AttachmentsIds = nil
		msg := models.Message {
			ID: mockPbMessages.Messages[idx].Id,
			DialogID: mockPbMessages.Messages[idx].DialogId,
			SenderID: mockPbMessages.Messages[idx].SenderId,
			ReceiverID: mockPbMessages.Messages[idx].ReceiverId,
			Body: mockPbMessages.Messages[idx].Body,
			CreatedAt: mockPbMessages.Messages[idx].CreatedAt.AsTime(),
			StickerID: mockPbMessages.Messages[idx].StickerId,
		}
		messages = append(messages, msg)
	}

	pbDialogIdError := chat.DialogId {
		Id: 2,
	}
	
	mockChatClient := chatMocks.NewChatClient(t)

	ctx := context.Background()

	selectErr := errors.New("error")

	mockChatClient.On("SelectMessages", ctx, &pbDialogId).Return(&mockPbMessages, nil)
	mockChatClient.On("SelectMessages", ctx, &pbDialogIdError).Return(nil, selectErr)

	repository := chatRep.New(mockChatClient)

	cases := map[string]TestCaseSelectMessages {
		"success": {
			ArgData:   pbDialogId.Id,
			ExpectedRes: messages,
			Error: nil,
		},
		"error": {
			ArgData:   pbDialogIdError.Id,
			ExpectedRes: nil,
			Error: selectErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			selectedMessages, err := repository.SelectMessages(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))

			if err == nil {
				assert.Equal(t, test.ExpectedRes, selectedMessages)
			}
		})
	}
	mockChatClient.AssertExpectations(t)
}

