package url

import (
	"backend/controller"

	"github.com/gofiber/fiber/v2"
)

func BeasiswaRoute(app *fiber.App) {

	app.Get("/api/beasiswa", controller.GetBeasiswa)

	app.Get("/api/beasiswa/:id", controller.GetDetailBeasiswa)

	app.Post("/api/beasiswa", controller.AddBeasiswa)

	app.Put("/api/beasiswa/:id", controller.UpdateBeasiswa)

	app.Delete("/api/beasiswa/:id", controller.DeleteBeasiswa)
}