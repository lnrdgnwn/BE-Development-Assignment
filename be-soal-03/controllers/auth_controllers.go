package controllers

import (
	"be-soal-03/database"
	"be-soal-03/models"
	"be-soal-03/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Register godoc
// @Summary Register new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body object{name=string,email=string,password=string} true "Register payload"
// @Success 201 {object} models.User
// @Router /auth/register [post]
func Register(c *fiber.Ctx) error {
	var body struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	var count int64
	database.DB.Model(&models.User{}).
		Where("email = ?", body.Email).
		Count(&count)

	if count > 0 {
		return c.Status(409).JSON(fiber.Map{
			"error": "Email already registered",
		})
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	user := models.User{
		Name:      body.Name,
		Email:     body.Email,
		Password:  string(hash),
		Role:      "CUSTOMER",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	database.DB.Create(&user)
	user.Password = ""

	return c.Status(201).JSON(user)
}


// Login godoc
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body object{email=string,password=string} true "Login payload"
// @Success 200 {object} object{token=string}
// @Router /auth/login [post]
func Login(c *fiber.Ctx) error {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	var user models.User
	if err := database.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	token, _ := utils.GenerateToken(user.ID, user.Role)

	return c.JSON(fiber.Map{
		"token": token,
	})
}
