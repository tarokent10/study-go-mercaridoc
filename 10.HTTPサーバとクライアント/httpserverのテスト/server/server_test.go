package server_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"study-go--mercaridoc/10.HTTPサーバとクライアント/httpserverのテスト/server"
	"testing"
)

func TestServer(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	server.SampleHandler(w, r)
	rw := w.Result()
	defer rw.Body.Close()
	if rw.StatusCode != http.StatusOK {
		t.Fatal("status code is not ok")
	}
	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal(err.Error())
	}

	var expected string = "Hello,World"
	if s := string(b); s != expected {
		t.Fatalf("unexpected reponse %s but expect %s", s, expected)
	}
}
