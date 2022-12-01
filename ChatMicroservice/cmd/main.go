package main

import (
	chatDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/ChatMicroservice/internal/chat/delivery"
	chatRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/ChatMicroservice/internal/chat/repository/postgres"
	chatUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/ChatMicroservice/internal/chat/usecase"
	chat "github.com/go-park-mail-ru/2022_2_TikTikIVProd/ChatMicroservice/proto"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
)

// var testCfgPg = postgres.Config{DSN: "host=localhost user=postgres password=postgres port=13080"}

var prodCfgPg = postgres.Config{DSN: "host=ws_pg user=postgres password=postgres port=5432"}

func main() {
	lis, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(postgres.New(prodCfgPg),
		&gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	chatDB := chatRep.NewChatRepository(db)
	chatUC := chatUsecase.New(chatDB)
	chat.RegisterChatServer(server, chatDelivery.New(chatUC))

	log.Println("starting server at :8083")
	server.Serve(lis)
}
