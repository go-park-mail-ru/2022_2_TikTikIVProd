package usecase

import (
	attachmentRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/AttachmentMicroservice/internal/attachment/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/AttachmentMicroservice/models"
	Attachment "github.com/go-park-mail-ru/2022_2_TikTikIVProd/AttachmentMicroservice/proto"
	"github.com/pkg/errors"
)

type UseCaseI interface {
	GetPostAttachments(*Attachment.GetPostAttachmentsRequest) (*Attachment.GetPostAttachmentsResponse, error)
	GetMessageAttachments(request *Attachment.GetMessageAttachmentsRequest) (*Attachment.GetMessageAttachmentsResponse, error)
	GetAttachment(*Attachment.AttachmentId) (*Attachment.Attachment, error)
	CreateAttachment(*Attachment.Attachment) (*Attachment.Nothing, error)
}

type useCase struct {
	attachmentRepository attachmentRep.RepositoryI
}

func New(attachmentRepository attachmentRep.RepositoryI) UseCaseI {
	return &useCase{
		attachmentRepository: attachmentRepository,
	}
}

func (uc *useCase) GetPostAttachments(pbPostId *Attachment.GetPostAttachmentsRequest) (*Attachment.GetPostAttachmentsResponse, error) {
	attachments, err := uc.attachmentRepository.GetPostAttachments(pbPostId.PostId)
	if err != nil {
		return nil, errors.Wrap(err, "Attachment repository postgres error")
	}

	pbAttachments := &Attachment.GetPostAttachmentsResponse{}

	for idx := range attachments {
		att := &Attachment.Attachment{
			Id:      attachments[idx].ID,
			AttLink: attachments[idx].AttLink,
			Type:    attachments[idx].Type,
		}
		pbAttachments.Attachments = append(pbAttachments.Attachments, att)
	}

	return pbAttachments, nil
}

func (uc *useCase) GetMessageAttachments(pbMessage *Attachment.GetMessageAttachmentsRequest) (*Attachment.GetMessageAttachmentsResponse, error) {
	attachments, err := uc.attachmentRepository.GetMessageAttachments(pbMessage.MessageId)
	if err != nil {
		return nil, errors.Wrap(err, "Attachment repository postgres error")
	}

	pbAttachments := &Attachment.GetMessageAttachmentsResponse{}

	for idx := range attachments {
		att := &Attachment.Attachment{
			Id:      attachments[idx].ID,
			AttLink: attachments[idx].AttLink,
			Type:    attachments[idx].Type,
		}
		pbAttachments.Attachments = append(pbAttachments.Attachments, att)
	}

	return pbAttachments, nil
}

func (uc *useCase) GetAttachment(pbAttachmentId *Attachment.AttachmentId) (*Attachment.Attachment, error) {
	att, err := uc.attachmentRepository.GetAttachment(pbAttachmentId.AttachmentId)
	if err != nil {
		return nil, errors.Wrap(err, "Attachment repository postgres error")
	}
	return &Attachment.Attachment{
		Id:      att.ID,
		AttLink: att.AttLink,
		Type:    att.Type,
	}, nil
}

func (uc *useCase) CreateAttachment(pbAttachment *Attachment.Attachment) (*Attachment.Nothing, error) {
	modelAttachment := models.Attachment{
		ID:      pbAttachment.Id,
		AttLink: pbAttachment.AttLink,
		Type:    pbAttachment.Type,
	}
	err := uc.attachmentRepository.CreateAttachment(&modelAttachment)
	if err != nil {
		return nil, errors.Wrap(err, "Attachment repository postgres error")
	}

	pbAttachment.Id = modelAttachment.ID

	return &Attachment.Nothing{Dummy: true}, nil
}

