package main

import (
	"log"
	"net"

	attachmentDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/AttachmentMicroservice/internal/attachment/delivery"
	attachmentRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/AttachmentMicroservice/internal/attachment/repository/postgres"
	attachmentUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/AttachmentMicroservice/internal/attachment/usecase"
	attachment "github.com/go-park-mail-ru/2022_2_TikTikIVProd/AttachmentMicroservice/proto"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	attachmentDB := attachmentRep.NewAttachmentRepository(db)
	attachmentUC := attachmentUsecase.New(attachmentDB)
	attachment.RegisterAttachmentsServer(server, attachmentDelivery.New(attachmentUC))

	log.Println("starting server at :8082")
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
