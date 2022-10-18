package main

import (
	"fmt"

	"github.com/ItaloCobains/api-rest-go/database"
	"github.com/ItaloCobains/api-rest-go/models"
	"github.com/ItaloCobains/api-rest-go/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 1", Historia: "Historia 1"},
	}

	database.ConectaComBancoDeDados()

	fmt.Println("Starting api rest with go")
	routes.HandleResponse()
}
