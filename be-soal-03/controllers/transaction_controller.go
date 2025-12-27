package controllers

import (
	"time"

	"be-soal-03/database"
	"be-soal-03/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CreateTransaction godoc
// @Summary Book ticket
// @Tags Transactions
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body object{event_id=int,quantity=int} true "Transaction payload"
// @Success 201 {object} models.Transaction
// @Router /transactions [post]
func CreateTransaction(c *fiber.Ctx) error {
	var body struct {
		EventID  uint `json:"event_id"`
		Quantity int  `json:"quantity"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	if body.EventID == 0 || body.Quantity <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Event ID and quantity are required",
		})
	}

	userID := c.Locals("user_id").(uint)
	var trx models.Transaction

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var event models.Event

		if err := tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&event, body.EventID).Error; err != nil {
			return fiber.NewError(fiber.StatusNotFound, "Event not found")
		}

		if event.AvailableTicket < body.Quantity {
			return fiber.NewError(fiber.StatusConflict, "Not enough tickets available")
		}

		event.AvailableTicket -= body.Quantity
		event.UpdatedAt = time.Now()
		if err := tx.Save(&event).Error; err != nil {
			return err
		}

		trx = models.Transaction{
			UserID:    userID,
			EventID:   event.ID,
			Quantity:  body.Quantity,
			Status:    "SUCCESS",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		return tx.Create(&trx).Error
	})

	if err != nil {
		if fiberErr, ok := err.(*fiber.Error); ok {
			return c.Status(fiberErr.Code).JSON(fiber.Map{
				"error": fiberErr.Message,
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create transaction",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(trx)
}


// GetMyTransactions godoc
// @Summary Get my transactions
// @Description Get all transactions made by the authenticated user
// @Tags Transactions
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.Transaction
// @Router /transactions [get]
func GetMyTransactions(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var transactions []models.Transaction
	if err := database.DB.
		Preload("Event").
		Where("user_id = ?", userID).
		Find(&transactions).Error; err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch transactions",
		})
	}

	return c.JSON(transactions)
}
