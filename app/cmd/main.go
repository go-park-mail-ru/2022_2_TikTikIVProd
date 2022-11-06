package main

import (
	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/cmd/server"
	_authDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/auth/delivery"
	authRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/auth/repository/redis"
	authUseCase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/auth/usecase"
	_friendsDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/friends/delivery"
	friendsRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/friends/repository/postgres"
	friendsUseCase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/friends/usecase"
	_imageDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/image/delivery"
	imagesRepository "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/image/repository/postgres"
	imageUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/image/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/middleware"
	_postsDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/delivery"
	postsRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/repository/postgres"
	postsUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/usecase"
	_usersDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/delivery"
	usersRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/repository/postgres"
	usersUseCase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/usecase"
)

// @title WS Swagger API
// @version 1.0
// @host 89.208.197.127:8080

//var testCfg = postgres.Config{DSN: "host=localhost user=postgres password=postgres port=13080"}

var prodCfg = postgres.Config{DSN: "host=ws_pg user=postgres password=postgres port=5432"}

func main() {
	db, err := gorm.Open(postgres.New(prodCfg),
		&gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:       "redis:6379",
		MaxRetries: 10,
	})

	err = redisClient.Ping().Err()
	if err != nil {
		log.Fatal(err)
		return
	}

	postDB := postsRep.NewPostRepository(db)
	authDB := authRep.New(redisClient)
	usersDB := usersRep.New(db)
	friendsDB := friendsRep.New(db)
	imageDB := imagesRepository.NewImageRepository(db)

	postsUC := postsUsecase.NewPostUsecase(postDB, imageDB, usersDB)
	authUC := authUseCase.New(authDB, usersDB)
	usersUC := usersUseCase.New(usersDB)
	friendsUC := friendsUseCase.New(friendsDB, usersDB)
	imageUC := imageUsecase.NewImageUsecase(imageDB)

	e := echo.New()

	e.Logger.SetHeader(`time=${time_rfc3339} level=${level} prefix=${prefix} ` +
		`file=${short_file} line=${line} message:`)
	e.Logger.SetLevel(log.INFO)

	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins:     []string{"http://89.208.197.127"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	e.Use(echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
		Format: `time=${time_custom} remote_ip=${remote_ip} ` +
			`host=${host} method=${method} uri=${uri} user_agent=${user_agent} ` +
			`status=${status} error="${error}" ` +
			`bytes_in=${bytes_in} bytes_out=${bytes_out}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))

	e.Use(echoMiddleware.Recover())

	authMiddleware := middleware.NewMiddleware(authUC)
	e.Use(authMiddleware.Auth)
	//e.Use(authMiddleware.CSRF)

	_postsDelivery.NewDelivery(e, postsUC)
	_authDelivery.NewDelivery(e, authUC)
	_usersDelivery.NewDelivery(e, usersUC)
	_imageDelivery.NewDelivery(e, imageUC)
	_friendsDelivery.NewDelivery(e, friendsUC)

	s := server.NewServer(e)
	if err := s.Start(); err != nil {
		e.Logger.Fatal(err)
	}
}
