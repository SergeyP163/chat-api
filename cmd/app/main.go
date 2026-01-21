package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/SergeyP163/chat-api/internal/db"
	"github.com/SergeyP163/chat-api/internal/handler"
	"github.com/SergeyP163/chat-api/internal/repository"
	"github.com/SergeyP163/chat-api/internal/service"
)

func main() {
	database, err := db.NewPostgres()
	if err != nil {
		log.Fatal("db connection error: ", err)
	}

	chatRepo := repository.NewChatRepository(database)
	messageRepo := repository.NewMessageRepository(database)

	chatService := service.NewChatService(chatRepo)
	messageService := service.NewMessageService(messageRepo, chatRepo)

	chatHandler := handler.NewChatHandler(chatService, messageService)
	messageHandler := handler.NewMessageHandler(messageService)

	mux := http.NewServeMux()

	mux.HandleFunc("/chats", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			chatHandler.Create(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/chats/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			chatHandler.Get(w, r)
		case http.MethodDelete:
			chatHandler.Delete(w, r)
		case http.MethodPost:
			if strings.HasSuffix(r.URL.Path, "/messages") {
				messageHandler.Create(w, r)
				return
			}
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server started on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
