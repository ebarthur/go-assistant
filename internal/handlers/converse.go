package handlers

import (
	"github.com/gin-gonic/gin"
	"groq-api/client/groq"
	"groq-api/types"
	"net/http"
	"os"
)

// Converse handles conversations and maintain context
//
// @POST /converse
func Converse(c *gin.Context) {
	// Get request from post body
	var userInput *types.ConverseBody

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
	userRequest := &types.ConverseBody{
		Text: userInput.Text,
	}

	systemPrompt := "You are a great friend and companion. Engage in a conversation, respond to user input " +
		"and maintain context. " +
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
