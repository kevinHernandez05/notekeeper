package main

import (
	database "kevo-codes/notekeeper/db"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func initDatabase() {
	var err error
	dsn := "notekeeper.db"

	database.DBConn, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}
}

func main() {
	app := fiber.New()

	app.Get("/", helloWorld)

	app.Listen(":8000")
}
