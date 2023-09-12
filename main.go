package main

import (
	"Backend_TA/controllers/authcontrollers"
	"Backend_TA/controllers/masyarakatcontrollers"
	"Backend_TA/middlewares"
	"Backend_TA/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	models.ConnectDB()
	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/")
	auth := api.Group("auth")
	profile := api.Group("/profile")

	auth.Post("/refresh", authcontrollers.RefreshToken)

	auth.Post("/register", authcontrollers.Register)
	auth.Post("/login", authcontrollers.Login)

	profile.Get("/", middlewares.Auth, masyarakatcontrollers.Show)
	profile.Get("/:nik", masyarakatcontrollers.ShowId)
	profile.Put("/:nik", masyarakatcontrollers.UpdateProfile)
	profile.Put("/password/:nik", masyarakatcontrollers.UpdatePassword)
	profile.Delete("/:nik", masyarakatcontrollers.DeleteProfile)
	app.Listen(":4000")
}
