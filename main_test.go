package main

import (
	"github.com/gorilla/sessions"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestLogin(t *testing.T) {
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/session", nil)


	var store = &StoreMock{Session: &sessions.Session{
		Values: map[interface{}]interface{}{
			"username" : "test",
		},
	}}


	route := initMux(store)
	route.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("%d", recorder.Code)
	}
}


type StoreMock struct {
	Session *sessions.Session
}

func (s StoreMock) Get(r *http.Request, name string) (*sessions.Session, error) {
	return s.Session, nil
}

func (s StoreMock) New(r *http.Request, name string) (*sessions.Session, error) {
	panic("implement me")
}

func (s StoreMock) Save(r *http.Request, w http.ResponseWriter, session *sessions.Session) error {
	panic("implement me")
}