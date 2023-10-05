package main

import (
	"Backend_TA/controllers/authcontrollers"
	"Backend_TA/controllers/ktpcontrollers"
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
	ktp := api.Group("/ktp")

	auth.Post("/refresh", authcontrollers.RefreshToken)
	auth.Post("/register", authcontrollers.Register)
	auth.Post("/login", authcontrollers.Login)

	profile.Get("/", middlewares.Auth, masyarakatcontrollers.Show)
	profile.Post("/", middlewares.Auth, authcontrollers.Register)
	profile.Get("/:nik", middlewares.Auth, masyarakatcontrollers.ShowId)
	profile.Put("/:nik", middlewares.Auth, masyarakatcontrollers.UpdateProfile)
	profile.Put("/password/:nik", middlewares.Auth, masyarakatcontrollers.UpdatePassword)
	profile.Delete("/:nik", middlewares.Auth, masyarakatcontrollers.DeleteProfile)

	ktp.Post("/", ktpcontrollers.Create)
	app.Listen(":4000")
}
