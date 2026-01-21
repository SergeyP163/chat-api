package repository

import (
	"github.com/SergeyP163/chat-api/internal/model"
	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) Create(msg *model.Message) error {
	return r.db.Create(msg).Error
}

func (r *MessageRepository) GetLast(chatID uint, limit int) ([]model.Message, error) {
	var msgs []model.Message
	err := r.db.
		Where("chat_id = ?", chatID).
		Order("created_at desc").
		Limit(limit).
		Find(&msgs).Error
	return msgs, err
}
