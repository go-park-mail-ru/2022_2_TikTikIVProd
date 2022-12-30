package main

import (
	"log"

	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	elog "github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/cmd/server"
	_attachmentDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/attachment/delivery"
	attachmentsRepository "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/attachment/repository/microservice"
	attachmentUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/attachment/usecase"
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
	_postsDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/post/delivery"
	postsRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/post/repository/postgres"
	postsUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/post/usecase"
	_stickersDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/stickers/delivery"
	stickersRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/stickers/repository/postgres"
	stickersUseCase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/stickers/usecase"
	_usersDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/user/delivery"
	usersRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/user/repository/microservice"
	usersUseCase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/user/usecase"
	attachment "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/proto/attachment"
	auth "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/proto/auth"
	chat "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/proto/chat"
	user "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/proto/user"
	"github.com/labstack/echo-contrib/prometheus"
)

// @title WS Swagger API
// @version 1.0
// @host 89.208.197.127:8080

// var testCfgPg = postgres.Config{DSN: "host=localhost user=ws password=postgres_ws database=postgres port=13080"}

var prodCfgPg = postgres.Config{DSN: "host=ws_pg user=ws password=postgres_ws database=postgres port=5432"}

func main() {
	db, err := gorm.Open(postgres.New(prodCfgPg),
		&gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	grpcConnAuth, err := grpc.Dial(
		"auth_mvs:8081",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer grpcConnAuth.Close()
	authManager := auth.NewAuthClient(grpcConnAuth)

	grpcConnAttachment, err := grpc.Dial(
		"attachment_mvs:8082",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer grpcConnAttachment.Close()
	attachmentManager := attachment.NewAttachmentsClient(grpcConnAttachment)

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
	attachmentDB := attachmentsRepository.New(attachmentManager)
	chatDB := chatRep.New(chatManager)
	communitiesDb := communitiesRep.NewCommunitiesRepository(db)
	stickersDb := stickersRep.New(db)

	postsUC := postsUsecase.NewPostUsecase(postDB, attachmentDB, usersDB)
	authUC := authUseCase.New(authDB, usersDB)
	usersUC := usersUseCase.New(usersDB)
	friendsUC := friendsUseCase.New(friendsDB, usersDB)
	attachmentUC := attachmentUsecase.NewAttachmentUsecase(attachmentDB)
	chatUC := chatUseCase.New(chatDB, attachmentDB)
	communitiesUC := communitiesUseCase.New(communitiesDb)
	stickersUC := stickersUseCase.NewStickerUsecase(stickersDb)

	e := echo.New()

	e.Logger.SetHeader(`time=${time_rfc3339} level=${level} prefix=${prefix} ` +
		`file=${short_file} line=${line} message:`)
	e.Logger.SetLevel(elog.INFO)

	p := prometheus.NewPrometheus("echo", nil)
	p.MetricsPath = "/prometheus"
	p.SetMetricsPath(e)
	p.Use(e)

	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins:     []string{"http://writesend.online", "http://89.208.197.127"},
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
	_attachmentDelivery.NewDelivery(e, attachmentUC)
	_friendsDelivery.NewDelivery(e, friendsUC)
	_chatDelivery.NewDelivery(e, chatUC)
	_communitiesDelivery.NewDelivery(e, communitiesUC)
	_stickersDelivery.NewDelivery(e, stickersUC)

	s := server.NewServer(e)
	if err := s.Start(); err != nil {
		e.Logger.Fatal(err)
	}
}
