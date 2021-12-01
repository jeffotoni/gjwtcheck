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

func (s StructConnect) Check2(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	var err error

	code := 400
	msgID := mw.GetUUID(c)
	if len(string(c.Body())) <= 0 {
		return c.Status(code).JSON(mErrors.Errors{ID: msgID, Msg: fmts.ConcatStr("Error: empty body")})
	}

	var user mLg.UserAuth
	if err := c.BodyParser(&user); err != nil {
		code = 400
		return c.Status(code).JSON(mErrors.Errors{ID: msgID, Msg: fmts.ConcatStr("Error: ", err.Error())})
	}

	jwtGen.SetExpires(60)
	if user.Key, user.Expires, err = jwtGen.TokenHS256(user.User, hd.IP(c)); err != nil {
		code = 401
		return c.Status(code).JSON(mErrors.Errors{ID: msgID, Msg: fmts.ConcatStr("Error: when generating jwt - ", err.Error())})
	}

	code = 200
	return c.Status(code).JSON(user)
}
