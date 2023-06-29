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
