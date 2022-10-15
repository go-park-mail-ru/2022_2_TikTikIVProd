package usecase

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	userRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
)

//go:generate mockgen -source=usecase.go -destination=mocks/mock.go

type UseCaseI interface {
	SelectUserByNickName(nickname string) (*models.User, error)
	SelectUserByEmail(email string) (*models.User, error)
	SelectUserById(id int) (*models.User, error)
	SignIn(user models.UserSignIn) (*models.User, *models.Cookie, error)
	SignUp(user models.User) (*models.User, *models.Cookie, error)
	CreateUser(user models.User) (*models.User, error)
	CreateCookie(userId int) (*models.Cookie, error)
	SelectCookie(value string) (*models.Cookie, error)
	DeleteCookie(value string) error
}

type useCase struct {
	repository userRep.RepositoryI
}

func New(rep userRep.RepositoryI) UseCaseI {
	return &useCase{
		repository: rep,
	}
}

func (uc *useCase) CreateCookie(userId int) (*models.Cookie, error) {
	cookie := models.Cookie{
		UserId:       userId,
		SessionToken: uuid.NewString(),
		Expires:      time.Now().AddDate(1, 0, 0)}

	newCookie, err := uc.repository.CreateCookie(cookie)
	if err != nil {
		return nil, errors.New("create cookie error")
	}

	return newCookie, nil
}

func (uc *useCase) SelectCookie(value string) (*models.Cookie, error) {
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

func (uc *useCase) SignIn(user models.UserSignIn) (*models.User, *models.Cookie, error) {
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

func (uc *useCase) SignUp(user models.User) (*models.User, *models.Cookie, error) {
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
