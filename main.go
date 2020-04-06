package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stianeikeland/go-rpio/v4"
	"net/http"
)

type status struct {
	Light bool `json:"light"`
}

func main() {
	rpio.Open()
	defer rpio.Close()

	e := echo.New()
	e.Use(middleware.CORS())

	e.Static("/", "./")

	api := e.Group("/api")
	{
		api.GET("/status", func(c echo.Context) error {
			pin := rpio.Pin(14)
			status := &status{
				Light: pin.Read() == 0,
			}
			return c.JSON(http.StatusOK, status)
		})

		api.POST("/on", func(c echo.Context) error {
			pin := rpio.Pin(14)
			pin.Output()
			return nil
		})

		api.POST("/off", func(c echo.Context) error {
			pin := rpio.Pin(14)
			pin.Input()
			return nil
		})
	}

	e.Logger.Fatal(e.Start(":1323"))
}
