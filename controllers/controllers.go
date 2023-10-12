package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vinicanuto/go_rest_api/database"
	"github.com/vinicanuto/go_rest_api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func ListaPersonalidades(w http.ResponseWriter, r *http.Request) {
	var personalidades []models.Personalidade
	database.DB.Find(&personalidades)
	ResponseJson(w, personalidades)
}

func GetPersonalidadeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade []models.Personalidade
	database.DB.First(&personalidade, id)
	ResponseJson(w, personalidade)
}

func DeletaPersonalidadeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade []models.Personalidade
	database.DB.Delete(&personalidade, id)
	w.WriteHeader(200)
}

func AtualizaPersonalidadeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade

	database.DB.First(&personalidade, id)

	json.NewDecoder(r.Body).Decode(&personalidade)

	database.DB.Updates(&personalidade)
	ResponseJson(w, personalidade)
}

func InserePersonalidade(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade

	json.NewDecoder(r.Body).Decode(&novaPersonalidade)

	database.DB.Create(&novaPersonalidade)

	ResponseJson(w, novaPersonalidade)
}

func ResponseJson(writer http.ResponseWriter, any interface{}) {
	json.NewEncoder(writer).Encode(any)
}
