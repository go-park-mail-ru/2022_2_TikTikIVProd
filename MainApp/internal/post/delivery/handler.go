package delivery

import (
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/pkg"
	"github.com/labstack/echo/v4"
	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"

	postsUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/post/usecase"
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
	idP, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	userId, ok := c.Get("user_id").(uint64)
	if !ok {
		c.Logger().Error(models.ErrInternalServerError)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	post, err := delivery.PUsecase.GetPostById(idP, userId)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, pkg.Response{Body: post})
}

// LikePost godoc
// @Summary      Like a post
// @Description  Like a post
// @Tags     	 posts
// @Produce  application/json
// @Param id  path int  true  "Post ID"
// @Success  204
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 403 {object} echo.HTTPError "invalid csrf"
// @Router   /post/like/{id} [put]
func (delivery *Delivery) LikePost(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	userId, ok := c.Get("user_id").(uint64)
	if !ok {
		c.Logger().Error(models.ErrInternalServerError)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	err = delivery.PUsecase.LikePost(id, userId)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	return c.NoContent(http.StatusNoContent)
}

// UnLikePost godoc
// @Summary      Unlike a post
// @Description  Unlike a post
// @Tags     	 posts
// @Produce  application/json
// @Param id  path int  true  "Post ID"
// @Success  204
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 403 {object} echo.HTTPError "invalid csrf"
// @Router   /post/unlike/{id} [put]
func (delivery *Delivery) UnLikePost(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	userId, ok := c.Get("user_id").(uint64)
	if !ok {
		c.Logger().Error(models.ErrInternalServerError)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	err = delivery.PUsecase.UnLikePost(id, userId)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	return c.NoContent(http.StatusNoContent)
}

func isRequestValid(entity interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(entity)
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

	userId, ok := c.Get("user_id").(uint64)
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

func requestSanitizeComment(comment *models.Comment) {
	sanitizer := bluemonday.UGCPolicy()

	comment.UserFirstName = sanitizer.Sanitize(comment.UserFirstName)
	comment.UserLastName = sanitizer.Sanitize(comment.UserLastName)
	comment.Message = sanitizer.Sanitize(comment.Message)
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

	userId, ok := c.Get("user_id").(uint64)
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
	idP, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusNotFound, "Post not found")
	}

	userId, ok := c.Get("user_id").(uint64)
	if !ok {
		c.Logger().Error(models.ErrInternalServerError)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
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
	userId, ok := c.Get("user_id").(uint64)
	if !ok {
		c.Logger().Error(models.ErrInternalServerError)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	posts, err := delivery.PUsecase.GetAllPosts(userId)

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
	idP, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusNotFound, "Post not found")
	}

	curUserId, ok := c.Get("user_id").(uint64)
	if !ok {
		c.Logger().Error(models.ErrInternalServerError)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	posts, err := delivery.PUsecase.GetUserPosts(idP, curUserId)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, pkg.Response{Body: posts})
}

// GetCommunityPosts godoc
// @Summary      Get community posts
// @Description  Get all community posts
// @Tags     	 posts
// @Param id path int  true  "Community ID"
// @Produce  application/json
// @Success  200 {object} pkg.Response{body=[]models.Post} "success get community posts"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 404 {object} echo.HTTPError "Community not found"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Router   /communities/{id}/posts [get]
func (delivery *Delivery) GetCommunityPosts(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusNotFound, models.ErrNotFound)
	}

	curUserId, ok := c.Get("user_id").(uint64)
	if !ok {
		c.Logger().Error(models.ErrInternalServerError)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	posts, err := delivery.PUsecase.GetCommunityPosts(id, curUserId)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError)
	}

	return c.JSON(http.StatusOK, pkg.Response{Body: posts})
}

