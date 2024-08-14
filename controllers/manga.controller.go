package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thitiphongD/go-backend/database"
	"github.com/thitiphongD/go-backend/models"
)

func GetMangas(c *fiber.Ctx) error {
	var mangas []models.Manga
	database.DB.Find(&mangas)
	return c.Status(200).JSON(mangas)
}

func AddManga(c *fiber.Ctx) error {
	manga := new(models.Manga)
	if err := c.BodyParser(manga); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	database.DB.Create(&manga)
	return c.Status(201).JSON(manga)
}

func UpdateManga(c *fiber.Ctx) error {
	manga := new(models.Manga)
	id := c.Params("id")
	if err := c.BodyParser(manga); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	database.DB.Where("id = ?", id).Updates(&manga)
	return c.Status(200).JSON(manga)
}

func RemoveManga(c *fiber.Ctx) error {
	id := c.Params("id")
	var manga models.Manga
	result := database.DB.Delete(&manga, id)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.SendStatus(200)
}
