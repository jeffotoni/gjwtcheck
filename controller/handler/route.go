// Go Api server
// @jeffotoni
package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	jwtware "github.com/gofiber/jwt/v2"
	mw "github.com/jeffotoni/gjwtcheck/apicore/middleware"
	hd "github.com/jeffotoni/gjwtcheck/apicore/pkg/headers"

	//certKey "github.com/jeffotoni/gjwtcheck/cert"
	certKey "github.com/jeffotoni/gjwtcheck/apicore/cert"
)

//AllRoutes todas as rotas
func AllRoutes(app *fiber.App, s StructConnect) {
	//MaxBody
	app.Use(mw.MaxBody(3 * 1024 * 1024)) //maximo para requests normais
	mw.Cors(app)
	mw.Logger(app)
	mw.Compress(app)
	mw.MsgUUID(app)

	app.Post("/auth/check", limiter.New(limiter.Config{
		Next:       nil,
		Max:        100,
		Expiration: 1 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return hd.IP(c)
		},
		LimitReached: func(c *fiber.Ctx) error {
			//ou 401
			return c.Status(429).SendString(`{"msg":"Much Request #bloqued"}`)
		}}), s.Check)

	app.Get("/auth/token", limiter.New(limiter.Config{
		Next:       nil,
		Max:        100,
		Expiration: 1 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return hd.IP(c)
		},
		LimitReached: func(c *fiber.Ctx) error {
			//ou 401
			return c.Status(429).SendString(`{"msg":"Much Request #bloqued"}`)
		}}), s.Token)

	app.Get("/auth/ping", limiter.New(limiter.Config{
		Next:       nil,
		Max:        100,
		Expiration: 1 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return hd.IP(c)
		},
		LimitReached: func(c *fiber.Ctx) error {
			//ou 401
			return c.Status(429).SendString(`{"msg":"Much Request #bloqued"}`)
		}}), s.Ping)

	app.Use(jwtware.New(jwtware.Config{
		SigningMethod: "RS256",
		SigningKey:    certKey.PublicKeyAuth,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			c.Set("Content-Type", "application/json")
			if err.Error() == "Missing or malformed JWT" {
				return c.Status(fiber.StatusUnauthorized).SendString(`{"msg":"Missing or malformed JWT"}`)
			}
			return c.Status(fiber.StatusUnauthorized).SendString(`{"msg":"token-invalid"}`)
		},
	}))

	app.Post("/auth/user", limiter.New(limiter.Config{
		Next:       nil,
		Max:        100,
		Expiration: 1 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return hd.IP(c)
		},
		LimitReached: func(c *fiber.Ctx) error {
			//ou 401
			return c.Status(429).SendString(`{"msg":"Much Request #bloqued"}`)
		}}), s.User)
}
