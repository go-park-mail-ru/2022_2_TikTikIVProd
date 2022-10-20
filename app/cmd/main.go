package main

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/cmd/router"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/cmd/server"
	imagesRepository "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/image/repository/postgres"
	postsDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/delivery"
	postsUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/usecase"
	postsRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/repository/postgres"
	usersDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/delivery"
	usersRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/repository/postgres"
	usersUseCase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/usecase"
	authDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/auth/delivery"
	authRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/auth/repository/postgres"
	authUseCase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/auth/usecase"
	friendsDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/friends/delivery"
	friendsRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/friends/repository/postgres"
	friendsUseCase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/friends/usecase"
)

// @title WS Swagger API
// @version 1.0
// @host 89.208.197.127:8080

func main() {
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: "host=ws_pg user=postgres password=postgres port=5432"}),
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

	postsDeliver := postsDelivery.NewDelivery(postsUC)
	usersDeliver := usersDelivery.New(usersUC)
	authDeliver := authDelivery.New(authUC)
	friendsDeliver := friendsDelivery.New(friendsUC)

	e := router.NewEchoRouter(usersDeliver, friendsDeliver, authDeliver, postsDeliver)

	s := server.NewServer(e)
	if err := s.Start(); err != nil {
		log.Fatal("ERROR", err)
	}
}
