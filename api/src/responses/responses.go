package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

//função que retornar respostas em formato json
func Json(w http.ResponseWriter, statusCode int, data interface{}) {
	//setar formato do conteudo como json/ dessa maneira a respostas serão entregues em json
	w.Header().Set("Content-Type", "application/json")
	//passar o status code na response
	w.WriteHeader(statusCode)
	//verificar se data pertence a um tipo de response 'no content'
	if data != nil {
		//criar json contendo as informações
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}

}

//função que retornar respostas em formato json para casos de erros
func Error(w http.ResponseWriter, statusCode int, err error) {
	Json(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}
