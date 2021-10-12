package controller

import (
	"api/database"
	"api/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Country struct {
	Name string `json:"name"`
}

func PostCountry(c *fiber.Ctx) error {

	db := database.DB // database connection
	var country model.Country

	if err := c.BodyParser(&country); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "data": err})
	}

	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	country.ID = id
	if err := db.Create(&country).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "data": err})
	}

	return c.JSON(fiber.Map{"status": "ok"})
}

func GetCountry(c *fiber.Ctx) error {

	id := c.Params("id")
	db := database.DB // database connection
	var country model.Country

	if err := db.Find(&country, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "data": err})
	}
	if country.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "data": "Id error"})
	}

	return c.JSON(fiber.Map{"status": "ok", "data": country})
}
