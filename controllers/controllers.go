package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Carlos-Lopes1985/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func GetArticles(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint Hit: GetArticles")

	json.NewEncoder(w).Encode(models.ReturnAllArticles())
}

func GetMilhas(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint Hit: GetMilhas")

	json.NewEncoder(w).Encode(models.ReturnTotalMilhas())
}
