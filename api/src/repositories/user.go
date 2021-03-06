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

//metodo de User que recebe um id e retorna um usuario e um erro
func (u User) FindByID(ID uint64) (models.User, error) {
	//criando query / struct user.db.query
	lines, err := u.db.Query(
		"SELECT id,name,email,createdat FROM users WHERE id = ?",
		ID,
	)
	if err != nil {
		//retornar uma instancia vazia
		return models.User{}, err
	}
	//fechando
	defer lines.Close()

	var user models.User
	// se tive uma linha a ser lida sera passada
	if lines.Next() {
		//ler a linha passando o endereco de memoria dos campos
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}

	}
	return user, nil
}

//metodo que recebe um id e struct com dados de usuario, atualiza usuario e retorna um err
func (u User) Update(ID uint64, user models.User) error {
	//criar statement
	statement, err := u.db.Prepare(
		"UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	//fechar statement
	defer statement.Close()
	//executar statement
	if _, err := statement.Exec(user.Name, user.Nick, user.Email, ID); err != nil {
		return err
	}
	return nil
}

// metodo que recebe um id, deleta usuario e retorna um erro
func (u User) Delete(ID uint64) error {
	//criar statement
	statement, err := u.db.Prepare(
		"DELETE FROM users WHERE id = ?",
	)
	if err != nil {
		return err
	}
	//fechar
	defer statement.Close()
	//executar
	if _, err = statement.Exec(ID); err != nil {
		return err
	}
	return nil
}

//funcao que recebe um email.string e retorna struct user e error
func (u User) FindByEmail(email string) (models.User, error) {
	line, err := u.db.Query(
		"SELECT id,password FROM users WHERE email = ?",
		email,
	)
	if err != nil {
		return models.User{}, err
	}
	defer line.Close()
	var user models.User
	if line.Next() {
		if err = line.Scan(
			&user.ID,
			&user.Password,
		); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

//metodo permite um usuario seguir outro /
func (u User) Follow(userID, FollowerID uint64) error {
	//clausula ignore: nao insere um dado na tabela caso ja exista
	statement, err := u.db.Prepare(
		"INSERT IGNORE INTO followers (user_id, follower_id) VALUES (?, ?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err := statement.Exec(FollowerID, userID); err != nil {
		return err
	}
	return nil
}

//metodo permite que um usuario pare de seguir outro
func (u User) Unfollow(userID, FollowerID uint64) error {
	statement, err := u.db.Prepare(
		"DELETE FROM followers WHERE user_id = ? AND follower_id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(userID, FollowerID); err != nil {
		return err
	}
	return nil
}

//metodo que retorna seguidores de determinado id informado
func (u User) Followers(userID uint64) ([]models.User, error) {
	//selecione/retorne na/da tabela USUARIOS valores onde o u.id é igual a f.follower_id
	lines, err := u.db.Query(
		`SELECT u.id, u.name,u.nick,u.email,u.createdat FROM users u INNER JOIN
		followers f on u.id = f.follower_id WHERE f.user_id = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()
	var followers []models.User
	for lines.Next() {
		var user models.User
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		followers = append(followers, user)
	}
	return followers, nil
}

// metodo que retorna lista de user que se está seguind
func (u User) Following(userID uint64) ([]models.User, error) {
	lines, err := u.db.Query(
		`SELECT u.id,u.name,u.nick,u.email,u.createdat FROM users u INNER JOIN
		followers f on u.id = f.user_id WHERE f.follower_id = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()
	var following []models.User
	for lines.Next() {
		var user models.User
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		following = append(following, user)
	}
	return following, nil
}

//metodo que busca hashpass salvo na tabeal users
func (u User) FindHashPass(userID uint64) (string, error) {
	line, err := u.db.Query(
		"SELECT password FROM users WHERE id = ?",
		userID,
	)
	if err != nil {
		return "", err
	}
	defer line.Close()
	var user models.User
	if line.Next() {
		if err = line.Scan(&user.Password); err != nil {
			return "", err
		}
	}
	return user.Password, nil
}

//metodo que atualiza a senha de um usuario/ 2 parametros id:uint64 hashPassword:string
func (u User) UpdatePassword(userID uint64, newPassword string) error {
	statement, err := u.db.Prepare(
		"UPDATE users SET password = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(newPassword, userID); err != nil {
		return err
	}
	return nil
}
