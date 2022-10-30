package mongo

import "gopkg.in/mgo.v2/bson"

type Message struct {
	ID        bson.ObjectId `bson:"_id"`
	Body      string        `bson:"body"`
	SenderID  string        `bson:"sender_id"`
	CreatedAt int64         `bson:"created_at"`
}

type Dialog struct {
	ID           bson.ObjectId `bson:"_id"`
	Name         string        `bson:"name"`
	Participants []string      `bson:"participants"`
	Messages     []Message     `bson:"messages"`
	CreatedAt    int64         `bson:"created_at"`
}
