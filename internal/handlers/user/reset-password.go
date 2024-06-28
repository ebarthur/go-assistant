package user

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"groq-api/db"
	"groq-api/initializers"
	"groq-api/types"
	"net/http"
)

// Change Password godoc
//
//	@Summary		Reset Password
//	@Description	Reset user password
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Security		JWT
//	@Success		200	{object}	types.APISuccessMessage
//	@Failure		400	{object}	types.APIErrorMessage
//	@Failure		500	{object}	types.APIErrorMessage
//	@Router			/auth/reset-password [post]
func ChangePassword(c *gin.Context) {
	// get current and new pass off body
	var body *types.ChangePasswordRequestBody

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// get user from context(cookie)
	userInterface, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	user := userInterface.(*db.ClientUsers)

	// fetch user record from the database
	var dbUser *db.ClientUsers
	initializers.DB.First(&dbUser, "id=?", user.ID)
	if dbUser.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Check if new password matches the current password
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(body.NewPass))
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "New password cannot match current one"})
		return
	}

	// compare current password with the stored password
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect current password"})
		return
	}

	// hash the new password
	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.NewPass), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash new password"})
		return
	}

	// update the password in the database
	initializers.DB.Model(&dbUser).Update("Password", string(newHashedPassword))

	// return success
	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}
