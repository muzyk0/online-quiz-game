package main

import (
	"log"

	"github.com/muzyk0/online-quiz-game/internal/app"
	"github.com/muzyk0/online-quiz-game/internal/app/config"

	"github.com/joho/godotenv"
)

//	@title			Online Quiz Game API
//	@version		1.0
//	@description	A RESTful API for a pair-based online quiz game.
//	@description	Players compete head-to-head answering 5 questions. Super admin manages questions and users.

//	@host		localhost:8080
//	@BasePath	/api

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and JWT token.

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	cfg := config.Load()

	application, err := app.New(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	if err := application.Run(); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}
