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
	r := router.Gen()
	fmt.Printf("Server running at port:%d...", config.ApiPort)
	//definições do servidor
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ApiPort), r))
}
