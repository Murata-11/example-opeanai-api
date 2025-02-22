package main

import (
	"app/router"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	router.Router(e)

	e.Logger.Fatal(e.Start(":1323"))
}
