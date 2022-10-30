package dto

import "gopkg.in/mgo.v2/bson"

type GetDialogsRequest struct {
	UserID string `json:"user_id"`
}

type GetDialogsResponseBody struct {
	Dialogs     []Dialog `json:"dialogs"`
	Total       int64    `json:"total"`
	AmountPages int64    `json:"amount_pages"`
}

type Dialog struct {
	ID           bson.ObjectId `json:"id"`
	Name         string        `json:"name"`
	Participants []string      `json:"participants"`
	CreatedAt    int64         `json:"created_at"`
}
