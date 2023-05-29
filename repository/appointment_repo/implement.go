package appointmentrepo

import (
	"appointment-api/internal/domain"
	"appointment-api/internal/entity"
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.AppointmentRepository {
	return &repository{db}
}

func (r *repository) GetAppointmentsByUserID(ctx context.Context, role entity.Role, userID int64, req entity.GetAppointmentsByUserIDRequest) ([]entity.Appointment, error) {
	appointments := []Appointment{}

	key := "user_id"
	if role == entity.ROLE_COACH {
		key = "coach_id"
	}

	q := r.db.Model(&Appointment{})

	if len(req.Status) > 0 {
		q.Where("status IN ?", req.Status)
	}

	// if !req.StartDate.IsZero() && !req.EndDate.IsZero() {
	// 	// q.Where("status", req.Status)
	// }

	if err := q.Where(key, userID).
		Order("start_at, end_at").
		Find(&appointments).Error; err != nil {
		return nil, err
	}

	res := make([]entity.Appointment, len(appointments))
	for i := range appointments {
		res[i] = appointments[i].ToEntity()
	}

	return res, nil
}

func (r *repository) GetActiveAppointmentsOnDaterange(ctx context.Context, start, end time.Time) ([]entity.Appointment, error) {
	appointments := []Appointment{}

	if err := r.db.Model(&Appointment{}).
		Where("status", entity.STATUS_SCHEDULED).
		Where("end_at >= ? AND end_at <= ?", start, end).Find(&appointments).Error; err != nil {
		return nil, err
	}

	res := make([]entity.Appointment, len(appointments))
	for i := range appointments {
		res[i] = appointments[i].ToEntity()
	}

	return res, nil
}

func (r *repository) GetAppointmentByID(ctx context.Context, ID int64) (entity.Appointment, error) {
	app := Appointment{}
	if err := r.db.Model(&Appointment{}).Where("id", ID).First(&app).Error; err != nil {
		return entity.Appointment{}, err
	}

	return app.ToEntity(), nil
}

func (r *repository) InsertAppointment(ctx context.Context, req entity.MakeAppointmentRequest) error {
	appointmentRow := Appointment{
		UserID:          req.UserID,
		CoachID:         req.CoachID,
		StartAt:         req.StartAt,
		DurationMinutes: req.DurationMinutes,
		EndAt:           req.GetEndAt(),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	if err := r.db.Create(&appointmentRow).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) AppointmentAction(ctx context.Context, req entity.AppointmentActionRequest) error {
	act := req.GetAppointmentStatusOfAction()

	if act == "" {
		return fmt.Errorf("invalid action")
	}

	app, err := r.GetAppointmentByID(ctx, req.AppointmentID)
	if err != nil {
		return err
	}

	if req.Action == entity.ACTION_APPROVE {
		if app.Status != entity.STATUS_SCHEDULING && app.Status != entity.STATUS_RESCHEDULING {
			return fmt.Errorf("can't update, the appointment is already %s", app.Status)
		}

		apps, err := r.GetActiveAppointmentsOnDaterange(ctx, app.StartAt, app.EndAt)
		if err != nil {
			return err
		}

		if len(apps) > 0 {
			return fmt.Errorf("can't approve, there's already active appointment in the time range")
		}
	}

	d := map[string]interface{}{
		"status": act,
	}

	if err := r.Update(ctx, req.AppointmentID, d); err != nil {
		return err
	}

	return nil
}

func (r *repository) Update(ctx context.Context, ID int64, data map[string]interface{}) error {
	if ID < 1 {
		return fmt.Errorf("appointment not found")
	}

	if err := r.db.Model(&Appointment{}).Where("id", ID).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
