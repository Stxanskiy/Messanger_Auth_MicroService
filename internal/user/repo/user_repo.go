package repo

import (
	"auth_sevice_microservice/internal/user/model"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/nevasik7/lg"
)

type UserRepo struct {
	repo *pgxpool.Pool
}

func NewUserRepo(repo *pgxpool.Pool) *UserRepo {
	return &UserRepo{repo: repo}
}

func (r *UserRepo) CreateUser(ctx context.Context, user *model.User) error {
	query := QueryCreateUser

	err := r.repo.QueryRow(ctx, query, user.Nickname, user.Email, user.PasswordHash).Scan(user.ID)
	if err != nil {
		lg.Error(err)
	}
	return nil
}

//запрос номер 1 SELECT 1 FROM user  WHERE nickname  = $1 Limit 1;

// Проверка уникальности NickName
func (r *UserRepo) IsNicknameTaken(ctx context.Context, nickname string) (bool, error) {
	query := QueryIsNicknameTaken
	row := r.repo.QueryRow(ctx, query, nickname)

	var dummy int
	if err := row.Scan(&dummy); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *UserRepo) GetUserByNickname(ctx context.Context, nickname string) (*model.User, error) {
	query := QueryGetUserByNickname

	var user model.User
	err := r.repo.QueryRow(ctx, query, nickname).Scan(&user.ID, &user.Nickname, &user.PasswordHash)
	if err != nil {
		return nil, err
	}

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, err // пользователь не найден
		}
		return nil, err //ошибка при запросе
	}
	return &user, nil

}

func (r *UserRepo) SearchUserByNickname(ctx context.Context, nickname string) ([]*model.User, error) {
	query := QuerySearchUserByNickname

	rows, err := r.repo.Query(ctx, query, "%"+nickname+"%")
	if err != nil {
		lg.Errorf("message: Не удалось найти пользователя с таким ником")
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Nickname, &user.CreatedAt); err != nil {
			lg.Errorf("Не удалось отсканировать пользователя")
			return nil, err
		}
		users = append(users, &user)

	}
	return users, nil
}
