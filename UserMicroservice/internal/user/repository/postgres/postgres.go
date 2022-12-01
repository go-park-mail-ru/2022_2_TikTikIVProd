package postgres

import (
	"gorm.io/gorm"
	"github.com/pkg/errors"

	userRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/UserMicroservice/internal/user/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/UserMicroservice/models"
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

func (dbUsers *dataBase) AddFriend(friends models.Friends) error {
	tx := dbUsers.db.Create(friends)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table friends)")
	}

	return nil
}

func (dbUsers *dataBase) DeleteFriend(friends models.Friends) error {
	tx := dbUsers.db.Where("user_id1 = ? AND user_id2 = ?", friends.Id1, friends.Id2).Delete(&models.Friends{})
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table friends)")
	}

	return nil
}

func (dbUsers *dataBase) CheckFriends(friends models.Friends) (bool, error) {
	var count int64
	tx := dbUsers.db.Model(&models.Friends{}).Where("user_id1 = ? AND user_id2 = ?",
									friends.Id1, friends.Id2).Count(&count)
	if tx.Error != nil {
		return false, errors.Wrap(tx.Error, "database error (table friends)")
	}
	return count > 0, nil
}

func (dbUsers *dataBase) SelectFriends(id uint64) ([]models.User, error) {
	friends := make([]models.User, 0, 10)
	tx := dbUsers.db.Omit("password").Joins("JOIN friends ON friends.user_id2 = users.id").Find(&friends, 
				"user_id1 = ?", id)

	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (tables friends, users)")
	}

	return friends, nil
}



