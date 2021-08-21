package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/phareal/ooni-app-backend/core/database"
	"github.com/phareal/ooni-app-backend/routing"
	"os"
)

func main() {
	app := fiber.New()
	//load the env variable
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Unable to load the environnement file")
	}
	database.InitDatabase()
	routing.InitApiRoutes(app)
	appPort := os.Getenv("PORT")
	err := app.Listen(":" + appPort)
	if err != nil {
		fmt.Println(err)
		panic("Unable start the app")
	}

}
