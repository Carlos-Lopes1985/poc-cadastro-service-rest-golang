package main

import (
	"fmt"
	"net/http"

	"github.com/Carlos-Lopes1985/go-rest-api/configs"
	"github.com/Carlos-Lopes1985/go-rest-api/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")

	err := configs.Load()

	fmt.Println("Valor arquivo de configuração", configs.Load())

	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Get("/milhas/{cpf}", handlers.GetTotalMilhasCliente)

	r.Post("/milhas", handlers.Create)

	http.ListenAndServe(":3000", r)

}
