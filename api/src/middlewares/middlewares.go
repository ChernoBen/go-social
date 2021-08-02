package middlewares

import (
	"api/src/authentication"
	"api/src/responses"
	"log"
	"net/http"
)

//func que gera logs no terminal das requisições executadas
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

//funcao que valida usuario/ recebe e retorna um handler
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//chama o validador
		if err := authentication.Validate(r); err != nil {
			responses.Error(w, http.StatusUnauthorized, err)
			return
		}
		//se valido então permite que o usuario prossiga
		next(w, r)
	}
}
