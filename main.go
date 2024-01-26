package main

import (
	"book-crud/pkg/containers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	containers.Serve(e)
}
