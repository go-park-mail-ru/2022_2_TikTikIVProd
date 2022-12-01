package usecase

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	userRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/user/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
)

type UseCaseI interface {
	SelectUserById(id uint64) (*models.User, error)
	UpdateUser(user models.User) error
	SelectUsers() ([]models.User, error)
	SearchUsers(name string) ([]models.User, error)
}

type useCase struct {
	userRepository userRep.RepositoryI
}

func New(rep userRep.RepositoryI) UseCaseI {
	return &useCase{
		userRepository: rep,
	}
}

func (uc *useCase) SelectUserById(id uint64) (*models.User, error) {
	user, err := uc.userRepository.SelectUserById(id)
	if err != nil {
		return nil, errors.Wrap(err, "user repository error")
	}
	user.Password = ""

	return user, nil
}

func (uc *useCase) UpdateUser(user models.User) error {
	_, err := uc.userRepository.SelectUserById(user.Id)
	if err != nil {
		return errors.Wrap(err, "user repository error")
	}

	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
		if err != nil {
			return errors.Wrap(err, "bcrypt error")
		}

		user.Password = string(hashedPassword)
	}

	err = uc.userRepository.UpdateUser(user)
	if err != nil {
		return errors.Wrap(err, "user repository error")
	}

	return nil
}

func (uc *useCase) SelectUsers() ([]models.User, error) {
	users, err := uc.userRepository.SelectAllUsers()
	if err != nil {
		return nil, errors.Wrap(err, "user repository error")
	}
	return users, nil
}

func (uc *useCase) SearchUsers(name string) ([]models.User, error) {
	users, err := uc.userRepository.SearchUsers(name)
	if err != nil {
		return nil, errors.Wrap(err, "user repository error")
	}
	return users, nil
}


