package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)
func handleLogin(c echo.Context) error {
	o := ""
	o+="<html>"
	o+="<body>"
	o+="<h1>Login Page</h1>"
	o+="<p>"
	o+="<label>Username</label><input type='text' name='username'>"
	o+="<label>Password</label><input type='text' name='password'>"
	o+="</p>"
	o+="</body>"
	
	o+="</html>"

	
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

	e.GET("/login", handleLogin)

	e.GET("/users/:id", func(c echo.Context) error {
		id := c.Param("id")

		return c.String(http.StatusOK, "/users/:id"+id)
	})


    e.Static("/", "web")

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}