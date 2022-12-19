package delivery

import (
	"context"

	attachmentUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/AttachmentMicroservice/internal/attachment/usecase"
	attachment "github.com/go-park-mail-ru/2022_2_TikTikIVProd/AttachmentMicroservice/proto"
)

type AttachmentManager struct {
	attachment.UnimplementedAttachmentsServer
	AttachmentUC attachmentUsecase.UseCaseI
}

func New(uc attachmentUsecase.UseCaseI) attachment.AttachmentsServer {
	return AttachmentManager{AttachmentUC: uc}
}

func (im AttachmentManager) GetPostAttachments(ctx context.Context, pbAttachments *attachment.GetPostAttachmentsRequest) (*attachment.GetPostAttachmentsResponse, error) {
	resp, err := im.AttachmentUC.GetPostAttachments(pbAttachments)
	return resp, err
}

func (im AttachmentManager) GetMessageAttachments(ctx context.Context, pbAttachments *attachment.GetMessageAttachmentsRequest) (*attachment.GetMessageAttachmentsResponse, error) {
	resp, err := im.AttachmentUC.GetMessageAttachments(pbAttachments)
	return resp, err
}

func (im AttachmentManager) GetAttachment(ctx context.Context, pbId *attachment.AttachmentId) (*attachment.Attachment, error) {
	resp, err := im.AttachmentUC.GetAttachment(pbId)
	return resp, err
}

func (im AttachmentManager) CreateAttachment(ctx context.Context, pbAttachment *attachment.Attachment) (*attachment.AttachmentId, error) {
	_, err := im.AttachmentUC.CreateAttachment(pbAttachment)
	return &attachment.AttachmentId{AttachmentId: pbAttachment.Id}, err
}

func (im AttachmentManager) AddAttachmentsToPost(ctx context.Context, pbAttachment *attachment.AddAttachmentsToMessageRequest) (*attachment.Nothing, error) {
	_, err := im.AttachmentUC.AddAttachmentsToMessage(pbAttachment)
	return &attachment.Nothing{Dummy: true}, err
}
