package delivery

import (
	"context"
	chatUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/ChatMicroservice/internal/chat/usecase"
	chat "github.com/go-park-mail-ru/2022_2_TikTikIVProd/ChatMicroservice/proto"
)

type ChatManager struct {
	chat.UnimplementedChatServer
	ChatUC chatUsecase.UseCaseI
}

func New(uc chatUsecase.UseCaseI) chat.ChatServer {
	return ChatManager{ChatUC: uc}
}

func (cm ChatManager) SelectDialog(ctx context.Context, dialogId *chat.DialogId) (*chat.Dialog, error) {
	resp, err := cm.ChatUC.SelectDialog(dialogId)
	return resp, err
}

func (cm ChatManager) SelectDialogByUsers(ctx context.Context, dialogUsers *chat.SelectDialogByUsersRequest) (*chat.Dialog, error) {
	resp, err := cm.ChatUC.SelectDialogByUsers(dialogUsers)
	return resp, err
}

func (cm ChatManager) SelectMessages(ctx context.Context, dialogId *chat.DialogId) (*chat.SelectMessagesResponse, error) {
	resp, err := cm.ChatUC.SelectMessages(dialogId)
	return resp, err
}

func (cm ChatManager) SelectAllDialogs(ctx context.Context, userId *chat.SelectAllDialogsRequest) (*chat.SelectAllDialogsResponse, error) {
	resp, err := cm.ChatUC.SelectAllDialogs(userId)
	return resp, err
}

func (cm ChatManager) CreateDialog(ctx context.Context, dialog *chat.Dialog) (*chat.Nothing, error) {
	resp, err := cm.ChatUC.CreateDialog(dialog)
	return resp, err
}

func (cm ChatManager) CreateMessage(ctx context.Context, msg *chat.Message) (*chat.Nothing, error) {
	resp, err := cm.ChatUC.CreateMessage(msg)
	return resp, err
}

