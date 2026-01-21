package repository

import (
	"errors"

	"github.com/SergeyP163/chat-api/internal/model"
	"gorm.io/gorm"
)

type ChatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) *ChatRepository {
	return &ChatRepository{db: db}
}

func (r *ChatRepository) Create(chat *model.Chat) error {
	if chat == nil {
		return errors.New("chat is nil")
	}
	return r.db.Create(chat).Error
}

func (r *ChatRepository) GetByID(id uint) (*model.Chat, error) {
	var chat model.Chat
	if err := r.db.First(&chat, id).Error; err != nil {
		return nil, err
	}
	return &chat, nil
}

func (r *ChatRepository) Delete(id uint) error {
	result := r.db.Delete(&model.Chat{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
