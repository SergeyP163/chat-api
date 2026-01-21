package model

import "time"

type Chat struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"not null;size:200"`
	Messages  []Message `gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt time.Time
}
