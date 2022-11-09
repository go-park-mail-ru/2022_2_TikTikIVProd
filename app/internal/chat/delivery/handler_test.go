package delivery_test

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	chatDelivery "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/chat/delivery"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/chat/usecase/mocks"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestCase struct {
	ArgData string
	Error error
	StatusCode int
}

func TestDeliveryGetDialog(t *testing.T) {
	var dialog models.Dialog
	err := faker.FakeData(&dialog)
	assert.NoError(t, err)

	mockUCase := mocks.NewUseCaseI(t)

	mockUCase.On("SelectDialog", dialog.Id).Return(&dialog, nil)
	handler := chatDelivery.Delivery{
		ChatUC: mockUCase,
	}

	userIdBadRequest := "jgsv"

	e := echo.New()
	chatDelivery.NewDelivery(e, mockUCase)

	cases := map[string]TestCase {
		"success": {
			ArgData: strconv.Itoa(dialog.Id),
			Error: nil,
			StatusCode: http.StatusOK,
		},
		"bad_request": {
			ArgData: userIdBadRequest,
			Error: &echo.HTTPError{
				Code: http.StatusBadRequest,
				Message: models.ErrBadRequest.Error(),
			},
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(echo.GET, "/users/:id", strings.NewReader(""))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/users/:id")
			c.SetParamNames("id")
			c.SetParamValues(test.ArgData)

			err := handler.GetDialog(c)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.StatusCode, rec.Code)
			}
		})
	}

	mockUCase.AssertExpectations(t)
}
