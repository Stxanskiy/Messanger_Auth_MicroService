package handler

import (
	"auth_sevice_microservice/internal/user/uc"
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

type RegisterRequest struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegiterRespone struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

func HeakthCHeckHandler(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := db.Ping(context.Background()); err != nil {
			w.Write([]byte(`{"status":"error","db":"unavailable"}`))
			return
		} else if err == nil {
			w.Write([]byte(`{"status OK":"База данных работает корректно"` + "\n"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	}
}

func RegisterHandler(uc *uc.UserUC) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid Requst", http.StatusBadRequest)
			return
		}

		//регистрация пользователя
		user, err := uc.RefisterUser(r.Context(), req.Nickname, req.Email, req.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//возвращаем резульат
		resp := RegiterRespone{
			ID:       user.ID,
			Nickname: user.Nickname,
			Email:    user.Email,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resp)

	}

}
