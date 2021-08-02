package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
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
	if err := user.Prepare("register"); err != nil {
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

//Update an User by passing his ID / este metodo nao deve atualizar a senha do usuario
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	//primeiro obter id no parametro da rota
	param := mux.Vars(r)
	//obter id e converter p/ Uint64
	userID, err := strconv.ParseUint(param["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	tokenID, err := authentication.GetID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	if userID != tokenID {
		responses.Error(w, http.StatusForbidden, err)
		return
	}
	//obter dados do corpo da request
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	//inserir corpo da requisição dentro da estrutura models.Usuario
	var user models.User
	if err = json.Unmarshal(reqBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	//tratar dados
	if err = user.Prepare("edit"); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	// abrir conexão com banco
	db, err := database.Connect()
	if err != nil {
		//devolver internal error
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	//intanciar repo passando conexão com banco
	repo := repositories.NewUserRepository(db)
	if err = repo.Update(userID, user); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.Json(w, http.StatusNoContent, nil)
}

//Delete an User by passing his ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	//obter parametros da rota
	params := mux.Vars(r)
	// obter id do ususario e converter para Uint64
	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	tokenID, err := authentication.GetID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	if tokenID != userID {
		responses.Error(w, http.StatusForbidden, err)
		return
	}
	//abrir conexao
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	//fechando con
	defer db.Close()
	//instanciar repo passando conexão com db
	repo := repositories.NewUserRepository(db)
	if err = repo.Delete(userID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.Json(w, http.StatusNoContent, nil)
}

//Follow User allows that an user follow/id no parametro é quem será seguido
func FollowUser(w http.ResponseWriter, r *http.Request) {
	// extrair id do token
	follower, err := authentication.GetID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	//obter id do parametro
	params := mux.Vars(r)
	paramID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	if follower == paramID {
		responses.Error(w, http.StatusForbidden, errors.New("Can not follow your own account"))
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	repo := repositories.NewUserRepository(db)
	if err = repo.Follow(follower, paramID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.Json(w, http.StatusNoContent, nil)

}
