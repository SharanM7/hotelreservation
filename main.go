package main

import (
	"github.com/SharanM7/hotelreservation/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	appv1 := app.Group("/api/v1")
	appv2 := app.Group("/api/v2")

	appv1.Get("/users", api.HandleGetUsers)
	appv1.Get("/user", api.HandleGetUser)
	app.Get("/", api.HandleHome)

	appv2.Get("/users", api.HandleGetUsers_v2)
	app.Listen(":5000")
}
