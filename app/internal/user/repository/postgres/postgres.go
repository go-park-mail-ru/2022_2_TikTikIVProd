package postgres

import (
	"errors"

	"gorm.io/gorm"

	userRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
)

type dataBase struct {
	db *gorm.DB
}

func New(db *gorm.DB) userRep.RepositoryI {
	return &dataBase{
		db: db,
	}
}

func (dbUsers *dataBase) SelectUserById(id int) (*models.User, error) {
	user := models.User{}

	tx := dbUsers.db.Table("users").Where("id = ?", id).Take(&user)
	if tx.Error == gorm.ErrRecordNotFound {
		return nil, errors.New("user with such id doesn't exist")
	} else if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (dbUsers *dataBase) SelectUserByNickName(nickname string) (*models.User, error) {
	user := models.User{}

	tx := dbUsers.db.Table("users").Where("nick_name = ?", nickname).Take(&user)
	if tx.Error == gorm.ErrRecordNotFound {
		return nil, errors.New("user with such nickname doesn't exist")
	} else if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (dbUsers *dataBase) SelectUserByEmail(email string) (*models.User, error) {
	user := models.User{}
	tx := dbUsers.db.Table("users").Where("email = ?", email).Take(&user)
	if tx.Error == gorm.ErrRecordNotFound {
		return nil, errors.New("user with such email doesn't exist")
	} else if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (dbUsers *dataBase) CreateUser(u models.User) (*models.User, error) {
	user := models.User{}

	tx := dbUsers.db.Table("users").Raw("INSERT INTO users (first_name, last_name, nick_name, email, passhash) VALUES (?, ?, ?, ?, ?) RETURNING *",
		u.FirstName, u.LastName, u.NickName, u.Email, u.Password).Scan(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}
