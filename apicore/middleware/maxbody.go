package middleware

import "github.com/gofiber/fiber/v2"

//MaxBody tamanho maximo da request
func MaxBody(size int) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if len(c.Body()) >= size {
			return fiber.ErrRequestEntityTooLarge
		}
		return c.Next()
	}
}
