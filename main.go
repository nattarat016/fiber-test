package main

import (
	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func main() {

	app := fiber.New()

	books = append(books, Book{ID: 1, Title: "firts", Author: "dat"})
	books = append(books, Book{ID: 2, Title: "secon", Author: "nom"})
	books = append(books, Book{ID: 3, Title: "tird", Author: "kol"})

	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("books/:id", deleteBook)
	app.Post("/uploade", uplodeFile)

	app.Listen(":8080")
}

func uplodeFile(c *fiber.Ctx) error {
	file, err := c.FormFile("image")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = c.SaveFile(file, "./uploades/"+file.Filename)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString("uplode complete")
}
