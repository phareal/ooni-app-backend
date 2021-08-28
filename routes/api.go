package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"ooni/backend/controllers"
)

var (
	authController  = controllers.AuthenticationController()
)

func InitApiRoutes(app *fiber.App) {
	api := app.Group("/api", func( ctx *fiber.Ctx) error {
		return ctx.Next()
	})
	authentication(api)
}

func authentication(apiRouter fiber.Router) {
	fmt.Print("sdsds")
	authRouter := apiRouter.Group("/auth",func( ctx *fiber.Ctx) error {
		return ctx.Next()
	})
	authRouter.Post("/login",authController.Login)

}

func photos()  {
	
}
