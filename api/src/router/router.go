package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

//função que gera e retorna router contendo as rotas da api
func Gen() *mux.Router {
	r := mux.NewRouter()
	//func config recebe instacia de mux, carrega internamente as rotas existentes e retorna mux carregado
	return routes.Config(r)
}
