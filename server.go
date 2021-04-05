package main

import (
	"app/database"
	config "app/routes"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
	database.Connect()
	e := echo.New()
	config.InitRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
