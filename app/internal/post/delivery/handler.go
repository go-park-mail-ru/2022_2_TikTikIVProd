package delivery

import (
	"fmt"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/pkg"
	"github.com/labstack/echo/v4"
	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"

	postsUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/usecase"
)

type Delivery struct {
	PUsecase postsUsecase.PostUseCaseI
}

// GetPost godoc
// @Summary      Show a post
// @Description  Get post by id
// @Tags     	 posts
// @Accept	 application/json
// @Produce  application/json
// @Param id  path int  true  "Post ID"
// @Success  200 {object} pkg.Response{body=models.Post} "success get post"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Router   /post/{id} [get]
func (delivery *Delivery) GetPost(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusNotFound, "Post not found")
	}

	post, err := delivery.PUsecase.GetPostById(idP)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, pkg.Response{Body: post})
}

func isRequestValid(p *models.Post) (bool, error) {
	validate := validator.New()
	err := validate.Struct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CreatePost godoc
// @Summary      Create a post
// @Description  Create a post
// @Tags     	 posts
// @Accept	 application/json
// @Produce  application/json
// @Param    post body models.Post true "post info"
// @Success  200 {object} pkg.Response{body=models.Post} "success get post"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 403 {object} echo.HTTPError "invalid csrf"
// @Router   /post/create [post]
func (delivery *Delivery) CreatePost(c echo.Context) error {
	var post models.Post
	err := c.Bind(&post)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&post); !ok {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	requestSanitizePost(&post)

	userId, ok := c.Get("user_id").(int)
	if !ok {
		c.Logger().Error(models.ErrInternalServerError)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	post.UserID = userId
	err = delivery.PUsecase.CreatePost(&post)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, pkg.Response{Body: post})
}

func requestSanitizePost(post *models.Post) {
	sanitizer := bluemonday.UGCPolicy()

	post.UserFirstName = sanitizer.Sanitize(post.UserFirstName)
	post.UserLastName = sanitizer.Sanitize(post.UserLastName)
	post.Message = sanitizer.Sanitize(post.Message)
}

// UpdatePost godoc
// @Summary      Update a post
// @Description  Update a post
// @Tags     	 posts
// @Accept	 application/json
// @Produce  application/json
// @Param    post body models.Post true "post info"
// @Success  200 {object} pkg.Response{body=models.Post} "success get post"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 403 {object} echo.HTTPError "invalid csrf"
// @Router   /post/edit [post]
func (delivery *Delivery) UpdatePost(c echo.Context) error {
	var post models.Post
	err := c.Bind(&post)

	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&post); !ok {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	requestSanitizePost(&post)

	userId, ok := c.Get("user_id").(int)
	if !ok {
		c.Logger().Error(models.ErrInternalServerError)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	post.UserID = userId
	err = delivery.PUsecase.UpdatePost(&post)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, pkg.Response{Body: post})
}

// DeletePost godoc
// @Summary      Delete a post
// @Description  Delete a post
// @Tags     	 posts
// @Accept	 application/json
// @Param id path int  true  "Post ID"
// @Success  204
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 403 {object} echo.HTTPError "invalid csrf"
// @Router   /post/{id} [delete]
func (delivery *Delivery) DeletePost(c echo.Context) error {
	fmt.Println("param", c.Param("id"))
	idP, err := strconv.Atoi(c.Param("id"))

	userId, ok := c.Get("user_id").(int)
	if !ok {
		c.Logger().Error(models.ErrInternalServerError)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusNotFound, "Post not found")
	}

	err = delivery.PUsecase.DeletePost(idP, userId)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}
	return c.NoContent(http.StatusNoContent)
}

// Feed godoc
// @Summary      Feed
// @Description  Feed
// @Tags     	 posts
// @Produce  application/json
// @Success  200 {object} pkg.Response{body=[]models.Post} "success get feed"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Router   /feed [get]
func (delivery *Delivery) Feed(c echo.Context) error {
	posts, err := delivery.PUsecase.GetAllPosts()

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, pkg.Response{Body: posts})
}

// GetUserPosts godoc
// @Summary      Get users posts
// @Description  Get all users posts
// @Tags     	 posts
// @Param id path int  true  "Post ID"
// @Produce  application/json
// @Success  200 {object} pkg.Response{body=[]models.Post} "success get feed"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 404 {object} echo.HTTPError "Post not found"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Router   /users/{id}/posts [get]
func (delivery *Delivery) GetUserPosts(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusNotFound, "Post not found")
	}

	posts, err := delivery.PUsecase.GetUserPosts(idP)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, pkg.Response{Body: posts})
}

func NewDelivery(e *echo.Echo, up postsUsecase.PostUseCaseI) {
	handler := &Delivery{
		PUsecase: up,
	}

	e.POST("/post/create", handler.CreatePost)
	e.POST("/post/edit", handler.UpdatePost)
	e.GET("/post/:id", handler.GetPost)
	e.GET("/users/:id/posts", handler.GetUserPosts)
	e.GET("/feed", handler.Feed)
	e.DELETE("/post/:id", handler.DeletePost)
}
