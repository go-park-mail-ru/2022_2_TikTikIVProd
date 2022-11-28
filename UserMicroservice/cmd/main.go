package main

import (
	userDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/UserMicroservice/internal/user/delivery"
	userRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/UserMicroservice/internal/user/repository/postgres"
	userUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/UserMicroservice/internal/user/usecase"
	user "github.com/go-park-mail-ru/2022_2_TikTikIVProd/UserMicroservice/proto"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
)

//var testCfgPg = postgres.Config{DSN: "host=localhost user=postgres password=postgres port=13080"}

var prodCfgPg = postgres.Config{DSN: "host=ws_pg user=postgres password=postgres port=5432"}

func main() {
	lis, err := net.Listen("tcp", ":8084")
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(postgres.New(prodCfgPg),
		&gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	userDB := userRep.New(db)
	userUC := userUsecase.New(userDB)
	user.RegisterUsersServer(server, userDelivery.New(userUC))

	log.Println("starting server at :8084")
	server.Serve(lis)
}

