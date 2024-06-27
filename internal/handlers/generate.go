package handlers

import (
	"github.com/gin-gonic/gin"
	"groq-api/client/groq"
	"groq-api/types"
	"net/http"
	"os"
)

// Generate handles questions that require a response
//
// @POST /generate
func Generate(c *gin.Context) {
	// Get request from post body
	var userInput *types.GenerateBody

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
	userRequest := &types.GenerateBody{
		Text: userInput.Text,
	}

	systemPrompt := "You are a general purpose AI. Read prompt and respond accordingly. " +
		"Keep response concise and relevant."

	prompt := userRequest.Text
	AIResponse, err := groqClient.ChatCompletion(groq.LLMModelLlama370b, systemPrompt, prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to respond to request"})
		return
	}

	users, _ := c.Get("user")

	// return response text to user
	c.JSON(http.StatusOK, gin.H{"response": *AIResponse, "user": users})
}
