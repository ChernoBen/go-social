package authentication

//importar json web token e atribuir um alias
import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

//funcao que valida token
func Validate(r *http.Request) error {
	token := getToken(r)
	tk, err := jwt.Parse(token, getKey)
	if err != nil {
		return err
	}
	//validar token
	if _, ok := tk.Claims.(jwt.MapClaims); ok && tk.Valid {
		return nil
	}
	return errors.New("Invalid Token")
}

//funcao local que extrai token da request e retorna uma string
func getToken(r *http.Request) string {
	// pegando Bearer token da requisição
	token := r.Header.Get("authorization")
	//processando e verificando token
	if len(strings.Split(token, " ")) == 2 {
		// se existir 2 palavras entao retorne a indice 1
		return strings.Split(token, " ")[1]
	}
	return ""
}

//funcao local que obtem metodo de assinatura do token
func getKey(token *jwt.Token) (interface{}, error) {
	//verifica metodo de assinatura
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected algoithm %s", token.Header["alg"])
	}
	return config.Secret, nil
}

//função que estrai ID do token/retorna um uint64 e um erro
func GetID(r *http.Request) (uint64, error) {
	token := getToken(r)
	tk, err := jwt.Parse(token, getKey)
	if err != nil {
		return 0, err
	}
	if permissions, ok := tk.Claims.(jwt.MapClaims); ok && tk.Valid {
		userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["id"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return userID, nil
	}
	return 0, errors.New("Invalid Token")
}
