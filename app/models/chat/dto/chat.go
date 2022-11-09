package dto

import (
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models/chat/entity"
	"time"
)

type CreateDialogRequest struct {
	UserID          int    `json:"-"`
	Name            string `json:"name"`
	ParticipantsIDs []int  `json:"participants_ids" validate:"required"`
}

func (dialog CreateDialogRequest) ToDialogEntities() *entity.Dialog {
	return &entity.Dialog{
		Name: dialog.Name,
		//CreatedAt: time.Now(),
	}
}

type CreateDialogResponse struct {
	DialogID int `json:"dialog_id"`
}

type GetDialogsRequest struct {
	UserID int `json:"-"`
}

type GetDialogsInfoResponseBody struct {
	DialogInfo []DialogInfo `json:"dialog_info"`
}

type Message struct {
	ID        int       `json:"id"`
	DialogID  int       `json:"dialog_id"`
	AuthorID  int       `json:"author_id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

type GetDialogRequest struct {
	DialogID int `json:"id"`
}

type GetDialogResponse struct {
	DialogInfo DialogInfo `json:"dialog"`
	Messages   []Message  `json:"messages"`
	//Total       int      `json:"total"`
	//AmountPages int      `json:"amount_pages"`
}

type Participants struct {
	UserID    int    `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type DialogInfo struct {
	DialogID     int            `json:"dialog_id"`
	Name         string         `json:"name"`
	Participants []Participants `json:"participants"`
	CreatedAt    time.Time      `json:"created_at"`
}
