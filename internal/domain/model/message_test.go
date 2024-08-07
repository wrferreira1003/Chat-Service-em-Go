package model

import (
	"testing"
	"time"
)

func TestMessageCreation(t *testing.T) {
	// Criando uma mensagem com os campos obrigatórios
	msg := Message{
		ID:             "1",
		SenderID:       "user1",
		ReceiverID:     "user2",
		Content:        "Hello, World!",
		ConversationID: "conversation1",
		CreatedAt:      time.Now(),
	}

	// Verificando se o ID foi gerado
	if msg.ID != "1" {
		t.Errorf("expected ID to be '1', got %s", msg.ID)
	}

	// Verificando se o senderID foi setado
	if msg.SenderID != "user1" {
		t.Errorf("expected senderID to be 'user1', got %s", msg.SenderID)
	}

	// Verificando se o receiverID foi setado
	if msg.ReceiverID != "user2" {
		t.Errorf("expected receiverID to be 'user2', got %s", msg.ReceiverID)
	}

	// Verificando se o content foi setado
	if msg.Content != "Hello, World!" {
		t.Errorf("expected content to be 'Hello, World!', got %s", msg.Content)
	}

	// Verificando se o conversationID foi setado
	if msg.ConversationID != "conversation1" {
		t.Errorf("expected conversationID to be 'conversation1', got %s", msg.ConversationID)
	}

	// Verificando se o createdAt foi setado
	if msg.CreatedAt.IsZero() {
		t.Errorf("expected createdAt to be set, got %s", msg.CreatedAt)
	}
}
