package handler

import (
	"github.com/gofiber/fiber/v2"
	mw "github.com/jeffotoni/gjwtcheck/apicore/middleware"
)

//Ping pong
func (s StructConnect) Ping(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	msgID := mw.GetUUID(c)
	return c.Status(200).SendString(`{"pong":"🏓"}`)
}
