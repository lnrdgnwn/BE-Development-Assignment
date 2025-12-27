package models

import "time"

type Event struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title           string    `gorm:"type:varchar(255);not null" json:"title"`
	Description     string    `gorm:"type:text" json:"description"`
	EventDate       time.Time `gorm:"not null" json:"event_date"`
	Location        string    `gorm:"type:varchar(255);not null" json:"location"`
	TotalTicket     int       `gorm:"not null" json:"total_ticket"`
	AvailableTicket int       `gorm:"not null" json:"available_ticket"`
	OrganizerID     uint      `gorm:"not null" json:"-"`
	Status          string    `gorm:"type:varchar(50);not null" json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`

	Organizer User `gorm:"foreignKey:OrganizerID" json:"organizer"`
}
