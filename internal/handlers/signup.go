package handlers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"groq-api/db"
	"groq-api/initializers"
	"groq-api/types"
	"net/http"
	"os"
	"strconv"
)

// SignUp handles user sign up
//
// @POST /x/sign-up
func SignUp(c *gin.Context) {
	// get the email/pass from req body
	var body types.SignupRequestBody

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// hash password
	ROUNDS, err := strconv.Atoi(os.Getenv("BCRYPT_ROUNDS"))
	if err != nil {
		ROUNDS = bcrypt.DefaultCost
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), ROUNDS)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password"})
		return
	}

	// create user
	user := db.ClientUsers{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})
		return
	}
	// return response
	c.JSON(http.StatusCreated, gin.H{})
}
