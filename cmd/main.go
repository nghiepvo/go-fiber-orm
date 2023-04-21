package main

import (
	"fmt"

	"github.com/nghiepvo/go-fiber-orm/internal/book"
	"github.com/nghiepvo/go-fiber-orm/internal/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/nghiepvo/go-fiber-orm/docs"
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
	// // // app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Put("/api/v1/book/:id", book.EditBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

// @title BOOK API
// @version 1.0
// @description A first look on Fiber and ORM on Golang
// @termsOfService http://swagger.io/terms/
// @contact.name Nghiep Vo
// @contact.email nghiep.vo@gmail.com
// @license.name None
// @license.url None
// @host localhost:3000
// @BasePath  /api/v1
func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	verifyConnectToDb()
	defer func() {
		dbInstance, _ := database.DBConn.DB()
		_ = dbInstance.Close()
	}()

	setRoutes(app)

	app.Get("/swagger/*", swagger.HandlerDefault)
	// app.Get("/swagger/*", swagger.New(swagger.Config{
	// 	URL:         "http://example.com/doc.json",
	// 	DeepLinking: false,
	// 	// Expand ("list") or Collapse ("none") tag groups by default
	// 	DocExpansion: "none",
	// 	// Prefill OAuth ClientId on Authorize popup
	// 	OAuth: &swagger.OAuthConfig{
	// 		AppName:  "OAuth Provider",
	// 		ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
	// 	},
	// 	// Ability to change OAuth2 redirect uri location
	// 	OAuth2RedirectUrl: "http://localhost:3000/swagger/oauth2-redirect.html",
	// }))

	app.Listen(":3000")
}
