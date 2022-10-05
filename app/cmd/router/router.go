package router

import (
	"github.com/gorilla/mux"

	usersDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/delivery"
	postsDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/delivery"
)

type Router struct {
	*mux.Router
	usersD usersDelivery.DeliveryI
	pd *postsDelivery.Delivery
}

func NewRouter(usersD usersDelivery.DeliveryI, pd* postsDelivery.Delivery) *Router {
	r := &Router {
		Router: mux.NewRouter(),
		usersD: usersD,
		pd: pd,
	}

	r.HandleFunc("/signin", usersD.SignIn)
	r.HandleFunc("/signup", usersD.SignUp)
	r.HandleFunc("/auth", usersD.Auth)
	r.HandleFunc("/logout", usersD.Logout)
	r.HandleFunc("/feed", pd.Feed)
	return r
}

