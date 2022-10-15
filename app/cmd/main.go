package main

import (
	postsRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/repository/postgres"
	"log"

	imagesRepository "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/image/repository/postgres"
	postsDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/delivery"
	postsUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/usecase"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/cmd/router"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/cmd/server"
	usersDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/delivery"
	usersPg "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/repository/postgres"
	usersUseCase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/usecase"
)

// @title WS Swagger API
// @version 1.0
// @host 89.208.197.127:8080

func main() {
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: "host=ws_pg user=postgres password=postgres port=8080"}),
		&gorm.Config{})

	if err != nil {
		log.Fatal(err)
		returngit
	}

	postDB := postsRep.NewPostRepository(db)
	imageDB := imagesRepository.NewImageRepository(db)
	postsUC := postsUsecase.NewPostUsecase(postDB, imageDB)
	postsDeliver := postsDelivery.NewDelivery(postsUC)

	usersDB := usersPg.New(db)
	usersUC := usersUseCase.New(usersDB)
	usersDeliver := usersDelivery.New(usersUC)

	r := router.NewRouter(usersDeliver, postsDeliver)

	s := server.NewServer(r)
	if err := s.Start(); err != nil {
		log.Fatal("ERROR", err)
	}
}
