package handler

import (
	"auth_sevice_microservice/internal/user/uc"
	"encoding/json"
	"net/http"

	"time"
)

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    string `json:"expires_at"`
}

func RefreshHandler(useCase *uc.TokenRefreshUC) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RefreshRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Обновление токенов
		access, refresh, err := useCase.RefreshTokens(r.Context(), req.RefreshToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		resp := RefreshResponse{
			AccessToken:  access.ID,
			RefreshToken: refresh.ID,
			ExpiresAt:    access.ExpiresAt.Time.Format(time.RFC3339),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}
