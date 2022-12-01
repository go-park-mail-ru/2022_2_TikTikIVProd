package main

import (
	imageDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/ImageMicroservice/internal/image/delivery"
	imageRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/ImageMicroservice/internal/image/repository/postgres"
	imageUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/ImageMicroservice/internal/image/usecase"
	image "github.com/go-park-mail-ru/2022_2_TikTikIVProd/ImageMicroservice/proto"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
)

// var testCfgPg = postgres.Config{DSN: "host=localhost user=postgres password=postgres port=13080"}

var prodCfgPg = postgres.Config{DSN: "host=ws_pg user=postgres password=postgres port=5432"}

func main() {
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(postgres.New(prodCfgPg),
		&gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	imageDB := imageRep.NewImageRepository(db)
	imageUC := imageUsecase.New(imageDB)
	image.RegisterImagesServer(server, imageDelivery.New(imageUC))

	log.Println("starting server at :8082")
	server.Serve(lis)
}
