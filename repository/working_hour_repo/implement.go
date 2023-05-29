package workinghourrepo

import (
	"appointment-api/internal/domain"
	"appointment-api/internal/entity"
	"context"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.WorkingHourRepository {
	return &repository{db}
}

func (r *repository) GetWorkingHourByUserID(ctx context.Context, userID int64) ([]entity.WorkingHour, error) {
	workingHourRows := []WorkingHour{}

	if userID < 1 {
		return nil, gorm.ErrRecordNotFound
	}

	if err := r.db.Model(&WorkingHour{}).Where("user_id = ?", userID).
		Order("day asc, start asc").
		Find(&workingHourRows).Error; err != nil {
		return nil, err
	}

	res := make([]entity.WorkingHour, len(workingHourRows))
	for i := range workingHourRows {
		res[i] = workingHourRows[i].ToEntity()
	}

	return res, nil
}
