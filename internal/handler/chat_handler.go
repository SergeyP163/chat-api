package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/SergeyP163/chat-api/internal/service"
	"gorm.io/gorm"
)

type ChatHandler struct {
	chatService    *service.ChatService
	messageService *service.MessageService
}

func NewChatHandler(chatService *service.ChatService, messageService *service.MessageService) *ChatHandler {
	return &ChatHandler{
		chatService:    chatService,
		messageService: messageService,
	}
}

func (h *ChatHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Title string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	chat, err := h.chatService.Create(req.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(chat)
}

func (h *ChatHandler) Get(w http.ResponseWriter, r *http.Request) {
	chatID, err := ParseChatID(r)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	chat, err := h.chatService.GetByID(chatID)
	if err != nil {
		http.Error(w, "chat not found", http.StatusNotFound)
		return
	}

	limit := 20
	if v := r.URL.Query().Get("limit"); v != "" {
		if l, err := strconv.Atoi(v); err == nil && l > 0 && l <= 100 {
			limit = l
		}
	}

	messages, err := h.messageService.GetLast(chatID, limit)
	if err != nil {
		http.Error(w, "could not get messages", http.StatusInternalServerError)
		return
	}

	resp := struct {
		Chat     any `json:"chat"`
		Messages any `json:"messages"`
	}{
		Chat:     chat,
		Messages: messages,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

func (h *ChatHandler) Delete(w http.ResponseWriter, r *http.Request) {
	chatID, err := ParseChatID(r)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := h.chatService.Delete(chatID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "chat not found", http.StatusNotFound)
			return
		}
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
