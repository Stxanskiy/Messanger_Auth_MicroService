package uc

import (
	"auth_sevice_microservice/internal/user/repo"
	"auth_sevice_microservice/pkg/jwt"
	"context"

	"errors"
)

type TokenRefreshUC struct {
	refreshTokenRepo *repo.RefreshTokenRepo
	jwtManager       *jwt.JWTManager
}

func NewTokenRefreshTokenUC(repo *repo.RefreshTokenRepo, jwtManager *jwt.JWTManager) *TokenRefreshUC {
	return &TokenRefreshUC{refreshTokenRepo: repo, jwtManager: jwtManager}
}

func (uc *TokenRefreshUC) RefreshTokens(ctx context.Context, refreshToken string) (*jwt.Claims, *jwt.Claims, error) {
	// Проверяем валидность токена
	isValid, err := uc.refreshTokenRepo.IsTokenValid(ctx, refreshToken)
	if err != nil || !isValid {
		return nil, nil, errors.New("токен недествителен либо устарел")
	}

	// Парсим токен
	claims, err := uc.jwtManager.ParseToken(refreshToken)
	if err != nil {
		return nil, nil, err
	}

	// Генерируем новые токены
	access, refresh, err := uc.jwtManager.GenerateTokens(claims.UserID)
	if err != nil {
		return nil, nil, err
	}

	// Обновляем refresh-токен в базе данных
	if err := uc.refreshTokenRepo.DeleteToken(ctx, refreshToken); err != nil {
		return nil, nil, err
	}
	if err := uc.refreshTokenRepo.SaveToken(ctx, claims.UserID, refresh.ID, refresh.ExpiresAt.Time); err != nil {
		return nil, nil, err
	}

	return access, refresh, nil
}
