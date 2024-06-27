package api

import (
	"github.com/gin-gonic/gin"
	"groq-api/internal/handlers"
	"groq-api/internal/middleware"
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
	r.POST("/translate/:language", middleware.AuthMiddleware, handlers.Translate)
	r.POST("/generate", middleware.AuthMiddleware, handlers.Generate)
	r.POST("/summarize", middleware.AuthMiddleware, handlers.Summarize)
	r.POST("/evaluate", middleware.AuthMiddleware, handlers.Evaluate)
	r.POST("/converse", middleware.AuthMiddleware, handlers.Converse)

	r.POST("/x/sign-up", handlers.SignUp)
	r.POST("/x/sign-in", handlers.Login)

	return r
}
