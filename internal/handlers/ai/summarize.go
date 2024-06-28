package ai

import (
	"github.com/gin-gonic/gin"
	"groq-api/client/groq"
	"groq-api/db"
	"groq-api/initializers"
	"groq-api/types"
	"net/http"
	"os"
)

// Summarize godoc
//
//	@Summary		AI Summarize
//	@Description	Summarize text to a concise version
//	@Tags			summarize
//	@Accept			json
//	@Produce		json
//	@Security		JWT
//	@Success		200	{object}	types.APISuccessMessage
//	@Failure		400	{object}	types.APIErrorMessage
//	@Failure		500	{object}	types.APIErrorMessage
//	@Router			/ai/summarize [post]
func Summarize(c *gin.Context) {
	// Get request from post body
	var userInput *types.SummarizeBody

	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// load api key from env
	var apiKey = os.Getenv("GROQ_API_KEY")
	if apiKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get GROQ_API_KEY"})
		return
	}
	groqClient := &groq.Client{ApiKey: apiKey}

	// Pass the request and prompt to the service
	userRequest := &types.SummarizeBody{
		Text: userInput.Text,
	}

	systemPrompt := "You are a text summarizer. Read prompt and return a concise summarized version."

	prompt := userRequest.Text
	AIResponse, err := groqClient.ChatCompletion(groq.LLMModelLlama370b, systemPrompt, prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to respond to request"})
		return
	}

	// get user making request
	user, exists := c.Get("user")

	// get url path
	endpoint := c.Request.URL.Path

	// get request and response body
	requestBody := userRequest.Text
	responseBody := *AIResponse

	//store in db as history
	if !exists {
		// return response text along w user details and just response when no user exists: for debugging purposes
		c.JSON(http.StatusOK, gin.H{"summarized_text": *AIResponse})
	} else {
		userID := user.(*db.ClientUsers).ID

		history := db.History{
			UserID:   userID,
			Request:  requestBody,
			Response: responseBody,
			Endpoint: endpoint,
		}

		// store history in the database
		result := initializers.DB.Create(&history)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to record history"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"summarized_text": *AIResponse})

	}

}
