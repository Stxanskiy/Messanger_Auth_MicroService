package uc

import (
	"auth_sevice_microservice/internal/user/model"
	"auth_sevice_microservice/internal/user/repo"
	"auth_sevice_microservice/pkg/jwt"
	"context"
	"golang.org/x/crypto/bcrypt"
)

type LoginUC struct {
	repo             *repo.UserRepo
	jwtManager       *jwt.JWTManager
	refreshTokenRepo *repo.RefreshTokenRepo
}

func NewLoginUC(repo *repo.UserRepo, jwtManager *jwt.JWTManager, refreshTokenRepo *repo.RefreshTokenRepo) *LoginUC {
	return &LoginUC{
		repo:             repo,
		jwtManager:       jwtManager,
		refreshTokenRepo: refreshTokenRepo,
	}
}

func (uc *LoginUC) Login(ctx context.Context, nickname, password string) (*model.Token, error) {
	//проверка существует ли пользватель
	user, err := uc.repo.GetUserByNickname(ctx, nickname)
	if err != nil {
		return nil, err

	}

	//проверяем пароль
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, err //errors.New("Неправильно введен пароль! Попробуйте еще раз.")
	}

	// генерация токенов
	access, refresh, err := uc.jwtManager.GenerateTokens(user.ID)
	if err != nil {
		return nil, err
	}

	// сохранение токенов в базу данных
	if err := uc.refreshTokenRepo.SaveToken(ctx, user.ID, refresh.ID, refresh.ExpiresAt.Time); err != nil {
		return nil, err
	}

	return &model.Token{
		AccessToken:  access.ID,
		RefreshToken: refresh.ID,
		ExpiresAt:    access.ExpiresAt.Time,
	}, nil
}
