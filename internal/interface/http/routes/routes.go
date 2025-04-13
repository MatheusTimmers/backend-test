package routes

import (
	"time"

	appErr "github.com/MatheusTimmers/backend-test/pkg/errors"

	"github.com/MatheusTimmers/backend-test/internal/application/usecase/user"
	"github.com/MatheusTimmers/backend-test/internal/interface/http/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RegisterRoutes(app *fiber.App, userService user.UserService) {
	handler := handlers.NewUserHandler(userService)

	app.Use(appErr.ErrorMiddleware)

	app.Use(limiter.New(limiter.Config{
		Max:        10,
		Expiration: 1 * time.Minute,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests. Please wait a bit.",
			})
		},
	}))

	app.Post("/register", handler.Register)
	app.Get("/ranking", handler.Ranking)
	app.Post("/notify-winners", handler.NotifyWinners)
}
