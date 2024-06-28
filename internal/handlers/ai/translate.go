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

// Translate godoc
//
//	@Summary		AI Translate
//	@Description	Translate text to a specified language
//	@Tags			translate
//	@Accept			json
//	@Produce		json
//	@Security		JWT
//	@Success		200	{object}	types.APISuccessMessage
//	@Failure		400	{object}	types.APIErrorMessage
//	@Failure		500	{object}	types.APIErrorMessage
//	@Router			/ai/translate/:language [post]
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
		"Translate input to " + language + ". Only respond with the translated text."
	prompt := textToTranslate

	translatedText, err := groqClient.ChatCompletion(groq.LLMModelLlama370b, systemPrompt, prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to translate text"})
		return
	}
	// get user making request
	user, exists := c.Get("user")

	// get url path
	endpoint := c.Request.URL.Path

	// get request and response body
	requestBody := translateText.Text
	responseBody := *translatedText

	//store in db as history
	if !exists {
		// return response text along w user details and just response when no user exists: for debugging purposes
		c.JSON(http.StatusOK, gin.H{"translated_text": *translatedText})
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
		c.JSON(http.StatusOK, gin.H{"translated_text": *translatedText})

	}

}
