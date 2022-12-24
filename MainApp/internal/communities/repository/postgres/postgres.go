package postgres

import (
	"time"

	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/communities/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Community struct {
	ID          uint64
	OwnerID     uint64
	AvatarID    uint64 `gorm:"column:avatar_att_id"`
	Name        string
	Description string
	CreateDate  time.Time `gorm:"column:created_at"`
}

type CommunityUserRelation struct {
	CommunityID uint64 `gorm:"column:community_id"`
	UserID      uint64 `gorm:"column:user_id"`
}

func (CommunityUserRelation) TableName() string {
	return "communities_users"
}

func toPostgresCommunity(c *models.Community) *Community {
	return &Community{
		ID:          c.ID,
		OwnerID:     c.OwnerID,
		AvatarID:    c.AvatarID,
		Name:        c.Name,
		Description: c.Description,
		CreateDate:  c.CreateDate,
	}
}

func toModelCommunity(c *Community) *models.Community {
	return &models.Community{
		ID:          c.ID,
		OwnerID:     c.OwnerID,
		AvatarID:    c.AvatarID,
		Name:        c.Name,
		Description: c.Description,
		CreateDate:  c.CreateDate,
	}
}

func toModelCommunities(communities []*Community) []*models.Community {
	out := make([]*models.Community, len(communities))

	for i, b := range communities {
		out[i] = toModelCommunity(b)
	}

	return out
}

func (Community) TableName() string {
	return "communities"
}

type communitiesRepository struct {
	db *gorm.DB
}

func (dbcomm *communitiesRepository) CheckSubscriptionCommunity(id uint64, userID uint64) (bool, error) {
	var count int64
	tx := dbcomm.db.Model(&CommunityUserRelation{}).Where(&CommunityUserRelation{UserID: userID, CommunityID: id}).Count(&count)

	if tx.Error != nil {
		return false, errors.Wrap(tx.Error, "database error (table communities_users) on check")
	}

	return count > 0, nil
}

func (dbcomm *communitiesRepository) GetCommunity(id uint64) (*models.Community, error) {
	var comm Community

	tx := dbcomm.db.Where("id = ?", id).Take(&comm)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, models.ErrNotFound
	} else if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table communities)")
	}

	return toModelCommunity(&comm), nil
}

func (dbcomm *communitiesRepository) SearchCommunities(searchString string) ([]*models.Community, error) {
	comms := make([]*Community, 0, 10)

	tx := dbcomm.db.Where("lower(name) LIKE lower(?)", "%"+searchString+"%").Find(&comms)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, models.ErrNotFound
	} else if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table communities)")
	}

	return toModelCommunities(comms), nil
}

func (dbcomm *communitiesRepository) UpdateCommunity(comm *models.Community) error {
	postgresComm := toPostgresCommunity(comm)

	tx := dbcomm.db.Omit("id").Updates(postgresComm)

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table communities)")
	}

	return nil
}

func (dbcomm *communitiesRepository) CreateCommunity(comm *models.Community) error {
	postgresComm := toPostgresCommunity(comm)

	tx := dbcomm.db.Create(postgresComm)

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table communities)")
	}

	comm.ID = postgresComm.ID
	comm.CreateDate = time.Now()
	return nil
}

func (dbcomm *communitiesRepository) DeleteCommunity(id uint64) error {
	tx := dbcomm.db.Delete(&Community{}, id)

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table communities)")
	}

	return nil
}

func (dbcomm *communitiesRepository) JoinCommunity(id uint64, userId uint64) error {
	tx := dbcomm.db.Create(&CommunityUserRelation{CommunityID: id, UserID: userId})

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table communities_users) on create")
	}

	return nil
}

func (dbcomm *communitiesRepository) LeaveCommunity(id uint64, userId uint64) error {
	tx := dbcomm.db.Where(&CommunityUserRelation{CommunityID: id, UserID: userId}).Delete(&CommunityUserRelation{})

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "\"database error (table communities_users) on delete")
	}

	return nil
}

func (dbcomm *communitiesRepository) GetCountUserCommunity(id uint64) (uint64, error) {
	var count int64
	tx := dbcomm.db.Model(&CommunityUserRelation{}).Where("community_id = ?", id).Count(&count)

	if tx.Error != nil {
		return 0, errors.Wrap(tx.Error, "database error (table communities_users) on count")
	}

	return uint64(count), nil
}

func (dbcomm *communitiesRepository) GetAllCommunities() ([]*models.Community, error) {
	communities := make([]*Community, 0, 10)
	tx := dbcomm.db.Find(&communities)

	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table communities) on GetAllCommunities")
	}

	return toModelCommunities(communities), nil
}

func (dbcomm *communitiesRepository) GetAllUserCommunities(userID uint64) ([]*models.Community, error) {
	communitiesUsersRel := make([]*CommunityUserRelation, 0, 10)
	tx := dbcomm.db.Where(&CommunityUserRelation{UserID: userID}).Find(&communitiesUsersRel)

	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table communities) on GetAllCommunities")
	}

	communities := make([]*Community, 0, 10)
	for idx := range communitiesUsersRel {
		var comm Community

		tx := dbcomm.db.Where("id = ?", communitiesUsersRel[idx].CommunityID).Take(&comm)
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, models.ErrNotFound
		} else if tx.Error != nil {
			return nil, errors.Wrap(tx.Error, "database error (table communities)")
		}
		communities = append(communities, &comm)
	}

	return toModelCommunities(communities), nil
}

func NewCommunitiesRepository(db *gorm.DB) repository.RepositoryI {
	return &communitiesRepository{
		db: db,
	}
}
