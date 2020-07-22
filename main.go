package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jnprogrammer/go_fiber_restapi/book"
	"github.com/jnprogrammer/go_fiber_restapi/database"
)

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect to database connection")
	}
	fmt.Println("Database connection opened!")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBooks)

}
func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(8420)
}
