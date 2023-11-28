//54:00
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/aprendagolang/api-pgsql/models"
	"github.com/go-chi/chi/v5"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao Fazer parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Delete(int64(id))

	if err != nil {
		log.Printf("Erro ao Remover o Registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if rows > 1 {

		log.Printf("Erro: Foram Removidos %d registros", rows)

	}

	resp := map[string]any{

		"Error":   false,
		"Message": "Registro Removido  com sucesso",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
