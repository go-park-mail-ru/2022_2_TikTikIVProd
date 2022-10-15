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
			User:   `{"first_name":"Nastya", "last_name":"Kuznetsova", "nick_name":"kuzkus", "email":"aaa@gmail.com", "password":"password1"}`,
			Method: "POST",
			StatusCode: http.StatusCreated,
		},
		{
			User:       `{""}`,
			Method: "POST",
			StatusCode: http.StatusBadRequest,
		},
		{
			User:       `{""}`,
			Method: "GET",
			StatusCode: http.StatusMethodNotAllowed,
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
	}
}

func TestDelivery_signInFailure(t *testing.T) {
	cases := []TestCase{
		{
			User:       `{""}`,
			Method: "POST",
			StatusCode: http.StatusBadRequest,
		},
		{
			User:       `{""}`,
			Method: "GET",
			StatusCode: http.StatusMethodNotAllowed,
		},
		{
			User:       `{"email":"aaa@gmail.com", "password":"password1"}`,
			Method: "POST",
			StatusCode: http.StatusNotFound,
		},
	}
	for caseNum, item := range cases {
		url := "http://89.208.197.127:8080/signin"
		req := httptest.NewRequest(item.Method, url, strings.NewReader(item.User))
		w := httptest.NewRecorder()

		usersLocalStorage := localstorage.New()
		usersUC := usecase.New(usersLocalStorage)
		usersDeliver := New(usersUC)
		usersDeliver.SignIn(w, req)

		if w.Code != item.StatusCode {
			t.Errorf("[%d] wrong StatusCode: got %d, expected %d",
				caseNum, w.Code, item.StatusCode)
		}
	}
}

func TestDelivery_LogoutFailure(t *testing.T) {
	cases := []TestCase{
		{
			User:       `{""}`,
			Method: "POST",
			StatusCode: http.StatusBadRequest,
		},
		{
			User:       `{""}`,
			Method: "GET",
			StatusCode: http.StatusMethodNotAllowed,
		},
		{
			User:       `{"email":"aaa@gmail.com", "password":"password1"}`,
			Method: "POST",
			StatusCode: http.StatusNotFound,
		},
	}
	for caseNum, item := range cases {
		url := "http://89.208.197.127:8080/signin"
		req := httptest.NewRequest(item.Method, url, strings.NewReader(item.User))
		w := httptest.NewRecorder()

		usersLocalStorage := localstorage.New()
		usersUC := usecase.New(usersLocalStorage)
		usersDeliver := New(usersUC)
		usersDeliver.SignIn(w, req)

		if w.Code != item.StatusCode {
			t.Errorf("[%d] wrong StatusCode: got %d, expected %d",
				caseNum, w.Code, item.StatusCode)
		}
	}
}


