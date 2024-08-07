package repository

import (
	"context"
	"log"
	"time"

	"github.com/wrferreira1003/chat-service/internal/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ChatRepository interface { // @Summary Salva uma mensagem no banco de dados
	SaveMessage(message *model.Message) error
	GetMessages(conversationID string) ([]*model.Message, error)
	GetOrCreateConversation(userID1, userID2 string) (*model.Conversation, error) // Obtem ou cria uma conversa entre dois usuarios
}

// ChatRepositoryImpl é uma implementacao da interface ChatRepository
type ChatRepositoryImpl struct {
	messageCollection      *mongo.Collection // Coleção de mensagens
	conversationCollection *mongo.Collection // Coleção de conversas
}

// NewChatRepository cria uma nova instância de ChatRepository
func NewChatRepository(db *mongo.Database) ChatRepository {
	return &ChatRepositoryImpl{
		messageCollection:      db.Collection("messages"),      // Criando a coleção de mensagens
		conversationCollection: db.Collection("conversations"), // Criando a coleção de conversas
	}
}

// SaveMessage salva uma mensagem no banco de dados
func (r *ChatRepositoryImpl) SaveMessage(message *model.Message) error {
	_, err := r.messageCollection.InsertOne(context.Background(), message)
	if err != nil {
		log.Printf("Error saving message with id %s: %v", message.ID, err)
		return err
	}
	log.Printf("Message saved successfully: %v", message)
	return nil
}

// GetMessages obtém todas as mensagens de uma conversa específica
func (r *ChatRepositoryImpl) GetMessages(conversationID string) ([]*model.Message, error) {
	// Criando o filtro para obter as mensagens de uma conversa especifica
	log.Printf("Getting messages for conversation with id %s", conversationID)
	filter := bson.M{"conversation_id": conversationID}

	// Obtendo as mensagens de uma conversa especifica
	cursor, err := r.messageCollection.Find(context.Background(), filter, options.Find().SetSort(bson.D{{Key: "created_at", Value: 1}}))
	if err != nil {
		log.Printf("Error finding messages for conversation with id %s: %v", conversationID, err)
		return nil, err
	}

	// obtendo as mensagens do cursor
	var messages []*model.Message
	err = cursor.All(context.Background(), &messages)
	if err != nil {
		log.Printf("Error getting messages from cursor for conversation with id %s: %v", conversationID, err)
		return nil, err
	}
	log.Printf("Messages found: %v", messages)
	return messages, nil
}

// GetOrCreateConversation obtém ou cria uma conversa entre dois usuários
func (r *ChatRepositoryImpl) GetOrCreateConversation(userID1, userID2 string) (*model.Conversation, error) {
	// Gerando o ID da conversa entre os dois usuarios
	conversationID := generateConversationID(userID1, userID2)
	log.Printf("Getting or creating conversation for users with ids %s and %s", userID1, userID2)

	// Criando o filtro para obter a conversa entre os dois usuarios
	filter := bson.M{"_id": conversationID} // Filtro para obter a conversa entre os dois usuarios

	// Criando a conversa
	conversation := model.Conversation{}

	// Obtendo a conversa entre os dois usuarios
	err := r.conversationCollection.FindOne(context.Background(), filter).Decode(&conversation)
	if err == mongo.ErrNoDocuments {
		conversation = model.Conversation{
			ID:           conversationID,
			Participants: []string{userID1, userID2},
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}
		_, err = r.conversationCollection.InsertOne(context.Background(), conversation)
		if err != nil {
			log.Printf("Error creating conversation with id %s: %v", conversationID, err)
			return nil, err
		}
		log.Printf("Conversation created: %v", conversation)
	} else if err != nil {
		log.Printf("Error finding conversation with id %s: %v", conversationID, err)
		return nil, err
	}
	log.Printf("Conversation found: %v", conversation)
	return &conversation, nil
}

// generateConversationID gera o ID da conversa entre os dois usuários
func generateConversationID(userID1, userID2 string) string {
	log.Printf("Generating conversation id for users with ids %s and %s", userID1, userID2)
	if userID1 < userID2 {
		return userID1 + "-" + userID2
	}
	return userID2 + "-" + userID1
}
