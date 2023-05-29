package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"appointment-api/graph/model"
	"appointment-api/internal/entity"
	"context"
	"fmt"
	"time"

	"github.com/thoas/go-funk"
)

// MakeAppointment is the resolver for the makeAppointment field.
func (r *mutationResolver) MakeAppointment(ctx context.Context, input model.MakeAppointmentRequest) (*bool, error) {
	start, err := time.Parse(time.DateTime, input.StartAt)
	if err != nil {
		return nil, err
	}

	if err := r.AppointmentUc.MakeAppointment(ctx, entity.MakeAppointmentRequest{
		UserID:          input.UserID,
		CoachID:         input.CoachID,
		StartAt:         start,
		DurationMinutes: int32(input.DurationMunites),
	}); err != nil {
		return nil, err
	}

	return nil, nil
}

// ApproveAppointment is the resolver for the approveAppointment field.
func (r *mutationResolver) ApproveAppointment(ctx context.Context, appointmentID int64) (*bool, error) {
	if err := r.AppointmentUc.AppointmentAction(ctx, entity.AppointmentActionRequest{
		AppointmentID: appointmentID,
		Action:        entity.ACTION_APPROVE,
	}); err != nil {
		return nil, err
	}

	return nil, nil
}

// DeclineAppointment is the resolver for the declineAppointment field.
func (r *mutationResolver) DeclineAppointment(ctx context.Context, appointmentID int64) (*bool, error) {
	if err := r.AppointmentUc.AppointmentAction(ctx, entity.AppointmentActionRequest{
		AppointmentID: appointmentID,
		Action:        entity.ACTION_DECLINE,
	}); err != nil {
		return nil, err
	}

	return nil, nil
}

// RescheduleAppointment is the resolver for the rescheduleAppointment field.
func (r *mutationResolver) RescheduleAppointment(ctx context.Context, input model.RescheduleAppointmentRequest) (*bool, error) {
	start, err := time.Parse(time.DateTime, input.StartAt)
	if err != nil {
		return nil, err
	}

	if err := r.AppointmentUc.RescheduleAppointment(ctx, entity.RescheduleAppointmentRequest{
		AppointmentID:   input.AppointmentID,
		StartAt:         start,
		DurationMinutes: int32(input.DurationMunites),
	}); err != nil {
		return nil, err
	}

	return nil, nil
}

// UserList is the resolver for the userList field.
func (r *queryResolver) UserList(ctx context.Context, page *int, pageSize *int) (*model.UserListResponse, error) {
	users, pi, err := r.UserUc.GetUsers(ctx, entity.GetUsersRequest{
		Page:     *page,
		PageSize: *pageSize,
	})
	if err != nil {
		return nil, err
	}

	fields := GetPreloads(ctx)
	withWorkingHours := funk.ContainsString(fields, "items.workingHours")

	res := make([]*model.User, len(users))
	for i := range users {
		res[i] = &model.User{
			ID:       users[i].ID,
			Name:     users[i].Name,
			Role:     model.Role(users[i].Role),
			Timezone: users[i].Timezone,
		}

		if withWorkingHours {
			works, err := r.UserWorkingHour(ctx, users[i].ID)
			if err != nil {
				return nil, err
			}

			res[i].WorkingHours = works
		}
	}

	return &model.UserListResponse{
		Items:     res,
		Page:      pi.Page,
		PageSize:  pi.PageSize,
		Total:     pi.Total,
		TotalPage: pi.TotalPage,
	}, nil
}

// UserWorkingHour is the resolver for the userWorkingHour field.
func (r *queryResolver) UserWorkingHour(ctx context.Context, userID int64) ([]*model.WorkingHour, error) {
	if userID < 1 {
		return nil, fmt.Errorf("user not found")
	}

	workingHours, err := r.WorkingHourUc.GetWorkingHourByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	res := make([]*model.WorkingHour, len(workingHours))
	for i := range workingHours {
		res[i] = &model.WorkingHour{
			ID:     workingHours[i].ID,
			UserID: workingHours[i].UserID,
			Day:    workingHours[i].Day.String(),
			Start:  workingHours[i].Start,
			End:    workingHours[i].End,
		}
	}

	return res, nil
}

// UserAppointments is the resolver for the userAppointments field.
func (r *queryResolver) UserAppointments(ctx context.Context, userID int64, role model.Role, params model.UserAppointmentsRequest) ([]*model.Appointment, error) {
	var loc *time.Location = nil
	if params.Tz != nil {
		locc, err := time.LoadLocation(*params.Tz)
		if err != nil {
			return nil, fmt.Errorf("failed to load tz %s", *params.Tz)
		}

		loc = locc
	}

	req := entity.GetAppointmentsByUserIDRequest{
		Tz: loc,
	}

	if params.Status != nil {
		for i := range params.Status {
			req.Status = append(req.Status, entity.AppointmentStatus(params.Status[i]))
		}
	}

	appointments, err := r.AppointmentUc.GetAppointmentsByUserID(ctx, entity.Role(role), userID, req)
	if err != nil {
		return nil, err
	}

	res := make([]*model.Appointment, len(appointments))
	for i := range appointments {
		res[i] = &model.Appointment{
			ID:              appointments[i].ID,
			UserID:          appointments[i].UserID,
			CoachID:         appointments[i].CoachID,
			Status:          string(appointments[i].Status),
			Rescheduled:     appointments[i].Rescheduled,
			StartAt:         appointments[i].StartAt.Format(time.RFC3339),
			EndAt:           appointments[i].EndAt.Format(time.RFC3339),
			DurationMunites: int(appointments[i].DurationMinutes),
		}
	}

	return res, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
