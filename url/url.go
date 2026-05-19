package url

import (
	"backend/controller"

	"github.com/gofiber/fiber/v2"
)

// Web maps routes to controllers
func Web(app *fiber.App) {

	app.Get("/", controller.Homepage)

	app.Get("/ip", controller.IPServer)

	BeasiswaRoute(app)

	PendaftaranRoute(app)
}