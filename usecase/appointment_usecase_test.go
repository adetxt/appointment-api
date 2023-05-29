package usecase

import (
	"appointment-api/internal/domain/mocks"
	"appointment-api/internal/entity"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AppointmentUsecaseTestSuite struct {
	suite.Suite
	ctx context.Context
	di  struct {
		appointmentrepo *mocks.AppointmentRepository
		workingHourRepo *mocks.WorkingHourRepository
		userRepo        *mocks.UserRepository
	}
}

// this function executes before the test suite begins execution
func (s *AppointmentUsecaseTestSuite) SetupSuite() {
	s.ctx = context.Background()
	s.di.appointmentrepo = mocks.NewAppointmentRepository(s.T())
	s.di.workingHourRepo = mocks.NewWorkingHourRepository(s.T())
	s.di.userRepo = mocks.NewUserRepository(s.T())
}

// this function executes after each test case
func (s *AppointmentUsecaseTestSuite) TearDownTest() {
	s.ctx = context.Background()
	s.di.appointmentrepo = mocks.NewAppointmentRepository(s.T())
	s.di.workingHourRepo = mocks.NewWorkingHourRepository(s.T())
	s.di.userRepo = mocks.NewUserRepository(s.T())
}

func TestAppointmentUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(AppointmentUsecaseTestSuite))
}

func (s *AppointmentUsecaseTestSuite) Test_appointmentUsecase_MakeAppointment() {
	dataUser := s.GetUsersDataTest()
	dataUserWorkingHours := s.GetUserWorkingHoursDataTest()
	dataUserAppointments := s.GetUserAppointmentsDataTest()
	dataMakeAppointment := s.GetMakeAppointmentRequestDataTest()

	type args struct {
		ctx context.Context
		req entity.MakeAppointmentRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		f       func()
	}{
		{
			name: "normal flow",
			args: args{
				ctx: context.Background(),
				req: dataMakeAppointment["data_1_2_normal"],
			},
			f: func() {
				userKey := "user_1"
				coachKey := "coach_2"

				coachLoc, _ := time.LoadLocation(dataUser[coachKey].Timezone)

				s.di.userRepo.EXPECT().GetUserByID(s.ctx, dataUser[userKey].ID).Return(dataUser[userKey], nil).Once()
				s.di.userRepo.EXPECT().GetUserByID(s.ctx, dataUser[coachKey].ID).Return(dataUser[coachKey], nil).Once()
				s.di.workingHourRepo.EXPECT().GetWorkingHourByUserID(s.ctx, dataUser[coachKey].ID).Return(dataUserWorkingHours[coachKey], nil).Once()
				s.di.appointmentrepo.EXPECT().GetActiveAppointmentsOnDaterange(s.ctx, mock.Anything, mock.Anything).Return(dataUserAppointments["empty"], nil).Maybe()

				d := dataMakeAppointment["data_1_2_normal"]
				d.StartAt = d.StartAt.In(coachLoc).UTC()
				s.di.appointmentrepo.EXPECT().InsertAppointment(s.ctx, d).Return(nil).Maybe()
			},
		},
		{
			name: "failed to insert",
			args: args{
				ctx: context.Background(),
				req: dataMakeAppointment["data_1_2_normal"],
			},
			wantErr: true,
			f: func() {
				userKey := "user_1"
				coachKey := "coach_2"

				coachLoc, _ := time.LoadLocation(dataUser[coachKey].Timezone)

				s.di.userRepo.EXPECT().GetUserByID(s.ctx, dataUser[userKey].ID).Return(dataUser[userKey], nil).Once()
				s.di.userRepo.EXPECT().GetUserByID(s.ctx, dataUser[coachKey].ID).Return(dataUser[coachKey], nil).Once()
				s.di.workingHourRepo.EXPECT().GetWorkingHourByUserID(s.ctx, dataUser[coachKey].ID).Return(dataUserWorkingHours[coachKey], nil).Once()
				s.di.appointmentrepo.EXPECT().GetActiveAppointmentsOnDaterange(s.ctx, mock.Anything, mock.Anything).Return(dataUserAppointments["empty"], nil).Once()

				d := dataMakeAppointment["data_1_2_normal"]
				d.StartAt = d.StartAt.In(coachLoc).UTC()
				s.di.appointmentrepo.EXPECT().InsertAppointment(s.ctx, d).Return(errors.New("")).Once()
			},
		},
		{
			name: "has conflict appointments",
			args: args{
				ctx: context.Background(),
				req: dataMakeAppointment["data_1_2_normal"],
			},
			wantErr: true,
			f: func() {
				userKey := "user_1"
				coachKey := "coach_2"

				s.di.userRepo.EXPECT().GetUserByID(s.ctx, dataUser[userKey].ID).Return(dataUser[userKey], nil).Once()
				s.di.userRepo.EXPECT().GetUserByID(s.ctx, dataUser[coachKey].ID).Return(dataUser[coachKey], nil).Once()
				s.di.workingHourRepo.EXPECT().GetWorkingHourByUserID(s.ctx, dataUser[coachKey].ID).Return(dataUserWorkingHours[coachKey], nil).Once()
				s.di.appointmentrepo.EXPECT().GetActiveAppointmentsOnDaterange(s.ctx, mock.Anything, mock.Anything).Return(dataUserAppointments["many"], nil).Once()
			},
		},
		{
			name: "outside working hour",
			args: args{
				ctx: context.Background(),
				req: dataMakeAppointment["data_1_2_outside_working_hour"],
			},
			wantErr: true,
			f: func() {
				userKey := "user_1"
				coachKey := "coach_2"

				s.di.userRepo.EXPECT().GetUserByID(s.ctx, dataUser[userKey].ID).Return(dataUser[userKey], nil).Once()
				s.di.userRepo.EXPECT().GetUserByID(s.ctx, dataUser[coachKey].ID).Return(dataUser[coachKey], nil).Once()
				s.di.workingHourRepo.EXPECT().GetWorkingHourByUserID(s.ctx, dataUser[coachKey].ID).Return(dataUserWorkingHours[coachKey], nil).Once()
			},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			uc := &appointmentUsecase{
				appointmentrepo: s.di.appointmentrepo,
				workingHourRepo: s.di.workingHourRepo,
				userRepo:        s.di.userRepo,
			}

			tt.f()

			if err := uc.MakeAppointment(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("appointmentUsecase.MakeAppointment() error = %v, wantErr %v", err, tt.wantErr)
			}

			s.TearDownTest()
		})
	}
}

