package weather

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"strconv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

type Weather struct {
	City     string  `json:"city" xml:"city" form:"city"`
	Month    int     `json:"month" xml:"month" form:"month"`
	Tempture float64 `json:"tempture" xml:"tempture" form:"tempture"`
}

type Weathers struct {
	City string  `json:"City"`
	Jan  float64 `json:"Jan"`
	Feb  float64 `json:"Feb"`
	Mar  float64 `json:"Mar"`
	Apr  float64 `json:"Apr"`
	Jun  float64 `json:"Jun"`
	Jul  float64 `json:"Jul"`
	Aug  float64 `json:"Aug"`
	Sep  float64 `json:"Sep"`
	Oct  float64 `json:"Oct"`
	Nov  float64 `json:"Nov"`
	Dec  float64 `json:"Dec"`
}

func Getweathers(c *fiber.Ctx) error {
	db, err := sql.Open("sqlite3", "weather.db")
	checkErr(err)
	rows, err := db.Query("SELECT City,jan,feb,march,apr,june,july,Aug,sep,oct,nov,dec FROM weatherss")
	checkErr(err)
	var b []Weathers

	for rows.Next() {
		a := new(Weathers)
		rows.Scan(&a.City, &a.Jan, &a.Feb, &a.Mar, &a.Apr, &a.Jun, &a.Jul, &a.Aug, &a.Sep, &a.Oct, &a.Nov, &a.Dec)
		//fmt.Printf("%v %v %v %v %v %v %v\n", a.City, a.Month, a.Tempture)
		//	fmt.Println(a)
		b = append(b, *a)

	}

	defer rows.Close()
	return c.JSON(b)
}

func Getweather(c *fiber.Ctx) error {
	db, err := sql.Open("sqlite3", "weather.db")
	checkErr(err)

	a := c.Params("month")
	//	month, _ := strconv.Atoi(a)
	city := c.Params("City")
	//	fmt.Printf(city)
	text, _ := url.QueryUnescape(city)
	//	fmt.Printf("%v %T", month, month)
	str := "select tempture from Weather where City='" + text + "' AND month=" + a
	//	fmt.Println(str)
	rows, err := db.Query(str)
	//	fmt.Println(rows)
	checkErr(err)
	//	fmt.Println(rows)
	var temp float64

	for rows.Next() {
		rows.Scan(&temp)
		fmt.Println(temp)
	}
	strKm := strconv.FormatFloat(temp, 'f', 1, 64)
	return c.SendString(string(strKm))
}

/*
func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	c.JSON(book)
}
*/

func Newtemp(c *fiber.Ctx) error {
	db, err := sql.Open("sqlite3", "weather.db")
	checkErr(err)
	fmt.Println(c.Get("authorization"))
	weather := new(Weather)
	//fmt.Println(c.Body())
	if err := c.BodyParser(weather); err != nil {
		//		fmt.Printf(weather.City)

		log.Fatal(weather)
		//		fmt.Printf(weather.City)

	}
	//fmt.Println(user.Name)
	/*
		month := c.Params("month")
		city := c.Params("City")
		tempture := c.Params("tempture")

		text, _ := url.QueryUnescape(city)

		row, err := db.Query("insert into weather (city,tempture,month) values (" + text + "," + month + "',month=" + month + ",tempture=" + tempture)
	*/
	s := fmt.Sprintf("%f", weather.Tempture) // s == "123.456000"
	row, err := db.Exec("Insert into Weather (city,tempture,month) values ('" + weather.City + "'," + s + "," + strconv.Itoa(weather.Month) + ")")

	if row != nil {
		return c.SendString("成功新增")
	} else {
		return c.SendString("新增失敗")
	}

}

func Deltemp(c *fiber.Ctx) error {
	db, err := sql.Open("sqlite3", "weather.db")
	checkErr(err)
	month := c.Params("month")
	city := c.Params("City")
	text, _ := url.QueryUnescape(city)

	row, err := db.Prepare("delete from Weather where city='" + text + "'and month=" + month)
	checkErr(err)
	_, err = row.Exec()
	if err != nil {
		return c.SendString("成功刪除")
	} else {
		return c.SendString("刪除失敗")
	}

}

func Newtemp_api(c *fiber.Ctx) error {
	db, err := sql.Open("sqlite3", "weather.db")
	checkErr(err)

	month := c.Params("month")
	city := c.Params("City")
	text, _ := url.QueryUnescape(city)
	s := c.Params("tempture")
	str := "Insert into Weather (City,tempture,month) values ('" + text + "'," + s + "," + month + ")"
	fmt.Println(str)
	_, err = db.Exec("Insert into Weather (City,tempture,month) values ('" + text + "'," + s + "," + month + ")")
	if err == nil {
		return c.SendString("成功新增")
	} else {
		return c.SendString("新增失敗,已有")
	}

}

func Mod_temp(c *fiber.Ctx) error {
	db, err := sql.Open("sqlite3", "weather.db")
	checkErr(err)

	month := c.Params("month")
	city := c.Params("City")
	text, _ := url.QueryUnescape(city)
	s := c.Params("tempture")
	str := "Update Weather set tempture=" + s + " where City='" + text + "' and month=" + month
	fmt.Println(str)
	_, err = db.Exec("Update Weather set tempture=" + s + " where City='" + text + "'and month=" + month)
	checkErr(err)
	fmt.Println(err)

	if err == nil {
		return c.SendString("成功修改")
	} else {
		return c.SendString("修改失敗")
	}

}

/*
func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("No book found with given ID")
		return
	}
	db.Delete(&book)
	c.Send("Book successfully deleted")
}
*/

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
