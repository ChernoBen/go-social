package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

//Insert an User in DB
func CreateUser(w http.ResponseWriter, r *http.Request) {
	//lendo corpo da requisição
	bodyContent, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//devolver error status entidade improcessavel
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	//inserir body na estrutura de usuário
	var user models.User
	if err = json.Unmarshal(bodyContent, &user); err != nil {
		//devolver um bad request
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	//verifica se dados de entrada são válidos
	if err := user.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	//abrir conexão com banco
	db, err := database.Connect()
	if err != nil {
		//devolver internal error
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	//apos aberta a con com banco, criar repo
	repo := repositories.NewUserRepository(db)
	//agora podemos chamar um metodo da struct dentro de repo
	user.ID, err = repo.Create(user)
	if err != nil {
		//devolver internal error
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.Json(w, http.StatusCreated, user)
	//w.Write([]byte(fmt.Sprintf("User inserted ID:%d", userID)))
}

//Get all User on DB /parametros name or nick ex:/users or /users?name=benja or /users?nick=benja
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// ler parametros passados na url e obter usuario
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))
	// a partir desse ponto posso abrir a conexão com o banco
	db, err := database.Connect()
	if err != nil {
		//devolver internal error
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	//instanciar repo para obter usuario
	repo := repositories.NewUserRepository(db)
	// executar busca
	users, err := repo.FindByNameOrNick(nameOrNick)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.Json(w, http.StatusOK, users)

}

//Get one user by ID
func GetUserById(w http.ResponseWriter, r *http.Request) {
	//obter parametros da rota
	params := mux.Vars(r)
	//obtendo id e convertendo para Uint64
	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	//abrir conexão com banco
	db, err := database.Connect()
	if err != nil {
		//devolver internal error
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	//intanciar repo passando a conexão com db
	repo := repositories.NewUserRepository(db)
	//executar busca por id
	user, err := repo.FindByID(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.Json(w, http.StatusOK, user)
}

//Update an User by passing his ID
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating an User"))
}

//Delete an User by passing his ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting an User"))
}