func (s *AppointmentUsecaseTestSuite) Test_appointmentUsecase_RescheduleAppointment() {
	dataUser := s.GetUsersDataTest()
	dataUserWorkingHours := s.GetUserWorkingHoursDataTest()
	// dataUserAppointments := s.GetUserAppointmentsDataTest()
	dataRescheduleAppointment := s.GetRescheduleAppointmentDataTest()

	type args struct {
		ctx context.Context
		req entity.RescheduleAppointmentRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		f       func()
	}{
		{
			name: "normal flow",
			args: args{
				ctx: context.Background(),
				req: dataRescheduleAppointment["data_1_2_normal"],
			},
			f: func() {
				coachKey := "coach_2"
				d := dataRescheduleAppointment["data_1_2_normal"]

				coachLoc, _ := time.LoadLocation(dataUser[coachKey].Timezone)

				s.di.appointmentrepo.EXPECT().GetAppointmentByID(s.ctx, d.AppointmentID).Return(entity.Appointment{
					ID:      d.AppointmentID,
					CoachID: dataUser[coachKey].ID,
					Status:  entity.STATUS_SCHEDULING,
				}, nil).Once()
				s.di.userRepo.EXPECT().GetUserByID(s.ctx, dataUser[coachKey].ID).Return(dataUser[coachKey], nil).Once()
				s.di.workingHourRepo.EXPECT().GetWorkingHourByUserID(s.ctx, dataUser[coachKey].ID).Return(dataUserWorkingHours[coachKey], nil).Once()

				d.StartAt = d.StartAt.In(coachLoc).UTC()

				s.di.appointmentrepo.EXPECT().Update(s.ctx, d.AppointmentID, map[string]interface{}{
					"status":           entity.STATUS_RESCHEDULING,
					"rescheduled":      true,
					"start_at":         d.StartAt,
					"end_at":           d.GetEndAt(),
					"duration_minutes": d.DurationMinutes,
				}).Return(nil).Maybe()
			},
		},
		{
			name: "appointment is declined or canceled",
			args: args{
				ctx: context.Background(),
				req: dataRescheduleAppointment["data_1_2_normal"],
			},
			wantErr: true,
			f: func() {
				coachKey := "coach_2"
				d := dataRescheduleAppointment["data_1_2_normal"]

				s.di.appointmentrepo.EXPECT().GetAppointmentByID(s.ctx, d.AppointmentID).Return(entity.Appointment{
					CoachID: dataUser[coachKey].ID,
					Status:  entity.STATUS_DECLINED,
				}, nil).Once()
			},
		},
		{
			name: "outside coch working hour",
			args: args{
				ctx: context.Background(),
				req: dataRescheduleAppointment["data_1_2_outside_working_hour"],
			},
			wantErr: true,
			f: func() {
				coachKey := "coach_2"
				d := dataRescheduleAppointment["data_1_2_outside_working_hour"]

				s.di.appointmentrepo.EXPECT().GetAppointmentByID(s.ctx, d.AppointmentID).Return(entity.Appointment{
					CoachID: dataUser[coachKey].ID,
					Status:  entity.STATUS_SCHEDULING,
				}, nil).Once()
				s.di.userRepo.EXPECT().GetUserByID(s.ctx, dataUser[coachKey].ID).Return(dataUser[coachKey], nil).Once()
				s.di.workingHourRepo.EXPECT().GetWorkingHourByUserID(s.ctx, dataUser[coachKey].ID).Return(dataUserWorkingHours[coachKey], nil).Once()
			},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			uc := &appointmentUsecase{
				appointmentrepo: s.di.appointmentrepo,
				workingHourRepo: s.di.workingHourRepo,
				userRepo:        s.di.userRepo,
			}

			tt.f()

			if err := uc.RescheduleAppointment(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("appointmentUsecase.RescheduleAppointment() error = %v, wantErr %v", err, tt.wantErr)
			}

			s.TearDownTest()
		})
	}
}

