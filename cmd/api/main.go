package main

import (
	"log"

	"github.com/MatheusTimmers/backend-test/internal/db"
	appErr "github.com/MatheusTimmers/backend-test/internal/errors"
	"github.com/MatheusTimmers/backend-test/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db.Connect()

	app := fiber.New()

  app.Use(appErr.ErrorMiddleware)

  app.Post("/register", handlers.Register)
  app.Get("/ranking", handlers.Ranking)

	log.Fatal(app.Listen(":8080"))
}
