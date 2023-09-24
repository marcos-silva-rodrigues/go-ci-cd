package main

import (
	"github.com/marcos-silva-rodrigues/go-ci-cd/database"
	"github.com/marcos-silva-rodrigues/go-ci-cd/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
