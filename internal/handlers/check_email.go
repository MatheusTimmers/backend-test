package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/MatheusTimmers/backend-test/internal/db"
	appError "github.com/MatheusTimmers/backend-test/internal/errors"
)

func CheckEmail(c *fiber.Ctx) error {
	email := c.Query("email")
		if email == "" {
		return appError.BadRequest("error: missing email parameter")
	}

	exist, err := db.EmailExist(email)
	if err != nil {
		return appError.DBError(err)
	}

	if exist {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "email already registered"})
	}

	return c.SendStatus(fiber.StatusNoContent) 
}
