// Go Api server
// @jeffotoni
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jeffotoni/gjwtcheck/config"
	route "github.com/jeffotoni/gjwtcheck/controller/handler"
)

func main() {
	app := fiber.New(fiber.Config{BodyLimit: 10 * 1024 * 1024})
	route.AllRoutes(app, route.StructConnect{})
	app.Listen(config.HTTPPORT)
}
