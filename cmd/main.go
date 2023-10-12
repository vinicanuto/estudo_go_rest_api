package main

import (
	"fmt"

	"github.com/vinicanuto/go_rest_api/database"
	"github.com/vinicanuto/go_rest_api/models"
	"github.com/vinicanuto/go_rest_api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Primeiro", Historia: "Historia 1"},
		{Id: 2, Nome: "Segundo", Historia: "Historia 2"},
	}
	database.ConnectWithDatabase()
	fmt.Println("Iniciando api")
	routes.HandleRequest()
}
