package postgres

import (
	"errors"

	"gorm.io/gorm"

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

func (dbCookies *dataBase) CreateCookie(c models.Cookie) (*models.Cookie, error) {
	cookie := models.Cookie{}

	tx := dbCookies.db.Table("cookies").Raw("INSERT INTO cookies VALUES (?, ?, ?) RETURNING *",
		c.SessionToken, c.UserId, c.Expires).Scan(&cookie)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &cookie, nil
}

func (dbCookies *dataBase) SelectCookie(value string) (*models.Cookie, error) {
	cookie := models.Cookie{}

	tx := dbCookies.db.Table("cookies").Where("value = ?", value).Scan(&cookie)
	if tx.Error != nil {
		return nil, tx.Error
	} else if cookie.SessionToken == "" {
		return nil, errors.New("cookie doesn't exist")
	}

	return &cookie, nil
}

func (dbCookies *dataBase) DeleteCookie(value string) error {
	tx := dbCookies.db.Table("cookies").Exec("DELETE FROM cookies WHERE value = ?", value)
	return tx.Error
}
