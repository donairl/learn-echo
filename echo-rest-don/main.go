package main

import (
	"net/http"

	m "./models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var db *gorm.DB
var err error

// GetUsers get all user
func GetUsers(c echo.Context) error {
	//db := db.DbManager()
	users := []m.User{}
	db.Find(&users)

	return c.JSON(http.StatusOK, users)
}

// CreateUser new user
func CreateUser(c echo.Context) error {
	u := &m.User{}
	if err := c.Bind(u); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, u)
}

func isidb() {
	db.AutoMigrate(&m.User{}, &m.Article{})

	user := m.User{Username: "donny", Password: "cobasekali", Email: "yahmo@sekali.com", Role: 1}

	db.NewRecord(user) // => returns `true` as primary key is blank

	db.Create(&user)
}

func main() {

	db, err = gorm.Open(
		"postgres",
		"host=127.0.0.1 port=5432 user=postgres  dbname=dblearn sslmode=disable")

	if err != nil {

		panic(err)
	}

	// Echo instance

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/about", func(c echo.Context) error {

		return c.String(http.StatusOK, "Hello, Indonesia!\n")
	})

	e.GET("/users", GetUsers)

	//e.GET("/help", controllers.help)

	e.Static("/", "web")

	// Start server
	e.Logger.Fatal(e.Start(":8080"))

}
