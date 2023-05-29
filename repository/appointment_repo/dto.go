package appointmentrepo

import (
	"appointment-api/internal/entity"
	"time"
)

type Appointment struct {
	ID              int64 `gorm:"primaryKey"`
	UserID          int64
	CoachID         int64
	Status          string `gorm:"default:null"`
	Rescheduled     bool
	StartAt         time.Time
	EndAt           time.Time
	DurationMinutes int32
	CreatedAt       time.Time `gorm:"default:null"`
	UpdatedAt       time.Time
}

func (Appointment) TableName() string {
	return "appointments"
}

func (a Appointment) ToEntity() entity.Appointment {
	return entity.Appointment{
		ID:              a.ID,
		UserID:          a.UserID,
		CoachID:         a.CoachID,
		Status:          entity.AppointmentStatus(a.Status),
		Rescheduled:     a.Rescheduled,
		StartAt:         a.StartAt,
		EndAt:           a.EndAt,
		DurationMinutes: a.DurationMinutes,
	}
}
