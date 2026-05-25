package url

import (
	"backend/controller"
	"backend/controller/modul9"

	"github.com/gofiber/fiber/v2"
)

// Web maps routes to controllers
func Web(app *fiber.App) {

	app.Get("/", controller.Homepage)

	app.Get("/ip", controller.IPServer)

	modul9.RegisterRoutes(app)

}
