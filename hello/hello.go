package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//internal html file
func handleHelp(c echo.Context) error {
	o := ""
	o += "<html>"
	o += `<head><script src="js\main.js" ></script></head>`
	o += `<body onload="init()">`
	o += "<h1>Help Page</h1>"
	o += "<p>"
	o += "<label>Search</label><input type='text' name='username'>"
	o += "<button id='find'>Find</button>"
	o += "</p>"
	o += `<div id="result">
	       Result 
	      </div>
		 
		 `
	o += "</body>"

	o += "</html>"

	return c.HTML(http.StatusOK, o)
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/about", func(c echo.Context) error {

		return c.String(http.StatusOK, "Hello, Indonesia!\n")
	})

	e.GET("/help", handleHelp)

	e.GET("/users/:id", func(c echo.Context) error {
		id := c.Param("id")

		return c.String(http.StatusOK, "Anda memanggil lewat /users/"+id)
	})

	e.Static("/", "web")

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
