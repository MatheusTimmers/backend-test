package main

import (
	"time"

	"github.com/MatheusTimmers/backend-test/internal/logger"

	"github.com/MatheusTimmers/backend-test/internal/db"
	appErr "github.com/MatheusTimmers/backend-test/internal/errors"
	"github.com/MatheusTimmers/backend-test/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	logger.Init()
	defer logger.Log.Sync()

	db.Connect()

	app := fiber.New()

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

	app.Post("/register", handlers.Register)
	app.Get("/ranking", handlers.Ranking)
	app.Get("/check-email", handlers.CheckEmail)

	logger.Log.Info("Server started on :8080")
	logger.Log.Error(app.Listen(":8080"))
}
