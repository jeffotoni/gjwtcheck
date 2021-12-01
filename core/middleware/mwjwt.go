package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	cert "github.com/jeffotoni/gjwtcheck/core/cert"
)

//Mwjwt Fiber jwt middleware call
func Mwjwt(app *fiber.App) {
	app.Use(jwtware.New(jwtware.Config{
		SigningMethod: "RS256",
		SigningKey:    cert.PublicKey,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			c.Set("Content-Type", "application/json")
			if err.Error() == "Missing or malformed JWT" {
				return c.Status(fiber.StatusBadRequest).SendString(`{"msg":"Missing or malformed JWT"}`)
			} else {
				return c.Status(fiber.StatusUnauthorized).SendString(`{"msg":"token-invalid"}`)
			}
		},
	}))
	return
}
