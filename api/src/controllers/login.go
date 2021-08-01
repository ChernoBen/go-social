package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//função que lida com a requisição de autenticação
func Login(w http.ResponseWriter, r *http.Request) {
	//obter corpo da requisição
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	// obter dados do body
	var user models.User
	if err = json.Unmarshal(reqBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	// abrir conexão com db
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	//fechar
	defer db.Close()
	//instanciar repo de usuario passando db conn
	repo := repositories.NewUserRepository(db)
	userData, err := repo.FindByEmail(user.Email)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	// verificando senha da entrada com hash no banco
	if err = security.Verify(user.Password, userData.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	//gerando token
	token, err := authentication.GenToken(userData.ID)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
	}
	responses.Json(w, http.StatusCreated, token)
}
