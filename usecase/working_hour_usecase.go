package usecase

import (
	"appointment-api/internal/domain"
	"appointment-api/internal/entity"
	"context"
)

type workingHourUsecase struct {
	workinghourrepo domain.WorkingHourRepository
}

func NewWorkingHourUsecase(workinghourrepo domain.WorkingHourRepository) domain.WorkingHourUsecase {
	return &workingHourUsecase{workinghourrepo}
}

func (uc *workingHourUsecase) GetWorkingHourByUserID(ctx context.Context, userID int64) ([]entity.WorkingHour, error) {
	workingHours, err := uc.workinghourrepo.GetWorkingHourByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return workingHours, nil
}
