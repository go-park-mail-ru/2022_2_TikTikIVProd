package main

import (
	"log"

	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/router"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/server"
	usersStore "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/store"
	usersRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/repository"
)



func main() {
	// инициализаци бд

	db, err := newdb() //TODO

	// создание хранилищ данных (называется репозиториями): юзеры, посты...
	dbUsers := usersStore.NewDataBaseUsers(db)
	userRep := usersRep.NewUsersRep(dbUsers)

	r := router.NewRouter(userRep)

	s := server.NewServer(r)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}

