package main

import (
	"log"

	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/router"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/server"
)



func main() {
	// инициализаци бд

	// создание хранилищ данных (называется репозиториями): юзеры, посты...

	h := router.NewRouter(/*сюда передадим данные какие то*/)

	s := server.NewServer(h)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}

