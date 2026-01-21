package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SergeyP163/chat-api/internal/handler"
	"github.com/SergeyP163/chat-api/internal/model"
	"github.com/SergeyP163/chat-api/internal/repository"
	"github.com/SergeyP163/chat-api/internal/service"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/glebarez/sqlite"
)

func TestCreateChatHandler(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(&model.Chat{}, &model.Message{}); err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	chatRepo := repository.NewChatRepository(db)
	msgRepo := repository.NewMessageRepository(db)
	chatService := service.NewChatService(chatRepo)
	msgService := service.NewMessageService(msgRepo, chatRepo)

	chatHandler := handler.NewChatHandler(chatService, msgService)

	reqBody := map[string]string{"title": "Test Chat"}
	jsonBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/chats", bytes.NewReader(jsonBody))
	w := httptest.NewRecorder()

	chatHandler.Create(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", w.Code)
	}

	var resp map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}

	if resp["Title"] != "Test Chat" && resp["title"] != "Test Chat" {
		t.Fatalf("expected title 'Test Chat', got %v", resp["title"])
	}
}
