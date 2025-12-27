package controllers

import (
	"time"

	"be-soal-03/database"
	"be-soal-03/models"

	"github.com/gofiber/fiber/v2"
)

// GetEvents godoc
// @Summary Get all events
// @Tags Events
// @Produce json
// @Success 200 {array} models.Event
// @Router /api/events [get]
func GetEvents(c *fiber.Ctx) error {
	var events []models.Event

	if err := database.DB.
		Where("status = ?", "PUBLISHED").
		Preload("Organizer").
		Find(&events).Error; err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch events",
		})
	}

	return c.JSON(events)
}

// GetEventByID godoc
// @Summary Get event by ID
// @Tags Events
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} models.Event
// @Router /api/events/{id} [get]
func GetEventByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var event models.Event
	if err := database.DB.
		Where("id = ? AND status = ?", id, "PUBLISHED").
		Preload("Organizer").
		First(&event).Error; err != nil {

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Event not found",
		})
	}

	return c.JSON(event)
}


// CreateEvent godoc
// @Summary Create new event
// @Description Create a new event (Admin only)
// @Tags Events
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body object{title=string,description=string,event_date=string,location=string,total_ticket=int} true "Event payload"
// @Success 201 {object} models.Event
// @Router /api/events [post]
func CreateEvent(c *fiber.Ctx) error {
	var body struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		EventDate   time.Time `json:"event_date"`
		Location    string    `json:"location"`
		TotalTicket int       `json:"total_ticket"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	if body.Title == "" || body.Location == "" || body.TotalTicket <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Title, location, and total_ticket are required",
		})
	}

	adminID := c.Locals("user_id").(uint)

	event := models.Event{
		Title:           body.Title,
		Description:     body.Description,
		EventDate:       body.EventDate,
		Location:        body.Location,
		TotalTicket:     body.TotalTicket,
		AvailableTicket: body.TotalTicket,
		OrganizerID:     adminID,
		Status:          "PUBLISHED",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	if err := database.DB.
		Preload("Organizer").
		First(&event, event.ID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load organizer",
	})
}

	return c.Status(fiber.StatusCreated).JSON(event)
}

// UpdateEvent godoc
// @Summary Update event
// @Description Update event information (Admin only)
// @Tags Events
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Param body body object{title=string,description=string,event_date=string,location=string,status=string,total_ticket=int} true "Update event payload"
// @Success 200 {object} models.Event
// @Router /api/events/{id} [put]
func UpdateEvent(c *fiber.Ctx) error {
	id := c.Params("id")

	var body struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		EventDate   time.Time `json:"event_date"`
		Location    string    `json:"location"`
		Status      string    `json:"status"`
		TotalTicket int       `json:"total_ticket"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	var event models.Event
	if err := database.DB.First(&event, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Event not found",
		})
	}

	sold := event.TotalTicket - event.AvailableTicket

	if body.Title != "" {
		event.Title = body.Title
	}
	if body.Description != "" {
		event.Description = body.Description
	}
	if body.Location != "" {
		event.Location = body.Location
	}
	if !body.EventDate.IsZero() {
		event.EventDate = body.EventDate
	}
	if body.Status != "" {
		event.Status = body.Status
	}

	if body.TotalTicket > 0 {
		if body.TotalTicket < sold {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "total_ticket cannot be less than sold tickets",
			})
		}

		event.TotalTicket = body.TotalTicket
		event.AvailableTicket = body.TotalTicket - sold
	}

	event.UpdatedAt = time.Now()

	if err := database.DB.
		Preload("Organizer").
		First(&event, event.ID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load organizer",
	})
}

	return c.JSON(event)
}
