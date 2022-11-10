package models

import "time"

type Dialog struct {
	Id       int       `json:"dialog_id" readonly:"true" gorm:"column:id"`
	UserId1  int       `gorm:"column:user_id1"`
	UserId2  int       `gorm:"column:user_id2"`
	Messages []Message `json:"messages,omitempty" readonly:"true" gorm:"-"`
}

type Message struct {
	ID         int       `json:"id" readonly:"true" gorm:"column:id"`
	DialogID   int       `json:"dialog_id" gorm:"column:chat_id"`
	SenderID   int       `json:"sender_id" readonly:"true" gorm:"column:sender_id"`
	ReceiverID int       `json:"receiver_id" gorm:"column:receiver_id"`
	Body       string    `json:"body" gorm:"column:body"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`
}
