package http

import (
	"auth_sevice_microservice/internal/user/handler"
	repo2 "auth_sevice_microservice/internal/user/repo"
	uc2 "auth_sevice_microservice/internal/user/uc"
	"auth_sevice_microservice/pkg/jwt"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

func RegisterRoutes(r chi.Router, db *pgxpool.Pool) {
	repo := repo2.NewUserRepo(db)

	uc := uc2.NewUserUC(repo)
	refreshRepo := repo2.NewRefreshTokenRepo(db)
	jwtManager := jwt.NewJWTManager("salt_secret", 15*time.Minute, 24*time.Hour)

	loginUC := uc2.NewLoginUC(repo, jwtManager, refreshRepo)
	refreshUC := uc2.NewTokenRefreshTokenUC(refreshRepo, jwtManager)
	logoutUC := uc2.NewLogoutUC(refreshRepo)

	r.Route("/users", func(r chi.Router) {
		r.Post("/register", handler.RegisterHandler(uc))
		r.Post("/login", handler.LoginHandler(loginUC))
		r.Post("/refresh", handler.RefreshHandler(refreshUC))
		r.Post("/logout", handler.LogoutHandler(logoutUC))
	})

	r.Route("/health_service_check", func(r chi.Router) {
		r.Get("/health", handler.HeakthCHeckHandler(db))

	})

}
