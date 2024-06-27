package handlers

import (
	"github.com/gin-gonic/gin"
	"groq-api/client/groq"
	"groq-api/types"
	"net/http"
	"os"
)

// Summarize handles the summarization of text
//
// @POST /summarize
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

	users, _ := c.Get("user")

	// return response text to user
	c.JSON(http.StatusOK, gin.H{"summarized_text": *AIResponse, "user": users})
}
