package usecase

import (
	"appointment-api/internal/domain"
	"appointment-api/internal/entity"
	"context"
	"fmt"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{userRepo}
}

func (uc *userUsecase) GetUsers(ctx context.Context, req entity.GetUsersRequest) ([]entity.User, entity.Pagination, error) {
	users, pi, err := uc.userRepo.GetUsers(ctx, req)
	if err != nil {
		fmt.Printf("[GetUsers] failed get users : %v", err)
		return nil, entity.Pagination{}, err
	}

	return users, pi, nil
}

func (uc *userUsecase) GetUserByID(ctx context.Context, ID int64) (entity.User, error) {
	return uc.userRepo.GetUserByID(ctx, ID)
}
