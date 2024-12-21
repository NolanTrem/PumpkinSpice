package server

import (
	"log"
	"os"

	openai "github.com/NolanTrem/PumpkinSpice/internal/llm/openai"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	// Initialize OpenAI client
	openaiAPIKey := os.Getenv("OPENAI_API_KEY")
	if openaiAPIKey == "" {
		log.Fatal("OPENAI_API_KEY is required")
	}

	llmClient := openai.NewClient(openaiAPIKey)

	handler := NewHandler(llmClient)

	setupRoutes(router, handler)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("server failed to start: ", err)
	}
}
