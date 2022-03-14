package appointment

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var errUnableToBookAppt = errors.New("Appointment Slot is Unavailable or Booked")

type Repository interface {
	GetAppointmentsByTrainerAndDates(ctx context.Context, trainerID int64, startsAt, endsAt time.Time) ([]Appointment, error)
	GetBookedAppointments(ctx context.Context, trainerID int64) ([]Appointment, error)
	UpdateAppointment(ctx context.Context, apptRequest Request) error
}

type repository struct {
	dbPool *pgxpool.Pool
	loc    *time.Location
}

func NewRepository(opts ...Option) Repository {
	r := &repository{}

	for _, fn := range opts {
		fn(r)
	}

	return r
}

func (r *repository) GetAppointmentsByTrainerAndDates(ctx context.Context, trainerID int64, startsAt, endsAt time.Time) ([]Appointment, error) {
	rows, err := r.dbPool.Query(
		ctx,
		"SELECT * FROM appointments WHERE trainer_id = $1 AND status = $2 AND starts_at >= $3 AND ends_at <= $4",
		trainerID,
		Open,
		startsAt.Format(time.RFC3339),
		endsAt.Format(time.RFC3339),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appointments []Appointment

	for rows.Next() {
		var appt Appointment
		err := rows.Scan(
			&appt.ID,
			&appt.TrainerID,
			&appt.UserID,
			&appt.StartsAt,
			&appt.EndsAt,
			&appt.Status,
		)
		if err != nil {
			return nil, err
		}

		appt.StartsAt = appt.StartsAt.In(r.loc)
		appt.EndsAt = appt.EndsAt.In(r.loc)

		appointments = append(appointments, appt)
	}

	return appointments, nil
}

func (r *repository) GetBookedAppointments(ctx context.Context, trainerID int64) ([]Appointment, error) {
	rows, err := r.dbPool.Query(
		ctx,
		"SELECT * FROM appointments WHERE trainer_id = $1 AND status = $2",
		trainerID,
		Booked,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appointments []Appointment

	for rows.Next() {
		var appt Appointment
		err := rows.Scan(
			&appt.ID,
			&appt.TrainerID,
			&appt.UserID,
			&appt.StartsAt,
			&appt.EndsAt,
			&appt.Status,
		)
		if err != nil {
			return nil, err
		}

		appt.StartsAt = appt.StartsAt.In(r.loc)
		appt.EndsAt = appt.EndsAt.In(r.loc)

		appointments = append(appointments, appt)
	}

	return appointments, nil
}

func (r *repository) UpdateAppointment(ctx context.Context, apptRequest Request) error {
	tx, err := r.dbPool.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.RepeatableRead})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	row := tx.QueryRow(ctx, "SELECT * FROM appointments WHERE id = $1", apptRequest.AppointmentID)

	var appointment Appointment
	row.Scan(
		&appointment.ID,
		&appointment.TrainerID,
		&appointment.UserID,
		&appointment.StartsAt,
		&appointment.EndsAt,
		&appointment.Status,
	)

	if appointment.Status == Booked || appointment.Status == Unavailable {
		return errUnableToBookAppt
	}

	if _, err := tx.Exec(
		ctx,
		"UPDATE appointments SET user_id = $1, status = $2 WHERE id = $3",
		apptRequest.UserID,
		Booked,
		apptRequest.AppointmentID,
	); err != nil {
		return errUnableToBookAppt
	}

	if err := tx.Commit(ctx); err != nil {
		return errUnableToBookAppt
	}

	return nil
}
