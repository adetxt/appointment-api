package domain

import (
	"appointment-api/internal/entity"
	"context"
)

type UserUsecase interface {
	GetUsers(ctx context.Context, req entity.GetUsersRequest) ([]entity.User, entity.Pagination, error)
	GetUserByID(ctx context.Context, ID int64) (entity.User, error)
}

//go:generate mockery --name UserRepository --with-expecter=true
type UserRepository interface {
	GetUsers(ctx context.Context, req entity.GetUsersRequest) ([]entity.User, entity.Pagination, error)
	GetUserByID(ctx context.Context, ID int64) (entity.User, error)
}
