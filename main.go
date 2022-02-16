package main

import (
	"github.com/gofiber/fiber/v2"

	"fmt"

	"github.com/zhan860127/golang_test_for_brahma.im/database"
	"github.com/zhan860127/golang_test_for_brahma.im/weather"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(9000)
	fmt.Printf("hellow")
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "weather.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")

	database.DBConn.AutoMigrate(&weather.Weather{})
	fmt.Println("Database Migrated")

}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/weather", weather.Getweathers)
	app.Get("/api/v1/weather/:id", weather.Getweather)
	app.Post("/api/v1/weather", weather.Newweather)
	app.Delete("/api/v1/weather/:id/:month", weather.Deleteweather)
}
