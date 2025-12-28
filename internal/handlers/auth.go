package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/kripesh12/my-notes/internal/auth"
	"github.com/kripesh12/my-notes/internal/db"
	"github.com/kripesh12/my-notes/internal/handlers/dto"
	"github.com/kripesh12/my-notes/internal/response"
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

	hashedPassword, err := auth.HashPassword(requestBody.Password)
	if err != nil {
		http.Error(w, "server error", 500)
	}

	var userId int
	query := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id"

	err = db.DB.QueryRow(context.Background(), query, requestBody.Email, hashedPassword).Scan(&userId)
	if db.IsDuplicateError(err) {
		http.Error(w, "user with this email already exist", http.StatusBadRequest)
		return
	}

	response.WriteJson(w, http.StatusCreated, dto.RegisterResponse{
		ID:    userId,
		Email: requestBody.Email,
	})
}
