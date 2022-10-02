package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestFeed(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/feed", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()

	r := NewRouter()
	r.Feed(res, req)

	exp, _ := json.Marshal("Feed") // прост потому что в хендлере тоже json от строки делает "Feed"
	act := res.Body.Bytes()
	if strings.Trim(string(act), "\n") != string(exp) {
		t.Fatalf("Expected %s got %s", string(exp), string(act))
	}
}

func TestSignIn(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/feed", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()

	r := NewRouter()
	r.SignIn(res, req)

	exp, _ := json.Marshal("SignIn") // прост потому что в хендлере тоже json от строки делает "Feed"
	act := res.Body.Bytes()
	if strings.Trim(string(act), "\n") != string(exp) {
		t.Fatalf("Expected %s got %s", string(exp), string(act))
	}
}

func TestSignUp(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/feed", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()

	r := NewRouter()
	r.SignUp(res, req)

	exp, _ := json.Marshal("SignUp") // прост потому что в хендлере тоже json от строки делает "Feed"
	act := res.Body.Bytes()
	if strings.Trim(string(act), "\n") != string(exp) {
		t.Fatalf("Expected %s got %s", string(exp), string(act))
	}
}

// func TestRouter_Feed(t *testing.T) {
// 	type fields struct {
// 		Router *mux.Router
// 	}
// 	type args struct {
// 		w http.ResponseWriter
// 		r *http.Request
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			router := &Router{
// 				Router: tt.fields.Router,
// 			}
// 			router.Feed(tt.args.w, tt.args.r)
// 		})
// 	}
// }

