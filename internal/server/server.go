package server

import (
	"auth_sevice_microservice/config"
	http2 "auth_sevice_microservice/internal/user/delivery/http"
	"auth_sevice_microservice/pkg/database"
	"github.com/go-chi/chi/v5"
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

	http2.RegisterRoutes(r, db)

	//запуск сервера
	address := "127.0.0.1:" + cfg.Server.Port
	lg.Infof("Сервер запущен по адресу http://" + address)
	return http.ListenAndServe(address, r)
}
