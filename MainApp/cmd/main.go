package main

import (
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/cmd/server"
	_authDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/auth/delivery"
	authRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/auth/repository/microservice"
	authUseCase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/auth/usecase"
	_chatDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/chat/delivery"
	chatRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/chat/repository/microservice"
	chatUseCase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/chat/usecase"
	_communitiesDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/communities/delivery"
	communitiesRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/communities/repository/postgres"
	communitiesUseCase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/communities/usecase"
	_friendsDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/friends/delivery"
	friendsRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/friends/repository/microservice"
	friendsUseCase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/friends/usecase"
	_imageDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/image/delivery"
	imagesRepository "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/image/repository/microservice"
	imageUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/image/usecase"
	_postsDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/post/delivery"
	postsRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/post/repository/postgres"
	postsUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/post/usecase"
	_usersDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/user/delivery"
	usersRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/user/repository/microservice"
	usersUseCase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/user/usecase"
	auth "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/proto/auth"
	image "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/proto/image"
	chat "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/proto/chat"
	user "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/proto/user"
)

// @title WS Swagger API
// @version 1.0
// @host 89.208.197.127:8080

var testCfgPg = postgres.Config{DSN: "host=localhost user=postgres password=postgres port=13080"}

// var prodCfgPg = postgres.Config{DSN: "host=ws_pg user=postgres password=postgres port=5432"}

func main() {
	db, err := gorm.Open(postgres.New(testCfgPg),
		&gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	grpcConnAuth, err := grpc.Dial(
		"auth_mvs:8081",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer grpcConnAuth.Close()
	authManager := auth.NewAuthClient(grpcConnAuth)

	grpcConnImage, err := grpc.Dial(
		"image_mvs:8082",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer grpcConnImage.Close()
	imageManager := image.NewImagesClient(grpcConnImage)

	grpcConnChat, err := grpc.Dial(
		"chat_mvs:8083",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer grpcConnChat.Close()
	chatManager := chat.NewChatClient(grpcConnChat)

	grpcConnUser, err := grpc.Dial(
		"user_mvs:8084",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer grpcConnUser.Close()
	userManager := user.NewUsersClient(grpcConnUser)

	postDB := postsRep.NewPostRepository(db)
	authDB := authRep.New(authManager)
	usersDB := usersRep.New(userManager)
	friendsDB := friendsRep.New(userManager)
	imageDB := imagesRepository.New(imageManager)
	chatDB := chatRep.New(chatManager)
	communitiesDb := communitiesRep.NewCommunitiesRepository(db)

	postsUC := postsUsecase.NewPostUsecase(postDB, imageDB, usersDB)
	authUC := authUseCase.New(authDB, usersDB)
	usersUC := usersUseCase.New(usersDB)
	friendsUC := friendsUseCase.New(friendsDB, usersDB)
	imageUC := imageUsecase.NewImageUsecase(imageDB)
	chatUC := chatUseCase.New(chatDB)
	communitiesUC := communitiesUseCase.New(communitiesDb)

	e := echo.New()

	e.Logger.SetHeader(`time=${time_rfc3339} level=${level} prefix=${prefix} ` +
		`file=${short_file} line=${line} message:`)
	e.Logger.SetLevel(log.INFO)

	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins:     []string{"http://89.208.197.127"},
		AllowHeaders:     []string{"Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"X-CSRF-Token"},
	}))

	e.Use(echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
		Format: `time=${time_custom} remote_ip=${remote_ip} ` +
			`host=${host} method=${method} uri=${uri} user_agent=${user_agent} ` +
			`status=${status} error="${error}" ` +
			`bytes_in=${bytes_in} bytes_out=${bytes_out}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))

	e.Use(echoMiddleware.Recover())

	authMiddleware := middleware.NewMiddleware(authUC)
	e.Use(authMiddleware.Auth)
	e.Use(authMiddleware.CSRF)

	_postsDelivery.NewDelivery(e, postsUC)
	_authDelivery.NewDelivery(e, authUC)
	_usersDelivery.NewDelivery(e, usersUC)
	_imageDelivery.NewDelivery(e, imageUC)
	_friendsDelivery.NewDelivery(e, friendsUC)
	_chatDelivery.NewDelivery(e, chatUC)
	_communitiesDelivery.NewDelivery(e, communitiesUC)

	s := server.NewServer(e)
	if err := s.Start(); err != nil {
		e.Logger.Fatal(err)
	}
}
