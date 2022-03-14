package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jason-adam/appointment-service/internal/handler"
)

func HealthRouter(app fiber.Router) {
	app.Get("/", handler.Health())
}
