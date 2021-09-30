package main

import (
	"github.com/gorilla/sessions"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("asd"))
type WithSession struct {
	NextHandler func(w http.ResponseWriter, r *http.Request, session *sessions.Session)
	store sessions.Store
}

func main(){
	mux := initMux(store)
	http.ListenAndServe(":8090", mux)
}

func initMux(store sessions.Store) * http.ServeMux{
	mux := &http.ServeMux{}
	mux.Handle("/session", WithSession{LogedIn, store})

	return mux
}

func LogedIn(w http.ResponseWriter, request *http.Request, sesion *sessions.Session) {
	if sesion.Values["username"] == "test" {
		w.WriteHeader(http.StatusOK)
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (w WithSession) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	session, _ := w.store.Get(request, "session-name")
	w.NextHandler(writer, request, session)
}





