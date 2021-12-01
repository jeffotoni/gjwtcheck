package handler

import (
	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
	mw "github.com/jeffotoni/gjwtcheck/core/middleware"
	mErrors "github.com/jeffotoni/gjwtcheck/core/models/errors"
	fmts "github.com/jeffotoni/gjwtcheck/core/pkg/fmts"
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

	userx := c.Locals("user").(*jwt.Token)
	umap := userx.Claims.(jwt.MapClaims)
	var u mLg.User
	code = 200
	u.Name = "HS256"
	u.AvatarURL = "https://logodix.com/logo/1989600.png"
	u.Message = "Welcome JWT HS256"
	u.User = umap["user"].(string)
	u.Id = umap["id"].(string)
	u.Iss = umap["iss"].(string)
	return c.Status(code).JSON(u)
}
