package book

import (
	"github.com/gofiber/fiber"
)

func GetBooks(c *fiber.Ctx) {
	c.Send("All books")
}

func GetBook(c *fiber.Ctx) {
	c.Send("Single book")
}

func NewBook(c *fiber.Ctx) {
	c.Send("Adds a book")
}

func DeleteBooks(c *fiber.Ctx) {
	c.Send("Deletes a book")
}
