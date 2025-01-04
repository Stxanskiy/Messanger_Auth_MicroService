package handler

import (
	"auth_sevice_microservice/internal/user/uc"
	"encoding/json"
	"net/http"
)

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func LogoutHandler(uc *uc.LogoutUC) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LogoutRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		//выполнение logout
		if err := uc.Logout(r.Context(), req.RefreshToken); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Logout successful"}`))
	}

}
