package routes

import (
	"github.com/gin-gonic/gin"
	"groq-api/internal/handlers/user"
	"groq-api/internal/middleware"
)

func AuthRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	auth.POST("/sign-up", user.SignUp)
	auth.POST("/sign-in", user.Login)
	auth.POST("/reset-password", middleware.AuthMiddleware, user.ChangePassword)

}
