package handler

import (
	"github.com/gofiber/fiber/v2"
	jwtGen "github.com/jeffotoni/gjwtcheck/core/pkg/jwt"
	mLg "github.com/jeffotoni/gjwtcheck/models/user"

	mw "github.com/jeffotoni/gjwtcheck/core/middleware"
	mErrors "github.com/jeffotoni/gjwtcheck/core/models/errors"
	"github.com/jeffotoni/gjwtcheck/core/pkg/fmts"
	. "github.com/jeffotoni/gjwtcheck/core/pkg/headers"
)

func (s StructConnect) HS256(c *fiber.Ctx) error {
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

	if len(user.User) == 0 {
		return c.Status(code).JSON(mErrors.Errors{ID: msgID, Msg: fmts.ConcatStr("Error: user required")})
	}

	if len(user.Password) == 0 {
		return c.Status(code).JSON(mErrors.Errors{ID: msgID, Msg: fmts.ConcatStr("Error: password required")})
	}

	if user.Time > 0 {
		jwtGen.SetExpires(user.Time)
	} else {
		jwtGen.SetExpires(3000)
	}

	if user.Secret, user.Key, user.Expires, err = jwtGen.TokenHS256(user.User, IP(c)); err != nil {
		code = 401
		return c.Status(code).JSON(mErrors.Errors{ID: msgID, Msg: fmts.ConcatStr("Error: when generating jwt - ", err.Error())})
	}

	code = 200
	return c.Status(code).JSON(user)
}
