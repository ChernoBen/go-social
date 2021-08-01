package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//uri mysql
	DbUri = ""
	//api port
	ApiPort = 0
	//chave passa assinar o token
	Secret []byte
)

//função inicializa variaveis de ambiente
func Load() {
	var err error
	//carregando variaveis do arquivo .env
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	//obtendo api_port do arquivo env e transformado em int
	ApiPort, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		// se err entao defina porta padrão como...
		ApiPort = 9000
	}
	//carregando e formatando uri de conexão
	DbUri = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	Secret = []byte(os.Getenv("SECRET_KEY"))
}
