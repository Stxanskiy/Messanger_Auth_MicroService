package uc

import (
	"auth_sevice_microservice/internal/user/model"
	"auth_sevice_microservice/internal/user/repo"
	"context"
)

type SearchUsersUC struct {
	userRepo *repo.UserRepo
}

func NewSearchUserUC(userRepo *repo.UserRepo) *SearchUsersUC {
	return &SearchUsersUC{userRepo: userRepo}
}

func (uc *SearchUsersUC) Search(ctx context.Context, nickname string) ([]*model.User, error) {
	return uc.userRepo.SearchUserByNickname(ctx, nickname)

}
