package repositories

import (
	"api/src/models"
	"database/sql"
)

//conexão será aberta no controller e passada para essa struct
// e a partir dela é dada a interação com banco de dados
//User representação do repo de user
type User struct {
	db *sql.DB
}

//funcao que recebe uma conn aberta e retorna uma instancia de user
// com a conn instanciada
func NewUserRepository(db *sql.DB) *User {
	//dentro dessa struct terá os metodos para manipula o banco de dados
	return &User{db}
}

//metodo de User que insere um novo user no banco re retorna um (id)uint64 e um erro
func (u User) Create(user models.User) (uint64, error) {
	//preparando statement de inserção de usuario
	statement, err := u.db.Prepare(
		"INSERT INTO users (name,nick,email,password) VALUES(?,?,?,?)",
	)
	if err != nil {
		return 0, err
	}
	//fechar statement
	defer statement.Close()
	//executar o statement com os paramentos do usuario
	res, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}
	//retornar Id do usuario
	userID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	//retornar id int convertido para uint64
	return uint64(userID), nil
}
