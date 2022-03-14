package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jason-adam/appointment-service/internal/appointment"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_GetAppointmentsByTrainerAndDates(t *testing.T) {
	t.Run("Success_OnlyTrainerID", func(t *testing.T) {
		var (
			trainerID = int64(1)
			svc       = &MockAppointmentService{}
			app       = setupTestApp(svc)
		)

		svc.
			On(
				"GetAppointmentsByTrainerAndDates",
				mock.AnythingOfType("*fasthttp.RequestCtx"),
				trainerID,
				mock.AnythingOfType("time.Time"),
				mock.AnythingOfType("time.Time"),
			).
			Return([]appointment.Appointment{}, nil)

		resp, err := app.Test(httptest.NewRequest(http.MethodGet, "/api/appointments?trainerID=1", nil))

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("Fail_MissingTrainerID", func(t *testing.T) {
		var (
			svc = &MockAppointmentService{}
			app = setupTestApp(svc)
		)

		resp, err := app.Test(httptest.NewRequest(http.MethodGet, "/api/appointments", nil))

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}

func setupTestApp(svc appointment.Service) *fiber.App {
	app := fiber.New()
	api := app.Group("/api")

	api.Get("/appointments/:trainerID", GetAppointmentsByTrainer(svc))
	api.Get("/appointments", GetAppointmentsByTrainerAndDates(svc))
	api.Post("/appointments", BookAppointment(svc))

	return app
}

// Mock Appointment Service
var _ appointment.Service = &MockAppointmentService{}

type MockAppointmentService struct {
	mock.Mock
}

func (m *MockAppointmentService) GetAppointmentsByTrainerAndDates(ctx context.Context, trainerID int64, startsAt, endsAt time.Time) ([]appointment.Appointment, error) {
	args := m.Called(ctx, trainerID, startsAt, endsAt)

	return args.Get(0).([]appointment.Appointment), args.Error(1)
}

func (m *MockAppointmentService) GetBookedAppointments(ctx context.Context, trainerID int64) ([]appointment.Appointment, error) {
	args := m.Called(ctx, trainerID)

	return args.Get(0).([]appointment.Appointment), args.Error(1)
}

func (m *MockAppointmentService) UpdateAppointment(ctx context.Context, apptRequest appointment.Request) error {
	args := m.Called(ctx, apptRequest)

	return args.Error(0)
}
