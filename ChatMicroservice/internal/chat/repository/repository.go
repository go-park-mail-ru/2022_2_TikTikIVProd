package repository

import (
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/ChatMicroservice/models"
)

type RepositoryI interface {
	SelectDialog(id uint64) (*models.Dialog, error)
	SelectDialogByUsers(userId, friendId uint64) (*models.Dialog, error)
	SelectMessages(id uint64) ([]models.Message, error)
	CreateDialog(dialog *models.Dialog) error
	CreateMessage(message *models.Message) error
	SelectAllDialogs(userId uint64) ([]models.Dialog, error)
}
