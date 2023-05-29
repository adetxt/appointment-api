package entity

import (
	"appointment-api/utils"
	"time"
)

type AppointmentStatus string
type AppointmentAction string

const (
	STATUS_SCHEDULING   AppointmentStatus = "SCHEDULING"
	STATUS_SCHEDULED    AppointmentStatus = "SCHEDULED"
	STATUS_CANCELED     AppointmentStatus = "CANCELED"
	STATUS_RESCHEDULING AppointmentStatus = "RESCHEDULING"
	STATUS_DECLINED     AppointmentStatus = "DECLINED"
)

const (
	ACTION_APPROVE    AppointmentAction = "APPROVE"
	ACTION_DECLINE    AppointmentAction = "DECLINE"
	ACTION_RESCHEDULE AppointmentAction = "RESCHEDULE"
)

type Appointment struct {
	ID              int64
	UserID          int64
	CoachID         int64
	Status          AppointmentStatus
	Rescheduled     bool
	StartAt         time.Time
	EndAt           time.Time
	DurationMinutes int32
}

type GetAppointmentsByUserIDRequest struct {
	Status    []AppointmentStatus
	StartDate time.Time
	EndDate   time.Time
	Tz        *time.Location
}

type MakeAppointmentRequest struct {
	UserID          int64
	CoachID         int64
	StartAt         time.Time
	DurationMinutes int32
	tz              *time.Location
}

type AppointmentActionRequest struct {
	AppointmentID int64
	Action        AppointmentAction
}

type RescheduleAppointmentRequest struct {
	AppointmentID   int64
	StartAt         time.Time
	DurationMinutes int32
}

func (r MakeAppointmentRequest) GetEndAt() time.Time {
	return r.StartAt.Add(time.Minute * time.Duration(r.DurationMinutes))
}

func (r *MakeAppointmentRequest) SetTz(tz *time.Location) {
	r.tz = tz
	r.StartAt = utils.RepresentTimeAs(r.StartAt, tz)
}

func (r MakeAppointmentRequest) GetTz() *time.Location {
	return r.tz
}

func (r AppointmentActionRequest) GetAppointmentStatusOfAction() AppointmentStatus {
	switch r.Action {
	case ACTION_APPROVE:
		return STATUS_SCHEDULED
	case ACTION_DECLINE:
		return STATUS_DECLINED
	}

	return ""
}

func (r RescheduleAppointmentRequest) GetEndAt() time.Time {
	return r.StartAt.Add(time.Minute * time.Duration(r.DurationMinutes))
}

func (r *RescheduleAppointmentRequest) SetTz(tz *time.Location) {
	r.StartAt = utils.RepresentTimeAs(r.StartAt, tz)
}
