package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jeffronworks/shop-credit/database"
)

func home(c *fiber.Ctx) error {
	return c.SendString("shop-credit default")
}

func Myroutes(app *fiber.App) {
	// default route
	app.Get("/v0/api", home)
}

func main() {
	// instantiate fiber app
	app := fiber.New()

	//connect to the database
	database.ConnectDb()

	// calling the route list function
	Myroutes(app)

	// Listen on port 8000 and log
	log.Fatal(app.Listen(":8000"))
}
