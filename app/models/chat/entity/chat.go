package entity

import (
	"time"
)

type Message struct {
	ID        int
	Body      string
	SenderID  int
	CreatedAt time.Time
}

func (Message) TableName() string {
	return "message"
}

type Dialog struct {
	ID int
	//OwnerID   int //todo
	Name      string
	CreatedAt time.Time
}

func (Dialog) TableName() string {
	return "chat"
}

type MessageDialogRelation struct {
	ChatID    int
	MessageID int
}

func (MessageDialogRelation) TableName() string {
	return "message_chat"
}

type UserDialogRelation struct {
	ChatID int
	UserID int
}

func (UserDialogRelation) TableName() string {
	return "user_chat"
}
