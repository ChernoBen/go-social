package database

//lembrar de importar manualmente de forma implicita sql driver
import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//função que retorna uma instancia de conexão mysql e um erro
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DbUri)
	if err != nil {
		return nil, err
	}

	//se algum erro na conexão...
	if err = db.Ping(); err != nil {
		//...feche a conexão e retorne vazio e o erro
		db.Close()
		return nil, err
	}
	return db, nil
}
