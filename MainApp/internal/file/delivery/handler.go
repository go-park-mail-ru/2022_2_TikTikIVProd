package delivery

import (
	"github.com/gabriel-vasile/mimetype"
	fileUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/file/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/pkg"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"strconv"
)

type DeliveryI interface {
	GetFileByID(c echo.Context) error
	UploadFile(c echo.Context) error
}

type delivery struct {
	fileUsecase fileUsecase.FileUseCaseI
}

// GetFileByID godoc
// @Summary      Get file by id
// @Description  Get file by id
// @Tags     	 file
// @Param id path int  true  "file ID"
// @Success  200 "success get file"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Router   /file/{id} [get]
func (delivery *delivery) GetFileByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	file, err := delivery.fileUsecase.GetFileById(id)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}
	path := "files/" + file.FileLink
	f, err := os.Open(path)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	mtype, err := mimetype.DetectFile(path)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	return c.Stream(http.StatusOK, mtype.Extension(), f)
}

// UploadFile godoc
// @Summary      Upload file
// @Description  Upload file
// @Tags     	 file
// @Param file formData file  true  "file"
// @Accept multipart/form-data
// @Produce  application/json
// @Success  200 "success upload file"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 403 {object} echo.HTTPError "invalid csrf"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Router   /file/upload [post]
func (delivery *delivery) UploadFile(c echo.Context) error {
	// Source
	f, err := c.FormFile("file")
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Not file in form")
	}
	src, err := f.Open()
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("files/" + f.Filename)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")

	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	file := models.File{FileLink: f.Filename}
	err = delivery.fileUsecase.CreateFile(&file)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")

	}

	return c.JSON(http.StatusOK, pkg.Response{Body: file})
}

func NewDelivery(e *echo.Echo, iu fileUsecase.FileUseCaseI) {
	handler := &delivery{
		fileUsecase: iu,
	}

	e.POST("/file/upload", handler.UploadFile)
	e.GET("/file/:id", handler.GetFileByID)
}
