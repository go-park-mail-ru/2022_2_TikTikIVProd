package main

import (
	"github.com/labstack/echo/v4"
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

var testCfg = postgres.Config{DSN: "host=localhost user=postgres password=postgres port=13080"}

//var prod_cfg = postgres.Config{DSN: "host=ws_pg user=postgres password=postgres port=5432"}

func main() {
	db, err := gorm.Open(postgres.New(testCfg),
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
	friendsUC := friendsUseCase.New(friendsDB)
	imageUC := imageUsecase.NewImageUsecase(imageDB)

	e := echo.New()

	_postsDelivery.NewDelivery(e, postsUC)
	_usersDelivery.NewDelivery(e, usersUC)
	_imageDelivery.NewDelivery(e, imageUC)
	_authDelivery.NewDelivery(e, authUC)
	_friendsDelivery.NewDelivery(e, friendsUC)

	s := server.NewServer(e)
	if err := s.Start(); err != nil {
		log.Fatal("ERROR", err)
	}
}
