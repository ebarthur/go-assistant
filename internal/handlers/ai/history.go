package ai

import (
	"github.com/gin-gonic/gin"
	"groq-api/db"
	"groq-api/initializers"
	"groq-api/types"
	"net/http"
)

// Get User History godoc
//
//	@Summary		Get User History
//	@Description	Handles the retrieval of user history
//	@Tags			generate
//	@Accept			json
//	@Produce		json
//	@Security		JWT
//	@Success		200	{object}	types.APISuccessMessage
//	@Failure		400	{object}	types.APIErrorMessage
//	@Failure		500	{object}	types.APIErrorMessage
//	@Router			/ai/history [get]
func GetHistory(c *gin.Context) {
	// get userid from cookie
	user, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You must be logged in to see history"})
		c.Abort()
		return
	}

	// search history table where UserID equals user.ID
	var histories []db.History
	userID := user.(*db.ClientUsers).ID
	result := initializers.DB.Where("user_id = ?", userID).Find(&histories)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get history"})
		c.Abort()
		return
	}

	// Prepare the response
	var historyResponses []types.HistoryResponse
	for _, history := range histories {
		historyResponses = append(historyResponses, types.HistoryResponse{
			ID:       history.ID,
			Request:  history.Request,
			Response: history.Response,
			Endpoint: history.Endpoint,
		})
	}

	userHistoryResponse := types.UserHistoryResponse{
		UserID:    userID,
		Histories: historyResponses,
	}

	// return history
	c.JSON(http.StatusOK, gin.H{"history": userHistoryResponse})
}
