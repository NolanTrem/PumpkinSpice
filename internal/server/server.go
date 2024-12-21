package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	setupRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("server failed to start: ", err)
	}
}
