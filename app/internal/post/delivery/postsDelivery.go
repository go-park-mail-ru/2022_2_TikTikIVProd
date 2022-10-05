package postsRouter

import (
	"encoding/json"
	"fmt"
	postsUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/usecase"
	"net/http"
)

//func wrapJSON(name string, item interface{}) ([]byte, error) {
//	wrapped := map[string]interface{}{
//		name: item,
//	}
//	converted, err := json.Marshal(wrapped)
//	return converted, err
//}

type Delivery struct {
	pUsecase *postsUsecase.PostsUsecase
}

func (delivery *Delivery) Feed(w http.ResponseWriter, r *http.Request) {
	posts, _ := delivery.pUsecase.SelectAllPosts() //TODO ошибки
	postJson, _ := json.Marshal(posts)             // TODO ошибки
	resp := fmt.Sprintf("{body: {posts: %s}}", string(postJson))
	//fmt.Println(resp)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(resp)
}

func NewDelivery(pUsecase *postsUsecase.PostsUsecase) *Delivery {
	delivery := &Delivery{
		pUsecase: pUsecase,
	}

	return delivery
}
