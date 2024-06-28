package routes

import (
	"github.com/gin-gonic/gin"
	"groq-api/internal/handlers/ai"
	"groq-api/internal/middleware"
)

func AssistantRoutes(r *gin.RouterGroup) {
	assistant := r.Group("/ai")
	assistant.POST("/translate/:language", middleware.AuthMiddleware, ai.Translate)
	assistant.POST("/generate", middleware.AuthMiddleware, ai.Generate)
	assistant.POST("/summarize", middleware.AuthMiddleware, ai.Summarize)
	assistant.POST("/evaluate", middleware.AuthMiddleware, ai.Evaluate)
	assistant.POST("/converse", middleware.AuthMiddleware, ai.Converse)
	assistant.GET("/history", middleware.AuthMiddleware, ai.GetHistory)
}
