package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/wrferreira1003/chat-service/internal/domain/model"
	"github.com/wrferreira1003/chat-service/internal/usecase"
)

// WebsocketHandler manipulates websocket messages.
type WebsocketHandler struct {
	useCase usecase.ChatUsecase      // Caso de uso do chat
	clients map[*websocket.Conn]bool // Mapa de conexoes do websocket que esta ativa
}

// NewWebsocketHandler initializes the websocket handler.
func NewWebsocketHandler(useCase usecase.ChatUsecase) *WebsocketHandler {
	return &WebsocketHandler{
		useCase: useCase,
		clients: make(map[*websocket.Conn]bool),
	}
}

// upgrader is the object that upgrades the HTTP connection to a WebSocket connection.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024, // Tamanho do buffer para leitura
	WriteBufferSize: 1024, // Tamanho do buffer para escrita
	CheckOrigin: func(r *http.Request) bool { // Funcao para verificar a origem da conexao
		return true
	},
}

// HandlerConnection handles websocket connections.
func (h *WebsocketHandler) HandlerConnection(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling connection: %v", r)
	// Upgrade a conexao HTTP para uma conexao WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading connection: %v", err)
		return
	}
	defer ws.Close()
	log.Printf("client connected: %v", ws.LocalAddr())

	// Adiciona a conexao ao mapa de conexoes ativas
	h.clients[ws] = true

	// Loop infinito para escutar as mensagens do cliente
	for {
		var msg model.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			delete(h.clients, ws)
			log.Printf("client disconnected: %v", ws.LocalAddr())
			break
		}

		// Gera o conversationID baseado nos IDs dos usuários
		msg.ConversationID = generateConversationID(msg.SenderID, msg.ReceiverID)

		// Envia a mensagem para o caso de uso
		err = h.useCase.SendMessage(&msg)
		if err != nil {
			log.Printf("Error sending message: %v", err)
			continue
		}

		// Envia a mensagem para todos os clientes conectados
		for client := range h.clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Error sending message to client: %v", err)
				delete(h.clients, client)
				log.Printf("client disconnected: %v", client.LocalAddr())
			}
		}
	}
}

// generateConversationID generates a conversation ID based on user IDs.
func generateConversationID(userID1, userID2 string) string {
	log.Printf("Generating conversation id for users with ids %s and %s", userID1, userID2)
	if userID1 < userID2 {
		return userID1 + "-" + userID2
	}
	return userID2 + "-" + userID1
}
