package main

import (
	"net/http"

	cnt "./controller"
	rice "github.com/GeertJohan/go.rice"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// the file server for rice. "web" is the folder where the files come from.
	assetHandler := http.FileServer(rice.MustFindBox("web").HTTPBox())

	// serves the index.html from rice
	e.GET("/", echo.WrapHandler(assetHandler))
	e.GET("/css/*", echo.WrapHandler(assetHandler))

	// servers other static files
	e.GET("/page/*", echo.WrapHandler(http.StripPrefix("/page/", assetHandler)))
	e.GET("/site", cnt.Myhandler)
	e.Logger.Fatal(e.Start(":8181"))

}
