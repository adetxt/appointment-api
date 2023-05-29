package usecase

import (
	"appointment-api/internal/domain"
	"appointment-api/internal/entity"
	"appointment-api/utils"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/thoas/go-funk"
	"golang.org/x/sync/errgroup"
)

type appointmentUsecase struct {
	appointmentrepo domain.AppointmentRepository
	workingHourRepo domain.WorkingHourRepository
	userRepo        domain.UserRepository
}

func NewAppointmentUsecase(
	appointmentrepo domain.AppointmentRepository,
	workingHourRepo domain.WorkingHourRepository,
	userRepo domain.UserRepository,
) domain.AppointmentUsecase {
	return &appointmentUsecase{appointmentrepo, workingHourRepo, userRepo}
}

func (uc *appointmentUsecase) GetAppointmentsByUserID(ctx context.Context, role entity.Role, userID int64, req entity.GetAppointmentsByUserIDRequest) ([]entity.Appointment, error) {
	eg := &errgroup.Group{}

	appointmentsCh := make(chan []entity.Appointment)
	eg.Go(func() error {
		defer close(appointmentsCh)
		appointments, err := uc.appointmentrepo.GetAppointmentsByUserID(ctx, role, userID, req)
		if err != nil {
			return err
		}

		appointmentsCh <- appointments
		return nil
	})

	if req.Tz == nil {
		userLocCh := make(chan *time.Location)
		eg.Go(func() error {
			defer close(userLocCh)
			user, err := uc.userRepo.GetUserByID(ctx, userID)
			if err != nil {
				return err
			}

			userLoc, err := time.LoadLocation(user.Timezone)
			if err != nil {
				return fmt.Errorf("failed to load tz %s : %v", user.Timezone, err)
			}

			userLocCh <- userLoc
			return nil
		})
		req.Tz = <-userLocCh
	}

	appointments := <-appointmentsCh

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	for i := range appointments {
		appointments[i].StartAt = utils.RepresentTimeAs(appointments[i].StartAt, time.UTC)
		appointments[i].EndAt = utils.RepresentTimeAs(appointments[i].EndAt, time.UTC)

		startConv := appointments[i].StartAt.In(req.Tz)
		endConv := appointments[i].EndAt.In(req.Tz)

		appointments[i].StartAt = startConv
		appointments[i].EndAt = endConv
	}

	return appointments, nil
}

func (uc *appointmentUsecase) MakeAppointment(ctx context.Context, req entity.MakeAppointmentRequest) error {
	eg := &errgroup.Group{}

	userLocCh := make(chan *time.Location)
	eg.Go(func() error {
		defer close(userLocCh)
		user, err := uc.userRepo.GetUserByID(ctx, req.UserID)
		if err != nil {
			return err
		}

		userLoc, err := time.LoadLocation(user.Timezone)
		if err != nil {
			return fmt.Errorf("failed to load tz %s : %v", user.Timezone, err)
		}

		userLocCh <- userLoc
		return nil
	})

	coachLocCh := make(chan *time.Location)
	eg.Go(func() error {
		defer close(coachLocCh)
		coach, err := uc.userRepo.GetUserByID(ctx, req.CoachID)
		if err != nil {
			return err
		}

		coachLoc, err := time.LoadLocation(coach.Timezone)
		if err != nil {
			return fmt.Errorf("failed to load tz %s : %v", coach.Timezone, err)
		}

		coachLocCh <- coachLoc
		return nil
	})

	workingHourCh := make(chan entity.WorkingHour)
	eg.Go(func() error {
		defer close(workingHourCh)
		woringHours, err := uc.workingHourRepo.GetWorkingHourByUserID(ctx, req.CoachID)
		if err != nil {
			return err
		}

		dayWorkIntf := funk.ToMap(woringHours, "Day")

		dayWorkMap, ok := dayWorkIntf.(map[entity.Day]entity.WorkingHour)
		if !ok {
			return errors.New("failed to map working days")
		}

		weekday := int(req.StartAt.Weekday()) - 1
		woringHour := dayWorkMap[entity.Day(weekday)]

		workingHourCh <- woringHour
		return nil
	})

	userLoc := <-userLocCh
	coachLoc := <-coachLocCh
	workingHour := <-workingHourCh

	if err := eg.Wait(); err != nil {
		return err
	}

	req.SetTz(userLoc)

	if isValid, err := workingHour.IsOnWorkingTime(req.StartAt, req.GetEndAt(), coachLoc); err != nil || !isValid {
		return fmt.Errorf("outside coach working hour, coach available at %s %s til %s %s", workingHour.Day, workingHour.Start, workingHour.End, coachLoc.String())
	}

	startUTC := req.StartAt.In(coachLoc).UTC()
	endUTC := req.GetEndAt().UTC()

	schedules, err := uc.appointmentrepo.GetActiveAppointmentsOnDaterange(ctx, startUTC, endUTC)
	if err != nil {
		return err
	}

	if len(schedules) > 0 {
		return fmt.Errorf("coach already has schedules around the given time")
	}

	if err := uc.appointmentrepo.InsertAppointment(ctx, entity.MakeAppointmentRequest{
		UserID:          req.UserID,
		CoachID:         req.CoachID,
		StartAt:         startUTC,
		DurationMinutes: req.DurationMinutes,
	}); err != nil {
		return err
	}

	return nil
}

