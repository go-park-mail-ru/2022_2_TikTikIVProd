package repository

import (
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models/chat/dto"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models/chat/entity"
)

type RepositoryI interface {
	CreateDialog(req *dto.CreateDialogRequest) (*entity.Dialog, error)
	AddUserDialogRelations(dialogId int, userIds []int) error
}
