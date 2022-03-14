package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jason-adam/appointment-service/internal/appointment"
	"github.com/jason-adam/appointment-service/internal/handler"
)

func AppointmentRouter(app fiber.Router, svc appointment.Service) {
	app.Get("/appointments", handler.GetAppointmentsByTrainerAndDates(svc))
	app.Get("/appointments/:trainerID", handler.GetAppointmentsByTrainer(svc))
	app.Post("/appointments", handler.BookAppointment(svc))
}
