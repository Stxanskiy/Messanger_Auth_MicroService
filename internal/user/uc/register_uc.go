package uc

import (
	"auth_sevice_microservice/internal/user/model"
	"auth_sevice_microservice/internal/user/repo"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserUC struct {
	repo *repo.UserRepo
}

func NewUserUC(repo *repo.UserRepo) *UserUC {
	return &UserUC{repo: repo}
}

func (uc *UserUC) RefisterUser(ctx context.Context, nickname, email, password string) (*model.User, error) {
	//Проверка уникальности nickname
	isTaken, err := uc.repo.IsNicknameTaken(ctx, nickname)
	if err != nil {
		return nil, err
	}

	if isTaken {
		return nil, errors.New("Пользователь с таким ником уже существую\n Попробуйте ввести другой ник")
	}

	//хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	//Создаем пользователя
	user := &model.User{
		Nickname:     nickname,
		Email:        email,
		PasswordHash: string(hashedPassword),
	}

	if err := uc.repo.CreateUser(ctx, user); err != nil {
		return nil, err
	}
	return user, nil

}
