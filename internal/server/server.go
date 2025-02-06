package server

import (
	"auth_sevice_microservice/config"
	http2 "auth_sevice_microservice/internal/user/delivery/http"
	"auth_sevice_microservice/pkg/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"gitlab.com/nevasik7/lg"
	"net/http"
)

func Run() error {
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		lg.Errorf("Нет данного файла в указанной директории, %v", err)
	}
	db, err := database.NewPostgresDB(
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.SSLMode,
	)
	if err != nil {
		lg.Errorf("Не удалось загрузить конфигурацию: %v", err)
	}
	defer db.Close()

	//
	r := chi.NewRouter()

	// CORS middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://your-frontend.com"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	http2.RegisterRoutes(r, db)

	//запуск сервера
	address := "127.0.0.1:" + cfg.Server.Port
	lg.Infof("Сервер запущен по адресу http://" + address)
	return http.ListenAndServe(address, r)
}
