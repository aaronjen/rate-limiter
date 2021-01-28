package main

import (
	"ratelimiter/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	h := handler.Handler{
		Config: &handler.Config{
			LimitTimes:       60,
			ExpirationSecond: 60,
		},
	}

	app.Get("/", h.GetHandler)

	app.Listen(":3000")
}
