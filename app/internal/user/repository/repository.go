package repository

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/model"
)

type RepositoryI interface {
	SelectUserByNickName(name string) (*model.User, error)
	SelectUserByEmail(email string) (*model.User, error)
	CreateUser(u model.User) (*model.User, error)
	CreateCookie(c model.Cookie) (*model.Cookie, error)
	SelectCookie(value string) (*model.Cookie, error)
	DeleteCookie(value string) (error)
	SelectUserById(id int) (*model.User, error)
}

type dataBase struct {
	db *gorm.DB
}

func New(db *gorm.DB) RepositoryI {
	return &dataBase{
		db: db,
	}
}

func (dbUsers *dataBase) SelectUserById(id int) (*model.User, error) {
	user := model.User{}

	row := dbUsers.db.Table("users").Where("id = ?", id).Row()
	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.NickName, &user.Avatar, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (dbUsers *dataBase) SelectUserByNickName(nickname string) (*model.User, error) {
	user := model.User{}

	row := dbUsers.db.Table("users").Where("nick_name = ?", nickname).Row()
	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.NickName, &user.Avatar, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (dbUsers *dataBase) SelectUserByEmail(email string) (*model.User, error) {
	user := model.User{}
	//row := dbUsers.db.Table("users").Select("").Where("email = ?", email).Joins("JOIN images ON users.avatar_img_id=images.id").Row()
	row := dbUsers.db.Table("users").Select("id, first_name, last_name," +
		"nick_name, COALESCE(avatar_img_id, 0), email, passhash").Where("email = ?", email).Row()
	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.NickName, &user.Avatar, &user.Email, &user.Password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}

func (dbUsers *dataBase) CreateUser(u model.User) (*model.User, error) {
	user := model.User{}

	tx := dbUsers.db.Table("users").Exec("INSERT INTO users (first_name, last_name, nick_name, email, passhash) VALUES (?, ?, ?, ?, ?) RETURNING *",
			u.FirstName, u.LastName, u.NickName, u.Email, u.Password).Scan(&user)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return &user, nil
}

func (dbUsers *dataBase) CreateCookie(c model.Cookie) (*model.Cookie, error) {
	cookie := model.Cookie{}

	row := dbUsers.db.Table("cookies").Create(&c).Clauses(clause.Returning{}).Row()
	err := row.Scan(&cookie.SessionToken, &cookie.UserId, &cookie.Expires)
	if err != nil {
		return nil, err
	}
	
	return &cookie, nil
}

func (dbUsers *dataBase) SelectCookie(value string) (*model.Cookie, error) {
	cookie := model.Cookie{}

	row := dbUsers.db.Table("cookies").Where("value = ?", value).Row()
	err := row.Scan(&cookie.SessionToken, &cookie.UserId, &cookie.Expires)
	if err != nil {
		return nil, err
	}

	return &cookie, nil
}

func (dbUsers *dataBase) DeleteCookie(value string) (error) {
	err := dbUsers.db.Table("cookies").Delete(&model.Cookie {
											SessionToken: value,
										}).Error
	if err != nil {
		return err
	}
	
	return nil
}

