package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/cmd/router"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/cmd/server"
	authDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/auth/delivery"
	authRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/auth/repository/postgres"
	authUseCase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/auth/usecase"
	friendsDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/friends/delivery"
	friendsRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/friends/repository/postgres"
	friendsUseCase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/friends/usecase"
	imageDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/image/delivery"
	imagesRepository "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/image/repository/postgres"
	imageUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/image/usecase"
	postsDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/delivery"
	postsRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/repository/postgres"
	postsUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/usecase"
	usersDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/delivery"
	usersRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/repository/postgres"
	usersUseCase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/usecase"
)

// @title WS Swagger API
// @version 1.0
// @host 89.208.197.127:8080

//var testCfg = postgres.Config{DSN: "host=localhost user=postgres password=postgres port=13080"}

var prod_cfg = postgres.Config{DSN: "host=ws_pg user=postgres password=postgres port=5432"}

func main() {
	db, err := gorm.Open(postgres.New(prod_cfg),
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

	postsDeliver := postsDelivery.NewDelivery(postsUC)
	usersDeliver := usersDelivery.New(usersUC)
	authDeliver := authDelivery.New(authUC)
	friendsDeliver := friendsDelivery.New(friendsUC)
	imageDeliver := imageDelivery.NewDelivery(imageUC)

	e := router.NewEchoRouter(usersDeliver, friendsDeliver, authDeliver, postsDeliver, imageDeliver)

	s := server.NewServer(e)
	if err := s.Start(); err != nil {
		log.Fatal("ERROR", err)
	}
}
