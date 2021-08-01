package routes

import (
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
	//usar handlefunc para configurar cada rota
	for _, route := range routes {
		//para cada rota em routes crie:
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}
	return r
}
