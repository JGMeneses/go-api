package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/JGMeneses/go-crud/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.Get_All()
	if err != nil {
		log.Println("Erro ao obter registros: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
