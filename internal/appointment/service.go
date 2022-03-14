package appointment

import (
	"context"
	"time"
)

type Service interface {
	GetAppointmentsByTrainerAndDates(ctx context.Context, trainerID int64, startsAt, endsAt time.Time) ([]Appointment, error)
	GetBookedAppointments(ctx context.Context, trainerID int64) ([]Appointment, error)
	UpdateAppointment(ctx context.Context, apptRequest Request) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetAppointmentsByTrainerAndDates(ctx context.Context, trainerID int64, startsAt, endsAt time.Time) ([]Appointment, error) {
	return s.repo.GetAppointmentsByTrainerAndDates(ctx, trainerID, startsAt, endsAt)
}

func (s *service) GetBookedAppointments(ctx context.Context, trainerID int64) ([]Appointment, error) {
	return s.repo.GetBookedAppointments(ctx, trainerID)
}

func (s *service) UpdateAppointment(ctx context.Context, apptRequest Request) error {
	return s.repo.UpdateAppointment(ctx, apptRequest)
}
