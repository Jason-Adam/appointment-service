package appointment

import "time"

type AppointmentStatus string

const (
	Open        AppointmentStatus = "OPEN"
	Booked      AppointmentStatus = "BOOKED"
	Unavailable AppointmentStatus = "UNAVAILABLE"
)

// Appointment holds relevant appointment information
type Appointment struct {
	ID        int64             `json:"id"`
	TrainerID int64             `json:"trainer_id"`
	UserID    *int64            `json:"user_id"`
	StartsAt  time.Time         `json:"starts_at"`
	EndsAt    time.Time         `json:"ends_at"`
	Status    AppointmentStatus `json:"appointment_status"`
}

// Request represents a minimal body for booking an appointmeent.
// Only the appointmentI ID and UserID are needed to attempt a booking.
type Request struct {
	AppointmentID int64 `json:"id"`
	UserID        int64 `json:"user_id"`
}
