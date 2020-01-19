package main

import (
	"flag"
	"fmt"
	"suggest-tree/handler"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	docroot := flag.String("docroot", "./assets", "specify document root path")
	port := flag.Int("port", 5000, "specify listening port")

	flag.Parse()
	// APIs
	e.GET("/api/suggests/:q/:depth", handler.Suggest)

	e.Static("/", *docroot)
	e.Start(fmt.Sprintf(":%d", *port))
}
