package controllers

import (
	"time"

	"loadboard/database"
	"loadboard/models"

	"github.com/gofiber/fiber/v2"
)

func CreateLoad(c *fiber.Ctx) error {
	userID := uint(c.Locals("user_id").(float64))

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(403).JSON(fiber.Map{"error": "Unauthorized"})
	}

	if user.Role != "broker" {
		return c.Status(403).JSON(fiber.Map{"error": "Only brokers can post loads"})
	}

	type LoadRequest struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Weight      string    `json:"weight"`
		Dimensions  string    `json:"dimensions"`
		Pickup      string    `json:"pickup"`
		Dropoff     string    `json:"dropoff"`
		Date        time.Time `json:"date"`
	}

	var data LoadRequest
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	load := models.Load{
		Title:       data.Title,
		Description: data.Description,
		Weight:      data.Weight,
		Dimensions:  data.Dimensions,
		Pickup:      data.Pickup,
		Dropoff:     data.Dropoff,
		Date:        data.Date,
		CreatedBy:   uint(userID),
	}

	result := database.DB.Create(&load)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create load"})
	}

	return c.JSON(load)
}

func GetAllLoads(c *fiber.Ctx) error {
	var loads []models.Load
	query := database.DB

	// Optional filters
	pickup := c.Query("pickup")
	if pickup != "" {
		query = query.Where("pickup = ?", pickup)
	}

	status := c.Query("status")
	if status != "" {
		query = query.Where("status = ?", status)
	}

	sort := c.Query("sort")
	if sort == "date" {
		query = query.Order("date ASC")
	} else {
		query = query.Order("created_at DESC")
	}

	if err := query.Find(&loads).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch loads"})
	}

	return c.JSON(loads)
}

func ClaimLoad(c *fiber.Ctx) error {
	userID := uint(c.Locals("user_id").(float64))
	loadID := c.Params("id")

	// Check user is a carrier
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil || user.Role != "carrier" {
		return c.Status(403).JSON(fiber.Map{"error": "Only carriers can claim loads"})
	}

	// Check load exists
	var load models.Load
	if err := database.DB.First(&load, loadID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Load not found"})
	}

	// Create claim
	claim := models.LoadClaim{
		LoadID: load.ID,
		UserID: userID,
	}
	if err := database.DB.Create(&claim).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to claim load"})
	}

	// Update load status
	load.Status = "claimed"
	database.DB.Save(&load)

	return c.JSON(fiber.Map{"message": "Load claimed"})
}

func MarkDelivered(c *fiber.Ctx) error {
	userID := uint(c.Locals("user_id").(float64))
	loadID := c.Params("id")

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil || user.Role != "broker" {
		return c.Status(403).JSON(fiber.Map{"error": "Only brokers can mark as delivered"})
	}

	var load models.Load
	if err := database.DB.First(&load, loadID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Load not found"})
	}

	load.Status = "delivered"
	database.DB.Save(&load)

	return c.JSON(fiber.Map{"message": "Load marked as delivered"})
}
