package uc

import (
	"auth_sevice_microservice/internal/user/repo"
	"context"
	"gitlab.com/nevasik7/lg"
)

type LogoutUC struct {
	refreshTokenRepo *repo.RefreshTokenRepo
}

func NewLogoutUC(refreshTokenRepo *repo.RefreshTokenRepo) *LogoutUC {
	return &LogoutUC{refreshTokenRepo: refreshTokenRepo}
}

func (uc *LogoutUC) Logout(ctx context.Context, refreshToken string) error {
	// Проверяем, существует ли токен
	isValid, err := uc.refreshTokenRepo.IsTokenValid(ctx, refreshToken)
	if err != nil {
		return err
	}
	if !isValid {
		lg.Errorf("token is invalid or already expired")
		return err
	}

	// Удаляем токен
	return uc.refreshTokenRepo.DeleteToken(ctx, refreshToken)
}
