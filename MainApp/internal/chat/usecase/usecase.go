package usecase

import (
	"time"

	chatRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/chat/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	"github.com/pkg/errors"
)

type UseCaseI interface {
	SelectDialog(id uint64) (*models.Dialog, error)
	SelectDialogByUsers(userId, friendId uint64) (*models.Dialog, error)
	SelectAllDialogs(userId uint64) ([]models.Dialog, error)
	SendMessage(message *models.Message) error
}

type useCase struct {
	chatRepository chatRep.RepositoryI
}

func New(rep chatRep.RepositoryI) UseCaseI {
	return &useCase{
		chatRepository: rep,
	}
}

func (uc *useCase) SelectDialog(id uint64) (*models.Dialog, error) {
	dialog, err := uc.chatRepository.SelectDialog(id)
	if err != nil {
		return nil, errors.Wrap(err, "chat repository error")
	}

	messages, err := uc.chatRepository.SelectMessages(id)
	if err != nil {
		return nil, errors.Wrap(err, "chat repository error")
	}

	dialog.Messages = messages

	return dialog, nil
}

func (uc *useCase) SelectDialogByUsers(userId, friendId uint64) (*models.Dialog, error) {
	dialog, err := uc.chatRepository.SelectDialogByUsers(userId, friendId)
	if err != nil {
		return nil, errors.Wrap(err, "chat repository error")
	}

	messages, err := uc.chatRepository.SelectMessages(dialog.Id)
	if err != nil {
		return nil, errors.Wrap(err, "chat repository error")
	}

	dialog.Messages = messages

	return dialog, nil
}

func (uc *useCase) SelectAllDialogs(userId uint64) ([]models.Dialog, error) {
	dialogs, err := uc.chatRepository.SelectAllDialogs(userId)
	if err != nil {
		return nil, errors.Wrap(err, "chat repository error")
	}

	return dialogs, nil
}

func (uc *useCase) SendMessage(message *models.Message) error {
	if _, err := uc.chatRepository.SelectDialog(message.DialogID); err != nil {
		dialog := models.Dialog {
			UserId1: message.SenderID,
			UserId2: message.ReceiverID,
		}
		dialog.Messages = append(dialog.Messages, *message)
		err := uc.chatRepository.CreateDialog(&dialog)

		if err != nil {
			return errors.Wrap(err, "chat repository error")
		}

		message.DialogID = dialog.Id
	}

	message.CreatedAt = time.Now()
	err := uc.chatRepository.CreateMessage(message)
	if err != nil {
		return errors.Wrap(err, "message repository error")
	}

	return nil
}
