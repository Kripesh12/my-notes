package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kripesh12/my-notes/internal/auth"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "invalid request", 400)
		return
	}

	hashedPassword, err := auth.HashPassword(requestBody.Email)
	if err != nil {
		http.Error(w, "server error", 500)
	}
	fmt.Println(hashedPassword)
}
