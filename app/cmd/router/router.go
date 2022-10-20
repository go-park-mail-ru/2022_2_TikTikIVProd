package router

import (
	imageDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/image/delivery"
	postsDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/delivery"
	usersDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/delivery"
	"github.com/labstack/echo/v4"
)

type EchoRouter struct {
	*echo.Echo
	usersD usersDelivery.DeliveryI
	pd     postsDelivery.DeliveryI
	imgD   imageDelivery.DeliveryI
}

func NewEchoRouter(usersD usersDelivery.DeliveryI, pD postsDelivery.DeliveryI, iD imageDelivery.DeliveryI) *EchoRouter {
	e := &EchoRouter{
		Echo:   echo.New(),
		usersD: usersD,
		pd:     pD,
		imgD:   iD,
	}

	//e.HandleFunc("/signin", usersD.SignIn)
	//e.HandleFunc("/signup", usersD.SignUp)
	//e.HandleFunc("/auth", usersD.Auth)
	//e.HandleFunc("/logout", usersD.Logout)
	e.GET("/feed", pD.Feed)
	e.GET("/image/:id", iD.GetImageByID)
	return e
}
