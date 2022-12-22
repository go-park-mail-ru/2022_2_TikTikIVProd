package main

import (
	authDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/AuthMicroservice/internal/auth/delivery"
	authRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/AuthMicroservice/internal/auth/repository/redis"
	authUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/AuthMicroservice/internal/auth/usecase"
	auth "github.com/go-park-mail-ru/2022_2_TikTikIVProd/AuthMicroservice/proto"
	"github.com/go-redis/redis"
	"google.golang.org/grpc"
	"log"
	"net"
)

// var testCfgRedis = &redis.Options{Addr: ":6379", Password: "ws_redis_password"}

var prodCfgRedis = &redis.Options{Addr: "redis:6379", Password: "ws_redis_password"}

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	redisClient := redis.NewClient(prodCfgRedis)

	err = redisClient.Ping().Err()
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	authDB := authRep.New(redisClient)
	authUC := authUsecase.New(authDB)
	auth.RegisterAuthServer(server, authDelivery.New(authUC))

	log.Println("starting server at :8081")
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
