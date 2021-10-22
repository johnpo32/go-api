package controller

import (
	"api/database"
	"api/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Country struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
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
	var country Country

	if err := db.Find(&country, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "data": err})
	}
	if country.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "data": "Id error"})
	}

	return c.JSON(fiber.Map{"status": "ok", "data": country})
}

func GetCountries(c *fiber.Ctx) error {

	db := database.DB
	var countries []model.Country

	if err := db.Find(&countries).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "data": err})
	}

	var val []interface{}

	for _, row := range countries {
		cont := Country{
			ID:   row.ID,
			Name: row.Name,
		}
		val = append(val, cont)
	}

	return c.JSON(fiber.Map{"status": "ok", "data": val})
}

func PatchCountry(c *fiber.Ctx) error {

	id := c.Params("id")
	db := database.DB

	type UpdateInput struct {
		Name string `json:"name"`
	}

	var upd UpdateInput

	if err := c.BodyParser(&upd); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	var country model.Country
	db.Find(&country, "id = ?", id)

	country.Name = upd.Name

	if err := db.Save(&country).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "data": err})
	}

	newCountry := Country{
		ID:   country.ID,
		Name: country.Name,
	}

	return c.JSON(fiber.Map{"status": "ok", "data": newCountry})
}

func DeleteCountry(c *fiber.Ctx) error {

	id := c.Params("id")
	db := database.DB

	var country model.Country

	if err := db.Where("id = ?", id).Delete(&country).Error; err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"status": "error", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Country deleted", "data": nil})
}
