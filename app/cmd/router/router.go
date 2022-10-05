package router

import (
	"github.com/gorilla/mux"

	usersDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/delivery"
)

type Router struct {
	*mux.Router
	usersD usersDelivery.DeliveryI
}

func NewRouter(usersD usersDelivery.DeliveryI) *Router {
	r := &Router {
		Router: mux.NewRouter(),
		usersD: usersD,
	}

	//r.HandleFunc("/feed", r.Feed)
	r.HandleFunc("/signin", usersD.SignIn)
	r.HandleFunc("/auth", usersD.Auth)
	r.HandleFunc("/signup", usersD.SignUp)
	return r
}

