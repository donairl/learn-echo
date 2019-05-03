package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

//Myhandler : internal html file
func Myhandler(c echo.Context) error {
	o := "<h1>Iyes</h1>"
	return c.HTML(http.StatusOK, o)
}