// AddComment godoc
// @Summary      Add a comment
// @Description  Add a comment
// @Tags     	 posts
// @Accept	 application/json
// @Produce  application/json
// @Param    comment body models.Comment true "comment info"
// @Success  200 {object} pkg.Response{body=models.Comment} "success get post"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 403 {object} echo.HTTPError "invalid csrf"
// @Router   /post/comment/add [post]
func (delivery *Delivery) AddComment(c echo.Context) error {
	var comment models.Comment
	err := c.Bind(&comment)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	if ok, err := isRequestValid(&comment); !ok {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	requestSanitizeComment(&comment)

	userId, ok := c.Get("user_id").(uint64)
	if !ok {
		c.Logger().Error(models.ErrInternalServerError)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	comment.UserID = userId
	err = delivery.PUsecase.AddComment(&comment)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	return c.JSON(http.StatusOK, pkg.Response{Body: comment})
}

// UpdateComment godoc
// @Summary      Update a comment
// @Description  Update a comment
// @Tags     	 posts
// @Accept	 application/json
// @Produce  application/json
// @Param    comment body models.Comment true "comment info"
// @Success  200 {object} pkg.Response{body=models.Comment} "success update comment"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 403 {object} echo.HTTPError "invalid csrf"
// @Router   /post/comment/edit [post]
// func (delivery *Delivery) UpdateComment(c echo.Context) error {
// 	var comment models.Comment
// 	err := c.Bind(&comment)
// 	if err != nil {
// 		c.Logger().Error(err)
// 		return c.JSON(http.StatusBadRequest, models.ErrBadRequest.Error())
// 	}

// 	requestSanitizeComment(&comment)

// 	userId, ok := c.Get("user_id").(uint64)
// 	if !ok {
// 		c.Logger().Error(models.ErrInternalServerError)
// 		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
// 	}

// 	comment.UserID = userId

// 	err = delivery.PUsecase.UpdateComment(&comment)
// 	if err != nil {
// 		c.Logger().Error(err)
// 		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
// 	}

// 	return c.JSON(http.StatusOK, pkg.Response{Body: comment})
// }

// DeleteComment godoc
// @Summary      Delete a comment
// @Description  Delete a comment
// @Tags     	 posts
// @Accept	 application/json
// @Param id path int  true  "Comment ID"
// @Success  204
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 404 {object} echo.HTTPError "item not found"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 403 {object} echo.HTTPError "invalid csrf"
// @Router   /post/comment/{id} [delete]
func (delivery *Delivery) DeleteComment(c echo.Context) error {
	idC, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	userId, ok := c.Get("user_id").(uint64)
	if !ok {
		c.Logger().Error(models.ErrInternalServerError)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	err = delivery.PUsecase.DeleteComment(idC, userId)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	return c.NoContent(http.StatusNoContent)
}

// GetComments godoc
// @Summary      Get comments
// @Description  Get all comments by post
// @Tags     	 posts
// @Param id path int  true  "Post ID"
// @Produce  application/json
// @Success  200 {object} pkg.Response{body=[]models.Comment} "success get post comments"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 404 {object} echo.HTTPError "Post not found"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Router   /post/{id}/comments [get]
func (delivery *Delivery) GetComments(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	comments, err := delivery.PUsecase.GetComments(id)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	return c.JSON(http.StatusOK, pkg.Response{Body: comments})
}

func NewDelivery(e *echo.Echo, up postsUsecase.PostUseCaseI) {
	handler := &Delivery{
		PUsecase: up,
	}

	e.POST("/post/create", handler.CreatePost)
	e.POST("/post/edit", handler.UpdatePost)
	e.PUT("/post/like/:id", handler.LikePost)
	e.PUT("/post/unlike/:id", handler.UnLikePost)
	e.GET("/post/:id", handler.GetPost)
	e.GET("/users/:id/posts", handler.GetUserPosts)
	e.GET("/communities/:id/posts", handler.GetCommunityPosts)
	e.GET("/feed", handler.Feed)
	e.DELETE("/post/:id", handler.DeletePost)
	e.GET("/post/:id/comments", handler.GetComments)
	e.POST("/post/comment/add", handler.AddComment)
	//e.POST("/post/comment/edit", handler.UpdateComment)
	e.DELETE("/post/comment/:id", handler.DeleteComment)
}
