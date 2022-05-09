package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	"github.com/gangleri/fiberswag/api"
	_ "github.com/gangleri/fiberswag/docs"
)

// @title           Swagger Fiber Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
func main() {
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)

	api.RegisterRoutes(app)
	app.Listen(":8080")
}
