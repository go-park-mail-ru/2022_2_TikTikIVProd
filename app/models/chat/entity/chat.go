package entity

import (
	"time"
)

type Message struct {
	ID        int
	Body      string
	SenderID  int
	ChatID    int
	CreatedAt time.Time
}

func (Message) TableName() string {
	return "Message_"
}

type Dialog struct {
	ID int
	//OwnerID   int //todo
	Name string
	//CreatedAt time.Time
}

func (Dialog) TableName() string {
	return "chat"
}

type UserDialogRelation struct {
	ChatID int
	UserID int
}

func (UserDialogRelation) TableName() string {
	return "user_chat"
}
