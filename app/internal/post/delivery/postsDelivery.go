package postsRouter

import (
	"net/http"

	postsUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/pkg"
)

//func wrapJSON(name string, item interface{}) ([]byte, error) {
//	wrapped := map[string]interface{}{
//		name: item,
//	}
//	converted, err := json.Marshal(wrapped)
//	return converted, err
//}

type DeliveryI interface {
	Feed(w http.ResponseWriter, r *http.Request)
}

type delivery struct {
	pUsecase postsUsecase.UseCaseI
}

func (delivery *delivery) Feed(w http.ResponseWriter, r *http.Request) {
	posts, _ := delivery.pUsecase.SelectAllPosts() //TODO ошибки

	err := pkg.JSONresponse(w, http.StatusOK, posts)
	if err != nil {
		pkg.ErrorResponse(w, http.StatusInternalServerError, err.Error())
	}
}

func NewDelivery(pUsecase postsUsecase.UseCaseI) DeliveryI {
	return &delivery{
		pUsecase: pUsecase,
	}
}
