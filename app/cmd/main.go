package main

import (
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"log"

	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/cmd/server"
	_authDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/auth/delivery"
	authRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/auth/repository/postgres"
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
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/pkg/logger"
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

	postDB := postsRep.NewPostRepository(db)
	usersDB := usersRep.New(db)
	authDB := authRep.New(db)
	friendsDB := friendsRep.New(db)
	imageDB := imagesRepository.NewImageRepository(db)

	postsUC := postsUsecase.NewPostUsecase(postDB, imageDB, usersDB)
	usersUC := usersUseCase.New(usersDB)
	authUC := authUseCase.New(authDB, usersDB)
	friendsUC := friendsUseCase.New(friendsDB, usersDB)
	imageUC := imageUsecase.NewImageUsecase(imageDB)

	e := echo.New()
	authMiddleware := middleware.NewMiddleware(authUC)
	// e.Use(middleware.Auth)

	log := logger.New()
	e.Use(echoMiddleware.RequestLoggerWithConfig(echoMiddleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values echoMiddleware.RequestLoggerValues) error {
			log.Logrus.WithFields(logrus.Fields{
				"URI":    values.URI,
				"method": c.Request().Method,
				"status": values.Status,
			}).Info("request")
			return nil
		},
	}))

	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.Secure())
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	// e.Use(echoMiddleware.CSRF())

	_postsDelivery.NewDelivery(e, postsUC, authMiddleware)
	_usersDelivery.NewDelivery(e, usersUC, authMiddleware)
	_imageDelivery.NewDelivery(e, imageUC, authMiddleware)
	_authDelivery.NewDelivery(e, authUC, authMiddleware)
	_friendsDelivery.NewDelivery(e, friendsUC, authMiddleware)

	s := server.NewServer(e)
	if err := s.Start(); err != nil {
		e.Logger.Fatal(err)
	}
}
