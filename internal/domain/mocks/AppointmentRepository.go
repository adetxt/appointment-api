// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "appointment-api/internal/entity"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// AppointmentRepository is an autogenerated mock type for the AppointmentRepository type
type AppointmentRepository struct {
	mock.Mock
}

type AppointmentRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *AppointmentRepository) EXPECT() *AppointmentRepository_Expecter {
	return &AppointmentRepository_Expecter{mock: &_m.Mock}
}

// AppointmentAction provides a mock function with given fields: ctx, req
func (_m *AppointmentRepository) AppointmentAction(ctx context.Context, req entity.AppointmentActionRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.AppointmentActionRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AppointmentRepository_AppointmentAction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AppointmentAction'
type AppointmentRepository_AppointmentAction_Call struct {
	*mock.Call
}

// AppointmentAction is a helper method to define mock.On call
//   - ctx context.Context
//   - req entity.AppointmentActionRequest
func (_e *AppointmentRepository_Expecter) AppointmentAction(ctx interface{}, req interface{}) *AppointmentRepository_AppointmentAction_Call {
	return &AppointmentRepository_AppointmentAction_Call{Call: _e.mock.On("AppointmentAction", ctx, req)}
}

func (_c *AppointmentRepository_AppointmentAction_Call) Run(run func(ctx context.Context, req entity.AppointmentActionRequest)) *AppointmentRepository_AppointmentAction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.AppointmentActionRequest))
	})
	return _c
}

func (_c *AppointmentRepository_AppointmentAction_Call) Return(_a0 error) *AppointmentRepository_AppointmentAction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AppointmentRepository_AppointmentAction_Call) RunAndReturn(run func(context.Context, entity.AppointmentActionRequest) error) *AppointmentRepository_AppointmentAction_Call {
	_c.Call.Return(run)
	return _c
}

// GetActiveAppointmentsOnDaterange provides a mock function with given fields: ctx, start, end
func (_m *AppointmentRepository) GetActiveAppointmentsOnDaterange(ctx context.Context, start time.Time, end time.Time) ([]entity.Appointment, error) {
	ret := _m.Called(ctx, start, end)

	var r0 []entity.Appointment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, time.Time) ([]entity.Appointment, error)); ok {
		return rf(ctx, start, end)
	}
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, time.Time) []entity.Appointment); ok {
		r0 = rf(ctx, start, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Appointment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, time.Time, time.Time) error); ok {
		r1 = rf(ctx, start, end)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AppointmentRepository_GetActiveAppointmentsOnDaterange_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetActiveAppointmentsOnDaterange'
type AppointmentRepository_GetActiveAppointmentsOnDaterange_Call struct {
	*mock.Call
}

// GetActiveAppointmentsOnDaterange is a helper method to define mock.On call
//   - ctx context.Context
//   - start time.Time
//   - end time.Time
func (_e *AppointmentRepository_Expecter) GetActiveAppointmentsOnDaterange(ctx interface{}, start interface{}, end interface{}) *AppointmentRepository_GetActiveAppointmentsOnDaterange_Call {
	return &AppointmentRepository_GetActiveAppointmentsOnDaterange_Call{Call: _e.mock.On("GetActiveAppointmentsOnDaterange", ctx, start, end)}
}

func (_c *AppointmentRepository_GetActiveAppointmentsOnDaterange_Call) Run(run func(ctx context.Context, start time.Time, end time.Time)) *AppointmentRepository_GetActiveAppointmentsOnDaterange_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(time.Time), args[2].(time.Time))
	})
	return _c
}

func (_c *AppointmentRepository_GetActiveAppointmentsOnDaterange_Call) Return(_a0 []entity.Appointment, _a1 error) *AppointmentRepository_GetActiveAppointmentsOnDaterange_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AppointmentRepository_GetActiveAppointmentsOnDaterange_Call) RunAndReturn(run func(context.Context, time.Time, time.Time) ([]entity.Appointment, error)) *AppointmentRepository_GetActiveAppointmentsOnDaterange_Call {
	_c.Call.Return(run)
	return _c
}

