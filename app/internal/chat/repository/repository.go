package repository

import (
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
)

type RepositoryI interface {
	SelectDialog(id int) (*models.Dialog, error)
	SelectDialogByUsers(userId, friendId int) (*models.Dialog, error)
	SelectMessages(id int) ([]models.Message, error)
	CreateDialog(dialog *models.Dialog) error
	CreateMessage(message *models.Message) error
	SelectAllDialogs(userId int) ([]models.Dialog, error)
}
