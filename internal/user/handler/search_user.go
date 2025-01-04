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
		var req SearchRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Не получилось найти пользоваеля", http.StatusBadRequest)
			lg.Error(err)
			return
		}
		users, err := uc.Search(r.Context(), req.Nickname)
		if err != nil {
			lg.Error(err)
			http.Error(w, "Не Удалось найти пользоателя", http.StatusInternalServerError)
			return
		}
		var response SearchResponse
		for _, user := range users {
			response.User = append(response.User, UserDTO{
				ID:         user.ID,
				Nicknmae:   user.Nickname,
				Created_at: user.CreatedAt.Format("2006-01-02 15:04:05"),
			})
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
