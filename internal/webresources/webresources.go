package webresources

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// ErrorDTO message
type ErrorDTO struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func readBodyFromJSON(w http.ResponseWriter, r *http.Request, entity interface{}) error {
	if r.Header.Get("Content-Type") != "application/json" {
		handleError(w, "Content-Type not supported", http.StatusUnsupportedMediaType)
		return fmt.Errorf("Content-Type not supported")
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("Error when read request body: %v\n", err)
		handleError(w, "Erro when try read body", http.StatusInternalServerError)
		return fmt.Errorf("Erro when try read body")
	}

	if !json.Valid(body) {
		handleError(w, "Bad Request", http.StatusBadRequest)
		return fmt.Errorf("Bad Request")
	}

	json.Unmarshal(body, entity)

	return nil
}

func handleError(w http.ResponseWriter, message string, httpCode int) {
	errorDto := ErrorDTO{
		Message: message,
		Code:    httpCode,
	}

	writeError(w, errorDto)
}

func writeError(w http.ResponseWriter, errorDto ErrorDTO) {
	w.WriteHeader(errorDto.Code)
	json.NewEncoder(w).Encode(errorDto)
}
