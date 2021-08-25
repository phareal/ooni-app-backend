package routes

import (
	"github.com/gofiber/fiber/v2"
)

func InitApiRoutes(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Tchalley tu veux quoi la conf est bonne")
	})
	
}

func authentication(app *fiber.App)  {

}

func photos()  {
	
}
