package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON seta o statuscode e transforma os dados em json
func JSON(w http.ResponseWriter, statusCode int, datas interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(datas); err != nil {
		log.Fatal(err)
	}
}

// Erro retorna um erro em formato JSON
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}