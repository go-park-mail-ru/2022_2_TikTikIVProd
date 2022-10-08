package usecase

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/model"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/repository"
)

//go:generate mockgen -source=usecase.go -destination=mocks/mock.go

type UseCaseI interface {
	SelectUserByNickName(nickname string) (*model.User, error)
	SelectUserByEmail(email string) (*model.User, error)
	SelectUserById(id int) (*model.User, error)
	SignIn(user model.UserSignIn) (*model.User, *model.Cookie, error)
	SignUp(user model.User) (*model.User, *model.Cookie, error)
	CreateUser(user model.User) (*model.User, error)
	CreateCookie(userId int) (*model.Cookie, error)
	SelectCookie(value string) (*model.Cookie, error)
	DeleteCookie(value string) error
}

type useCase struct {
	repository repository.RepositoryI
}

func New(rep repository.RepositoryI) UseCaseI {
	return &useCase{
		repository: rep,
	}
}

func (uc *useCase) CreateCookie(userId int) (*model.Cookie, error) {
	cookie := model.Cookie{
		UserId:       userId,
		SessionToken: uuid.NewString(),
		Expires:      time.Now().AddDate(1, 0, 0)}

	newCookie, err := uc.repository.CreateCookie(cookie)
	if err != nil {
		return nil, errors.New("create cookie error")
	}

	return newCookie, nil
}

func (uc *useCase) SelectCookie(value string) (*model.Cookie, error) {
	cookie, err := uc.repository.SelectCookie(value)
	if err != nil {
		return nil, errors.New("cookie doesn't exist")
	}

	return cookie, nil
}

func (uc *useCase) DeleteCookie(value string) error {
	err := uc.repository.DeleteCookie(value)
	if err != nil {
		return errors.New("cookie doesn't exist")
	}

	return nil
}

func (uc *useCase) SelectUserById(id int) (*model.User, error) {
	user, err := uc.repository.SelectUserById(id)
	if err != nil {
		return nil, errors.New("can't find user with such id")
	}

	return user, nil
}

func (uc *useCase) SelectUserByNickName(nickname string) (*model.User, error) {
	user, err := uc.repository.SelectUserByNickName(nickname)
	if err != nil {
		return nil, errors.New("can't find user with such nickname")
	}

	return user, nil
}

func (uc *useCase) SelectUserByEmail(email string) (*model.User, error) {
	user, err := uc.repository.SelectUserByEmail(email)
	if err != nil {
		return nil, errors.New("can't find user with such email")
	}

	return user, nil
}

func (uc *useCase) SignIn(user model.UserSignIn) (*model.User, *model.Cookie, error) {
	u, err := uc.SelectUserByEmail(user.Email)
	if err != nil {
		return nil, nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password)); err != nil {
		return nil, nil, errors.New("invalid password")
	}

	cookie, err := uc.CreateCookie(u.Id)
	if err != nil {
		return nil, nil, err
	}

	return u, cookie, nil
}

func (uc *useCase) SignUp(user model.User) (*model.User, *model.Cookie, error) {
	createdUser, err := uc.CreateUser(user)
	if err != nil {
		return nil, nil, err
	}

	cookie, err := uc.CreateCookie(createdUser.Id)
	if err != nil {
		return nil, nil, err
	}

	return createdUser, cookie, nil
}

func (uc *useCase) CreateUser(user model.User) (*model.User, error) {
	if _, err := uc.repository.SelectUserByNickName(user.NickName); err == nil {
		return nil, errors.New("nickname already in use")
	}

	if _, err := uc.repository.SelectUserByEmail(user.Email); err == nil {
		return nil, errors.New("user with such email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		return nil, errors.New("hash error")
	}

	user.Password = string(hashedPassword)

	newUser, err := uc.repository.CreateUser(user)
	if err != nil {
		return nil, errors.New("create user error")
	}

	return newUser, nil
}
