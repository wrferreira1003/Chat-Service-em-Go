package model

import (
	"testing"
	"time"
)

func TestConversationCreation(t *testing.T) {
	// Criando uma conversa com dois usuários
	conversation := Conversation{
		ID:           "conv1",
		Participants: []string{"user1", "user2"}, // Criando uma conversa com dois usuários
		CreatedAt:    time.Now(),                 // Criando a data de criação da conversa
		UpdatedAt:    time.Now(),                 // Criando a data de atualização da conversa
	}

	// Verificando se o ID foi gerado
	if conversation.ID != "conv1" {
		t.Errorf("expected ID to be set, got %s", conversation.ID)
	}

	// Verificando se os participantes foram setados
	if len(conversation.Participants) != 2 || conversation.Participants[0] != "user1" || conversation.Participants[1] != "user2" {
		t.Errorf("expected Participants to be ['user1', 'user2'], got %v", conversation.Participants)
	}

	// Verificando se a data de criação foi setada
	if conversation.CreatedAt.IsZero() {
		t.Errorf("expected CreatedAt to be set, got %s", conversation.CreatedAt)
	}

	// Verificando se a data de atualização foi setada
	if conversation.UpdatedAt.IsZero() {
		t.Errorf("expected UpdatedAt to be set, got %s", conversation.UpdatedAt)
	}
}
