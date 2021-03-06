package handler

import (
	"github.com/gofiber/fiber/v2"
	jwtGen "github.com/jeffotoni/gjwtcheck/core/pkg/jwt"
	mLg "github.com/jeffotoni/gjwtcheck/models/user"

	mw "github.com/jeffotoni/gjwtcheck/core/middleware"
	mErrors "github.com/jeffotoni/gjwtcheck/core/models/errors"
	"github.com/jeffotoni/gjwtcheck/core/pkg/fmts"
)

func (s StructConnect) RS256(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	var err error

	code := 400
	msgID := mw.GetUUID(c)
	if len(string(c.Body())) <= 0 {
		return c.Status(code).JSON(mErrors.Errors{ID: msgID, Msg: fmts.ConcatStr("Error: empty body")})
	}

	// var claim jwtGen.Claims
	// if err := c.BodyParser(&claim); err != nil {
	// 	return c.Status(code).JSON(mErrors.Errors{ID: msgID, Msg: fmts.ConcatStr("Error: ", err.Error())})
	// }

	var user mLg.UserAuth
	if err := c.BodyParser(&user); err != nil {
		return c.Status(code).JSON(mErrors.Errors{ID: msgID, Msg: fmts.ConcatStr("Error: ", err.Error())})
	}

	if len(user.User) == 0 {
		return c.Status(code).JSON(mErrors.Errors{ID: msgID, Msg: fmts.ConcatStr("Error: user required")})
	}

	if user.Time > 0 {
		jwtGen.SetExpires(user.Time)
	} else {
		jwtGen.SetExpires(3600)
	}

	if _, user.Public, user.Key, user.Expires, err = jwtGen.Token(user.User); err != nil {
		code = 401
		return c.Status(code).JSON(mErrors.Errors{ID: msgID, Msg: fmts.ConcatStr("Error: when generating jwt - ", err.Error())})
	}

	code = 200
	return c.Status(code).JSON(user)
}
