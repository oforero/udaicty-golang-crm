package backend

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Url     string
	Handler http.HandlerFunc
	Method  string
}

func BuildRouter(routes []Route) *mux.Router {
	router := mux.NewRouter()
	for _, r := range routes {
		router.HandleFunc(r.Url, r.Handler).Methods(r.Method)
	}
	return router
}
