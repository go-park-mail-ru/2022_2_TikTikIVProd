package postgres

import (
	attachmentRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/AttachmentMicroservice/internal/attachment/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/AttachmentMicroservice/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Attachment struct {
	ID      uint64
	AttLink string
	Type    int  `gorm:"column:ttype"`
}

type MessageAttachmentRelation struct {
	MessageID uint64
	AttID     uint64
}

func (Attachment) TableName() string {
	return "attachments"
}

type attachmentRepository struct {
	db *gorm.DB
}

func NewAttachmentRepository(db *gorm.DB) attachmentRep.RepositoryI {
	return &attachmentRepository{
		db: db,
	}
}

func toPostgresAttachment(im *models.Attachment) *Attachment {
	ttype := 0

	switch im.Type {
	case models.FileAtt:
		ttype = 1
	case models.ImageAtt:
		ttype = 0
	}

	return &Attachment{
		ID:      im.ID,
		AttLink: im.AttLink,
		Type:    ttype,
	}
}

func toModelAttachment(im *Attachment) *models.Attachment {
	ttype := models.ImageAtt

	switch im.Type {
	case 1:
		ttype = models.FileAtt
	}

	return &models.Attachment{
		ID:      im.ID,
		AttLink: im.AttLink,
		Type:    ttype,
	}
}

func toModelAttachments(attachments []*Attachment) []*models.Attachment {
	out := make([]*models.Attachment, len(attachments))

	for i, b := range attachments {
		out[i] = toModelAttachment(b)
	}

	return out
}

func (dbAttachment *attachmentRepository) GetPostAttachments(postID uint64) ([]*models.Attachment, error) {
	var attachments []*Attachment
	tx := dbAttachment.db.Model(Attachment{}).Joins("JOIN user_posts_attachments upi ON upi.att_id = attachments.id AND upi.user_post_id = ?", postID).Scan(&attachments)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return toModelAttachments(attachments), nil
}

func (dbAttachment *attachmentRepository) GetMessageAttachments(messageID uint64) ([]*models.Attachment, error) {
	var attachments []*Attachment
	tx := dbAttachment.db.Model(Attachment{}).Joins("JOIN message_attachments upi ON upi.att_id = attachments.id AND upi.message_id = ?", messageID).Scan(&attachments)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return toModelAttachments(attachments), nil
}

func (dbAttachment *attachmentRepository) GetAttachment(attachmentID uint64) (*models.Attachment, error) {
	var att Attachment
	tx := dbAttachment.db.Table("attachments").First(&att, attachmentID)

	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "Get attachment repository error")
	}

	return toModelAttachment(&att), nil
}

func (dbAttachment *attachmentRepository) CreateAttachment(attachment *models.Attachment) error {
	att := toPostgresAttachment(attachment)

	tx := dbAttachment.db.Create(att)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "attachmentRepository.CreateAttachment error while insert attachment")
	}

	attachment.ID = att.ID

	return nil
}

