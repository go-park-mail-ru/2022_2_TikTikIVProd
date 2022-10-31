package mongo

import (
	"context"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/chat/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models/chat/dto"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models/chat/entity"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func dialogRequestToEntity(d *dto.CreateDialogRequest) *entity.Dialog {
	return &entity.Dialog{
		ID:           bson.NewObjectId(),
		Participants: d.Participants,
		Name:         d.Name,
		CreatedAt:    time.Now(),
	}
}

type chatRepository struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func (dbChat *chatRepository) CreateDialog(d *dto.CreateDialogRequest) (*dto.CreateDialogResponse, error) {
	dialog := dialogRequestToEntity(d)
	_, err := dbChat.coll.InsertOne(context.TODO(), dialog)

	if err != nil {
		return nil, errors.Wrap(err, "chatRepository.CreateDialog error while create")
	}

	return nil, nil
}

func NewChatRepository(db *mongo.Database, coll *mongo.Collection) repository.RepositoryI {
	return &chatRepository{
		db:   db,
		coll: coll,
	}
}
