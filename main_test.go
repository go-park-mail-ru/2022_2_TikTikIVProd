package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFeed(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/feed", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	Feed(res, req)

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
	SignIn(res, req)

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
	SignUp(res, req)

	exp, _ := json.Marshal("SignUp") // прост потому что в хендлере тоже json от строки делает "Feed"
	act := res.Body.Bytes()
	if strings.Trim(string(act), "\n") != string(exp) {
		t.Fatalf("Expected %s got %s", string(exp), string(act))
	}
}
