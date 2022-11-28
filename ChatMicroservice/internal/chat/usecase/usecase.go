package usecase

import (
	chatRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/ChatMicroservice/internal/chat/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/ChatMicroservice/models"
	chat "github.com/go-park-mail-ru/2022_2_TikTikIVProd/ChatMicroservice/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UseCaseI interface {
	SelectDialog(*chat.DialogId) (*chat.Dialog, error)
	SelectDialogByUsers(*chat.SelectDialogByUsersRequest) (*chat.Dialog, error)
	SelectAllDialogs(*chat.SelectAllDialogsRequest) (*chat.SelectAllDialogsResponse, error)
	SelectMessages(*chat.DialogId) (*chat.SelectMessagesResponse, error)
	CreateDialog(*chat.Dialog) (*chat.Nothing, error)
	CreateMessage(*chat.Message) (*chat.Nothing, error)
}

type useCase struct {
	chatRepository chatRep.RepositoryI
}

func New(chatRepository chatRep.RepositoryI) UseCaseI {
	return &useCase{
		chatRepository: chatRepository,
	}
}

func (uc *useCase) SelectDialog(id *chat.DialogId) (*chat.Dialog, error) {
	dialog, err := uc.chatRepository.SelectDialog(id.Id)

	pbDialog := &chat.Dialog {
		Id: dialog.Id,
		UserId1: dialog.UserId1,
		UserId2: dialog.UserId2,
	}

	for idx := range dialog.Messages {
		ts := timestamppb.New(dialog.Messages[idx].CreatedAt)
		msg := &chat.Message {
			Id: dialog.Messages[idx].ID,
			DialogId: dialog.Messages[idx].DialogID,
			SenderId: dialog.Messages[idx].SenderID,
			ReceiverId: dialog.Messages[idx].ReceiverID,
			Body: dialog.Messages[idx].Body,
			CreatedAt: ts,
		}
		pbDialog.Messages = append(pbDialog.Messages, msg)
	}
	
	return pbDialog, err
}

func (uc *useCase) SelectDialogByUsers(usersId *chat.SelectDialogByUsersRequest) (*chat.Dialog, error) {
	dialog, err := uc.chatRepository.SelectDialogByUsers(usersId.UserId, usersId.FriendId)

	pbDialog := &chat.Dialog {
		Id: dialog.Id,
		UserId1: dialog.UserId1,
		UserId2: dialog.UserId2,
	}

	for idx := range dialog.Messages {
		ts := timestamppb.New(dialog.Messages[idx].CreatedAt)
		msg := &chat.Message {
			Id: dialog.Messages[idx].ID,
			DialogId: dialog.Messages[idx].DialogID,
			SenderId: dialog.Messages[idx].SenderID,
			ReceiverId: dialog.Messages[idx].ReceiverID,
			Body: dialog.Messages[idx].Body,
			CreatedAt: ts,
		}
		pbDialog.Messages = append(pbDialog.Messages, msg)
	}
	
	return pbDialog, err
}

func (uc *useCase) SelectAllDialogs(pbUserId *chat.SelectAllDialogsRequest) (*chat.SelectAllDialogsResponse, error) {
	dialogs, err := uc.chatRepository.SelectAllDialogs(pbUserId.UserId)

	pbDialogs := &chat.SelectAllDialogsResponse {}

	for idx := range dialogs {
		dialog := &chat.Dialog {
			Id: dialogs[idx].Id,
			UserId1: dialogs[idx].UserId1,
			UserId2: dialogs[idx].UserId2,
		}

		for idx2 := range dialogs[idx].Messages {
			ts := timestamppb.New(dialogs[idx].Messages[idx2].CreatedAt)
			msg := &chat.Message {
				Id: dialogs[idx].Messages[idx2].ID,
				DialogId: dialogs[idx].Messages[idx2].DialogID,
				SenderId: dialogs[idx].Messages[idx2].SenderID,
				ReceiverId: dialogs[idx].Messages[idx2].ReceiverID,
				Body: dialogs[idx].Messages[idx2].Body,
				CreatedAt: ts,
			}
			dialog.Messages = append(dialog.Messages, msg)
		}

		pbDialogs.Dialogs = append(pbDialogs.Dialogs, dialog)
	}
	
	return pbDialogs, err
}

func (uc *useCase) SelectMessages(pbDialogId *chat.DialogId) (*chat.SelectMessagesResponse, error) {
	messages, err := uc.chatRepository.SelectMessages(pbDialogId.Id)

	pbMessages := &chat.SelectMessagesResponse {}

	for idx := range messages {
		ts := timestamppb.New(messages[idx].CreatedAt)
		msg := &chat.Message {
			Id: messages[idx].ID,
			DialogId: messages[idx].DialogID,
			SenderId: messages[idx].SenderID,
			ReceiverId: messages[idx].ReceiverID,
			Body: messages[idx].Body,
			CreatedAt: ts,
		}
		pbMessages.Messages = append(pbMessages.Messages, msg)
	}
	
	return pbMessages, err
}

func (uc *useCase) CreateDialog(pbDialog *chat.Dialog) (*chat.Nothing, error) {
	dialog := &models.Dialog {
		Id: pbDialog.Id,
		UserId1: pbDialog.UserId1,
		UserId2: pbDialog.UserId2,
	}

	for idx := range pbDialog.Messages {
		msg := models.Message {
			ID: dialog.Messages[idx].ID,
			DialogID: dialog.Messages[idx].DialogID,
			SenderID: dialog.Messages[idx].SenderID,
			ReceiverID: dialog.Messages[idx].ReceiverID,
			Body: dialog.Messages[idx].Body,
			CreatedAt: pbDialog.Messages[idx].CreatedAt.AsTime(),
		}
		dialog.Messages = append(dialog.Messages, msg)
	}

	err := uc.chatRepository.CreateDialog(dialog)

	pbDialog.Id = dialog.Id
	
	return &chat.Nothing{Dummy: true}, err
}

func (uc *useCase) CreateMessage(pbMessage *chat.Message) (*chat.Nothing, error) {
	msg := &models.Message {
		ID: pbMessage.Id,
		DialogID: pbMessage.DialogId,
		SenderID: pbMessage.SenderId,
		ReceiverID: pbMessage.ReceiverId,
		Body: pbMessage.Body,
		CreatedAt: pbMessage.CreatedAt.AsTime(),
	}

	err := uc.chatRepository.CreateMessage(msg)
	
	return &chat.Nothing{Dummy: true}, err
}
