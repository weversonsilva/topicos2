//55:40
package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aprendagolang/api-pgsql/models"
)

func List(w http.ResponseWriter, r *http.Request) {

	todos, err := models.GetAll()

	if err != nil {

		log.Printf("Erro ao Obter Registros: %v", err)

	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)

}
