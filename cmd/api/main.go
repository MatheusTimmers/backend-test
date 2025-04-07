package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/MatheusTimmers/backend-test/internal/db"
  "github.com/MatheusTimmers/backend-test/internal/handlers"
)

func main() {
	db.Connect()

	app := fiber.New()

  app.Post("/register", handlers.Register)

	log.Fatal(app.Listen(":8080"))
}
