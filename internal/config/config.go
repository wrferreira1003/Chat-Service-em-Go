package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func Load() {
	// Obtém o diretório atual
	currentDir, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting current directory: %v", err)
	}

	// Tenta carregar o .env do diretório atual
	err = godotenv.Load()
	if err != nil {
		rootDir := filepath.Dir(filepath.Dir(currentDir))
		envPath := filepath.Join(rootDir, ".env")
		err = godotenv.Load(envPath)
		if err != nil {
			log.Printf("Warning: .env file not found at %s. Using system environment variables.", envPath)
		} else {
			log.Printf(".env file loaded from project root: %s", envPath)
		}
	} else {
		log.Printf(".env file loaded from current directory: %s", currentDir)
	}

	// Verifica se a variável MONGO_URI foi carregada
	if uri, exists := os.LookupEnv("MONGO_URI"); !exists {
		log.Println("Warning: MONGO_URI not set in .env or system environment")
	} else {
		log.Printf("MONGO_URI loaded: %s", uri)
	}

	log.Println("Configuration loaded")
}

func GetMongoURI() string {
	Load()
	uri, exists := os.LookupEnv("MONGO_URI")
	if !exists {
		log.Println("MONGO_URI not found in environment variables")
		log.Println("Current environment variables:")
		for _, env := range os.Environ() {
			log.Println(env)
		}
		log.Fatal("MONGO_URI environment variable not set")
	}
	return uri
}

func GetMongoUser() string {
	Load()
	uri, exists := os.LookupEnv("MONGO_USER")
	if !exists {
		log.Fatal("MONGO_USER environment variable not set")
	}
	return uri
}

func GetMongoPassword() string {
	Load()
	uri, exists := os.LookupEnv("MONGO_PASSWORD")
	if !exists {
		log.Fatal("MONGO_PASSWORD environment variable not set")
	}
	return uri
}

func GetToken() string {
	Load()
	uri, exists := os.LookupEnv("JWT_SECRET")
	if !exists {
		log.Fatal("JWT_SECRET environment variable not set")
	}
	return uri
}
