package models

import "time"

type Dialog struct {
	Id       uint64    `gorm:"column:id"`
	UserId1  uint64    `gorm:"column:user_id1"`
	UserId2  uint64    `gorm:"column:user_id2"`
	Messages []Message `gorm:"-"`
}

type Message struct {
	ID         uint64    `gorm:"column:id"`
	DialogID   uint64    `gorm:"column:chat_id"`
	SenderID   uint64    `gorm:"column:sender_id"`
	ReceiverID uint64    `gorm:"column:receiver_id"`
	Body       string    `gorm:"column:text"`
	CreatedAt  time.Time `gorm:"column:created_at"`
}
