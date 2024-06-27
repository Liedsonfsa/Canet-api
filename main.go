package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	fmt.Println("Rodando a API.")

	r := router.Gerar()
	fmt.Println(r)

	fmt.Println(config.StringConexao)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%d:", config.Porta), r))
}