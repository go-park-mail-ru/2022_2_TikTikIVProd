package postgres

import (
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/chat/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type chatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) repository.RepositoryI {
	return &chatRepository{
		db: db,
	}
}

func (dbChat *chatRepository) CreateDialog(dialog *models.Dialog) error {
	tx := dbChat.db.Table("chat").Create(dialog)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table chat)")
	}

	return nil
}

func (dbChat *chatRepository) CreateMessage(message *models.Message) error {
	tx := dbChat.db.Table("message").Create(message)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table message)")
	}

	return nil
}

func (dbChat *chatRepository) SelectDialog(id uint64) (*models.Dialog, error) {
	dialog := models.Dialog{}

	tx := dbChat.db.Table("chat").Where("id = ?", id).Take(&dialog)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, models.ErrNotFound
	} else if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table chat)")
	}

	return &dialog, nil
}

func (dbChat *chatRepository) SelectDialogByUsers(userId, friendId uint64) (*models.Dialog, error) {
	dialog := models.Dialog{}

	tx := dbChat.db.Table("chat").
	Where("(user_id1 = ? AND user_id2 = ?) OR (user_id1 = ? AND user_id2 = ?)",
	userId, friendId, friendId, userId).Take(&dialog)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, models.ErrNotFound
	} else if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table chat)")
	}

	return &dialog, nil
}

func (dbChat *chatRepository) SelectMessages(id uint64) ([]models.Message, error) {
	messages := make([]models.Message, 0, 10)
	tx := dbChat.db.Table("message").Order("created_at").Find(&messages, "chat_id = ?", id)

	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (tables message)")
	}

	return messages, nil
}

func (dbChat *chatRepository) SelectAllDialogs(userId uint64) ([]models.Dialog, error) {
	dialogs := make([]models.Dialog, 0, 10)
	tx := dbChat.db.Table("chat").Omit("messages").Find(&dialogs, "user_id1 = ? OR user_id2 = ?", userId, userId)

	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (tables chat)")
	}

	return dialogs, nil
}

