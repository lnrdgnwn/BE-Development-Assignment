package controllers

import (
	"be-soal-03/database"
	"be-soal-03/models"

	"github.com/gofiber/fiber/v2"
)

// GetMyProfile godoc
// @Summary Get current user profile
// @Description Get profile of authenticated user
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Success 200 {object} models.User
// @Router /api/users/me [get]
func GetMyProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	user.Password = ""
	return c.JSON(user)
}


// GetUsers godoc
// @Summary Get all users
// @Description Get list of all users (Admin only)
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.User
// @Router /api/users [get]
func GetUsers(c *fiber.Ctx) error {
	var users []models.User

	if err := database.DB.Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch users",
		})
	}

	for i := range users {
		users[i].Password = ""
	}

	return c.JSON(users)
}

