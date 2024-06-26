package api

import (
	"github.com/gin-gonic/gin"
	"groq-api/internal/handlers"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// TODO: Add more routes here
	r.POST("/translate/:language", handlers.Translate)
	r.POST("/general", handlers.OpenAI)

	return r
}
