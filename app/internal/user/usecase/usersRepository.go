package usersRep

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UsersStore interface {
	SelectUser(name string) (*User, error)
	CreateUser(u User) (*User, error)
}

type UsersRep struct {
	usersStore UsersStore
}

type User struct {
	Id       int    `json:"id"`
	LastName string `json:"last_name"`
	NickName string `json:"nick_name"`
	Email    string `json:"email"`
	Avatar   int    `json:"avatar"`
	Password string `json:"password"`
}

func NewUsersRep(us UsersStore) *UsersRep {
	return &UsersRep{
		usersStore: us,
	}
}

func (ur *UsersRep) SelectUserByNickName(nickname string) (*User, error) {
	user, err := ur.usersStore.SelectUser(nickname)
	if err != nil {
		return nil, errors.New("can't find user with nickname " + nickname)
	}

	return user, nil
}

func (ur *UsersRep) SelectUserByEmail(email string) (*User, error) {
	user, err := ur.usersStore.SelectUser(email)
	if err != nil {
		return nil, errors.New("can't find user with email " + email)
	}

	return user, nil
}

func (ur *UsersRep) SignIn(user User) (*User, error) {
	u, err := ur.SelectUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password)); err != nil {
		return nil, errors.New("incorrect password")
	}

	return u, nil
}

func (ur *UsersRep) CreateUser(user User) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		return nil, errors.New("hash error")
	}

	user.Password = string(hashedPassword)
	
	if _, err := ur.usersStore.SelectUser(user.NickName); err != nil {
		return nil, errors.New("user with nickname " + user.NickName + "already exists.")
	}

	newUser, err := ur.usersStore.CreateUser(user)
	if err != nil {
		return nil, errors.New("create user error")
	}

	return newUser, nil
}