func (s *AppointmentUsecaseTestSuite) GetUsersDataTest() map[string]entity.User {
	return map[string]entity.User{
		"user_1": {
			ID:       1,
			Name:     "Fulan",
			Role:     entity.ROLE_USER,
			Timezone: "Asia/Jakarta",
		},
		"coach_2": {
			ID:       2,
			Name:     "Christy Schumm",
			Role:     entity.ROLE_COACH,
			Timezone: "America/North_Dakota/New_Salem",
		},
	}
}

func (s *AppointmentUsecaseTestSuite) GetUserWorkingHoursDataTest() map[string][]entity.WorkingHour {
	return map[string][]entity.WorkingHour{
		"coach_2": {
			{
				Day:   entity.Monday,
				Start: "09:00:00",
				End:   "17:30:00",
			},
		},
	}
}

func (s *AppointmentUsecaseTestSuite) GetUserAppointmentsDataTest() map[string][]entity.Appointment {
	return map[string][]entity.Appointment{
		"empty": {},
		"many": {
			{
				ID: 1,
			},
		},
	}
}

func (s *AppointmentUsecaseTestSuite) GetMakeAppointmentRequestDataTest() map[string]entity.MakeAppointmentRequest {
	u := s.GetUsersDataTest()

	return map[string]entity.MakeAppointmentRequest{
		"data_1_2_normal": {
			UserID:          u["user_1"].ID,
			CoachID:         u["coach_2"].ID,
			StartAt:         time.Date(2023, 5, 29, 21, 0, 0, 0, s.getLoc(u["user_1"].Timezone)),
			DurationMinutes: 15,
		},
		"data_1_2_outside_working_hour": {
			UserID:          u["user_1"].ID,
			CoachID:         u["coach_2"].ID,
			StartAt:         time.Date(2023, 5, 29, 9, 0, 0, 0, s.getLoc(u["user_1"].Timezone)),
			DurationMinutes: 15,
		},
	}
}

func (s *AppointmentUsecaseTestSuite) GetRescheduleAppointmentDataTest() map[string]entity.RescheduleAppointmentRequest {
	u := s.GetUsersDataTest()

	return map[string]entity.RescheduleAppointmentRequest{
		"data_1_2_normal": {
			AppointmentID:   1,
			StartAt:         time.Date(2023, 5, 29, 9, 0, 0, 0, s.getLoc(u["coach_2"].Timezone)),
			DurationMinutes: 15,
		},
		"data_1_2_outside_working_hour": {
			AppointmentID:   1,
			StartAt:         time.Date(2023, 5, 29, 18, 0, 0, 0, s.getLoc(u["coach_2"].Timezone)),
			DurationMinutes: 15,
		},
	}
}

func (s *AppointmentUsecaseTestSuite) GetErrDataTest() map[string]error {
	return map[string]error{
		"error": errors.New("error"),
	}
}

func (s *AppointmentUsecaseTestSuite) getLoc(v string) *time.Location {
	l, _ := time.LoadLocation(v)
	return l
}
