package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "internet_service/docs"
	"internet_service/internal/controller"
)

// @title Iran Net Packages
// @version 1.0
// @description
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func Server() *fiber.App {
	//conf := config.GetConfToml()
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/api/irancell", controller.GetIrancell)
	app.Get("/api/irancell/fetch")
	return app
}
