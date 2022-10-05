package postgres

import (
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

	row := dbUsers.db.Table("users").Where("email = ?", email).Row()
	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.NickName, &user.Avatar, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (dbUsers *dataBase) CreateUser(u model.User) (*model.User, error) {
	user := model.User{}

	row := dbUsers.db.Exec("INSERT INTO users (first_name, last_name, nick_name, avatar_img_id, email, passhash) VALUES (?, ?, ?, ?, ?, ?) RETURNING *",
			u.FirstName, u.LastName, u.NickName, u.Avatar, u.Email, u.Password).Row()
	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.NickName, &user.Avatar, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (dbUsers *dataBase) CreateCookie(c model.Cookie) (*model.Cookie, error) {
	cookie := model.Cookie{}

	row := dbUsers.db.Exec("INSERT INTO cookies (value, user_id, expires) VALUES (?, ?, ?) RETURNING *",
				   c.SessionToken, c.UserId, c.Expires).Row()
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

