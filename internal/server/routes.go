package server

import (
	"github.com/gin-gonic/gin"
)

func setupRoutes(r *gin.Engine, h *Handler) {
	r.GET("/health", h.healthCheck)
	r.POST("/completion", h.handleCompletion)
}
