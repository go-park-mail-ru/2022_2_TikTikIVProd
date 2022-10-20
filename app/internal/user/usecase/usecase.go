package usecase

import (
	"errors"
	"golang.org/x/crypto/bcrypt"

	userRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
)

//go:generate mockgen -source=usecase.go -destination=mocks/mock.go

type UseCaseI interface {
	SelectUserByNickName(nickname string) (*models.User, error)
	SelectUserByEmail(email string) (*models.User, error)
	SelectUserById(id int) (*models.User, error)
	CreateUser(user models.User) (*models.User, error)
}

type useCase struct {
	repository userRep.RepositoryI
}

func New(rep userRep.RepositoryI) UseCaseI {
	return &useCase{
		repository: rep,
	}
}

func (uc *useCase) SelectUserById(id int) (*models.User, error) {
	user, err := uc.repository.SelectUserById(id)
	if err != nil {
		return nil, errors.New("can't find user with such id")
	}

	return user, nil
}

func (uc *useCase) SelectUserByNickName(nickname string) (*models.User, error) {
	user, err := uc.repository.SelectUserByNickName(nickname)
	if err != nil {
		return nil, errors.New("can't find user with such nickname")
	}

	return user, nil
}

func (uc *useCase) SelectUserByEmail(email string) (*models.User, error) {
	user, err := uc.repository.SelectUserByEmail(email)
	if err != nil {
		return nil, errors.New("can't find user with such email")
	}

	return user, nil
}

func (uc *useCase) CreateUser(user models.User) (*models.User, error) {
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
