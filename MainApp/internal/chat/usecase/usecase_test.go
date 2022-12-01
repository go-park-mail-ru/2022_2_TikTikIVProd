package usecase_test

import (
	"testing"

	"github.com/bxcodec/faker"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/chat/repository/mocks"
	chatUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/chat/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type TestCaseSendMessage struct {
	ArgData models.Message
	ExpectedRes uint64
	Error error
}

type TestCaseSelectAllDialogs struct {
	ArgData uint64
	ExpectedRes []models.Dialog
	Error       error
}

type TestCaseSelectDialog struct {
	ArgData uint64
	ExpectedRes *models.Dialog
	Error       error
}

type TestCaseSelectDialogByUsers struct {
	ArgDataUser1 uint64
	ArgDataUser2 uint64
	ExpectedRes *models.Dialog
	Error error
}

func TestUsecaseSelectDialog(t *testing.T) {
	var mockDialog models.Dialog
	err := faker.FakeData(&mockDialog)
	assert.NoError(t, err)

	mockChatRepo := mocks.NewRepositoryI(t)

	mockChatRepo.On("SelectDialog", mockDialog.Id).Return(&mockDialog, nil)
	mockChatRepo.On("SelectMessages", mockDialog.Id).Return(mockDialog.Messages, nil)

	useCase := chatUsecase.New(mockChatRepo)

	cases := map[string]TestCaseSelectDialog{
		"success": {
			ArgData:     mockDialog.Id,
			ExpectedRes: &mockDialog,
			Error:       nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			dialog, err := useCase.SelectDialog(test.ArgData)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.ExpectedRes, dialog)
			}
		})
	}
	mockChatRepo.AssertExpectations(t)
}

func TestUsecaseSelectDialogByUsers(t *testing.T) {
	var mockDialog models.Dialog
	err := faker.FakeData(&mockDialog)
	assert.NoError(t, err)

	mockChatRepo := mocks.NewRepositoryI(t)

	mockChatRepo.On("SelectDialogByUsers", mockDialog.UserId1, mockDialog.UserId2).Return(&mockDialog, nil)
	mockChatRepo.On("SelectMessages", mockDialog.Id).Return(mockDialog.Messages, nil)

	useCase := chatUsecase.New(mockChatRepo)

	cases := map[string]TestCaseSelectDialogByUsers{
		"success": {
			ArgDataUser1: mockDialog.UserId1,
			ArgDataUser2: mockDialog.UserId2,
			ExpectedRes:  &mockDialog,
			Error:        nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			dialog, err := useCase.SelectDialogByUsers(test.ArgDataUser1, test.ArgDataUser2)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.ExpectedRes, dialog)
			}
		})
	}
	mockChatRepo.AssertExpectations(t)
}

func TestUsecaseSelectAllDialogs(t *testing.T) {
	mockDialogs := make([]models.Dialog, 0, 10)
	err := faker.FakeData(&mockDialogs)
	assert.NoError(t, err)

	mockChatRepo := mocks.NewRepositoryI(t)

	var userId uint64 = 1

	mockChatRepo.On("SelectAllDialogs", userId).Return(mockDialogs, nil)

	useCase := chatUsecase.New(mockChatRepo)

	cases := map[string]TestCaseSelectAllDialogs{
		"success": {
			ArgData:     userId,
			ExpectedRes: mockDialogs,
			Error:       nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			dialogs, err := useCase.SelectAllDialogs(test.ArgData)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.ExpectedRes, dialogs)
			}
		})
	}
	mockChatRepo.AssertExpectations(t)
}

func TestUsecaseSendMessage(t *testing.T) {
	var mockMessage models.Message
	err := faker.FakeData(&mockMessage)
	assert.NoError(t, err)

	var mockMessageNewDialog models.Message
	err = faker.FakeData(&mockMessageNewDialog)
	assert.NoError(t, err)

	mockMessageNewDialog.DialogID = 0

	mockDialog := models.Dialog{
		UserId1:  mockMessageNewDialog.SenderID,
		UserId2:  mockMessageNewDialog.ReceiverID,
		Messages: []models.Message{mockMessageNewDialog},
	}

	mockChatRepo := mocks.NewRepositoryI(t)

	mockChatRepo.On("SelectDialog", mockMessage.DialogID).Return(nil, nil)
	mockChatRepo.On("CreateMessage", mock.AnythingOfType("*models.Message")).Return(nil)

	mockChatRepo.On("SelectDialog", mockMessageNewDialog.DialogID).Return(nil, models.ErrNotFound)
	mockChatRepo.On("CreateDialog", &mockDialog).Return(nil)
	mockChatRepo.On("CreateMessage", mock.AnythingOfType("*models.Message")).Return(nil)

	useCase := chatUsecase.New(mockChatRepo)

	cases := map[string]TestCaseSendMessage{
		"success": {
			ArgData:     mockMessage,
			ExpectedRes: mockMessage.ID,
			Error:       nil,
		},
		"success_new_chat": {
			ArgData:     mockMessageNewDialog,
			ExpectedRes: mockMessageNewDialog.ID,
			Error:       nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := useCase.SendMessage(&test.ArgData)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.ExpectedRes, test.ArgData.ID)
			}
		})
	}
	mockChatRepo.AssertExpectations(t)
}
