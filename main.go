package main

import (
	"fmt"

	"github.com/ItaloCobains/api-rest-go/database"
	"github.com/ItaloCobains/api-rest-go/routes"
)

func main() {

	database.ConectaComBancoDeDados()

	fmt.Println("Starting api rest with go")
	routes.HandleResponse()
}
