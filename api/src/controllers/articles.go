package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Create a new article
func CreateArticle(w http.ResponseWriter, r *http.Request) {
	tokenID, err := authentication.GetID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	bodyContent, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	var newArticle models.Articles
	if err = json.Unmarshal(bodyContent, &newArticle); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	if err = newArticle.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	newArticle.AuthorID = tokenID
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	//repositorio de publicações
	repo := repositories.NewArticleRepository(db)
	articleID, err := repo.Create(newArticle)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.Json(w, http.StatusCreated, articleID)
}

//List articles
func ListArticles(w http.ResponseWriter, r *http.Request) {

}

//Get Article by ID
func GetArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	articleID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
	}
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	var article models.Articles
	repo := repositories.NewArticleRepository(db)
	article, err = repo.FindByID(articleID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.Json(w, http.StatusOK, article)
}

//Update Article by ID
func UpdateArticle(w http.ResponseWriter, r *http.Request) {

}

//Delete Article by ID
func DeleteArticle(w http.ResponseWriter, r *http.Request) {

}
