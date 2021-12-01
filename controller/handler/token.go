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

func (s StructConnect) Token(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	var err error

	code := 400
	msgID := mw.GetUUID(c)

	var user mLg.UserAuth
	user.User = "TestUser"
	jwtGen.SetExpires(3600)
	if _, user.Public, user.Key, user.Expires, err = jwtGen.Token(user.User, hd.IP(c)); err != nil {
		code = 401
		return c.Status(code).JSON(mErrors.Errors{ID: msgID, Msg: fmts.ConcatStr("Error: when generating jwt - ", err.Error())})
	}

	code = 200
	return c.Status(code).JSON(user)
}
