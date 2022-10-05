package main

import (
	"log"

	imagesRepository "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/image/repository"
	postsDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/delivery"
	postsRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/repository"
	postsUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/usecase"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/cmd/router"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/cmd/server"
	usersDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/delivery"
	usersRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/repository"
	usersUseCase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/usecase"
)

// @title WS Swagger API
// @version 1.0
// @host localhost:8080

func main() {
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: "host=ws_pg user=postgres password=postgres port=5432"}),
		&gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	dbPosts := postsRep.NewDataBasePosts(db)
	dbImages := imagesRepository.NewDataBaseImages(db)
	postsUsecase := postsUsecase.NewPostsUsecase(dbPosts, dbImages)
	postsDeliver := postsDelivery.NewDelivery(postsUsecase)

	usersDB := usersRep.New(db)
	usersUC := usersUseCase.New(usersDB)
	usersDeliver := usersDelivery.New(usersUC)

	r := router.NewRouter(usersDeliver, postsDeliver)

	s := server.NewServer(r)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
