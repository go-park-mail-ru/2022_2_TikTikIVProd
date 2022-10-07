package delivery

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/repository/localstorage"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/usecase"
)

type TestCase struct {
	User       string
	Method string
	StatusCode int
}

func TestDelivery_signUp(t *testing.T) {
	cases := []TestCase{
		{
			User:       `{"first_name":"Nastya", "last_name":"Kuznetsova", "nick_name":"kuzkus", "email":"aaa@gmail.com", "password":"password1"}`,
			Method: "POST",
			StatusCode: http.StatusCreated,
		},
		{
			User:       `{""}`,
			Method: "POST",
			StatusCode: http.StatusBadRequest,
		},
	}
	for caseNum, item := range cases {
		url := "http://89.208.197.127:8080/signup"
		req := httptest.NewRequest(item.Method, url, strings.NewReader(item.User))
		w := httptest.NewRecorder()

		usersLocalStorage := localstorage.New()
		usersUC := usecase.New(usersLocalStorage)
		usersDeliver := New(usersUC)
		usersDeliver.SignUp(w, req)

		if w.Code != item.StatusCode {
			t.Errorf("[%d] wrong StatusCode: got %d, expected %d",
				caseNum, w.Code, item.StatusCode)
		}

		// resp := w.Result()
		// body, _ := ioutil.ReadAll(resp.Body)

		// bodyStr := string(body)
		// if bodyStr != item.Response {
		// 	t.Errorf("[%d] wrong Response: got %+v, expected %+v",
		// 		caseNum, bodyStr, item.Response)
		// }
	}
}
