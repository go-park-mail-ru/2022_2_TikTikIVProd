package postsRouter

import (
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

// Feed godoc
// @Summary      Feed
// @Description  Feed
// @Tags     	 posts
// @Accept	 application/json
// @Produce  application/json
// @Success  200 {object} []model.Post "success get feed"
// @Failure 405 {object} pkg.Error "invalid http method"
// @Failure 500 {object} pkg.Error "internal server error"
// @Router   /feed [get]
func (delivery *delivery) Feed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	if r.Method != http.MethodGet {
		pkg.ErrorResponse(w, http.StatusMethodNotAllowed, "invalid http method")
		return
	}

	posts, err := delivery.pUsecase.SelectAllPosts()

	if err != nil {
		pkg.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

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
