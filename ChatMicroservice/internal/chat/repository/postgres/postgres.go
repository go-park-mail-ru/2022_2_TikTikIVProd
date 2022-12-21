package postgres

import (
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/ChatMicroservice/internal/chat/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/ChatMicroservice/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type Message struct {
	ID         uint64
	DialogID   uint64
	SenderID   uint64
	ReceiverID uint64
	Body       string
	CreatedAt  time.Time
	StickerID  uint64
}

func (Message) TableName() string {
	return "message"
}

type MessageAttachmentsRelation struct {
	MessageID    uint64 `gorm:"column:message_id"`
	AttachmentID uint64 `gorm:"column:att_id"`
}

func (MessageAttachmentsRelation) TableName() string {
	return "message_attachments"
}

func toPostgresMessage(p *models.Message) *Message {
	return &Message{
		ID:         p.ID,
		DialogID:   p.DialogID,
		SenderID:   p.SenderID,
		ReceiverID: p.ReceiverID,
		CreatedAt:  p.CreatedAt,
		Body:       p.Body,
		StickerID:  p.StickerID,
	}
}

func toModelMessage(p *Message) *models.Message {
	return &models.Message{
		ID:         p.ID,
		DialogID:   p.DialogID,
		SenderID:   p.SenderID,
		ReceiverID: p.ReceiverID,
		CreatedAt:  p.CreatedAt,
		Body:       p.Body,
		StickerID:  p.StickerID,
	}
}

func toModelMessages(posts []*Message) []*models.Message {
	out := make([]*models.Message, len(posts))

	for i, b := range posts {
		out[i] = toModelMessage(b)
	}

	return out
}

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

	messageAttachments := make([]MessageAttachmentsRelation, 0, 10)
	for _, elem := range message.Attachments {
		messageAttachments = append(messageAttachments, MessageAttachmentsRelation{MessageID: message.ID, AttachmentID: elem.ID})
	}

	if len(messageAttachments) > 0 {
		tx = dbChat.db.Create(&messageAttachments)
		if tx.Error != nil {
			return errors.Wrap(tx.Error, "chatRepository.CreateMessage error while insert relation")
		}
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
	tx := dbChat.db.Table("message").Order("id").Find(&messages, "chat_id = ?", id)

	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table message)")
	}

	return messages, nil
}

func (dbChat *chatRepository) SelectAllDialogs(userId uint64) ([]models.Dialog, error) {
	dialogs := make([]models.Dialog, 0, 10)
	tx := dbChat.db.Table("chat").Omit("messages").Find(&dialogs, "user_id1 = ? OR user_id2 = ?", userId, userId)

	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table chat)")
	}

	return dialogs, nil
}
