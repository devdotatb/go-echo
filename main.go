package main

import (
	"golang-echo/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	s := service.NewServiceTest("test")

	echosetting(e, s)

	e.Logger.Fatal(e.Start(":1323"))
}

func echosetting(e *echo.Echo, s service.ServiceInterface) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, s.GetHello())
	})
	e.GET("/updatehello/:txt", func(c echo.Context) error {
		txt := c.Param("txt")
		s.UpdateHello(txt)
		return c.NoContent(http.StatusOK)
	})
}
