package router

import (
	"github.com/gorilla/mux"
	"github.com/jebus24/mus/api/router/routes"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutesWithMiddlewares(r)
}
