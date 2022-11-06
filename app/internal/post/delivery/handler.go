package delivery

import (
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/pkg"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
	"github.com/microcosm-cc/bluemonday"
	"net/http"
	"strconv"

	postsUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/post/usecase"
)

type DeliveryI interface {
	Feed(c echo.Context) error
	GetPost(c echo.Context) error
	GetUserPosts(c echo.Context) error
	CreatePost(c echo.Context) error
	UpdatePost(c echo.Context) error
	DeletePost(c echo.Context) error
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
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 403 {object} echo.HTTPError "invalid csrf"
// @Router   /post/{id} [get]
func (delivery *delivery) GetPost(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusNotFound, "Post not found") //TODO переделать на ошибки в файле
	}

	post, err := delivery.pUsecase.GetPostById(idP)

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
func (delivery *delivery) CreatePost(c echo.Context) error {
	var post models.Post
	err := c.Bind(&post)

	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&post); !ok {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	requestSanitizePost(&post)

	err = delivery.pUsecase.CreatePost(&post)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error") // TODO здесь тоже, нужно разграничить ошибки
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
// @Router   /post/update [post]
func (delivery *delivery) UpdatePost(c echo.Context) error {
	var post models.Post
	err := c.Bind(&post)

	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&post); !ok {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	requestSanitizePost(&post)

	err = delivery.pUsecase.UpdatePost(&post)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error") // TODO здесь тоже, нужно разграничить ошибки
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
func (delivery *delivery) DeletePost(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusNotFound, "Post not found") //TODO переделать на ошибки в файле
	}

	err = delivery.pUsecase.DeletePost(idP)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error") // TODO здесь тоже, нужно разграничить ошибки
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
// @Failure 403 {object} echo.HTTPError "invalid csrf"
// @Router   /feed [get]
func (delivery *delivery) Feed(c echo.Context) error {
	posts, err := delivery.pUsecase.GetAllPosts()

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
// @Failure 403 {object} echo.HTTPError "invalid csrf"
// @Router   /users/{id}/posts [get]
func (delivery *delivery) GetUserPosts(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusNotFound, "Post not found") //TODO переделать на ошибки в файле
	}

	posts, err := delivery.pUsecase.GetUserPosts(idP)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, pkg.Response{Body: posts})
}

func NewDelivery(e *echo.Echo, up postsUsecase.PostUseCaseI) {
	handler := &delivery{
		pUsecase: up,
	}

	e.POST("/post/create", handler.CreatePost)
	e.POST("/post/edit", handler.UpdatePost)
	e.GET("/post/:id", handler.GetPost)
	e.GET("/users/:id/posts", handler.GetUserPosts)
	e.GET("/feed", handler.Feed)
	e.DELETE("/post/:id", handler.DeletePost)
}
