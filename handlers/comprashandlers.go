package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Carlos-Lopes1985/go-rest-api/models"
	"github.com/go-chi/chi/v5"
)

func Create(w http.ResponseWriter, r *http.Request) {

	var compra models.Compra

	err := json.NewDecoder(r.Body).Decode(&compra)

	if err != nil {
		log.Printf("Erro ao fazer o decode json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	id, err := models.Insert(compra)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Compra inserida com sucesso: %v", id),
		}
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// Compras
var compras []models.Compra
var milhas models.Milhas

func GetTotalMilhasCliente(w http.ResponseWriter, r *http.Request) {

	cpf := chi.URLParam(r, "cpf")

	log.Printf("CPF GETTOTALMILHAS: %v", cpf)

	//if err != nil {
	//	log.Printf("Erro ao buscar registro de total de milhas por cpf: %d", err)
	//	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	//	return
	//}

	milhas = models.ReturnCalculoMilhas(string(cpf))

	log.Printf("OBJETO MILHAS: %v", milhas)

	// if err != nil {
	// 	log.Printf("Erro ao fazer o decode json: %v", err)
	// 	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	// 	return
	// }

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(milhas)
}
