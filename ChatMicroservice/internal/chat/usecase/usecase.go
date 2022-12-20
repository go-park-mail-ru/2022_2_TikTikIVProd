package usecase

import (
	chatRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/ChatMicroservice/internal/chat/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/ChatMicroservice/models"
	chat "github.com/go-park-mail-ru/2022_2_TikTikIVProd/ChatMicroservice/proto"
	"github.com/pkg/errors"
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
	if err != nil {
		return nil, errors.Wrap(err, "chat repository postgres error")
	}

	pbDialog := &chat.Dialog{
		Id:      dialog.Id,
		UserId1: dialog.UserId1,
		UserId2: dialog.UserId2,
	}

	for idx := range dialog.Messages {
		ts := timestamppb.New(dialog.Messages[idx].CreatedAt)
		msg := &chat.Message{
			Id:         dialog.Messages[idx].ID,
			DialogId:   dialog.Messages[idx].DialogID,
			SenderId:   dialog.Messages[idx].SenderID,
			ReceiverId: dialog.Messages[idx].ReceiverID,
			Body:       dialog.Messages[idx].Body,
			CreatedAt:  ts,
			StickerId:  dialog.Messages[idx].StickerID,
		}
		pbDialog.Messages = append(pbDialog.Messages, msg)
	}

	return pbDialog, nil
}

func (uc *useCase) SelectDialogByUsers(usersId *chat.SelectDialogByUsersRequest) (*chat.Dialog, error) {
	dialog, err := uc.chatRepository.SelectDialogByUsers(usersId.UserId, usersId.FriendId)
	if err != nil {
		return nil, errors.Wrap(err, "chat repository postgres error")
	}

	pbDialog := &chat.Dialog{
		Id:      dialog.Id,
		UserId1: dialog.UserId1,
		UserId2: dialog.UserId2,
	}

	for idx := range dialog.Messages {
		ts := timestamppb.New(dialog.Messages[idx].CreatedAt)
		msg := &chat.Message{
			Id:         dialog.Messages[idx].ID,
			DialogId:   dialog.Messages[idx].DialogID,
			SenderId:   dialog.Messages[idx].SenderID,
			ReceiverId: dialog.Messages[idx].ReceiverID,
			Body:       dialog.Messages[idx].Body,
			CreatedAt:  ts,
			StickerId:  dialog.Messages[idx].StickerID,
		}
		pbDialog.Messages = append(pbDialog.Messages, msg)
	}

	return pbDialog, nil
}

func (uc *useCase) SelectAllDialogs(pbUserId *chat.SelectAllDialogsRequest) (*chat.SelectAllDialogsResponse, error) {
	dialogs, err := uc.chatRepository.SelectAllDialogs(pbUserId.UserId)
	if err != nil {
		return nil, errors.Wrap(err, "chat repository postgres error")
	}

	pbDialogs := &chat.SelectAllDialogsResponse{}
	pbDialogs.Dialogs = make([]*chat.Dialog, 0)

	for idx := range dialogs {
		dialog := &chat.Dialog{
			Id:      dialogs[idx].Id,
			UserId1: dialogs[idx].UserId1,
			UserId2: dialogs[idx].UserId2,
		}

		dialog.Messages = make([]*chat.Message, 0)

		for idx2 := range dialogs[idx].Messages {
			ts := timestamppb.New(dialogs[idx].Messages[idx2].CreatedAt)
			msg := &chat.Message{
				Id:         dialogs[idx].Messages[idx2].ID,
				DialogId:   dialogs[idx].Messages[idx2].DialogID,
				SenderId:   dialogs[idx].Messages[idx2].SenderID,
				ReceiverId: dialogs[idx].Messages[idx2].ReceiverID,
				Body:       dialogs[idx].Messages[idx2].Body,
				CreatedAt:  ts,
				StickerId:  dialogs[idx].Messages[idx2].StickerID,
			}
			dialog.Messages = append(dialog.Messages, msg)
		}

		pbDialogs.Dialogs = append(pbDialogs.Dialogs, dialog)
	}

	return pbDialogs, nil
}

func (uc *useCase) SelectMessages(pbDialogId *chat.DialogId) (*chat.SelectMessagesResponse, error) {
	messages, err := uc.chatRepository.SelectMessages(pbDialogId.Id)
	if err != nil {
		return nil, errors.Wrap(err, "chat repository postgres error")
	}

	pbMessages := &chat.SelectMessagesResponse{}
	pbMessages.Messages = make([]*chat.Message, 0)

	for idx := range messages {
		ts := timestamppb.New(messages[idx].CreatedAt)
		msg := &chat.Message{
			Id:         messages[idx].ID,
			DialogId:   messages[idx].DialogID,
			SenderId:   messages[idx].SenderID,
			ReceiverId: messages[idx].ReceiverID,
			Body:       messages[idx].Body,
			CreatedAt:  ts,
			StickerId: messages[idx].StickerID,
		}
		pbMessages.Messages = append(pbMessages.Messages, msg)
	}

	return pbMessages, nil
}

func (uc *useCase) CreateDialog(pbDialog *chat.Dialog) (*chat.Nothing, error) {
	dialog := &models.Dialog{
		Id:      pbDialog.Id,
		UserId1: pbDialog.UserId1,
		UserId2: pbDialog.UserId2,
	}

	for idx := range pbDialog.Messages {
		msg := models.Message{
			ID:         pbDialog.Messages[idx].Id,
			DialogID:   pbDialog.Messages[idx].DialogId,
			SenderID:   pbDialog.Messages[idx].SenderId,
			ReceiverID: pbDialog.Messages[idx].ReceiverId,
			Body:       pbDialog.Messages[idx].Body,
			CreatedAt:  pbDialog.Messages[idx].CreatedAt.AsTime(),
			StickerID: pbDialog.Messages[idx].StickerId,
		}
		dialog.Messages = append(dialog.Messages, msg)
	}

	err := uc.chatRepository.CreateDialog(dialog)
	if err != nil {
		return nil, errors.Wrap(err, "chat repository postgres error")
	}

	pbDialog.Id = dialog.Id

	return &chat.Nothing{Dummy: true}, nil
}

func (uc *useCase) CreateMessage(pbMessage *chat.Message) (*chat.Nothing, error) {
	msg := &models.Message{
		ID:         pbMessage.Id,
		DialogID:   pbMessage.DialogId,
		SenderID:   pbMessage.SenderId,
		ReceiverID: pbMessage.ReceiverId,
		Body:       pbMessage.Body,
		CreatedAt:  pbMessage.CreatedAt.AsTime(),
		StickerID:  pbMessage.StickerId,
	}

	for idx := range pbMessage.AttachmentsIds {
		msg.Attachments = append(msg.Attachments, models.Attachment{ID: pbMessage.AttachmentsIds[idx]})
	}

	err := uc.chatRepository.CreateMessage(msg)
	if err != nil {
		return nil, errors.Wrap(err, "chat repository postgres error")
	}

	return &chat.Nothing{Dummy: true}, nil
}
