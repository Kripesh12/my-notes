package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)

	w.WriteHeader(status)
	if err := encoder.Encode(data); err != nil {
		log.Println("json encode error: ", err)
	}

}
