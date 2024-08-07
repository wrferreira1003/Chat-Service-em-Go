package usecase

import (
	"log"
	"time"

	"github.com/wrferreira1003/chat-service/internal/domain/model"
	"github.com/wrferreira1003/chat-service/internal/repository"
)

//Vamos definir as operacoes de negocio para o chat

type ChatUsecase interface {
	SendMessage(message *model.Message) error
	GetMessages(conversationID string) ([]*model.Message, error)
	GetOrCreateConversation(userID1, userID2 string) (*model.Conversation, error)
}

// ChatUsecaseImpl é uma implementacao da interface ChatUsecase
type ChatUsecaseImpl struct {
	chatRepo repository.ChatRepository
}

// NewChatUsecase cria uma nova instancia da interface ChatUsecase
func NewChatUsecase(chatRepo repository.ChatRepository) ChatUsecase {
	return &ChatUsecaseImpl{chatRepo: chatRepo} // Inicializando a struct ChatUsecase com o repositorio
}

// Envia mensagem no repositorio
func (u *ChatUsecaseImpl) SendMessage(message *model.Message) error {
	log.Printf("Sending message: %v", message)

	// Obtem ou cria uma conversa entre os dois usuarios
	conversation, err := u.chatRepo.GetOrCreateConversation(message.SenderID, message.ReceiverID)
	if err != nil {
		log.Printf("Error getting or creating conversation: %v", err)
		return err
	}
	message.ConversationID = conversation.ID
	message.CreatedAt = time.Now()
	err = u.chatRepo.SaveMessage(message)
	if err != nil {
		log.Printf("Error saving message: %v", err)
		return err
	}
	conversation.LastMessage = message
	conversation.UpdatedAt = time.Now()
	log.Printf("Message sent: %v", message)
	return nil
}

// Obtem as mensagens de uma conversa
func (u *ChatUsecaseImpl) GetMessages(conversationID string) ([]*model.Message, error) {
	return u.chatRepo.GetMessages(conversationID)
}

// Obtem ou cria uma conversa entre dois usuarios
func (u *ChatUsecaseImpl) GetOrCreateConversation(userID1, userID2 string) (*model.Conversation, error) {
	return u.chatRepo.GetOrCreateConversation(userID1, userID2)
}
