package main

import (
	"fmt"

	"github.com/Carlos-Lopes1985/go-rest-api/configs"
	"github.com/Carlos-Lopes1985/go-rest-api/routes"
)

func main() {

	fmt.Println("Iniciando o servidor Rest com Go...")

	err := configs.Load()

	if err != nil {
		panic(err)
	}

	routes.HandleRequest()

}
