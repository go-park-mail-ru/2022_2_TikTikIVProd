package postgres

import (
	"fmt"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/chat/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models/chat/dto"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models/chat/entity"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type chatRepository struct {
	db *gorm.DB
}

func (dbChat *chatRepository) CreateDialog(req *dto.CreateDialogRequest) (*entity.Dialog, error) {
	dialog := req.ToDialogEntities()

	tx := dbChat.db.Create(dialog)

	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "chatRepository.CreateDialog error while insert dialog")
	}

	return dialog, nil
}

func (dbChat *chatRepository) AddUserDialogRelations(dialogId int, userIds []int) error {
	relations := make([]entity.UserDialogRelation, 0, 10)

	for _, id := range userIds {
		relations = append(relations, entity.UserDialogRelation{ChatID: dialogId, UserID: id})
	}

	fmt.Println(relations)
	tx := dbChat.db.Create(&relations)

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "chatRepository.AddUserDialogRelations")
	}

	return nil
}

func NewChatRepository(db *gorm.DB) repository.RepositoryI {
	return &chatRepository{
		db: db,
	}
}
