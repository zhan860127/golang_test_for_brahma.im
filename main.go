package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	_ "moduledemo/docs"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/redirect/v2"

	"moduledemo/database"
	"moduledemo/weather"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// @title Weather CRUD API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email zhan860127.gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:7800
// @BasePath /

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
	app.Get("/api/v1/new/:City/:month/:tempture", weather.Newtemp_api)
	app.Get("/api/v1/weather/:City/:month", weather.Getweather)
	app.Get("/api/v1/modify/:City/:month/:tempture", weather.Mod_temp)
	app.Get("/api/v1/del/:City/:month", weather.Deltemp)
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
			"/post": "/	",
		},
		StatusCode: 301,
	}))
	//	app.Post("/api/v1/weather", weather.Newweather)
	//	app.Delete("/api/v1/weather/:id/:month", weather.Deleteweather)
	app.Listen(":7800")
}

// ShowAccount godoc
// @Summary Show a tempture
// @Description get string  by City and month
// @City get-string-by-string
// @Accept  json
// @Produce  json
// @Param City path int true "City"
// @Success 200 {object} tempture
// @Failure 404 {object} HTTPError
// @Router /api/v1/weather/{City}/{month} [get]

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
