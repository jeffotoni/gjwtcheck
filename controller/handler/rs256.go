package handler

import (
	"github.com/gofiber/fiber/v2"
	jwtGen "github.com/jeffotoni/gjwtcheck/core/pkg/jwt"
	mLg "github.com/jeffotoni/gjwtcheck/models/user"

	mw "github.com/jeffotoni/gjwtcheck/core/middleware"
	mErrors "github.com/jeffotoni/gjwtcheck/core/models/errors"
	"github.com/jeffotoni/gjwtcheck/core/pkg/fmts"
	hd "github.com/jeffotoni/gjwtcheck/core/pkg/headers"
)

func (s StructConnect) RS256(c *fiber.Ctx) error {
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

	jwtGen.SetExpires(180)
	if user.Public,user.Key, user.Expires, err = jwtGen.Token(user.User, hd.IP(c)); err != nil {
		code = 401
			return c.Status(code).JSON(mErrors.Errors{ID: msgID, Msg: fmts.ConcatStr("Error: when generating jwt - ", err.Error())})
	}

	code = 200
	return c.Status(code).JSON(user)
}