// GetAppointmentByID provides a mock function with given fields: ctx, ID
func (_m *AppointmentRepository) GetAppointmentByID(ctx context.Context, ID int64) (entity.Appointment, error) {
	ret := _m.Called(ctx, ID)

	var r0 entity.Appointment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (entity.Appointment, error)); ok {
		return rf(ctx, ID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) entity.Appointment); ok {
		r0 = rf(ctx, ID)
	} else {
		r0 = ret.Get(0).(entity.Appointment)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AppointmentRepository_GetAppointmentByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAppointmentByID'
type AppointmentRepository_GetAppointmentByID_Call struct {
	*mock.Call
}

// GetAppointmentByID is a helper method to define mock.On call
//   - ctx context.Context
//   - ID int64
func (_e *AppointmentRepository_Expecter) GetAppointmentByID(ctx interface{}, ID interface{}) *AppointmentRepository_GetAppointmentByID_Call {
	return &AppointmentRepository_GetAppointmentByID_Call{Call: _e.mock.On("GetAppointmentByID", ctx, ID)}
}

func (_c *AppointmentRepository_GetAppointmentByID_Call) Run(run func(ctx context.Context, ID int64)) *AppointmentRepository_GetAppointmentByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *AppointmentRepository_GetAppointmentByID_Call) Return(_a0 entity.Appointment, _a1 error) *AppointmentRepository_GetAppointmentByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AppointmentRepository_GetAppointmentByID_Call) RunAndReturn(run func(context.Context, int64) (entity.Appointment, error)) *AppointmentRepository_GetAppointmentByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetAppointmentsByUserID provides a mock function with given fields: ctx, role, userID, req
func (_m *AppointmentRepository) GetAppointmentsByUserID(ctx context.Context, role entity.Role, userID int64, req entity.GetAppointmentsByUserIDRequest) ([]entity.Appointment, error) {
	ret := _m.Called(ctx, role, userID, req)

	var r0 []entity.Appointment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Role, int64, entity.GetAppointmentsByUserIDRequest) ([]entity.Appointment, error)); ok {
		return rf(ctx, role, userID, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.Role, int64, entity.GetAppointmentsByUserIDRequest) []entity.Appointment); ok {
		r0 = rf(ctx, role, userID, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Appointment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.Role, int64, entity.GetAppointmentsByUserIDRequest) error); ok {
		r1 = rf(ctx, role, userID, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AppointmentRepository_GetAppointmentsByUserID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAppointmentsByUserID'
type AppointmentRepository_GetAppointmentsByUserID_Call struct {
	*mock.Call
}

// GetAppointmentsByUserID is a helper method to define mock.On call
//   - ctx context.Context
//   - role entity.Role
//   - userID int64
//   - req entity.GetAppointmentsByUserIDRequest
func (_e *AppointmentRepository_Expecter) GetAppointmentsByUserID(ctx interface{}, role interface{}, userID interface{}, req interface{}) *AppointmentRepository_GetAppointmentsByUserID_Call {
	return &AppointmentRepository_GetAppointmentsByUserID_Call{Call: _e.mock.On("GetAppointmentsByUserID", ctx, role, userID, req)}
}

func (_c *AppointmentRepository_GetAppointmentsByUserID_Call) Run(run func(ctx context.Context, role entity.Role, userID int64, req entity.GetAppointmentsByUserIDRequest)) *AppointmentRepository_GetAppointmentsByUserID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.Role), args[2].(int64), args[3].(entity.GetAppointmentsByUserIDRequest))
	})
	return _c
}

func (_c *AppointmentRepository_GetAppointmentsByUserID_Call) Return(_a0 []entity.Appointment, _a1 error) *AppointmentRepository_GetAppointmentsByUserID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AppointmentRepository_GetAppointmentsByUserID_Call) RunAndReturn(run func(context.Context, entity.Role, int64, entity.GetAppointmentsByUserIDRequest) ([]entity.Appointment, error)) *AppointmentRepository_GetAppointmentsByUserID_Call {
	_c.Call.Return(run)
	return _c
}

// InsertAppointment provides a mock function with given fields: ctx, req
func (_m *AppointmentRepository) InsertAppointment(ctx context.Context, req entity.MakeAppointmentRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.MakeAppointmentRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AppointmentRepository_InsertAppointment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InsertAppointment'
type AppointmentRepository_InsertAppointment_Call struct {
	*mock.Call
}

// InsertAppointment is a helper method to define mock.On call
//   - ctx context.Context
//   - req entity.MakeAppointmentRequest
func (_e *AppointmentRepository_Expecter) InsertAppointment(ctx interface{}, req interface{}) *AppointmentRepository_InsertAppointment_Call {
	return &AppointmentRepository_InsertAppointment_Call{Call: _e.mock.On("InsertAppointment", ctx, req)}
}

func (_c *AppointmentRepository_InsertAppointment_Call) Run(run func(ctx context.Context, req entity.MakeAppointmentRequest)) *AppointmentRepository_InsertAppointment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.MakeAppointmentRequest))
	})
	return _c
}

func (_c *AppointmentRepository_InsertAppointment_Call) Return(_a0 error) *AppointmentRepository_InsertAppointment_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AppointmentRepository_InsertAppointment_Call) RunAndReturn(run func(context.Context, entity.MakeAppointmentRequest) error) *AppointmentRepository_InsertAppointment_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, ID, data
func (_m *AppointmentRepository) Update(ctx context.Context, ID int64, data map[string]interface{}) error {
	ret := _m.Called(ctx, ID, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, map[string]interface{}) error); ok {
		r0 = rf(ctx, ID, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AppointmentRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type AppointmentRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - ID int64
//   - data map[string]interface{}
func (_e *AppointmentRepository_Expecter) Update(ctx interface{}, ID interface{}, data interface{}) *AppointmentRepository_Update_Call {
	return &AppointmentRepository_Update_Call{Call: _e.mock.On("Update", ctx, ID, data)}
}

func (_c *AppointmentRepository_Update_Call) Run(run func(ctx context.Context, ID int64, data map[string]interface{})) *AppointmentRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(map[string]interface{}))
	})
	return _c
}

func (_c *AppointmentRepository_Update_Call) Return(_a0 error) *AppointmentRepository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AppointmentRepository_Update_Call) RunAndReturn(run func(context.Context, int64, map[string]interface{}) error) *AppointmentRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewAppointmentRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewAppointmentRepository creates a new instance of AppointmentRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAppointmentRepository(t mockConstructorTestingTNewAppointmentRepository) *AppointmentRepository {
	mock := &AppointmentRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
