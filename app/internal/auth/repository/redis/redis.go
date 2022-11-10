package redis

import (
	authRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/auth/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

type dataBase struct {
	db *redis.Client
}

func New(redisClient *redis.Client) authRep.RepositoryI {
	return &dataBase {
		db: redisClient,
	}
}

func (dbCookies *dataBase) CreateCookie(cookie *models.Cookie) error {
	err := dbCookies.db.Do("SETEX", cookie.SessionToken, cookie.MaxAge, cookie.UserId).Err()
	if err != nil {
		return errors.Wrap(err, "redis error")
	}
	return nil
}

func (dbCookies *dataBase) GetCookie(value string) (string, error) {
	userIdStr, err := dbCookies.db.Get(value).Result()
	if errors.Is(err, redis.Nil) {
		return "", models.ErrNotFound
	} else if err != nil {
		return "", errors.Wrap(err, "redis error")
	}

	return userIdStr, nil
}

func (dbCookies *dataBase) DeleteCookie(value string) error {
	err := dbCookies.db.Del(value).Err()
	if err != nil {
		return errors.Wrap(err, "redis error")
	}

	return nil
}

