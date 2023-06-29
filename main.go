package main

import (
	"fmt"
	database "kevo-codes/notekeeper/db"
	models "kevo-codes/notekeeper/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func initDatabase() {
	var err error

	//Connection String
	dsn := "host=localhost user=postgres password=123456 dbname=notekeeper port=5432"

	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Database connection successfully opened!")
	database.DBConn.AutoMigrate(&models.Note{})
	fmt.Println("Database migrated!")
}

func setRoutes(app *fiber.App) {
	//get all
	app.Get("/api/v1/notes", models.GetNotes)
	
	//create
	app.Post("/api/v1/notes", models.CreateNotes)
	
	//get by id
	app.Get("/api/v1/notes/:id", models.GetNoteById)
	
	//update

	//delete
	// app.Delete("/api/v1/notes/:id", models.DeleteNote)
}

func main() {

	app := fiber.New()
	initDatabase()
	// app.Get("/", helloWorld)
	setRoutes(app)
	app.Listen(":8000")
}
