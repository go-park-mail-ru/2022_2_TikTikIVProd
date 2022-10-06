package postsRouter

import (
	"fmt"
	"net/http"

	postsUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/pkg"
)

type DeliveryI interface {
	Feed(w http.ResponseWriter, r *http.Request)
}

type delivery struct {
	pUsecase postsUsecase.UseCaseI
}

func (delivery *delivery) Feed(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header)
	if r.Method != http.MethodGet {
		pkg.ErrorResponse(w, http.StatusMethodNotAllowed, "invalid http method")
		return
	}

	posts, err := delivery.pUsecase.SelectAllPosts()

	if err != nil {
		pkg.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	err = pkg.JSONresponse(w, http.StatusOK, posts)
	if err != nil {
		pkg.ErrorResponse(w, http.StatusInternalServerError, err.Error())
	}
}

func NewDelivery(pUsecase postsUsecase.UseCaseI) DeliveryI {
	return &delivery{
		pUsecase: pUsecase,
	}
}
