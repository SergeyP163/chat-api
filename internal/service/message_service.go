package service

import (
	"errors"
	"strings"

	"github.com/SergeyP163/chat-api/internal/model"
	"github.com/SergeyP163/chat-api/internal/repository"
	"gorm.io/gorm"
)

type MessageService struct {
	repo     *repository.MessageRepository
	chatRepo *repository.ChatRepository
}

func NewMessageService(repo *repository.MessageRepository, chatRepo *repository.ChatRepository) *MessageService {
	return &MessageService{repo: repo, chatRepo: chatRepo}
}

func (s *MessageService) Create(chatID uint, text string) (*model.Message, error) {
	text = strings.TrimSpace(text)
	if len(text) == 0 {
		return nil, errors.New("text cannot be empty")
	}
	if len(text) > 5000 {
		return nil, errors.New("text too long")
	}

	if _, err := s.chatRepo.GetByID(chatID); errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}

	msg := &model.Message{ChatID: chatID, Text: text}
	if err := s.repo.Create(msg); err != nil {
		return nil, err
	}
	return msg, nil
}

func (s *MessageService) GetLast(chatID uint, limit int) ([]model.Message, error) {
	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}

	if _, err := s.chatRepo.GetByID(chatID); err != nil {
		return nil, err
	}
	return s.repo.GetLast(chatID, limit)
}
