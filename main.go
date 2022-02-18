package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gofiber/fiber/v2"

	"moduledemo/database"
	"moduledemo/weather"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	fmt.Printf("hellow")
}

func initDatabase() {
	//var err error
	_, err := sql.Open("sqlite3", "weather.db")
	checkErr(err)

	//	rows, err := db.Query("SELECT* FROM Weather")
	checkErr(err)

	//db := database.DBConn
	fmt.Println("Database connection successfully opened")
	//fmt.Println(rows)
	//var weather weather.Weather

}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/weather", weather.Getweathers)
	app.Static("/", "./weather.html")
	app.Post("/post", weather.Newtemp)
	app.Get("api/v1/new/:City/:month/:tempture", weather.Newtemp_api)
	app.Get("/api/v1/weather/:City/:month", weather.Getweather)
	app.Get("api/v1/modify/:City/:month/:tempture", weather.Mod_temp)
	app.Get("/api/v1/del/:City/:month", weather.Deltemp)
	//	app.Post("/api/v1/weather", weather.Newweather)
	//	app.Delete("/api/v1/weather/:id/:month", weather.Deleteweather)
	app.Listen(":7800")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
