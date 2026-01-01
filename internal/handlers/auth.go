package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kripesh12/my-notes/internal/auth"
	"github.com/kripesh12/my-notes/internal/db"
	"github.com/kripesh12/my-notes/internal/handlers/dto"
	"github.com/kripesh12/my-notes/internal/response"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var requestBody dto.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		response.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := requestBody.Validate(); err != nil {
		response.WriteError(w, http.StatusCreated, err.Error())
		return
	}

	hashedPassword, err := auth.HashPassword(requestBody.Password)
	if err != nil {
		fmt.Println()
		response.WriteError(w, http.StatusInternalServerError, "failed to create a user")
		http.Error(w, "server error", 500)
		return
	}

	var userId int
	query := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id"

	err = db.DB.QueryRow(context.Background(), query, requestBody.Email, hashedPassword).Scan(&userId)
	if err != nil {
		if db.IsDuplicateError(err) {
			response.WriteError(w, http.StatusBadRequest, "user with this email already exist")
			return
		}
		log.Printf("database error during registration: %v", err)
		response.WriteJson(w, http.StatusInternalServerError, "failed to create user")
	}

	response.WriteJson(w, http.StatusCreated, dto.RegisterResponse{
		ID:    userId,
		Email: requestBody.Email,
	})

}
