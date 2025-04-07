package errors

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func ErrorMiddleware(c *fiber.Ctx) error {
	err := c.Next()
	if err == nil {
		return nil
	}

	var appErr *AppError
	if errors.As(err, &appErr) {
		return c.Status(appErr.StatusCode).JSON(fiber.Map{
			"error": appErr.Message,
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "unexpected error",
	})
}
