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
	_chatDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/chat/delivery"
	chatRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/chat/repository/postgres"
	chatUseCase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/chat/usecase"
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

//var testCfgPg = postgres.Config{DSN: "host=localhost user=postgres password=postgres port=13080"}

var prodCfgPg = postgres.Config{DSN: "host=ws_pg user=postgres password=postgres port=5432"}

//var testCfgRedis = &redis.Options{Addr: ":6379"}

var prodCfgRedis = &redis.Options{Addr: "redis:6379"}

func main() {
	db, err := gorm.Open(postgres.New(prodCfgPg),
		&gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	redisClient := redis.NewClient(prodCfgRedis)

	err = redisClient.Ping().Err()
	if err != nil {
		log.Fatal(err)
	}

	postDB := postsRep.NewPostRepository(db)
	authDB := authRep.New(redisClient)
	usersDB := usersRep.New(db)
	friendsDB := friendsRep.New(db)
	imageDB := imagesRepository.NewImageRepository(db)
	chatDB := chatRep.NewChatRepository(db)

	postsUC := postsUsecase.NewPostUsecase(postDB, imageDB, usersDB)
	authUC := authUseCase.New(authDB, usersDB)
	usersUC := usersUseCase.New(usersDB)
	friendsUC := friendsUseCase.New(friendsDB, usersDB)
	imageUC := imageUsecase.NewImageUsecase(imageDB)
	chatUC := chatUseCase.New(chatDB)

	e := echo.New()

	e.Logger.SetHeader(`time=${time_rfc3339} level=${level} prefix=${prefix} ` +
		`file=${short_file} line=${line} message:`)
	e.Logger.SetLevel(log.INFO)

	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins:     []string{"http://89.208.197.127"},
		AllowHeaders:     []string{"Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"X-CSRF-Token"},
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
	e.Use(authMiddleware.CSRF)

	_postsDelivery.NewDelivery(e, postsUC)
	_authDelivery.NewDelivery(e, authUC)
	_usersDelivery.NewDelivery(e, usersUC)
	_imageDelivery.NewDelivery(e, imageUC)
	_friendsDelivery.NewDelivery(e, friendsUC)
	_chatDelivery.NewDelivery(e, chatUC)

	s := server.NewServer(e)
	if err := s.Start(); err != nil {
		e.Logger.Fatal(err)
	}
}
