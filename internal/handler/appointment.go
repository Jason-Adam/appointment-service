package handler

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jason-adam/appointment-service/internal/appointment"
)

func GetAppointmentsByTrainerAndDates(svc appointment.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse trainerID
		trainerID := c.Query("trainerID")
		if trainerID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(errorResponse(errMissingTrainerID))
		}

		tID, err := strconv.ParseInt(trainerID, 10, 64)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(errorResponse(errParseTrainerID))
		}

		// Parse start
		start := c.Query("start", time.Now().Format(time.RFC3339))
		startsAt, err := time.Parse(time.RFC3339, start)
		if err != nil {
			log.Println("missing or unable to parse start param - defaulting to now")
			startsAt = time.Now()
		}

		end := c.Query("end")
		endsAt, err := time.Parse(time.RFC3339, end)
		if err != nil {
			log.Println("missing or unable to parse end param - defaulting to 4 week window")
			endsAt = startsAt.Add(28 * 24 * time.Hour)
		}

		appointments, err := svc.GetAppointmentsByTrainerAndDates(c.Context(), tID, startsAt, endsAt)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(errorResponse(err))
		}

		return c.JSON(successResponse(appointments))
	}
}

func GetAppointmentsByTrainer(svc appointment.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse trainerID
		trainerID := c.Params("trainerID")
		if trainerID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(errorResponse(errMissingTrainerID))
		}

		tID, err := strconv.ParseInt(trainerID, 10, 64)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(errorResponse(errParseTrainerID))
		}

		appointments, err := svc.GetBookedAppointments(c.Context(), tID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(errorResponse(err))
		}

		return c.JSON(successResponse(appointments))
	}
}

func BookAppointment(svc appointment.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var apptRequest appointment.Request
		if err := c.BodyParser(&apptRequest); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(errorResponse(err))
		}

		if err := svc.UpdateAppointment(c.Context(), apptRequest); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(errorResponse(err))
		}

		return c.JSON(successResponse(&fiber.Map{
			"id":     apptRequest.AppointmentID,
			"status": appointment.Booked,
		}))
	}
}
