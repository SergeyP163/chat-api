package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/SergeyP163/chat-api/internal/service"
	"gorm.io/gorm"
)

type MessageHandler struct {
	service *service.MessageService
}

func NewMessageHandler(service *service.MessageService) *MessageHandler {
	return &MessageHandler{service: service}
}

func (h *MessageHandler) Create(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, "/message") {
		return
	}

	chatID, err := ParseChatID(r)
	if err != nil {
		http.Error(w, "invalid chat id", http.StatusBadRequest)
		return
	}

	var req struct {
		Text string `json:"text"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	msg, err := h.service.Create(chatID, req.Text)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "chat not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(msg)
}
