package postgres

import (
	"errors"

	"gorm.io/gorm"

	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/model"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/repository"
)

type dataBase struct {
	db *gorm.DB
}

func New(db *gorm.DB) repository.RepositoryI {
	return &dataBase{
		db: db,
	}
}

func (dbUsers *dataBase) SelectUserById(id int) (*model.User, error) {
	user := model.User{}

	tx := dbUsers.db.Table("users").Where("id = ?", id).Scan(&user)
	if tx.Error != nil {
		return nil, tx.Error
	} else if user.NickName == "" {
		return nil, errors.New("user with such id doesn't exist")
	}

	return &user, nil
}

func (dbUsers *dataBase) SelectUserByNickName(nickname string) (*model.User, error) {
	user := model.User{}

	tx := dbUsers.db.Table("users").Where("nick_name = ?", nickname).Scan(&user)
	if tx.Error != nil {
		return nil, tx.Error
	} else if user.NickName == "" {
		return nil, errors.New("user with such nickname doesn't exist")
	}

	return &user, nil
}

func (dbUsers *dataBase) SelectUserByEmail(email string) (*model.User, error) {
	user := model.User{}
	tx := dbUsers.db.Table("users").Where("email = ?", email).Scan(&user)
	if tx.Error != nil {
		return nil, tx.Error
	} else if user.NickName == "" {
		return nil, errors.New("user with such email doesn't exists")
	}

	return &user, nil
}

func (dbUsers *dataBase) CreateUser(u model.User) (*model.User, error) {
	user := model.User{}

	tx := dbUsers.db.Table("users").Raw("INSERT INTO users (first_name, last_name, nick_name, email, passhash) VALUES (?, ?, ?, ?, ?) RETURNING *",
			u.FirstName, u.LastName, u.NickName, u.Email, u.Password).Scan(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (dbUsers *dataBase) CreateCookie(c model.Cookie) (*model.Cookie, error) {
	cookie := model.Cookie{}

	tx := dbUsers.db.Table("cookies").Raw("INSERT INTO cookies VALUES (?, ?, ?) RETURNING *",
			c.SessionToken, c.UserId, c.Expires).Scan(&cookie)
	if tx.Error != nil {
		return nil, tx.Error
	}
	
	return &cookie, nil
}

func (dbUsers *dataBase) SelectCookie(value string) (*model.Cookie, error) {
	cookie := model.Cookie{}

	tx := dbUsers.db.Table("cookies").Where("value = ?", value).Scan(&cookie)
	if tx.Error != nil || cookie.SessionToken == "" {
		return nil, errors.New("cookie doesn't exist")
	}

	return &cookie, nil
}

func (dbUsers *dataBase) DeleteCookie(value string) (error) {
	tx := dbUsers.db.Table("cookies").Exec("DELETE FROM cookies WHERE value = ?", value)
	if tx.Error != nil {
		return tx.Error
	}
	
	return nil
}

