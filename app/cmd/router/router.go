package router

import (
	"github.com/gorilla/mux"

	postsDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/delivery"
	usersDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/delivery"
)

type Router struct {
	*mux.Router
	usersD usersDelivery.DeliveryI
	pd     postsDelivery.DeliveryI
}

func NewRouter(usersD usersDelivery.DeliveryI, pd postsDelivery.DeliveryI) *Router {
	r := &Router{
		Router: mux.NewRouter(),
		usersD: usersD,
		pd:     pd,
	}

	r.HandleFunc("/signin", usersD.SignIn)
	r.HandleFunc("/signup", usersD.SignUp)
	r.HandleFunc("/auth", usersD.Auth)
	r.HandleFunc("/logout", usersD.Logout)
	r.HandleFunc("/feed", pd.Feed)
	return r
}
