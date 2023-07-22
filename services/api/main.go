package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mehdisadeghidev/croniq/config"
	"github.com/mehdisadeghidev/croniq/services/queue"
	"log"
	"strings"
)

type QueueInput struct {
	Name     string                 `json:"name"`
	Endpoint string                 `json:"endpoint"`
	Data     map[string]interface{} `json:"data"`
}

var app *fiber.App

func Setup() {
	app = fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")

		if authorization == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		token := strings.TrimPrefix(authorization, "Bearer ")

		if token != config.Token {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		return c.Next()
	})

	app.Post("/v1/queue", func(c *fiber.Ctx) error {
		input := new(QueueInput)

		if err := c.BodyParser(&input); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		queue.MainChannel <- queue.MainQueue{
			Name:     input.Name,
			Endpoint: input.Endpoint,
			Data:     input.Data,
		}

		return nil
	})

	go func() {
		if err := app.Listen(":4000"); err != nil {
			log.Panic(err)
		}
	}()
}

func Shutdown() error {
	return app.Shutdown()
}
