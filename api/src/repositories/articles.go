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

//metodo um id e retorna um article e um erro
func (a Article) FindByID(articleID uint64) (models.Articles, error) {
	line, err := a.db.Query(
		`SELECT a.*,u.nick FROM articles a INNER JOIN users u ON u.id = a.author_id WHERE a.id = ?`,
		articleID,
	)
	if err != nil {
		return models.Articles{}, err
	}
	defer line.Close()
	var article models.Articles
	if line.Next() {
		if err = line.Scan(
			&article.ID,
			&article.Title,
			&article.Content,
			&article.AuthorID,
			&article.Likes,
			&article.CreatedAt,
			&article.AuthorNick,
		); err != nil {
			return models.Articles{}, err
		}
	}
	return article, nil
}

//metodo que recebe um ID uint64 e retorna slice de articles e um error/
func (a Article) FindArticlesByUser(userID uint64) ([]models.Articles, error) {
	lines, err := a.db.Query(
		`SELECT DISTINCT a.*, u.nick FROM articles a 
		INNER JOIN users u ON u.id = a.author_id 
		INNER JOIN followers f ON a.author_id = f.user_id 
		WHERE u.id = ? OR f.follower_id = ? ORDER BY 1 DESC`,
		userID, userID,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()
	var articles []models.Articles
	for lines.Next() {
		var article models.Articles
		if err = lines.Scan(
			&article.ID,
			&article.Title,
			&article.Content,
			&article.AuthorID,
			&article.Likes,
			&article.CreatedAt,
			&article.AuthorNick,
		); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil

}
