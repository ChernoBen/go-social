package authentication

//importar json web token e atribuir um alias
import (
	"api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//função que gera token
func GenToken(userID uint64) (string, error) {
	//definir permissoes do token
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	//definindo tempo de expiração
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["id"] = userID
	// gerar token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	//assinar token
	return token.SignedString(config.Secret)
}
