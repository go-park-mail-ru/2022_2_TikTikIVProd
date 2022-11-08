package usecase

import (
	"strconv"

	"github.com/pkg/errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	authRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/auth/repository"
	userRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
)

type UseCaseI interface {
	Auth(cookie string) (*models.User, error)
	SignIn(user models.UserSignIn) (*models.User, *models.Cookie, error)
	SignUp(user *models.User) (*models.Cookie, error)
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

func (uc *useCase) DeleteCookie(value string) error {
	_, err := uc.authRepository.GetCookie(value)
	if err != nil {
		return errors.Wrap(err, "auth repository error")
	}

	err = uc.authRepository.DeleteCookie(value)
	if err != nil {
		return errors.Wrap(err, "auth usecase error")
	}

	return nil
}

func (uc *useCase) SignIn(user models.UserSignIn) (*models.User, *models.Cookie, error) {
	u, err := uc.userRepository.SelectUserByEmail(user.Email)
	if err != nil {
		return nil, nil, errors.Wrap(err, "user repository error")
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return nil, nil, models.ErrInvalidPassword
	} else if err != nil {
		return nil, nil, errors.Wrap(err, "bcrypt error")
	}

	u.Password = ""

	cookie := models.Cookie{
		UserId:       u.Id,
		SessionToken: uuid.NewString(),
		MaxAge:      3600 * 24 * 365}

	err = uc.authRepository.CreateCookie(&cookie)
	if err != nil {
		return nil, nil, errors.Wrap(err, "auth repository error")
	}

	return u, &cookie, nil
}

func (uc *useCase) SignUp(user *models.User) (*models.Cookie, error) {
	_, err := uc.userRepository.SelectUserByNickName(user.NickName)
	if err != nil && !errors.Is(err, models.ErrNotFound) {
		return nil, errors.Wrap(err, "user repository error")
	} else if err == nil {
		return nil, models.ErrConflictNickname
	}

	_, err = uc.userRepository.SelectUserByEmail(user.Email)
	if err != nil && !errors.Is(err, models.ErrNotFound) {
		return nil, errors.Wrap(err, "user repository error")
	} else if err == nil {
		return nil, models.ErrConflictEmail
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		return nil, errors.Wrap(err, "bcrypt error")
	}

	user.Password = string(hashedPassword)

	err = uc.userRepository.CreateUser(user)
	if err != nil {
		return nil, errors.Wrap(err, "user repository error")
	}
	user.Password = ""

	cookie := models.Cookie{
		UserId:       user.Id,
		SessionToken: uuid.NewString(),
		MaxAge:      3600 * 24 * 365}

	err = uc.authRepository.CreateCookie(&cookie)
	if err != nil {
		return nil, errors.Wrap(err, "auth repository error")
	}

	return &cookie, nil
}

func (uc *useCase) Auth(cookie string) (*models.User, error) {
	userIdStr, err := uc.authRepository.GetCookie(cookie)
	if err != nil {
		return nil, errors.Wrap(err, "auth repository error")
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return nil, errors.Wrap(err, "auth repository error")
	}

	gotUser, err := uc.userRepository.SelectUserById(userId)
	if err != nil {
		return nil, errors.Wrap(err, "user repository error")
	}
	gotUser.Password = ""

	return gotUser, nil
}
