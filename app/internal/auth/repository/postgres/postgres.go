package postgres

import (
	"gorm.io/gorm"
	"github.com/pkg/errors"

	authRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/auth/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
)

type dataBase struct {
	db *gorm.DB
}

func New(db *gorm.DB) authRep.RepositoryI {
	return &dataBase{
		db: db,
	}
}

func (dbCookies *dataBase) CreateCookie(cookie *models.Cookie) error {
	tx := dbCookies.db.Create(cookie)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table cookies)")
	}

	return nil
}

func (dbCookies *dataBase) SelectCookie(value string) (*models.Cookie, error) {
	cookie := models.Cookie{}

	tx := dbCookies.db.Where("value = ?", value).Take(&cookie)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, models.ErrNotFound
	} else if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table cookies)")
	}
	
	return &cookie, nil
}

func (dbCookies *dataBase) DeleteCookie(value string) error {
	tx := dbCookies.db.Delete(&models.Cookie{}, "value = ?", value)

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table cookies)")
	}

	return tx.Error
}

