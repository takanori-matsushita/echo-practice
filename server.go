package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello world")
}

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")

	db, dbErr := gorm.Open("postgres", "host="+dbHost+" "+"port="+dbPort+" "+"dbname="+dbName+" "+" "+"sslmode=disable")
	if dbErr != nil {
		panic("データベースへの接続に失敗しました")
	}
	defer db.Close()

	e := echo.New()
	e.GET("/", hello)
	e.Logger.Fatal(e.Start(":1323"))
}
