package usecase

import (
	"github.com/pkg/errors"

	userRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
)

type UseCaseI interface {
	SelectUserById(id int) (*models.User, error)
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
		return nil, errors.Wrap(err, "user repository error")
	}
	user.Password = ""

	return user, nil
}

