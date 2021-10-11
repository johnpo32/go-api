package main

import (
	"api/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()
	log.Fatal(app.Listen("127.0.0.1:3000"))
}
