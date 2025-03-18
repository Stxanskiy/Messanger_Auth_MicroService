package repo

import (
	"context"
	"gitlab.com/nevasik7/lg"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RefreshTokenRepo struct {
	db *pgxpool.Pool
}

func NewRefreshTokenRepo(db *pgxpool.Pool) *RefreshTokenRepo {
	return &RefreshTokenRepo{db: db}
}

// сохранение токена в базу данных
func (r *RefreshTokenRepo) SaveToken(ctx context.Context, userID int, token string, expiresAt time.Time) error {
	query := QuerySaveToken
	_, err := r.db.Exec(ctx, query, userID, token, expiresAt)
	return err
}

// удаление токена перед обновлением токенов
func (r *RefreshTokenRepo) DeleteToken(ctx context.Context, token string) error {
	query := QueryDeleteToken
	_, err := r.db.Exec(ctx, query, token)
	if err != nil {
		lg.Errorf("произошла ошибка при удалении токена")
		return err
	}
	return err
}

// проверка валидности токена
func (r *RefreshTokenRepo) IsTokenValid(ctx context.Context, token string) (bool, error) {
	query := QueryIsTokenValid

	var isValid bool
	err := r.db.QueryRow(ctx, query, token).Scan(&isValid)
	return isValid, err
}

// удланеие токена пользователя при logout
func (r *RefreshTokenRepo) DeleteTokenByUserID(ctx context.Context, userID int) error {
	query := QueryDeleteTokenByUserID
	_, err := r.db.Exec(ctx, query, userID)
	if err != nil {
		lg.Errorf("Не удалось удалить токен")
		return err
	}
	return err
}
