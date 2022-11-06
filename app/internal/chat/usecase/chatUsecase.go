package usecase

import (
	chatRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/chat/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models/chat/dto"
	"github.com/pkg/errors"
)

type ChatUsecaseI interface {
	CreateDialog(req *dto.CreateDialogRequest, resp *dto.CreateDialogResponse) error
}

type chatUsecase struct {
	chatRepo chatRep.RepositoryI
}

func NewChatUsecase(cr chatRep.RepositoryI) ChatUsecaseI {
	return &chatUsecase{
		chatRepo: cr,
	}
}

func (c *chatUsecase) CreateDialog(req *dto.CreateDialogRequest, resp *dto.CreateDialogResponse) error {
	dialog, err := c.chatRepo.CreateDialog(req)

	if err != nil {
		return errors.Wrap(err, "Error in func chatUsecase.CreateDialog")
	}

	resp.DialogID = dialog.ID

	err = c.chatRepo.AddUserDialogRelations(dialog.ID, req.ParticipantsIDs)

	if err != nil {
		return errors.Wrap(err, "Error in func chatUsecase.CreateDialog")
	}

	return nil
}
