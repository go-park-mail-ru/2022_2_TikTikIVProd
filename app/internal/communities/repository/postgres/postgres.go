package postgres

import (
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/communities/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type Community struct {
	ID          int
	OwnerID     int
	AvatarID    int `gorm:"column:avatar_img_id"`
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

func toModelPCommunity(c *Community) *models.Community {
	return &models.Community{
		ID:          c.ID,
		OwnerID:     c.OwnerID,
		AvatarID:    c.AvatarID,
		Name:        c.Name,
		Description: c.Description,
		CreateDate:  c.CreateDate,
	}
}

func (Community) TableName() string {
	return "communities"
}

type communitiesRepository struct {
	db *gorm.DB
}

func (dbcomm *communitiesRepository) GetCommunity(id int) (*models.Community, error) {
	var comm Community

	tx := dbcomm.db.Where("id = ?", id).Take(&comm)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, models.ErrNotFound
	} else if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table communities)")
	}

	return toModelPCommunity(&comm), nil
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
		return errors.Wrap(tx.Error, "communities repository error")
	}

	comm.ID = postgresComm.ID
	comm.CreateDate = time.Now()
	return nil
}

func (dbcomm *communitiesRepository) DeleteCommunity(id int) error {
	tx := dbcomm.db.Delete(&Community{}, id)

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table communities)")
	}

	return nil
}

func NewCommunitiesRepository(db *gorm.DB) repository.RepositoryI {
	return &communitiesRepository{
		db: db,
	}
}
