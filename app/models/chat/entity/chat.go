package entity

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Message struct {
	ID        bson.ObjectId `bson:"_id"`
	Body      string        `bson:"body"`
	SenderID  string        `bson:"sender_id"`
	CreatedAt time.Time     `bson:"created_at"`
}

type Dialog struct {
	ID           bson.ObjectId `bson:"_id"`
	Name         string        `bson:"name"`
	Participants []string      `bson:"participants"`
	Messages     []Message     `bson:"messages"`
	CreatedAt    time.Time     `bson:"created_at"`
}