func (uc *appointmentUsecase) AppointmentAction(ctx context.Context, req entity.AppointmentActionRequest) error {
	if err := uc.appointmentrepo.AppointmentAction(ctx, req); err != nil {
		return err
	}

	return nil
}

func (uc *appointmentUsecase) RescheduleAppointment(ctx context.Context, req entity.RescheduleAppointmentRequest) error {
	appointment, err := uc.appointmentrepo.GetAppointmentByID(ctx, req.AppointmentID)
	if err != nil {
		return err
	}

	if appointment.Status == entity.STATUS_DECLINED || appointment.Status == entity.STATUS_CANCELED {
		return fmt.Errorf("can'r reschedule %s appointment", appointment.Status)
	}

	eg := &errgroup.Group{}

	coachLocCh := make(chan *time.Location)
	eg.Go(func() error {
		defer close(coachLocCh)
		coach, err := uc.userRepo.GetUserByID(ctx, appointment.CoachID)
		if err != nil {
			return err
		}

		coachLoc, err := time.LoadLocation(coach.Timezone)
		if err != nil {
			return fmt.Errorf("failed to load tz %s : %v", coach.Timezone, err)
		}

		coachLocCh <- coachLoc
		return nil
	})

	workingHourCh := make(chan entity.WorkingHour)
	eg.Go(func() error {
		defer close(workingHourCh)
		woringHours, err := uc.workingHourRepo.GetWorkingHourByUserID(ctx, appointment.CoachID)
		if err != nil {
			return err
		}

		dayWorkIntf := funk.ToMap(woringHours, "Day")

		dayWorkMap, ok := dayWorkIntf.(map[entity.Day]entity.WorkingHour)
		if !ok {
			return errors.New("failed to map working days")
		}

		weekday := int(req.StartAt.Weekday()) - 1
		woringHour := dayWorkMap[entity.Day(weekday)]

		workingHourCh <- woringHour
		return nil
	})

	coachLoc := <-coachLocCh
	workingHour := <-workingHourCh

	if err := eg.Wait(); err != nil {
		return err
	}

	req.SetTz(coachLoc)

	if isValid, err := workingHour.IsOnWorkingTime(req.StartAt, req.GetEndAt(), coachLoc); err != nil || !isValid {
		return fmt.Errorf("outside coach working hour, coach available at %s %s til %s %s", workingHour.Day, workingHour.Start, workingHour.End, coachLoc.String())
	}

	startUTC := req.StartAt.UTC()
	endUTC := req.GetEndAt().UTC()

	d := map[string]interface{}{
		"status":           entity.STATUS_RESCHEDULING,
		"rescheduled":      true,
		"start_at":         startUTC,
		"end_at":           endUTC,
		"duration_minutes": req.DurationMinutes,
	}

	if err := uc.appointmentrepo.Update(ctx, appointment.ID, d); err != nil {
		return err
	}

	return nil
}
