package routes

import (
	"net/http"

	"github.com/jebus24/mus/api/controllers"
)

var usersRoutes = []Route{
	Route{
		URI:          "/users",
		Method:       http.MethodGet,
		Handler:      controllers.GetUsers,
		AuthRequired: false,
	},
	Route{
		URI:          "/users",
		Method:       http.MethodPost,
		Handler:      controllers.CreateUser,
		AuthRequired: false,
	},
	Route{
		URI:          "/users/{id}",
		Method:       http.MethodGet,
		Handler:      controllers.GetUser,
		AuthRequired: false,
	},
	Route{
		URI:          "/users/{id}",
		Method:       http.MethodPut,
		Handler:      controllers.UpdateUser,
		AuthRequired: true,
	},
	Route{
		URI:          "/users/{id}",
		Method:       http.MethodDelete,
		Handler:      controllers.DeleteUser,
		AuthRequired: true,
	},
}
