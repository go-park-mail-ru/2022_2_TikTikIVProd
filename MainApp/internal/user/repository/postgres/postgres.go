package postgres

import (
	"gorm.io/gorm"
	"github.com/pkg/errors"

	userRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/user/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
)

type dataBase struct {
	db *gorm.DB
}

func New(db *gorm.DB) userRep.RepositoryI {
	return &dataBase{
		db: db,
	}
}

func (dbUsers *dataBase) SelectUserById(id uint64) (*models.User, error) {
	user := models.User{}

	tx := dbUsers.db.Where("id = ?", id).Take(&user)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, models.ErrNotFound
	} else if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table users)")
	}

	return &user, nil
}

func (dbUsers *dataBase) SelectUserByNickName(nickname string) (*models.User, error) {
	user := models.User{}

	tx := dbUsers.db.Where("nick_name = ?", nickname).Take(&user)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, models.ErrNotFound
	} else if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table users)")
	}

	return &user, nil
}

func (dbUsers *dataBase) SelectUserByEmail(email string) (*models.User, error) {
	user := models.User{}
	tx := dbUsers.db.Where("email = ?", email).Take(&user)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, models.ErrNotFound
	} else if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table users)")
	}

	return &user, nil
}

func (dbUsers *dataBase) CreateUser(user *models.User) error {
	tx := dbUsers.db.Omit("avatar_img_id").Create(user)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table users)")
	}

	return nil
}

func (dbUsers *dataBase) UpdateUser(user models.User) error {
	tx := dbUsers.db.Omit("id").Updates(user)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table users)")
	}

	return nil
}

func (dbUsers *dataBase) SelectAllUsers() ([]models.User, error) {
	users := make([]models.User, 0, 10)
	tx := dbUsers.db.Omit("password").Find(&users)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table users)")
	}

	return users, nil
}

func (dbUsers *dataBase) SearchUsers(name string) ([]models.User, error) {
	users := make([]models.User, 0, 10)

	tx := dbUsers.db.Omit("password").Find(&users,
		"lower(first_name || ' ' || last_name) LIKE lower(?) OR lower(last_name || ' ' || first_name) LIKE lower(?)",
		name + "%", name + "%")
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table users)")
	}

	return users, nil
}

