package router

import (
	friendsDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/friends/delivery"
	authDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/auth/delivery"
	imageDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/image/delivery"
	postsDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/delivery"
	usersDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/delivery"
	"github.com/labstack/echo/v4"
)

type EchoRouter struct {
	*echo.Echo
	ud usersDelivery.DeliveryI
	fd friendsDelivery.DeliveryI
	ad authDelivery.DeliveryI
	pd postsDelivery.DeliveryI
	imgd imageDelivery.DeliveryI
}

func NewEchoRouter(ud usersDelivery.DeliveryI, fd friendsDelivery.DeliveryI, ad authDelivery.DeliveryI, pd postsDelivery.DeliveryI, imgd imageDelivery.DeliveryI) *EchoRouter {
	e := &EchoRouter{
		Echo: echo.New(),
		ud:   ud,
		fd:   fd,
		ad:   ad,
		pd:   pd,
		imgd: imgd,
	}

	e.POST("/signin", ad.SignIn)
	e.POST("/signup", ad.SignUp)
	e.GET("/auth", ad.Auth)
	e.DELETE("/logout", ad.Logout)
	e.GET("/users/:id", ud.GetProfile)
	e.POST("/friends/add", fd.AddFriend)
	e.DELETE("/friends/delete", fd.DeleteFriend)
	e.GET("/feed", pd.Feed)
	e.GET("/image/:id", imgd.GetImageByID)
	return e
}
