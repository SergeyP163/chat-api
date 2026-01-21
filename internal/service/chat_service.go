package service

import (
	"errors"
	"strings"

	"github.com/SergeyP163/chat-api/internal/model"
	"github.com/SergeyP163/chat-api/internal/repository"
	"gorm.io/gorm"
)

type ChatService struct {
	repo *repository.ChatRepository
}

func NewChatService(repo *repository.ChatRepository) *ChatService {
	return &ChatService{repo: repo}
}

func (s *ChatService) Create(title string) (*model.Chat, error) {
	title = strings.TrimSpace(title)
	if len(title) == 0 {
		return nil, errors.New("title cannot be empty")
	}
	if len(title) > 200 {
		return nil, errors.New("title too long")
	}

	chat := &model.Chat{Title: title}
	if err := s.repo.Create(chat); err != nil {
		return nil, err
	}
	return chat, nil
}

func (s *ChatService) GetByID(id uint) (*model.Chat, error) {
	return s.repo.GetByID(id)
}

func (s *ChatService) Delete(id uint) error {
	err := s.repo.Delete(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}
	return err
}
