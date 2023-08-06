package api

import (
	"github.com/SharanM7/hotelreservation/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(router *fiber.Ctx) error {
	u := types.User{
		Id:   1,
		Name: "sha",
	}
	router.JSON(u)
	return nil
}

func HandleGetUser(router *fiber.Ctx) error {
	s := `{"result":"success"}`
	router.JSON(s)
	return nil
}

func HandleHome(router *fiber.Ctx) error {
	s := "Welcome to home"
	router.WriteString(s)
	return nil
}

func HandleGetUsers_v2(router *fiber.Ctx) error {
	u := types.User{
		Id:   100,
		Name: "sha_v2",
	}

	router.JSON(u)
	return nil
}
