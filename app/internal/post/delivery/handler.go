package postsRouter

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"

	postsUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/pkg"
)

type DeliveryI interface {
	Feed(c echo.Context) error
	//GetPost(w http.ResponseWriter, r *http.Request)
	//CreatePost(w http.ResponseWriter, r *http.Request)
	//UpdatePost(w http.ResponseWriter, r *http.Request)
	//DeletePost(w http.ResponseWriter, r *http.Request)
}

type delivery struct {
	pUsecase postsUsecase.PostUseCaseI
}

// GetPost godoc
// @Summary      Show a post
// @Description  Get post by id
// @Tags     	 posts
// @Accept	 application/json
// @Produce  application/json
// @Param id  path int  true  "Post ID"
// @Success  200 {object} pkg.Response{body=models.Post} "success get post"
// @Failure 405 {object} pkg.Error "invalid http method"
// @Failure 500 {object} pkg.Error "internal server error"
// @Router   /post/{id} [get]
func (delivery *delivery) GetPost(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

// CreatePost godoc
// @Summary      Create a post
// @Description  Create a post
// @Tags     	 posts
// @Accept	 application/json
// @Produce  application/json
// @Param    post body models.Post true "post info"
// @Success  200 {object} pkg.Response{body=models.Post} "success get post"
// @Failure 405 {object} pkg.Error "invalid http method"
// @Failure 500 {object} pkg.Error "internal server error"
// @Router   /post/create [post]
func (delivery *delivery) CreatePost(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

// UpdatePost godoc
// @Summary      Update a post
// @Description  Update a post
// @Tags     	 posts
// @Accept	 application/json
// @Produce  application/json
// @Param    post body models.Post true "post info"
// @Success  200 {object} pkg.Response{body=models.Post} "success get post"
// @Failure 405 {object} pkg.Error "invalid http method"
// @Failure 500 {object} pkg.Error "internal server error"
// @Router   /post/update [post]
func (delivery *delivery) UpdatePost(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

// DeletePost godoc
// @Summary      Delete a post
// @Description  Delete a post
// @Tags     	 posts
// @Accept	 application/json
// @Produce  application/json
// @Param id path int  true  "Post ID"
// @Success  200 {object} pkg.Response{body=models.Post} "success get post"
// @Failure 405 {object} pkg.Error "invalid http method"
// @Failure 500 {object} pkg.Error "internal server error"
// @Router   /post/delete/{id} [delete]
func (delivery *delivery) DeletePost(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

// Feed godoc
// @Summary      Feed
// @Description  Feed
// @Tags     	 posts
// @Accept	 application/json
// @Produce  application/json
// @Success  200 {object} pkg.Response{body=[]models.Post} "success get feed"
// @Failure 405 {object} pkg.Error "invalid http method"
// @Failure 500 {object} pkg.Error "internal server error"
// @Router   /feed [get]
func (delivery *delivery) Feed(c echo.Context) error {
	log.Println("/feed")
	c.Response().Header().Set("Access-Control-Allow-Methods", "GET")
	//if r.Method != http.MethodGet {
	//	pkg.ErrorResponse(w, http.StatusMethodNotAllowed, "invalid http method")
	//	return
	//}

	posts, err := delivery.pUsecase.GetAllPosts()

	if err != nil {
		//pkg.ErrorResponse(w, http.StatusInternalServerError, err.Error()) TODO сделать через echo
		return err
	}

	err = pkg.JSONresponse(c.Response(), http.StatusOK, posts)
	if err != nil {
		//pkg.ErrorResponse(w, http.StatusInternalServerError, err.Error()) TODO сделать через echo
		return err
	}

	return nil
}

func NewDelivery(pUsecase postsUsecase.PostUseCaseI) DeliveryI {
	return &delivery{
		pUsecase: pUsecase,
	}
}
