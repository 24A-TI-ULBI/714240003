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


	BeasiswaRoute(app)
	PendaftaranRoute(app)

	// Mahasiswa
	app.Get("/mahasiswa", controller.GetAllMahasiswa)
	app.Get("/mahasiswa/:npm", controller.GetMahasiswaByNPM)
	app.Post("/mahasiswa", controller.CreateMahasiswa)
	app.Put("/mahasiswa/:npm", controller.UpdateMahasiswa)
	app.Delete("/mahasiswa/:npm", controller.DeleteMahasiswa)

	// Auth
	app.Post("/auth/login", controller.Login)
	app.Get("/auth/profile/:phone", controller.GetProfile)
}
