package postgres

import (
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/communities/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type Community struct {
	ID          uint64
	OwnerID     uint64
	AvatarID    uint64 `gorm:"column:avatar_img_id"`
	Name        string
	Description string
	CreateDate  time.Time `gorm:"column:created_at"`
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

	tx := dbcomm.db.Where("name LIKE ?", "%"+searchString+"%").Find(&comms)
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

func (dbcomm *communitiesRepository) GetAllCommunities() ([]*models.Community, error) {
	communities := make([]*Community, 0, 10)
	tx := dbcomm.db.Find(&communities)

	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table communities) on GetAllCommunities")
	}

	return toModelCommunities(communities), nil
}

func NewCommunitiesRepository(db *gorm.DB) repository.RepositoryI {
	return &communitiesRepository{
		db: db,
	}
}
