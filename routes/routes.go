package routes

import (
	"fmt"
	"net/http"

	"github.com/Carlos-Lopes1985/go-rest-api/handlers"
	"github.com/go-chi/chi/v5"
)

func HandleRequest() {

	r := chi.NewRouter()

	r.Get("/milhas/{cpf}", handlers.GetTotalMilhasCliente)
	r.Post("/milhas", handlers.Create)

	fmt.Println("Servidor iniciado com sucesso!")

	http.ListenAndServe(":3000", r)
}
