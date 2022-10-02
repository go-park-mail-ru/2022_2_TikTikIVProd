package usersRep

import (
	"errors"
)

type UsersStore interface {
	SelectUser(name string) (*User, error)
	CreateUser(u User) (*User, error)
}

type UsersRep struct {
	usersStore UsersStore
}

type User struct {
	LastName  string `json:"last_name"`
	NickName  string `json:"nick_name"`
	Email     string `json:"email"`
	Avatar    int    `json:"avatar"`
	Password  string `json:"password"`
}

func NewUsersRep(us UsersStore) *UsersRep {
	return &UsersRep{
		usersStore: us,
	}
}

func (ur *UsersRep) SelectUser(nickname string) (*User, error) {
	user, err := ur.usersStore.SelectUser(nickname)
	if err != nil {
		return nil, errors.New("can't find user with nickname " + nickname)
	}

	return user, nil
}

func (ur *UsersRep) CreateUser(u User) (*User, error) {
	if _, err := ur.usersStore.SelectUser(u.NickName); err != nil {
		return nil, errors.New("user with nickname " + u.NickName + "already exists.")
	}

	newUser, err := ur.usersStore.CreateUser(u)
	if err != nil {
		return nil, errors.New("create user error")
	}

	return newUser, nil
}
