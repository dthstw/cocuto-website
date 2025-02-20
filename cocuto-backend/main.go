package main

import (
	"cocuto-backend/database"
	"fmt"

	"github.com/gofiber/fiber"
)

func main() {
	database.Connect()

	app := fiber.New()
	fmt.Println("Server running on http://localhost:8080")

	app.Listen("127.0.0.1:8080")
}
