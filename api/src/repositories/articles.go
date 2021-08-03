package repositories

import (
	"api/src/models"
	"database/sql"
)

type Article struct {
	db *sql.DB
}

//função que retorna instancia de Articles
func NewArticleRepository(db *sql.DB) *Article {
	return &Article{db}
}

//metodo que insere um novo artigo na tabela articles/recebe ID artigo;retorna id do artigo e erro
func (a Article) Create(newArticle models.Articles) (uint64, error) {
	statement, err := a.db.Prepare(
		"INSERT INTO articles (title,content,author_id) values(?,?,?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()
	result, err := statement.Exec(newArticle.Title, newArticle.Content, newArticle.AuthorID)
	if err != nil {
		return 0, err
	}
	articleID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(articleID), nil
}
