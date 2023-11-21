package main

import (
	"github.com/labstack/echo/v4"
)

type R struct {
	A int
	B int
}

func main() {
	e := echo.New()

	e.POST("/api/add", func(c echo.Context) error {
		rq := new(R)

		c.Bind(rq)

		return c.JSON(200, echo.Map{
			"result": rq.A + rq.B,
		})
	})

	e.POST("/api/subtract", func(c echo.Context) error {
		rq := new(R)

		c.Bind(rq)

		return c.JSON(200, echo.Map{
			"result": rq.A - rq.B,
		})
	})

	e.POST("/api/multiply", func(c echo.Context) error {
		rq := new(R)

		c.Bind(rq)

		return c.JSON(200, echo.Map{
			"result": rq.A * rq.B,
		})
	})

	e.POST("/api/div", func(c echo.Context) error {
		rq := new(R)

		c.Bind(rq)

		return c.JSON(200, echo.Map{
			"result": rq.A / rq.B,
		})
	})

	e.Start(":1323")
}
