package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	mw "github.com/jeffotoni/gjwtcheck/apicore/middleware"
	mErrors "github.com/jeffotoni/gjwtcheck/apicore/models/errors"
	fmts "github.com/jeffotoni/gjwtcheck/apicore/pkg/fmts"
	hd "github.com/jeffotoni/gjwtcheck/apicore/pkg/headers"
	jwtCore "github.com/jeffotoni/gjwtcheck/apicore/pkg/jwt"
	mLg "github.com/jeffotoni/gjwtcheck/models/user"
)

func (s StructConnect) User(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	msgID := mw.GetUUID(c)
	code := 400
	if len(string(c.Body())) <= 0 {
		return c.Status(code).JSON(mErrors.Errors{Msg: `Error: empty body`})
	}

	var user mLg.UserAuth
	if err := c.BodyParser(&user); err != nil {
		code = 400
		return c.Status(code).JSON(mErrors.Errors{Msg: fmts.ConcatStr("Error: ", err.Error())})
	}

	if len(user.Password) == 0 {
		code = 400
		return c.Status(code).JSON(mErrors.Errors{Msg: "password is mandatory"})
	}

	if len(user.User) == 0 {
		code = 400
		return c.Status(code).JSON(mErrors.Errors{Msg: "User is mandatory"})
	}

	if response.Token, response.Expires, err = jwtCore.Token(response.UserToken, hd.IP(c)); err != nil {
		code = 401
		return c.Status(code).JSON(mErrors.Errors{Msg: fmts.ConcatStr("Error: ", err.Error())})
	}

	var jsonstr string = fmts.Concat(`{"accesstype":1,"iduser":"`,
		response.UserToken, `","data":"`, time.Now().Format("2006-01-02"), `","hora":"`,
		time.Now().Format("15:04:05"), `","useragent":"`, hd.UserAgent(c), `","ip":"`, hd.IP(c), `"}`)

	println(jsonstr)

	code = 200

	response.UserToken = ""
	response.Message = "Welcome"
	return c.Status(code).JSON(response)
}
