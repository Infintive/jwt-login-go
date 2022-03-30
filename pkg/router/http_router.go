package router

import (
	"github.com/Infintive/predictive-go/app/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
)

type HttpRouter struct {
}

func (h HttpRouter) InstallRouter(app *fiber.App) {
	group := app.Group("", cors.New(), csrf.New())

	group.Get("/", controllers.RenderHello)

	group.Get("/name", controllers.RenderName)

	group.Post("/upload", controllers.UploadFile)

}

func NewHttpRouter() *HttpRouter {
	return &HttpRouter{}
}
