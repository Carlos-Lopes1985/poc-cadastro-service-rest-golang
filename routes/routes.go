package routes

import (
	"net/http"

	"github.com/Carlos-Lopes1985/go-rest-api/handlers"
)

func HandleRequest() {

	//r := chi.newRouter()
	//r.Post("/", handlers.Create)

	//	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)

	http.HandleFunc("/", handlers.Create)
	//	http.HandleFunc("/articles", controllers.GetArticles)
	//http.HandleFunc("/milhas", controllers.GetMilhas)
	//log.Fatal(http.ListenAndServe(":8085", nil))
}
