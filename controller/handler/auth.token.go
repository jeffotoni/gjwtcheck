package handler

import (
	"github.com/gofiber/fiber/v2"
	mw "github.com/jeffotoni/gjwtcheck/apicore/middleware"
	mErrors "github.com/jeffotoni/gjwtcheck/apicore/models/errors"
	mjwt "github.com/jeffotoni/gjwtcheck/apicore/models/jwt"
	fmts "github.com/jeffotoni/gjwtcheck/apicore/pkg/fmts"
	hd "github.com/jeffotoni/gjwtcheck/apicore/pkg/headers"
	jwtCore "github.com/jeffotoni/gjwtcheck/apicore/pkg/jwt"
)

func (s StructConnect) Token(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	msgID := mw.GetUUID(c)
	code := 400
	authorization := string(c.Request().Header.Peek("X-Authorization"))
	if len(string(authorization)) <= 0 {
		return c.Status(code).JSON(mErrors.Errors{Msg: `Error: failed to try to generate token`})
	}

	var response mjwt.ResponseToken
	if response.Token, response.Expires, err = jwtCore.Token(authorization, hd.IP(c)); err != nil {
		code = 401
		return c.Status(code).JSON(mErrors.Errors{Msg: fmts.ConcatStr("Error: ", err.Error())})
	}

	code = 200
	response.Message = "Welcome"
	return c.Status(code).JSON(response)
}
