package main

import (
	"github.com/gofiber/fiber"
	"github.com/jnprogrammer/go_fiber_restapi/book"
)

func hellofiber(c *fiber.Ctx) {
	c.Send("Hello, World")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBooks)
	app.Get("/api/v1/book", book.NewBook)

}
func main() {
	app := fiber.New()

	app.Get("/", hellofiber)

	app.Listen(420)
}
