package routes

import (
	"api/src/controllers"
	"net/http"
)

var articleRoutes = []Route{
	{
		URI:            "/articles",
		Method:         http.MethodPost,
		Function:       controllers.CreateArticle,
		Authentication: true,
	}, {
		URI:            "/articles",
		Method:         http.MethodGet,
		Function:       controllers.ListArticles,
		Authentication: true,
	}, {
		URI:            "/articles/{id}",
		Method:         http.MethodGet,
		Function:       controllers.GetArticle,
		Authentication: true,
	}, {
		URI:            "/articles/{id}",
		Method:         http.MethodPut,
		Function:       controllers.UpdateArticle,
		Authentication: true,
	}, {
		URI:            "/articles/{id}",
		Method:         http.MethodDelete,
		Function:       controllers.DeleteArticle,
		Authentication: true,
	},
}
