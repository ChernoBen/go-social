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

//metodo que atualiza dados de um article recebendo um articleID e um model Article e retorna um error/
func (a Article) UpdateArticle(articleID uint64, article models.Articles) error {
	statement, err := a.db.Prepare("UPDATE articles SET title = ?, content = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(article.Title, article.Content, articleID); err != nil {
		return err
	}
	return nil
}

//metodo que recebe um article id e user id/ retorna um error
func (a Article) DeleteArticle(articleID, userID uint64) error {
	statement, err := a.db.Prepare("DELETE FROM articles WHERE id =?")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(articleID); err != nil {
		return err
	}
	return nil
}

//metodo que recebe um id e retorna um slice de models.Articles e um err
func (a Article) FindUserArticles(userID uint64) ([]models.Articles, error) {
	lines, err := a.db.Query(
		`SELECT a.*,u.nick FROM articles a
		JOIN users u ON u.id = a.author_id
		WHERE a.author_id = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	var articles []models.Articles
	for lines.Next() {
		var item models.Articles
		if err = lines.Scan(
			&item.ID,
			&item.Title,
			&item.Content,
			&item.AuthorID,
			&item.Likes,
			&item.CreatedAt,
			&item.AuthorNick,
		); err != nil {
			return nil, err
		}
		articles = append(articles, item)
	}
	return articles, nil
}

//metodo que insere uma curtida na tabela articles
func (a Article) LikeArticle(articleID uint64) error {
	statement, err := a.db.Prepare("UPDATE articles SET likes = likes + 1 WHERE id = ? ")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(articleID); err != nil {
		return err
	}
	return nil
}
