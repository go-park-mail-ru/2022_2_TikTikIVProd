package delivery

import (
	"io"
	"net/http"
	"os"
	"strconv"

	attUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/attachment/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/pkg"
	"github.com/labstack/echo/v4"
)

type DeliveryI interface {
	GetAttachmentByID(c echo.Context) error
	UploadAttachment(c echo.Context) error
}

type delivery struct {
	attUsecase attUsecase.AttachmentUseCaseI
}

// GetAttachmentByID godoc
// @Summary      Get Attachment by id
// @Description  Get Attachment by id
// @Tags     	 Attachment
// @Param id path int  true  "Attachment ID"
// @Produce  Attachment/png
// @Success  200 "success get Attachment"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Router   /attachment/{id} [get]
func (delivery *delivery) GetAttachmentByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	att, err := delivery.attUsecase.GetAttachmentById(id)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	f, err := os.Open("attachments/" + att.AttLink)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	return c.Stream(http.StatusOK, "Attachment/png", f)
}

// UploadAttachmentImage godoc
// @Summary      Upload image
// @Description  Upload image
// @Tags     	 Attachment
// @Param Attachment formData image  true  ""
// @Accept multipart/form-data
// @Produce  application/json
// @Success  200 "success upload Attachment"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 403 {object} echo.HTTPError "invalid csrf"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Router   /attachment/image/upload [post]
func (delivery *delivery) UploadAttachmentImage(c echo.Context) error {
	// Source
	file, err := c.FormFile("Attachment")
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Not Attachment in form")
	}
	src, err := file.Open()
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}
	defer src.Close()

	// Destination
	path := "attachments/images" + file.Filename
	dst, err := os.Create(path)
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

	attachment := models.Attachment{AttLink: path, Type: models.ImageAtt}
	err = delivery.attUsecase.CreateAttachment(&attachment)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")

	}

	return c.JSON(http.StatusOK, pkg.Response{Body: attachment})
}

// UploadAttachmentFile godoc
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
// @Router   /attachment/file/upload [post]
func (delivery *delivery) UploadAttachmentFile(c echo.Context) error {
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
	path := "attachments/files" + f.Filename
	dst, err := os.Create(path)
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

	attachment := models.Attachment{AttLink: path, Type: models.FileAtt}
	err = delivery.attUsecase.CreateAttachment(&attachment)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")

	}

	return c.JSON(http.StatusOK, pkg.Response{Body: attachment})
}

func NewDelivery(e *echo.Echo, iu attUsecase.AttachmentUseCaseI) {
	handler := &delivery{
		attUsecase: iu,
	}

	e.POST("/attachment/image/upload", handler.UploadAttachmentImage)
	e.POST("/attachment/file/upload", handler.UploadAttachmentFile)
	e.GET("/attachment/:id", handler.GetAttachmentByID)
}
