package router

import (
	"api/controller"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Post("/country", controller.PostCountry)
	app.Get("/country/:id", controller.GetCountry)
}
