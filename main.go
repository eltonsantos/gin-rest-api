package main

import (
	"github.com/eltonsantos/gin-rest-api/database"
	"github.com/eltonsantos/gin-rest-api/routes"
)

func main() {
	database.ConectaComBancoDeDados()

	routes.HandleRequest()
}
