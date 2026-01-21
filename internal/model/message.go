package model

import "time"

type Message struct {
	ID        uint   `gorm:"primaryKey"`
	ChatID    uint   `gorm:"not null;index"`
	Text      string `gorm:"not null;size:5000"`
	CreatedAt time.Time
}
