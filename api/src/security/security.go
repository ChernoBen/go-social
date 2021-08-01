package security

import "golang.org/x/crypto/bcrypt"

//processo de criar um hash da senha dps comparar a senha inserida com o hash guardado no db

//função que recebe uma string e gera um hash a partir dela/ retorna um slice de bytes  um error
func Hash(password string) ([]byte, error) {
	//gerar um hsh a partir de uma senha/parametros: slice tipo byte e custo
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//funcao que 2 strings verifica se o valor de uma primeira string é igual ao de um determinado hashString/ retorna um erro
func Verify(password, hashString string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashString), []byte(password))
}
