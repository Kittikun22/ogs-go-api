package main

import (
	"ogs-create-api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	routes.UserRoute(e)

	e.Logger.Fatal(e.Start(":8000"))

}
