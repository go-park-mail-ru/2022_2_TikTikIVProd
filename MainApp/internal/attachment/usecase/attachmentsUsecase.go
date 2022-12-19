package attachmentUsecase

import (
	attachmentRepository "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/attachment/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	"github.com/pkg/errors"
)

type AttachmentUseCaseI interface {
	GetPostAttachments(postID uint64) ([]*models.Attachment, error)
	GetAttachmentById(attachmentID uint64) (*models.Attachment, error)
	CreateAttachment(att *models.Attachment) error
}

type attachmentUsecase struct {
	attachmentRep attachmentRepository.RepositoryI
}

func NewAttachmentUsecase(ir attachmentRepository.RepositoryI) AttachmentUseCaseI {
	return &attachmentUsecase{
		attachmentRep: ir,
	}
}

func (i *attachmentUsecase) GetPostAttachments(postID uint64) ([]*models.Attachment, error) {
	attachments, err := i.attachmentRep.GetPostAttachments(postID)

	if err != nil {
		return nil, err
	}

	return attachments, nil
}

func (i *attachmentUsecase) GetMessageAttachments(postID uint64) ([]models.Attachment, error) {
	attachments, err := i.attachmentRep.GetMessageAttachments(postID)

	if err != nil {
		return nil, err
	}

	return attachments, nil
}

func (i *attachmentUsecase) GetAttachmentById(attachmentID uint64) (*models.Attachment, error) {
	Attachment, err := i.attachmentRep.GetAttachment(attachmentID)

	if err != nil {
		return nil, errors.Wrap(err, "GetAttachment usecase error")
	}

	return Attachment, nil
}

func (i *attachmentUsecase) CreateAttachment(att *models.Attachment) error {
	err := i.attachmentRep.CreateAttachment(att)

	if err != nil {
		return errors.Wrap(err, "attachmentUsecase.CreateAttachment error")
	}

	return nil
}
