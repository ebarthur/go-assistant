package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"groq-api/db"
	"groq-api/initializers"
	"net/http"
	"os"
	"time"
)

func AuthMiddleware(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	// if token is present
	if err == nil {
		// decode/validate it
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// check the expiry
			if float64(time.Now().Unix()) > claims["expiry"].(float64) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
				c.Abort()
				return
			}

			// find the user with token subject
			var user *db.ClientUsers
			result := initializers.DB.First(&user, "id = ?", claims["subject"])
			if result.Error != nil || user.ID == 0 {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
				c.Abort()
				return
			}

			// attach to req
			c.Set("user", user)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
		}
	}

	// continue
	// also, when there is no authorization
	c.Next()
}
