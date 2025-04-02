package controllers

import (
	"loadboard/database"
	"loadboard/models"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	var body Request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	body.Email = strings.ToLower(body.Email) //

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 12)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Couldn't hash password"})
	}

	user := models.User{
		Email:    body.Email,
		Password: string(hash),
		Role:     body.Role,
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Email may already be taken"})
	}

	return c.JSON(fiber.Map{"message": "User created", "user": user})
}

func Login(c *fiber.Ctx) error {
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body Request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	body.Email = strings.ToLower(body.Email) //

	var user models.User
	result := database.DB.Where("email = ?", body.Email).First(&user)
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Wrong password"})
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{"token": t})

}

func Me(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(float64)

	var user models.User
	result := database.DB.First(&user, uint(userID))
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(fiber.Map{
		"email": user.Email,
		"role":  user.Role,
	})
}
