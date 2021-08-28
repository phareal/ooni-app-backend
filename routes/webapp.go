package routes

import "github.com/gofiber/fiber/v2"

func SetWebappRoutes(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return  ctx.Render("index",fiber.Map{})
	})
	app.Get("/login", func(ctx *fiber.Ctx) error {
		return  ctx.Render("login",fiber.Map{})
	})

}
