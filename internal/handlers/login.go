package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"groq-api/db"
	"groq-api/initializers"
	"groq-api/types"
	"net/http"
	"os"
	"time"
)

// Login handles user sign in
//
// @POST /x/sign-in
func Login(c *gin.Context) {
	// get email and password from request
	var ReqBody types.LoginRequestBody

	err := c.ShouldBindJSON(&ReqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// find user with email
	var user db.ClientUsers
	initializers.DB.First(&user, "email = ?", ReqBody.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		return
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(ReqBody.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		return
	}
	// generate jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subject": user.ID,
		"expiry":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create token"})
		return
	}

	// store token in a cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	// send statusOk
	c.JSON(http.StatusOK, gin.H{})
}
