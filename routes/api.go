package routes

import (
	"github.com/gofiber/fiber/v2"
)

func InitApiRoutes(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Yo start the go api frm app")
	})
}
