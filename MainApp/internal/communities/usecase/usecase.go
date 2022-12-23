package usecase

import (
	communitiesRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/communities/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	"github.com/pkg/errors"
)

type UseCaseI interface {
	CreateCommunity(p *models.Community) error
	GetCommunity(id uint64) (*models.Community, error)
	SearchCommunities(searchString string) ([]*models.Community, error)
	UpdateCommunity(p *models.Community) error
	DeleteCommunity(id uint64, userId uint64) error
	LeaveCommunity(id uint64, userId uint64) error
	JoinCommunity(id uint64, userId uint64) error
	GetAllCommunities() ([]*models.Community, error)
}

type useCase struct {
	communitiesRep communitiesRep.RepositoryI
}

func (u useCase) CreateCommunity(comm *models.Community) error {
	err := u.communitiesRep.CreateCommunity(comm)

	if err != nil {
		return errors.Wrap(err, "community repository error")
	}

	return nil
}

func (u useCase) GetCommunity(id uint64) (*models.Community, error) {
	community, err := u.communitiesRep.GetCommunity(id)

	err = addCountUsersCommunity(community, u.communitiesRep)

	if err != nil {
		return nil, errors.Wrap(err, "community repository error while add count subs")
	}

	if err != nil {
		return nil, errors.Wrap(err, "community repository error")
	}

	return community, nil
}

func (u useCase) SearchCommunities(searchString string) ([]*models.Community, error) {
	communities, err := u.communitiesRep.SearchCommunities(searchString)

	if err != nil {
		return nil, errors.Wrap(err, "community repository error")
	}

	return communities, nil
}

func (u useCase) UpdateCommunity(comm *models.Community) error {
	existedCommunity, err := u.communitiesRep.GetCommunity(comm.ID)
	if err != nil {
		return errors.Wrap(err, "community repository error")
	}

	if existedCommunity == nil {
		return errors.New("Community not found")
	}

	if existedCommunity.OwnerID != comm.OwnerID {
		return models.ErrPermissionDenied
	}

	err = u.communitiesRep.UpdateCommunity(comm)

	if err != nil {
		return errors.Wrap(err, "community repository error")
	}

	return nil
}

func (u useCase) DeleteCommunity(id uint64, userId uint64) error {
	existedCommunity, err := u.communitiesRep.GetCommunity(id)
	if err != nil {
		return errors.Wrap(err, "community repository error")
	}

	if existedCommunity == nil {
		return errors.New("Community not found")
	}

	if existedCommunity.OwnerID != userId {
		return models.ErrPermissionDenied
	}

	err = u.communitiesRep.DeleteCommunity(id)

	if err != nil {
		return errors.Wrap(err, "community repository error")
	}

	return nil
}

func (u useCase) JoinCommunity(id uint64, userId uint64) error {
	existedCommunity, err := u.communitiesRep.GetCommunity(id)
	if err != nil {
		return errors.Wrap(err, "community repository error")
	}

	if existedCommunity == nil {
		return errors.New("Community not found")
	}

	err = u.communitiesRep.JoinCommunity(id, userId)

	if err != nil {
		return errors.Wrap(err, "community repository error")
	}

	return nil
}

func (u useCase) LeaveCommunity(id uint64, userId uint64) error {
	existedCommunity, err := u.communitiesRep.GetCommunity(id)
	if err != nil {
		return errors.Wrap(err, "community repository error")
	}

	if existedCommunity == nil {
		return errors.New("Community not found")
	}

	err = u.communitiesRep.LeaveCommunity(id, userId)

	if err != nil {
		return errors.Wrap(err, "community repository error")
	}

	return nil
}

func addCountUsersCommunity(comm *models.Community, commRepo communitiesRep.RepositoryI) error {
	count, err := commRepo.GetCountUserCommunity(comm.ID)

	if err != nil {
		return errors.Wrap(err, "Post repository error in func addPostAttachmentsAuthors")
	}

	comm.CountSubs = count

	return nil
}

func (u *useCase) GetAllCommunities() ([]*models.Community, error) {
	communities, err := u.communitiesRep.GetAllCommunities()

	for idx := range communities {
		err = addCountUsersCommunity(communities[idx], u.communitiesRep)

		if err != nil {
			return nil, errors.Wrap(err, "community repository error while add count subs")
		}
	}

	if err != nil {
		return nil, errors.Wrap(err, "community repository error")
	}

	return communities, nil
}

func New(rep communitiesRep.RepositoryI) UseCaseI {
	return &useCase{
		communitiesRep: rep,
	}
}
