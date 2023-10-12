package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/vinicanuto/go_rest_api/controllers"
	"github.com/vinicanuto/go_rest_api/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home).Methods("GET")
	r.HandleFunc("/api/personalidades", controllers.ListaPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.GetPersonalidadeById).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaPersonalidadeById).Methods("DELETE")
	r.HandleFunc("/api/personalidades", controllers.InserePersonalidade).Methods("POST")
	r.HandleFunc("/api/personalidades/{id}", controllers.AtualizaPersonalidadeById).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"http://gorest-api:3000"}))(r)))
}
