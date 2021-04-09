package transport

import (
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Router() http.Handler  {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/hello-world", helloWorld).Methods(http.MethodGet)
	return logMiddleware(r)
}

func logMiddleware(h http.Handler) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.WithFields(log.Fields {
			"method": r.Method,
			"url": r.URL,
			"remoteAddr": r.RemoteAddr,
			"userAgent": r.UserAgent(),
		}).Info("got a new request")
		h.ServeHTTP(w, r)
	})
}

func helloWorld(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello world!")
}