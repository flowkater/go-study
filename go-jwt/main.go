package main

import (
	"net/http"

	_ "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	h := &handler{}
	e.POST("/login", h.login)
	e.GET("/isloggedin", h.restricted, IsLoggedIn)
	e.GET("/isadmin", h.restricted, IsLoggedIn, isAdmin)

	e.Logger.Fatal(e.Start(":1323"))
}
