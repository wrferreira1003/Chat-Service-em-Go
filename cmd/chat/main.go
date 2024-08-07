package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"    // http-swagger middleware
	_ "github.com/wrferreira1003/chat-service/docs" // Import generated docs
	"github.com/wrferreira1003/chat-service/internal/config"
	"github.com/wrferreira1003/chat-service/internal/controller"
	"github.com/wrferreira1003/chat-service/internal/middleware"
	"github.com/wrferreira1003/chat-service/internal/repository"
	"github.com/wrferreira1003/chat-service/internal/usecase"
	"github.com/wrferreira1003/chat-service/pkg/database"
)

// @title Chat Service API
// @version 1.0
// @description This is a sample chat service API.
// @termsOfService http://swagger.io/terms/

// @contact.name Wellington Ferreira
// @contact.url https://github.com/wrferreira1003
// @contact.email wellingtonferreira1003@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

func main() {
	// Inicializa o log
	config.InitLogging()
	log.Println("Starting the chat service...")

	// Carrega as variaveis de ambiente
	config.Load()

	db := database.Connect()            // Conexão com o banco de dados MongoDB
	chatDB := db.Database("chats")      // Conexão com o banco de dados MongoDB passando o nome do banco de dados
	defer db.Disconnect(context.TODO()) // Fecha a conexão com o banco de dados MongoDB

	// Criar uma instancia do ChatRepository passando o banco de dados
	chatRepo := repository.NewChatRepository(chatDB)
	log.Printf("ChatRepository created")

	// Criar uma instancia do ChatUsecase
	chatUseCase := usecase.NewChatUsecase(chatRepo)

	// Criar uma instancia do WebsocketHandler
	wsHandler := controller.NewWebsocketHandler(chatUseCase)

	// Criar uma instancia do ChatController
	chatController := controller.NewChatController(chatUseCase)

	//Configurar as rotas
	r := mux.NewRouter()
	r.HandleFunc("/ws", wsHandler.HandlerConnection) // Handler para a conexão do websocket

	// Protegendo as rotas de mensagens com o middleware de autenticação
	r.Handle("/messages", middleware.AuthMiddleware(http.HandlerFunc(chatController.SendMessage))).Methods("POST")
	r.Handle("/messages/{userID1}/{userID2}", middleware.AuthMiddleware(http.HandlerFunc(chatController.GetMessages))).Methods("GET")

	//Rota para servir a documentação da API
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Inicializa o servidor HTTP
	log.Printf("Server is running on port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
