package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	FIBER_TIMEOUT = 20
)

var App *fiber.App
var API fiber.Router

func init() {
	App = fiber.New(fiber.Config{
		ReadTimeout:  FIBER_TIMEOUT * time.Second,
		WriteTimeout: FIBER_TIMEOUT * time.Second,
		IdleTimeout:  FIBER_TIMEOUT * time.Second,
		BodyLimit:    100 * 1024 * 1024,
	})

	App.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

}
