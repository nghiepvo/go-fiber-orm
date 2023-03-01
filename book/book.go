package book

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nghiepvo/go-fiber-orm/database"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	return c.JSON(book)
}

func NewBook(c *fiber.Ctx) error {
	db := database.DBConn
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).Send([]byte(err.Error()))
	}
	book.CreatedAt = time.Now()
	db.Create(&book)
	return c.JSON(book)
}

func EditBook(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DBConn
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).Send([]byte(err.Error()))
	}
	var bookExited Book
	db.First(&bookExited, id)
	fmt.Println(bookExited)
	if bookId, err := strconv.ParseUint(id, 10, 32); err != nil || bookExited.ID != uint(bookId) || book.ID == 0 {
		return c.Status(400).Send(([]byte)("No book found with given ID"))
	}
	book.CreatedAt = bookExited.CreatedAt
	book.UpdatedAt = time.Now()
	db.Model(&bookExited).Updates(book)
	return c.JSON(bookExited)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.First(&book, id)
	if bookId, err := strconv.ParseUint(id, 10, 32); err != nil || book.ID != uint(bookId) {
		return c.Status(400).Send(([]byte)("No book found with given ID"))
	}
	db.Delete(&book)
	return c.Status(200).Send(([]byte)("Book successfully deleted."))
}
