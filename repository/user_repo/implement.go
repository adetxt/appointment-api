package userrepo

import (
	"appointment-api/internal/domain"
	"appointment-api/internal/entity"
	"appointment-api/utils"
	"context"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.UserRepository {
	return &repository{db}
}

func (r *repository) GetUsers(ctx context.Context, req entity.GetUsersRequest) ([]entity.User, entity.Pagination, error) {
	userRows := []User{}

	q := r.db.Model(&User{})

	if len(req.Fields) > 0 {
		q.Select(req.Fields)
	}

	if err := q.Limit(req.PageSize).
		Offset(utils.CalculateOffset(req.Page, req.PageSize)).
		Find(&userRows).
		Error; err != nil {
		return nil, entity.Pagination{}, err
	}

	var c int64

	if err := q.Select("*").Count(&c).Error; err != nil {
		return nil, entity.Pagination{}, err
	}

	result := make([]entity.User, len(userRows))

	for i := 0; i < len(userRows); i++ {
		result[i] = userRows[i].ToEntity()
	}

	return result, entity.Pagination{
		Page:      req.Page,
		PageSize:  req.PageSize,
		Total:     c,
		TotalPage: int(c / int64(req.PageSize)),
	}, nil
}

func (r *repository) GetUserByID(ctx context.Context, ID int64) (entity.User, error) {
	user := User{}

	if err := r.db.Model(&User{}).Where("id = ?", ID).First(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user.ToEntity(), nil
}
