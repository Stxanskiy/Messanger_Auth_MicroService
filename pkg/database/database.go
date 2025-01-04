package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/nevasik7/lg"
	"time"
)

func NewPostgresDB(host string, port int, user, password, dbname, sslmode string) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		user, password, host, port, dbname, sslmode)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("DSN ошибка строки плдключения %w", err)
	}
	config.MaxConns = 5
	config.ConnConfig.ConnectTimeout = 5 * time.Second

	db, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("Не удалось подключиться к базе данных %w", err)
	}

	//Проверка подклбчения
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.Ping(ctx); err != nil {
		return nil, fmt.Errorf("База даннвх недоступна %w", err)
	}

	lg.Infof("Подключение к базе данных по адресу %s", dsn)
	return db, nil

}
