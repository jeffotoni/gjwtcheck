package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

//Logger log
func Logger(app *fiber.App) {
	if os.Getenv("ENV_AMBI") != "PROD" {
		app.Use(logger.New(logger.Config{
			Format:     "\u001b[91m${pid}\u001b[0m ${time} \u001b[93m${method}\u001b[0m ${path} - ${ip} - \u001b[92m${status}\u001b[0m - \u001b[94m${latency}\u001b[0m\n",
			TimeFormat: "02-Jan-2006 15:04:05",
			Output:     os.Stdout,
		}))
	}
	return
}
