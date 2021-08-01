package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

//metodo de User que recebe uma string e retorna um slice de models.User e um erro/busca usuarios 'like' nameOrNick
func (u User) FindByNameOrNick(nameOrNick string) ([]models.User, error) {
	//formatando string para padrões de consulta
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)
	// criando query / struct user.db.query
	lines, err := u.db.Query(
		"SELECT id,name,nick,email,createdat FROM users WHERE name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)
	if err != nil {
		return nil, err
	}
	// fechar
	defer lines.Close()
	//iterar lines e armazenar cada item em um slice do tipo models.user
	var users []models.User
	for lines.Next() {
		var item models.User
		if err = lines.Scan(
			&item.ID,
			&item.Name,
			&item.Nick,
			&item.Email,
			&item.CreatedAt,
		); err != nil {
			return nil, err
		}
		//adicionar item ao slice
		users = append(users, item)
	}
	return users, nil
}
