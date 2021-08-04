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

//List all user articles
func UserArticles(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	userID, err := strconv.ParseUint(param["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	var userArticles []models.Articles
	repo := repositories.NewArticleRepository(db)
	userArticles, err = repo.FindUserArticles(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.Json(w, http.StatusOK, userArticles)
}

//List articles/lista publições de proprio usuario e de quem ele segue
func ListArticles(w http.ResponseWriter, r *http.Request) {
	tokenID, err := authentication.GetID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := repositories.NewArticleRepository(db)
	var articles []models.Articles
	articles, err = repo.FindArticlesByUser(tokenID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.Json(w, http.StatusOK, articles)
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
	tokenID, err := authentication.GetID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	//id da publicação a ser editada
	param := mux.Vars(r)
	articleID, err := strconv.ParseUint(param["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := repositories.NewArticleRepository(db)
	//verificar se o artigo pertence ao usuario
	actualArticle, err := repo.FindByID(articleID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	if actualArticle.AuthorID != tokenID {
		responses.Error(w, http.StatusForbidden, errors.New("Can not update an article that is not yours"))
		return
	}
	// obter nova edição do body
	bodyContent, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	// fortatar conteudo do body p/ tipo model.Article
	var newArticle models.Articles
	if err = json.Unmarshal(bodyContent, &newArticle); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	//verificar e formatar novo modelo de artigo
	if err = newArticle.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	if err = repo.UpdateArticle(articleID, newArticle); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.Json(w, http.StatusNoContent, nil)

}

//Delete Article by ID
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	tokenID, err := authentication.GetID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	param := mux.Vars(r)
	articleID, err := strconv.ParseUint(param["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := repositories.NewArticleRepository(db)
	var article models.Articles
	article, err = repo.FindByID(articleID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	if article.AuthorID != tokenID {
		responses.Error(w, http.StatusForbidden, errors.New("Can not delete an article that is not yours"))
		return
	}
	if err = repo.DeleteArticle(articleID, tokenID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.Json(w, http.StatusNoContent, nil)
}

//Like article
func Like(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	articleID, err := strconv.ParseUint(param["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := repositories.NewArticleRepository(db)
	if err = repo.LikeArticle(articleID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.Json(w, http.StatusNoContent, nil)
}
