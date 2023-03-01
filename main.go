package main

import (
	"fmt"

	"github.com/nghiepvo/go-fiber-orm/book"
	"github.com/nghiepvo/go-fiber-orm/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/gofiber/fiber/v2"
)

func verifyConnectToDb() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("book.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database.")
	}
	fmt.Println("Database connection successfully opened.")

	database.DBConn.AutoMigrate(&book.Book{})

	fmt.Println("Database Migrated")
}

func setRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Put("/api/v1/book/:id", book.EditBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func main() {
	app := fiber.New()

	verifyConnectToDb()
	defer func() {
		dbInstance, _ := database.DBConn.DB()
		_ = dbInstance.Close()
	}()

	setRoutes(app)

	app.Listen(":3000")
}
