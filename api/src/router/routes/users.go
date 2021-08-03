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
	}, {
		URI:            "/users/{id}/follow",
		Method:         http.MethodPost,
		Function:       controllers.FollowUser,
		Authentication: true,
	}, {
		URI:            "/users/{id}/unfollow",
		Method:         http.MethodPost,
		Function:       controllers.Unfollow,
		Authentication: true,
	}, {
		URI:            "/users/{id}/followers",
		Method:         http.MethodGet,
		Function:       controllers.Followers,
		Authentication: false,
	}, {
		URI:            "/users/{id}/following",
		Method:         http.MethodGet,
		Function:       controllers.Following,
		Authentication: false,
	}, {
		URI:            "/users/{id}/password",
		Method:         http.MethodPost,
		Function:       controllers.UpdatePassword,
		Authentication: true,
	},
}
