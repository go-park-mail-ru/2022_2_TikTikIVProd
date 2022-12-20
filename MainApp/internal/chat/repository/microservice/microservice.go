package microservice

import (
	"context"

	chatRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/chat/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	chat "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/proto/chat"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type microService struct {
	client chat.ChatClient
}

func New(client chat.ChatClient) chatRep.RepositoryI {
	return &microService{
		client: client,
	}
}

func (chatMS *microService) SelectDialog(id uint64) (*models.Dialog, error) {
	ctx := context.Background()

	pbDialogId := chat.DialogId{
		Id: id,
	}

	pbDialog, err := chatMS.client.SelectDialog(ctx, &pbDialogId)
	if err != nil {
		return nil, errors.Wrap(err, "chat microservice error")
	}

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
			StickerID:  pbDialog.Messages[idx].StickerId,
		}
		dialog.Messages = append(dialog.Messages, msg)
	}

	return dialog, nil
}

func (chatMS *microService) SelectDialogByUsers(userId, friendId uint64) (*models.Dialog, error) {
	ctx := context.Background()

	pbSelectDialogByUsersRequest := chat.SelectDialogByUsersRequest{
		UserId:   userId,
		FriendId: friendId,
	}

	pbDialog, err := chatMS.client.SelectDialogByUsers(ctx, &pbSelectDialogByUsersRequest)
	if err != nil {
		return nil, errors.Wrap(err, "chat microservice error")
	}

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
			StickerID:  pbDialog.Messages[idx].StickerId,
		}
		dialog.Messages = append(dialog.Messages, msg)
	}

	return dialog, nil
}

func (chatMS *microService) SelectAllDialogs(userId uint64) ([]models.Dialog, error) {
	ctx := context.Background()

	pbSelectAllDialogsRequest := chat.SelectAllDialogsRequest{
		UserId: userId,
	}

	pbDialogs, err := chatMS.client.SelectAllDialogs(ctx, &pbSelectAllDialogsRequest)
	if err != nil {
		return nil, errors.Wrap(err, "chat microservice error")
	}

	dialogs := make([]models.Dialog, 0)

	for idx := range pbDialogs.Dialogs {
		dialog := models.Dialog{
			Id:      pbDialogs.Dialogs[idx].Id,
			UserId1: pbDialogs.Dialogs[idx].UserId1,
			UserId2: pbDialogs.Dialogs[idx].UserId2,
		}

		for idx2 := range pbDialogs.Dialogs[idx].Messages {
			msg := models.Message{
				ID:         pbDialogs.Dialogs[idx].Messages[idx2].Id,
				DialogID:   pbDialogs.Dialogs[idx].Messages[idx2].DialogId,
				SenderID:   pbDialogs.Dialogs[idx].Messages[idx2].SenderId,
				ReceiverID: pbDialogs.Dialogs[idx].Messages[idx2].ReceiverId,
				Body:       pbDialogs.Dialogs[idx].Messages[idx2].Body,
				CreatedAt:  pbDialogs.Dialogs[idx].Messages[idx2].CreatedAt.AsTime(),
			}
			dialog.Messages = append(dialog.Messages, msg)
		}

		dialogs = append(dialogs, dialog)
	}

	return dialogs, nil
}

func (chatMS *microService) SelectMessages(id uint64) ([]models.Message, error) {
	ctx := context.Background()

	pbDialogId := chat.DialogId{
		Id: id,
	}

	pbMessages, err := chatMS.client.SelectMessages(ctx, &pbDialogId)
	if err != nil {
		return nil, errors.Wrap(err, "chat microservice error")
	}

	messages := make([]models.Message, 0)

	for idx := range pbMessages.Messages {
		msg := models.Message{
			ID:         pbMessages.Messages[idx].Id,
			DialogID:   pbMessages.Messages[idx].DialogId,
			SenderID:   pbMessages.Messages[idx].SenderId,
			ReceiverID: pbMessages.Messages[idx].ReceiverId,
			Body:       pbMessages.Messages[idx].Body,
			CreatedAt:  pbMessages.Messages[idx].CreatedAt.AsTime(),
		}
		messages = append(messages, msg)
	}

	return messages, nil
}

func (chatMS *microService) CreateDialog(dialog *models.Dialog) error {
	ctx := context.Background()

	pbDialog := chat.Dialog{
		Id:      dialog.Id,
		UserId1: dialog.UserId1,
		UserId2: dialog.UserId2,
	}

	for idx := range dialog.Messages {
		ts := timestamppb.New(dialog.Messages[idx].CreatedAt)
		pbMsg := chat.Message{
			Id:         dialog.Messages[idx].ID,
			DialogId:   dialog.Messages[idx].DialogID,
			SenderId:   dialog.Messages[idx].SenderID,
			ReceiverId: dialog.Messages[idx].ReceiverID,
			Body:       dialog.Messages[idx].Body,
			CreatedAt:  ts,
		}
		pbDialog.Messages = append(pbDialog.Messages, &pbMsg)
	}

	dialogId, err := chatMS.client.CreateDialog(ctx, &pbDialog)
	if err != nil {
		return errors.Wrap(err, "chat microservice error")
	}

	dialog.Id = dialogId.Id

	return nil
}

func (chatMS *microService) CreateMessage(message *models.Message) error {
	ctx := context.Background()

	ts := timestamppb.New(message.CreatedAt)
	pbMessage := chat.Message{
		Id:         message.ID,
		DialogId:   message.DialogID,
		SenderId:   message.SenderID,
		ReceiverId: message.ReceiverID,
		Body:       message.Body,
		CreatedAt:  ts,
	}

	for idx := range message.Attachments {
		pbMessage.AttachmentsIds = append(pbMessage.AttachmentsIds, message.Attachments[idx].ID)
	}

	_, err := chatMS.client.CreateMessage(ctx, &pbMessage)
	if err != nil {
		return errors.Wrap(err, "chat microservice error")
	}

	return nil
}
