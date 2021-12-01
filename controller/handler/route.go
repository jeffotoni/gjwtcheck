// Go Api server
// @jeffotoni
package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"

	jwtware "github.com/gofiber/jwt/v3"

	certKey "github.com/jeffotoni/gjwtcheck/core/cert"
	mw "github.com/jeffotoni/gjwtcheck/core/middleware"
	hd "github.com/jeffotoni/gjwtcheck/core/pkg/headers"
)

//AllRoutes todas as rotas
func AllRoutes(app *fiber.App, s StructConnect) {

	//MaxBody
	app.Use(mw.MaxBody(3 * 1024 * 1024)) //maximo para requests normais
	mw.Cors(app)
	mw.Logger(app)
	mw.Compress(app)
	mw.MsgUUID(app)

	app.Get("/ping", limiter.New(limiter.Config{
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

	// ------------------------------------------------
	auth1 := app.Group("rs256/")
	auth1.Post("/", limiter.New(limiter.Config{
		Next:       nil,
		Max:        100,
		Expiration: 1 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return hd.IP(c)
		},
		LimitReached: func(c *fiber.Ctx) error {
			//ou 401
			return c.Status(429).SendString(`{"msg":"Much Request #bloqued"}`)
		}}), s.RS256)

	auth1.Use(jwtware.New(jwtware.Config{
		SigningMethod: "RS256",
		SigningKey:    certKey.PublicKeyAuth,
	}))
	auth1.Post("/user", limiter.New(limiter.Config{
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

	// --------------------------------------------
	auth2 := app.Group("hs256/")
	auth2.Post("/", limiter.New(limiter.Config{
		Next:       nil,
		Max:        100,
		Expiration: 1 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return hd.IP(c)
		},
		LimitReached: func(c *fiber.Ctx) error {
			//ou 401
			return c.Status(429).SendString(`{"msg":"Much Request #bloqued"}`)
		}}), s.HS256)

	// JWT Middleware
	auth2.Use(jwtware.New(jwtware.Config{
		SigningMethod: "HS256",
		SigningKey:    []byte(certKey.SecretSH256),
	}))
	auth2.Post("/user", limiter.New(limiter.Config{
		Next:       nil,
		Max:        100,
		Expiration: 1 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return hd.IP(c)
		},
		LimitReached: func(c *fiber.Ctx) error {
			//ou 401
			return c.Status(429).SendString(`{"msg":"Much Request #bloqued"}`)
		}}), s.User2)

}
