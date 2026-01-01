package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func WriteJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	encoder := json.NewEncoder(w)

	if err := encoder.Encode(data); err != nil {
		log.Println("json encode error: ", err)
	}
}

func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJson(w, status, ErrorResponse{
		Error:   http.StatusText(status),
		Message: message,
	})
}
