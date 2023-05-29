package domain

import (
	"appointment-api/internal/entity"
	"context"
	"time"
)

type AppointmentUsecase interface {
	GetAppointmentsByUserID(ctx context.Context, role entity.Role, userID int64, req entity.GetAppointmentsByUserIDRequest) ([]entity.Appointment, error)
	MakeAppointment(ctx context.Context, req entity.MakeAppointmentRequest) error
	AppointmentAction(ctx context.Context, req entity.AppointmentActionRequest) error
	RescheduleAppointment(ctx context.Context, req entity.RescheduleAppointmentRequest) error
}

//go:generate mockery --name AppointmentRepository --with-expecter=true
type AppointmentRepository interface {
	GetAppointmentsByUserID(ctx context.Context, role entity.Role, userID int64, req entity.GetAppointmentsByUserIDRequest) ([]entity.Appointment, error)
	GetActiveAppointmentsOnDaterange(ctx context.Context, start, end time.Time) ([]entity.Appointment, error)
	GetAppointmentByID(ctx context.Context, ID int64) (entity.Appointment, error)
	InsertAppointment(ctx context.Context, req entity.MakeAppointmentRequest) error
	AppointmentAction(ctx context.Context, req entity.AppointmentActionRequest) error
	Update(ctx context.Context, ID int64, data map[string]interface{}) error
}
