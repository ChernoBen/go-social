package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

//definição e representação da estrutura das rotas
type Route struct {
	URI            string
	Method         string
	Function       func(w http.ResponseWriter, r *http.Request)
	Authentication bool
}

// Authentication: requer autenticação?

//funcao que retorna router com rotas já configuradas
func Config(r *mux.Router) *mux.Router {
	//carregar as rotas de usuarios
	routes := userRoutes
	//adicionando novas rotas
	routes = append(routes, loginRoute)
	// detructuring operator "..."
	routes = append(routes, articleRoutes...)
	//usar handlefunc para configurar cada rota
	for _, route := range routes {
		if route.Authentication {
			r.HandleFunc(route.URI, middlewares.Logger(
				middlewares.Authenticate(route.Function),
			)).Methods(route.Method)
		} else {
			//para cada rota em routes crie:
			r.HandleFunc(route.URI, middlewares.Logger(
				route.Function,
			)).Methods(route.Method)
		}
	}
	return r
}
