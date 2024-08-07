package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wrferreira1003/chat-service/internal/domain/model"
	"github.com/wrferreira1003/chat-service/internal/usecase"
)

type ChatController struct {
	chatUsecase usecase.ChatUsecase
}

// NewChatController cria um novo controller de chat
func NewChatController(chatUsecase usecase.ChatUsecase) *ChatController {
	return &ChatController{chatUsecase: chatUsecase}
}

// SendMessage envia uma mensagem para o chat
// @Summary Envia uma mensagem para o chat
// @Description Envia uma mensagem para o chat
// @Tags chat
// @Accept json
// @Produce json
// @Param message body model.Message true "Mensagem a ser enviada"
// @Success 201 {string} string "Mensagem enviada com sucesso"
// @Failure 400 {string} string "Erro na requisição"
// @Failure 500 {string} string "Erro interno do servidor"
// @Router /chat/send [post]
func (c *ChatController) SendMessage(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received message: %v", r)
	var message model.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		log.Printf("Error decoding message: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Gerar o convesationID baseado nos ids dos usuarios
	message.ConversationID = generateConversationID(message.SenderID, message.ReceiverID)

	// Enviar a mensagem para o usecase
	err = c.chatUsecase.SendMessage(&message)
	if err != nil {
		log.Printf("Error sending message: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retorna o status code 201 Created
	w.WriteHeader(http.StatusCreated)

}

// GetMessages obtem as mensagens de uma conversa
// @Summary Obtem as mensagens de uma conversa
// @Description Obtem as mensagens de uma conversa
// @Tags chat
// @Accept json
// @Produce json
// @Param userID1 path string true "ID do primeiro usuário"
// @Param userID2 path string true "ID do segundo usuário"
// @Success 200 {array} model.Message "Lista de mensagens"
// @Failure 500 {string} string "Erro interno do servidor"
// @Router /chat/messages/{userID1}/{userID2} [get]
func (c *ChatController) GetMessages(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID1 := params["userID1"]
	userID2 := params["userID2"]
	conversationID := generateConversationID(userID1, userID2)

	// Obtem as mensagens de uma conversa
	messages, err := c.chatUsecase.GetMessages(conversationID)
	if err != nil {
		log.Printf("Error getting messages: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(messages)
}
