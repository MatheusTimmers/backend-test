package handlers

import (

	"github.com/MatheusTimmers/backend-test/internal/db"
	"github.com/MatheusTimmers/backend-test/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var req models.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

  user, err := db.CreateUser(req)
  if err != nil {
    return err 
  }

	shareLink := "http://localhost:8080/register?invite=" + user.InviteCode

	return c.JSON(fiber.Map{
		"message":     "registered successfully",
		"user":        user,
		"share_link":  shareLink,
	})
}
