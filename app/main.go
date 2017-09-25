package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name" gorm:"size:255"`
	Email string `json:"email"`
	gorm.Model
}

var db *gorm.DB
var dbConnectionFailCount int = 0

func setUpDbConnection() {
	dbConnectionFailCount++
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/user?charset=utf8&parseTime=True", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST")))
	if err != nil {
		if dbConnectionFailCount >= 10 {
			panic(err)
		}
		fmt.Println(err)
		time.Sleep(5 * time.Second)
		setUpDbConnection()
	}
}

func main() {
	setUpDbConnection()
	e := echo.New()
	e.GET("/users", indexUser)
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":4000"))
}

func indexUser(c echo.Context) error {
	u := &[]User{}
	db.Find(&u)
	return c.JSONPretty(http.StatusCreated, u, "  ")
}

func createUser(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	u := &User{Name: name, Email: email}
	db.Save(&u)
	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	u := &User{}
	db.Find(&u, id)
	return c.JSON(http.StatusOK, u)
}

func updateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	u := &User{}
	db.Find(&u, id)
	db.Model(&u).Updates(map[string]interface{}{"name": c.FormValue("name"), "email": c.FormValue("email")})
	return c.JSON(http.StatusOK, u)
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	u := &User{}
	db.Find(&u, id)
	db.Delete(&u)
	return c.NoContent(http.StatusNoContent)
}
