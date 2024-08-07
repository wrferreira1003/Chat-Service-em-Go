package config

import (
	"log"
	"os"
)

// InitLogging initializes the logging configuration
func InitLogging() {
	// Abrindo o aquivo de log no diretório atual
	logFile, err := os.OpenFile("/root/logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetFlags(log.LstdFlags | log.Lshortfile) // aqui estamos setando as flags do log
	log.SetOutput(logFile)                       // aqui estamos setando o output do log

	// Verificando se o arquivo foi criado corretamente
	if _, err := os.Stat("logs/app.log"); os.IsNotExist(err) {
		log.Fatalf("log file does not exist: %v", err)
	}
}
