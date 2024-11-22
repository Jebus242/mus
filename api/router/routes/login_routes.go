package routes

import (
	"net/http"

	"github.com/jebus24/mus/api/controllers"
)

var loginRoutes = []Route{
	Route{
		URI:          "/login",
		Method:       http.MethodPost,
		Handler:      controllers.Login,
		AuthRequired: false,
	},
}
