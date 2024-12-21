package server

import (
	"net/http"

	openai "github.com/NolanTrem/PumpkinSpice/internal/llm/openai"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	llmClient *openai.Client
}

func NewHandler(llmClient *openai.Client) *Handler {
	return &Handler{
		llmClient: llmClient,
	}
}

// Simple health check handler
func (h *Handler) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (h *Handler) handleCompletion(c *gin.Context) {
	var req struct {
		Prompt string `json:"prompt"`
		Stream bool   `json:"stream"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Stream {
		// Set headers for plain text streaming
		c.Header("Content-Type", "text/plain")
		c.Header("Cache-Control", "no-cache")
		c.Header("Connection", "keep-alive")

		// Create a stream of completion chunks
		chunks, err := h.llmClient.StreamCompletion(c.Request.Context(), req.Prompt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Stream the chunks to the client
		for chunk := range chunks {
			_, err := c.Writer.WriteString(chunk)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.Writer.Flush()
		}
		return
	}

	// Create a single completion
	completion, err := h.llmClient.CreateCompletion(c.Request.Context(), req.Prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": completion})
}
