package weather

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinburke/nacl"
	"github.com/kevinburke/nacl/secretbox"
	_ "github.com/mattn/go-sqlite3"
)

type Weather struct {
	City     string  `json:"city" xml:"city" form:"city"`
	Month    int     `json:"month" xml:"month" form:"month"`
	Tempture float64 `json:"tempture" xml:"tempture" form:"tempture"`
}

type User struct {
	Username string `json:"username" xml:username"" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
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
	db.Close()
	return c.JSON(b)
}

// @Summary Get City_month Tempture
// @Tags Weather
// @version 1.0
// @produce text/plain
// @param City header string true "City"
// @param month path int true "month"
// @Success 200 {number} 0.0 "查詢月份之城市溫度"
// @Failure 404 {string} Sorry,not found! "查詢失敗"
// @Router /api/v1/weather/{City}/{month} [get]
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
	var b int
	b = 1
	for rows.Next() {
		fmt.Println(a)
		rows.Scan(&temp)
		fmt.Println(temp)
		b = b + 1
	}

	strKm := strconv.FormatFloat(temp, 'f', 1, 64)

	if b > 1 {
		return c.SendString(string(strKm))
	} else {
		return fiber.NewError(404, "Sorry, not found!")
	}
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
	err = c.BodyParser(weather)
	//fmt.Println(user.Name)
	/*
		month := c.Params("month")
		city := c.Params("City")
		tempture := c.Params("tempture")

		text, _ := url.QueryUnescape(city)

		row, err := db.Query("insert into weather (city,tempture,month) values (" + text + "," + month + "',month=" + month + ",tempture=" + tempture)
	*/
	if err == nil && weather != nil {

		s := fmt.Sprintf("%f", weather.Tempture) // s == "123.456000"
		_, err = db.Exec("Insert into Weather (city,tempture,month) values ('" + weather.City + "'," + s + "," + strconv.Itoa(weather.Month) + ")")
		fmt.Println(err)
	} else {
		return fiber.NewError(404, "值不可為空")
	}
	if err == nil {
		return c.SendString("成功新增")
	} else {
		return fiber.NewError(404, "Sorry,新增失敗")
	}

}

func Signin(c *fiber.Ctx) error {
	db, err := sql.Open("sqlite3", "weather.db")
	checkErr(err)
	fmt.Println(c.Get("authorization"))
	user := new(User)
	//fmt.Println(c.Body())
	err = c.BodyParser(user)
	//fmt.Println(user.Name)
	/*
		month := c.Params("month")
		city := c.Params("City")
		tempture := c.Params("tempture")

		text, _ := url.QueryUnescape(city)

		row, err := db.Query("insert into weather (city,tempture,month) values (" + text + "," + month + "',month=" + month + ",tempture=" + tempture)
	*/
	if err == nil && user != nil {

		println(user.Username)
		println(user.Password)
		//		str := "select Username from User where Username='" + user.Username + "'and Password=" + "'" + user.Password + "'"
		row, err := db.Query("select * from User where Username='" + user.Username + "'and Password=" + "'" + user.Password + "'")
		println(err)
		println(&row)
		j := 0
		for row.Next() {

			j = j + 1
		}

		if j > 0 && err == nil {

			return (c.Redirect("/"))
		} else {
			return (c.Redirect("./sign/"))

		}

	} else {
		return fiber.NewError(404, "值不可為空")
	}

}

// @Summary delect the tempture in city on month
// @Tags Weather
// @version 1.0
// @produce text/plain
// @param City header string true "City"
// @param month path int true "month"
// @Success 200 {string} Sorry,not found! "成功刪除"
// @Failure 404 {string} Sorry,not  "刪除失敗"
// @Router /api/v1/del/{City}/{month}/{tempture} [get]
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
		return fiber.NewError(404, "Sorry,刪除失敗")
	}

}

// @Summary create the tempture in city on month
// @Tags Weather
// @version 1.0
// @produce text/plain
// @param City header string true "City"
// @param month path int true "month"
// @param tempture path number true "tempture"
// @Success 200 {string} Sorry,not found! "成功新增"
// @Failure 404 {string} Sorry,not  "查詢失敗"
// @Router /api/v1/new/{City}/{month}/{tempture} [get]
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
		return fiber.NewError(404, "Sorry,新增失敗")
	}

}

// @Summary Modify the tempture in city on month
// @Tags Weather
// @version 1.0
// @produce text/plain
// @param City header string true "City"
// @param month path int true "month"
// @param tempture path number true "tempture"
// @Success 200 {string} Sorry,not found! "成功修改"
// @Failure 404 {string} Sorry,not  "修改失敗"
// @Router /api/v1/modify/{City}/{month}/{tempture} [get]
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
		return fiber.NewError(404, "Sorry,修改失敗")
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

// @Summary encrpt message
// @Tags Encrpt
// @version 1.0
// @produce text/plain
// @param message header string true "message"
// @Success 200 {string} 0.0 "成功加密訊息"
// @Router /encrpt/{message} [get]
func Encrpt(c *fiber.Ctx) error {
	key, err := nacl.Load("6368616e676520746869732070617373776f726420746f206120736563726574")
	if err != nil {
		panic(err)
	}
	message := c.Params("message")
	data := []byte(message)
	encrypted := secretbox.EasySeal(data, key)

	encrypted_message := string(base64.StdEncoding.EncodeToString(encrypted))
	fmt.Println(encrypted_message)
	return c.SendString(encrypted_message)

}
