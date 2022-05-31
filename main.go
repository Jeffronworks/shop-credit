package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	Myroutes(app)
}

func home(c *fiber.Ctx) error {
	return c.SendString("default home route!")
}

func Myroutes(app *fiber.App) {
	// home endpoint
	app.Get("/v0/api", home)

	log.Fatal(app.Listen(":8000"))
}
