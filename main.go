package main

import (
	"api/database"
	"api/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()
	router.Router(app)
	log.Fatal(app.Listen("127.0.0.1:3000"))

	defer database.DB.Close()
}
