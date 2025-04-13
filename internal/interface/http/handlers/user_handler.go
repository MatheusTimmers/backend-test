package handlers

import (
	"github.com/MatheusTimmers/backend-test/internal/application/usecase/user"
	"github.com/MatheusTimmers/backend-test/pkg/models"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service user.UserService
}

func NewUserHandler(s user.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req models.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	user, err := h.service.Register(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *UserHandler) Ranking(c *fiber.Ctx) error {
	ranking, total, err := h.service.Ranking()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"total":   total,
		"ranking": ranking,
	})
}

func (h *UserHandler) NotifyWinners(c *fiber.Ctx) error {
	if err := h.service.NotifyWinners(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

