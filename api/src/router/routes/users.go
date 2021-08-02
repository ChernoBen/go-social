package routes

import (
	"api/src/controllers"
	"net/http"
)

//definição do slice contendo/representado todas as rotas de user
var userRoutes = []Route{
	{
		URI:            "/users",
		Method:         http.MethodPost,
		Function:       controllers.CreateUser,
		Authentication: false,
	},
	{
		URI:            "/users",
		Method:         http.MethodGet,
		Function:       controllers.GetUsers,
		Authentication: false,
	}, {
		URI:            "/users/{id}",
		Method:         http.MethodGet,
		Function:       controllers.GetUserById,
		Authentication: true,
	}, {
		URI:            "/users/{id}",
		Method:         http.MethodPut,
		Function:       controllers.UpdateUser,
		Authentication: true,
	}, {
		URI:            "/users/{id}",
		Method:         http.MethodDelete,
		Function:       controllers.DeleteUser,
		Authentication: false,
	},
}
