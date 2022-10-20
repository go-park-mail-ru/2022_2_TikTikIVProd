package postgres

import (
	"gorm.io/gorm"

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

func (dbFriends *dataBase) AddFriend(friends models.Friends) (error) {
	tx := dbFriends.db.Table("friends").Exec("INSERT INTO friends VALUES (?, ?, ?, ?, ?)", friends.Id1, friends.Id2)
	return tx.Error
}

func (dbFriends *dataBase) DeleteFriend(friends models.Friends) (error) {
	tx := dbFriends.db.Table("friends").Exec("DELETE FROM friends WHERE id1 = ? AND id2 = ?", friends.Id1, friends.Id2)
	return tx.Error
}
