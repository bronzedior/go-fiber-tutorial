package main

import (
	"github.com/bronzedior/go-fiber-tutorial/routers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routers.RouterApp(app)

	app.Listen(":8000")
}
