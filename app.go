package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"ooni/backend/core/database"
	"ooni/backend/routes"
	"os"
)

func main() {
	app := fiber.New()
	//load the env variable
	errEnv := godotenv.Load()

	if os.Getenv("APP_ENV")!= "production" {
		if errEnv != nil {
			panic("Unable to load the environnement file")
		}
	}
	database.InitDatabase()
	routes.InitApiRoutes(app)
	appPort := os.Getenv("PORT")
	err := app.Listen(":" + appPort)
	if err != nil {
		fmt.Println(err)
		panic("Unable start the app")
	}

}
