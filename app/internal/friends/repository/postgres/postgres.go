package postgres

import (
	"gorm.io/gorm"
	"github.com/pkg/errors"

	friendsRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/friends/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
)

type dataBase struct {
	db *gorm.DB
}

func New(db *gorm.DB) friendsRep.RepositoryI {
	return &dataBase{
		db: db,
	}
}

func (dbFriends *dataBase) AddFriend(friends models.Friends) error {
	tx := dbFriends.db.Create(friends)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table friends)")
	}

	return nil
}

func (dbFriends *dataBase) DeleteFriend(friends models.Friends) error {
	tx := dbFriends.db.Where("id1 = ? AND id2 = ?", friends.Id1, friends.Id2).Delete(&models.Friends{})
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table friends)")
	}

	return nil
}

func (dbFriends *dataBase) CheckFriends(friends models.Friends) (bool, error) {
	var count int64
	tx := dbFriends.db.Model(&models.Friends{}).Where("id1 = ? AND id2 = ?",
									friends.Id1, friends.Id2).Count(&count)
	if tx.Error != nil {
		return false, errors.Wrap(tx.Error, "database error (table friends)")
	}
	return count > 0, nil
}

