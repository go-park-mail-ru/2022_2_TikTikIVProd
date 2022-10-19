package router

import (
	postsDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/delivery"
	usersDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/delivery"
	"github.com/labstack/echo/v4"
)

type EchoRouter struct {
	*echo.Echo
	usersD usersDelivery.DeliveryI
	pd     postsDelivery.DeliveryI
}

func NewEchoRouter(usersD usersDelivery.DeliveryI, pd postsDelivery.DeliveryI) *EchoRouter {
	e := &EchoRouter{
		Echo:   echo.New(),
		usersD: usersD,
		pd:     pd,
	}

	//e.HandleFunc("/signin", usersD.SignIn)
	//e.HandleFunc("/signup", usersD.SignUp)
	//e.HandleFunc("/auth", usersD.Auth)
	//e.HandleFunc("/logout", usersD.Logout)
	e.GET("/feed", pd.Feed)
	return e
}
