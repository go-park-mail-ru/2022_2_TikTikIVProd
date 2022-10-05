package main

import (
	imagesRepository "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/image/repository"
	postsDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/delivery"
	postsRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/repository"
	postsUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/usecase"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/server"
	_ "gorm.io/gorm"
)

func main() {
	// инициализаци бд
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: "host=localhost user=postgres password=postgres port=13080"}),
		&gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// создание хранилищ данных (называется репозиториями): юзеры, посты...
	//dbUsers := usersUsecase.NewDataBaseUsers(db)
	//userRep := usersUsecase.NewUsersRep(dbUsers)

	dbPosts := postsRep.NewDataBasePosts(db)
	dbImages := imagesRepository.NewDataBaseImages(db)
	postsUsecase := postsUsecase.NewPostsUsecase(dbPosts, dbImages)
	postsDeliv := postsDelivery.NewDelivery(postsUsecase)
	r := mux.NewRouter()
	r.HandleFunc("/feed", postsDeliv.Feed)
	s := server.NewServer(r)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
