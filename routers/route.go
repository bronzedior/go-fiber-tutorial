package routers

import (
	"github.com/bronzedior/go-fiber-tutorial/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouterApp(c *fiber.App) {
	c.Get("/", controllers.UserControllerShow)
}
