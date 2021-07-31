package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//obtendo variavel API_PORT
	config.Load()
	fmt.Println("Server running...")
	r := router.Gen()

	//definições do servidor
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ApiPort), r))
}
