package handler

import (
	"github.com/gofiber/fiber/v2"
	jwtGen "github.com/jeffotoni/gjwtcheck/apicore/pkg/jwt"
	mLg "github.com/jeffotoni/gjwtcheck/models/user"

	mw "github.com/jeffotoni/gjwtcheck/apicore/middleware"
	mErrors "github.com/jeffotoni/gjwtcheck/apicore/models/errors"
	"github.com/jeffotoni/gjwtcheck/apicore/pkg/fmts"
	hd "github.com/jeffotoni/gjwtcheck/apicore/pkg/headers"
)

func (s StructConnect) Check(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	code := 400
	msgID := mw.GetUUID(c)
	if len(string(c.Body())) <= 0 {
		return c.Status(code).JSON(`{"msg":"Error empty body"}`)
	}

	var user mLg.UserAuth
	if err := c.BodyParser(&user); err != nil {
		code = 400
		return c.Status(code).JSON(mErrors.Errors{Msg: fmts.ConcatStr("Error: ", err.Error())})
	}

	if user.Key, user.Expires, err = jwtGen.Token(user.User, hd.IP(c)); err != nil {
		code = 401
		return c.Status(code).JSON(mErrors.Errors{Msg: fmts.ConcatStr("Error: when generating jwt - ", err.Error())})
	}

	code = 200
	return c.Status(code).JSON(user)
}
