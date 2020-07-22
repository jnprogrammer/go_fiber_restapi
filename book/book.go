package book

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/jnprogrammer/go_fiber_restapi/database"
)

type Book struct {
	gorm.Model
	Title  string //'json:"title"'
	Author string //'json:"author'
	rating int    //'json:"rating"'
}

func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
	//c.Send("All books")
}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	c.JSON(book)
	//c.Send("Single book")
}

func NewBook(c *fiber.Ctx) {
	db := database.DBConn
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&book)
	c.JSON(book)
	c.Send("Adds a book")
}

func DeleteBooks(c *fiber.Ctx) {
	c.Send("Deletes a book")
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("No book found with that ID")
		return
	}
	db.Delete(&book)
	c.Send("Book has been deleted")

}
