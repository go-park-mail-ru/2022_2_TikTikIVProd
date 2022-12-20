package models

import "time"

type Dialog struct {
	Id       uint64    `json:"dialog_id" readonly:"true" gorm:"column:id"`
	UserId1  uint64    `gorm:"column:user_id1"`
	UserId2  uint64    `gorm:"column:user_id2"`
	Messages []Message `json:"messages,omitempty" readonly:"true" gorm:"-"`
}

type Message struct {
	ID          uint64       `json:"id" readonly:"true" gorm:"column:id"`
	DialogID    uint64       `json:"dialog_id" gorm:"column:chat_id"`
	SenderID    uint64       `json:"sender_id" readonly:"true" gorm:"column:sender_id"`
	ReceiverID  uint64       `json:"receiver_id" gorm:"column:receiver_id"`
	Body        string       `json:"body" gorm:"column:text"`
	CreatedAt   time.Time    `json:"created_at" gorm:"column:created_at"`
	Attachments []Attachment `json:"attachments"`
	StickerID   uint64       `json:"sticker"`
}
