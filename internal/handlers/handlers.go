package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"groq-api/client/groq"
	"groq-api/types"
	"net/http"
	"os"
)

func Translate(c *gin.Context) {

	// Get text from post body
	var translateText *types.TranslateBody

	err := c.ShouldBindJSON(&translateText)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// load api key from env
	var apiKey = os.Getenv("GROQ_API_KEY")
	if apiKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get GROQ_API_KEY"})
		return
	}
	groqClient := &groq.Client{ApiKey: apiKey}

	// Get language to translate from the request url
	language := c.Param("language")
	// Validate language parameter
	if language == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Language parameter is required"})
		return
	}

	// Pass the text to the translation service with prompt

	textToTranslate := translateText.Text
	systemPrompt := "you are a professional language translator. " +
		"translate the passage or input to " + language + "."
	prompt := fmt.Sprintf("Return a translation of this passage: %s", textToTranslate)

	translatedText, err := groqClient.ChatCompletion(groq.LLMModelLlama370b, systemPrompt, prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to translate text"})
		return
	}

	// return translated text to user
	c.JSON(http.StatusOK, gin.H{"translated_text": *translatedText})
}

// OpenAI because it is open to any prompt
func OpenAI(c *gin.Context) {
	// Get request from post body
	var userInput *types.OpenBody

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

	// Pass the request and prompt  to the service
	userRequest := &types.OpenBody{
		Request: userInput.Request,
	}

	systemPrompt := "you are a general purpose AI. Your name is Arthur and you're Ghanaian." +
		" Respond to requests with the most accurate information. " +
		"And please keep responses short and concise."

	prompt := fmt.Sprintf("Return a well curated response to this."+
		" Try to sound human if necessary: %s", userRequest.Request)

	AIResponse, err := groqClient.ChatCompletion(groq.LLMModelLlama370b, systemPrompt, prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to respond to request"})
		return
	}

	// return response text to user
	c.JSON(http.StatusOK, gin.H{"response": *AIResponse})
}
