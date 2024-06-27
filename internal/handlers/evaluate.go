package handlers

import (
	"github.com/gin-gonic/gin"
	"groq-api/client/groq"
	"groq-api/types"
	"net/http"
	"os"
)

// Evaluate handles questions that require a response
//
// @POST /evaluate
func Evaluate(c *gin.Context) {
	// Get request from post body
	var userInput *types.EvaluateBody

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
	userRequest := &types.EvaluateBody{
		Text: userInput.Text,
	}

	systemPrompt := "You are an evaluator. Read prompt and evaluate the accuracy and quality of text. " +
		"Provide feedback on the text. Check for grammar, spelling, and coherence."

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
