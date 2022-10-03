package main

import (
	imagesRepository "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/image/repository"
	postsRouter "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/delivery"
	postsRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/repository"
	postsUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/usecase"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/server"
	_ "gorm.io/gorm"
)

func main() {
	// инициализаци бд
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: "host=localhost user=postgres password=postgres port=8080"}),
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
	r := postsRouter.NewPostsRouter(postsUsecase)

	s := server.NewServer(r)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
