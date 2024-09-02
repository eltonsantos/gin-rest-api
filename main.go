package main

import (
	"github.com/eltonsantos/gin-rest-api/database"
	"github.com/eltonsantos/gin-rest-api/models"
	"github.com/eltonsantos/gin-rest-api/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "Elton Santos", CPF: "123456789", RG: "111222333444555"},
		{Nome: "Ana Santos", CPF: "78654321", RG: "999888777666555"},
	}
	routes.HandleRequest()
}
