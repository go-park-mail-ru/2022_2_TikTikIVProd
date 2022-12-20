package microservice

import (
	"context"
	attachmentRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/attachment/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	Attachment "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/proto/attachment"
	"github.com/pkg/errors"
)

type microService struct {
	client Attachment.AttachmentsClient
}

func New(client Attachment.AttachmentsClient) attachmentRep.RepositoryI {
	return &microService{
		client: client,
	}
}

func (attachmentMS *microService) GetPostAttachments(postID uint64) ([]*models.Attachment, error) {
	ctx := context.Background()

	pbGetPostAttachmentsRequest := Attachment.GetPostAttachmentsRequest{
		PostId: postID,
	}

	pbAttachments, err := attachmentMS.client.GetPostAttachments(ctx, &pbGetPostAttachmentsRequest)
	if err != nil {
		return nil, errors.Wrap(err, "Attachment microservice error")
	}

	attachments := make([]*models.Attachment, 0)

	for idx := range pbAttachments.Attachments {
		att := &models.Attachment{
			ID:      pbAttachments.Attachments[idx].Id,
			AttLink: pbAttachments.Attachments[idx].AttLink,
			Type:    pbAttachments.Attachments[idx].Type,
		}
		attachments = append(attachments, att)
	}

	return attachments, nil
}

func (attachmentMS *microService) GetMessageAttachments(msgId uint64) ([]models.Attachment, error) {
	ctx := context.Background()

	pbGetMessageAttachmentsRequest := Attachment.GetMessageAttachmentsRequest{
		MessageId: msgId,
	}

	pbAttachments, err := attachmentMS.client.GetMessageAttachments(ctx, &pbGetMessageAttachmentsRequest)
	if err != nil {
		return nil, errors.Wrap(err, "Attachment microservice error")
	}

	attachments := make([]models.Attachment, 0)

	for idx := range pbAttachments.Attachments {
		att := models.Attachment{
			ID:      pbAttachments.Attachments[idx].Id,
			AttLink: pbAttachments.Attachments[idx].AttLink,
			Type:    pbAttachments.Attachments[idx].Type,
		}
		attachments = append(attachments, att)
	}

	return attachments, nil
}

func (attachmentMS *microService) GetAttachment(attachmentID uint64) (*models.Attachment, error) {
	ctx := context.Background()

	pbGetAttachmentRequest := Attachment.AttachmentId{
		AttachmentId: attachmentID,
	}

	pbAttachment, err := attachmentMS.client.GetAttachment(ctx, &pbGetAttachmentRequest)
	if err != nil {
		return nil, errors.Wrap(err, "Attachment microservice error")
	}

	att := &models.Attachment{
		ID:      pbAttachment.Id,
		AttLink: pbAttachment.AttLink,
	}

	return att, nil
}

func (attachmentMS *microService) CreateAttachment(att *models.Attachment) error {
	ctx := context.Background()

	pbAttachment := Attachment.Attachment{
		AttLink: att.AttLink,
	}

	attId, err := attachmentMS.client.CreateAttachment(ctx, &pbAttachment)
	if err != nil {
		return errors.Wrap(err, "Attachment microservice error")
	}

	att.ID = attId.AttachmentId

	return nil
}
