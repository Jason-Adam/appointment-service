package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

var (
	errMissingTrainerID = errors.New("Missing trainerID path param")
	errParseTrainerID   = errors.New("Unable to parse trainerID provided")
	errMissingStart     = errors.New("Missing start path param")
)

func errorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}

func successResponse(arg interface{}) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   arg,
		"error":  nil,
	}
}
