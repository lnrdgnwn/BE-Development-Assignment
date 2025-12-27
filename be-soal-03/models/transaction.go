package models

import "time"

type Transaction struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	EventID   uint      `gorm:"not null" json:"event_id"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	Status    string    `gorm:"type:varchar(50);not null" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User  User  `gorm:"foreignKey:UserID" json:"user"`
	Event Event `gorm:"foreignKey:EventID" json:"event"`
}
