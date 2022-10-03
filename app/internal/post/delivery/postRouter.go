package postsRouter

import (
	"encoding/json"
	postsRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

type PostsRouter struct {
	*mux.Router
	pUsecase *postsRep.PostsUsecase
}

func (router *PostsRouter) Feed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Feed")
}

func NewPostsRouter(pUsecase *postsRep.PostsUsecase) *PostsRouter {
	r := &PostsRouter{
		Router:   mux.NewRouter(),
		pUsecase: pUsecase,
	}

	r.HandleFunc("/feed", r.Feed)
	return r
}
