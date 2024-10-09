package main

import (
	"net/http"

	"github.com/ebfe/scard"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	Context    *scard.Context
	CardReader *string = nil
	Card       *scard.Card
)

func main() {
	go ConnectCardReader()
	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/card-reader-connection", func(c echo.Context) error {
		readerResponse := GetCardReader()
		return c.JSON(http.StatusOK, readerResponse)
	})
	e.GET("/healthcard", func(c echo.Context) error {
		answer := GetHealthcardData()
		return c.JSON(http.StatusOK, answer)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
