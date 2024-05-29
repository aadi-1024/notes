package main

import (
	"log"

	"github.com/labstack/echo/v4"
)

var app = &Config{}

func main() {
	e := echo.New()

	app.debug = true

	SetupRouter(e)
	if err := e.Start("localhost:8080"); err != nil {
		log.Fatalln(err)
	}
}
