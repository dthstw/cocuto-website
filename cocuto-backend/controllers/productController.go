package controllers

import (
	"cocuto-backend/database"
	"cocuto-backend/models"

	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	var products []models.Product
	result := database.DB.Find(&products)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}
	return c.JSON(products)
}

func AddProduct(c *fiber.Ctx) error {
	product := &models.Product{}
	if err := c.BodyParser(product); err != nil { // Declare err correctly
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	database.DB.Create(product)
	return c.JSON(product)
}

func GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	result := database.DB.First(&product, id)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}

	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	// Find product
	if err := database.DB.First(&product, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}

	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	// Update only provided fields
	database.DB.Model(&product).Updates(data)

	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	result := database.DB.First(&product, id)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}
	database.DB.Delete(&product)

	return c.JSON(fiber.Map{"message": "Product deleted successfully"})
}
