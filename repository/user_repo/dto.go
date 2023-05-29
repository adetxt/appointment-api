package userrepo

import "appointment-api/internal/entity"

type User struct {
	ID       int64 `gorm:"primaryKey"`
	Name     string
	Role     string
	Timezone string
}

func (User) TableName() string {
	return "users"
}

func (u User) ToEntity() entity.User {
	return entity.User{
		ID:       u.ID,
		Name:     u.Name,
		Role:     entity.Role(u.Role),
		Timezone: u.Timezone,
	}
}
