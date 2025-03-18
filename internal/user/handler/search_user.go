package handler

import (
	"auth_sevice_microservice/internal/user/uc"
	"encoding/json"
	"gitlab.com/nevasik7/lg"
	"net/http"
)

type SearchRequest struct {
	Nickname string `json:"nickname"`
}

type SearchResponse struct {
	User []UserDTO `json:"users"`
}

type UserDTO struct {
	ID         int    `json:"id"`
	Nicknmae   string `json:"nicknmae"`
	Created_at string `json:"created_at"`
}

func SearchUserHandler(uc *uc.SearchUsersUC) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Считываем параметр из URL: /users/search?nickname=john
		nickname := r.URL.Query().Get("nickname")
		if nickname == "" {
			http.Error(w, "nickname param is required", http.StatusBadRequest)
			return
		}

		users, err := uc.Search(r.Context(), nickname)
		if err != nil {
			lg.Error(err)
			http.Error(w, "Не удалось найти пользователя", http.StatusInternalServerError)
			return
		}

		// Формируем ответ
		var resp SearchResponse
		for _, user := range users {
			resp.User = append(resp.User, UserDTO{
				ID:         user.ID,
				Nicknmae:   user.Nickname,
				Created_at: user.CreatedAt.Format("2006-01-02 15:04:05"),
			})
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
