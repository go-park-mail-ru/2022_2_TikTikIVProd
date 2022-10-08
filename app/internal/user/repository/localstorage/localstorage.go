package localstorage

import (
	"errors"
	"strconv"
	"sync"

	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/model"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/repository"
)

type UserStorage struct {
	users  []model.User
	cookie []model.Cookie
	mtx    sync.RWMutex
}

// type dataBase struct {
// 	db *gorm.DB
// }

func New() repository.RepositoryI {
	return &UserStorage{
		users:  make([]model.User, 0),
		cookie: make([]model.Cookie, 0),
	}
}

// var MyUserStorage repository.RepositoryI = &UserStorage{
// 	users: make([]model.User, 0),
// }

func (usersStorage *UserStorage) SelectUserById(id int) (*model.User, error) {
	usersStorage.mtx.RLock()
	defer usersStorage.mtx.RUnlock()
	for _, user := range usersStorage.users {
		if user.Id == id {
			return &user, nil
		}
	}

	return nil, errors.New("can't find user with id " + strconv.Itoa(id))
}

func (usersStorage *UserStorage) SelectUserByNickName(nickname string) (*model.User, error) {
	usersStorage.mtx.RLock()
	defer usersStorage.mtx.RUnlock()
	for _, user := range usersStorage.users {
		if user.NickName == nickname {
			return &user, nil
		}
	}

	return nil, errors.New("can't find user with nickname " + nickname)
}

func (usersStorage *UserStorage) SelectUserByEmail(email string) (*model.User, error) {
	usersStorage.mtx.RLock()
	defer usersStorage.mtx.RUnlock()
	for _, user := range usersStorage.users {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, errors.New("can't find user with email " + email)
}

func (usersStorage *UserStorage) CreateUser(u model.User) (*model.User, error) {
	usersStorage.mtx.RLock()
	if _, err := usersStorage.SelectUserByNickName(u.NickName); err == nil {
		usersStorage.mtx.RUnlock()
		return nil, errors.New("nickname " + u.NickName + "already in use.")
	}
	if _, err := usersStorage.SelectUserByEmail(u.Email); err == nil {
		usersStorage.mtx.RUnlock()
		return nil, errors.New("user with email " + u.Email + "already exists.")
	}
	usersStorage.mtx.RUnlock()

	u.Id = int(len(usersStorage.users)) + 1

	usersStorage.mtx.Lock()
	usersStorage.users = append(usersStorage.users, u)
	usersStorage.mtx.Unlock()
	return &usersStorage.users[u.Id-1], nil
}

func (usersStorage *UserStorage) CreateCookie(c model.Cookie) (*model.Cookie, error) {
	usersStorage.mtx.Lock()
	usersStorage.cookie = append(usersStorage.cookie, c)
	usersStorage.mtx.Unlock()
	return &usersStorage.cookie[len(usersStorage.cookie)-1], nil
}

func (usersStorage *UserStorage) SelectCookie(value string) (*model.Cookie, error) {
	usersStorage.mtx.RLock()
	defer usersStorage.mtx.RUnlock()
	for _, cookie := range usersStorage.cookie {
		if cookie.SessionToken == value {
			return &cookie, nil
		}
	}

	return nil, errors.New("cookie doesn't exist")
}

func (usersStorage *UserStorage) DeleteCookie(value string) error {
	usersStorage.mtx.Lock()
	defer usersStorage.mtx.Unlock()
	for i := 0; i < len(usersStorage.cookie); i++ {
		if usersStorage.cookie[i].SessionToken == value {
			usersStorage.cookie = append(usersStorage.cookie[:i], usersStorage.cookie[i+1:]...)
			return nil
		}
	}

	return errors.New("cookie doesn't exist")
}
