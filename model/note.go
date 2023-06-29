package models

import (
	database "kevo-codes/notekeeper/db"

	"github.com/gofiber/fiber/v2"
)

type Note struct {
	Id      uint   `gorm:"primarykey" json:"id"`
	Content string `json:"content"`
	Title   string `json:"title"`
}

func GetNotes(c *fiber.Ctx) error {
	db := database.DBConn
	var notes []Note
	db.Find(&notes)

	return c.JSON(&notes)
}

func CreateNotes(c *fiber.Ctx) error {
	db := database.DBConn
	note := new(Note)

	if err := c.BodyParser(note); err != nil {
		return c.Status(503).JSON(fiber.Map{"status": "error", "message": "Couldn't create note", "data": err})
	}

	db.Create(&note)
	return c.JSON(&note)
}

func GetNoteById(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var note Note

	db.Find(&note, id)

	return c.JSON(&note)
}

func UpdateNoteById(c *fiber.Ctx) error {
	type updateNoteInput struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	db := database.DBConn
	id := c.Params("id")

	var note Note
	err := db.Find(&note, id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Note not found", "data": err})
	}

	return c.JSON(&note)
	// var updateNote updateNoteInput

}
