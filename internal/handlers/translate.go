package handlers

import (
	"github.com/gin-gonic/gin"
	"groq-api/client/groq"
	"groq-api/types"
	"net/http"
	"os"
)

// Translate handles the translation of text to a specified language
//
// @POST /translate/:language
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
	systemPrompt := "You are a professional language translator. " +
		"Translate input to " + language + ". Only respond with the translated text. "
	prompt := textToTranslate

	translatedText, err := groqClient.ChatCompletion(groq.LLMModelLlama370b, systemPrompt, prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to translate text"})
		return
	}
	users, _ := c.Get("user")

	// return response text along w user details and just response when no user exists: for debugging purposes
	c.JSON(http.StatusOK, gin.H{"translated_text": *translatedText, "user": users})
}
