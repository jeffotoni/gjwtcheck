package handler

import (
	"github.com/gofiber/fiber/v2"
	mw "github.com/jeffotoni/gjwtcheck/apicore/middleware"
	mErrors "github.com/jeffotoni/gjwtcheck/apicore/models/errors"
	fmts "github.com/jeffotoni/gjwtcheck/apicore/pkg/fmts"
	mLg "github.com/jeffotoni/gjwtcheck/models/user"
)

func (s StructConnect) User2(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	var err error
	msgID := mw.GetUUID(c)
	code := 400
	if len(string(c.Body())) <= 0 {
		return c.Status(code).JSON(mErrors.Errors{ID: msgID, Msg: `Error: empty body`})
	}
	var user mLg.UserAuth
	err = c.BodyParser(&user)
	if err != nil {
		code = 400
		return c.Status(code).JSON(mErrors.Errors{ID: msgID, Msg: fmts.ConcatStr("Error: ", err.Error())})
	}

	if len(user.Password) == 0 {
		code = 400
		return c.Status(code).JSON(mErrors.Errors{ID: msgID, Msg: "password is mandatory"})
	}

	if len(user.User) == 0 {
		code = 400
		return c.Status(code).JSON(mErrors.Errors{ID: msgID, Msg: "User is mandatory"})
	}

	var u mLg.User
	code = 200
	u.Name = "Jeff-HS256"
	u.AvatarURL = "https://www.letsgophers.com/web/images/jeffotoni.png"
	u.Message = "seja bem vindo test jwt HS256"
	return c.Status(code).JSON(u)
}
