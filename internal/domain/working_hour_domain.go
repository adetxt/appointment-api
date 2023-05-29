package domain

import (
	"appointment-api/internal/entity"
	"context"
)

type WorkingHourUsecase interface {
	GetWorkingHourByUserID(ctx context.Context, userID int64) ([]entity.WorkingHour, error)
}

//go:generate mockery --name WorkingHourRepository --with-expecter=true
type WorkingHourRepository interface {
	GetWorkingHourByUserID(ctx context.Context, userID int64) ([]entity.WorkingHour, error)
}
