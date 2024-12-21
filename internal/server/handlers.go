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

func (h *Handler) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (h *Handler) handleCompletion(c *gin.Context) {
	var req struct {
		Prompt string `json:"prompt"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	completion, err := h.llmClient.CreateCompletion(c.Request.Context(), req.Prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": completion})
}
