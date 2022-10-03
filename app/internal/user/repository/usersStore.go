package usersStore

import (
	"database/sql"
	usersRrep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/usecase"
)

type DataBaseUsers struct {
	db *sql.DB
}

func NewDataBaseUsers(db *sql.DB) *DataBaseUsers {
	return &DataBaseUsers{
		db: db,
	}
}

func (dbUsers *DataBaseUsers) SelectUserByNickName(name string) (*usersRrep.User, error) {
	row, err := dbUsers.db.Query("SELECT * FROM users WHERE nickname=" + name)
	if err != nil {
		return nil, err
	}

	user := usersRrep.User{}

	//row.Scan(&user.Nickname, user.Fullname, user.About, user.Email) todo

	return &user, nil
}

func (dbUsers *DataBaseUsers) SelectUserByEmail(email string) (*usersRrep.User, error) {
	row, err := dbUsers.db.Query("SELECT * FROM users WHERE email=" + email)
	if err != nil {
		return nil, err
  	}
	
	user := usersRrep.User{}

	//row.Scan(&user.Nickname, user.Fullname, user.About, user.Email) todo

	return &user, nil
}

func (dbUsers *DataBaseUsers) CreateUser(u usersRrep.User) (*usersRrep.User, error) {
	return nil, nil
}
