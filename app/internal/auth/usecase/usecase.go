package usecase

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	userRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/repository"
	authRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/auth/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
)

//go:generate mockgen -source=usecase.go -destination=mocks/mock.go

type UseCaseI interface {
	Auth(cookie string) (*models.User, error)
	SignIn(user models.UserSignIn) (*models.User, *models.Cookie, error)
	SignUp(user models.User) (*models.User, *models.Cookie, error)
	CreateCookie(userId int) (*models.Cookie, error)
	SelectCookie(value string) (*models.Cookie, error)
	DeleteCookie(value string) error
}

type useCase struct {
	authRepository authRep.RepositoryI
	userRepository userRep.RepositoryI
}

func New(authRepository authRep.RepositoryI, userRepository userRep.RepositoryI) UseCaseI {
	return &useCase{
		authRepository: authRepository,
		userRepository: userRepository,
	}
}

func (uc *useCase) CreateCookie(userId int) (*models.Cookie, error) {
	cookie := models.Cookie{
		UserId:       userId,
		SessionToken: uuid.NewString(),
		Expires:      time.Now().AddDate(1, 0, 0)}

	newCookie, err := uc.authRepository.CreateCookie(cookie)
	if err != nil {
		return nil, errors.New("create cookie error")
	}

	return newCookie, nil
}

func (uc *useCase) SelectCookie(value string) (*models.Cookie, error) {
	cookie, err := uc.authRepository.SelectCookie(value)
	if err != nil {
		return nil, errors.New("cookie doesn't exist")
	}

	return cookie, nil
}

func (uc *useCase) DeleteCookie(value string) error {
	err := uc.authRepository.DeleteCookie(value)
	if err != nil {
		return errors.New("cookie doesn't exist")
	}

	return nil
}

func (uc *useCase) SignIn(user models.UserSignIn) (*models.User, *models.Cookie, error) {
	u, err := uc.userRepository.SelectUserByEmail(user.Email)
	if err != nil {
		return nil, nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password)); err != nil {
		return nil, nil, errors.New("invalid password")
	}

	cookie := models.Cookie{
		UserId:       u.Id,
		SessionToken: uuid.NewString(),
		Expires:      time.Now().AddDate(1, 0, 0)}

	newCookie, err := uc.authRepository.CreateCookie(cookie)
	if err != nil {
		return nil, nil, err
	}

	return u, newCookie, nil
}

func (uc *useCase) SignUp(user models.User) (*models.User, *models.Cookie, error) {
	if _, err := uc.userRepository.SelectUserByNickName(user.NickName); err == nil {
		return nil, nil, errors.New("nickname already in use")
	}

	if _, err := uc.userRepository.SelectUserByEmail(user.Email); err == nil {
		return nil, nil, errors.New("user with such email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		return nil, nil, errors.New("hash error")
	}

	user.Password = string(hashedPassword)

	createdUser, err := uc.userRepository.CreateUser(user)
	if err != nil {
		return nil, nil, errors.New("create user error")
	}

	cookie := models.Cookie{
		UserId:       createdUser.Id,
		SessionToken: uuid.NewString(),
		Expires:      time.Now().AddDate(1, 0, 0)}

	newCookie, err := uc.authRepository.CreateCookie(cookie)
	if err != nil {
		return nil, nil, err
	}

	return createdUser, newCookie, nil
}

func (uc *useCase) Auth(cookie string) (*models.User, error) {
	gotCookie, err := uc.authRepository.SelectCookie(cookie)
	if err != nil {
		return nil, err
	}

	gotUser, err := uc.userRepository.SelectUserById(gotCookie.UserId)
	if err != nil {
		return nil, err
	}

	return gotUser, nil
}

