package groq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"groq-api/types"
	"io"
	"net/http"
)

const (
	apiBaseUrl = "https://api.groq.com/openai"
	SYSTEM     = "system"
	USER       = "user"

	LLMModelLlama38b  = "llama3-8b-8192"
	LLMModelLlama370b = "llama3-70b-8192"
	//LLMModelMixtral8x7b32k = "mixtral-8x7b-32768"
	//LLMModelGemma7b        = "gemma-7b-it"
)

type Client struct {
	ApiKey string
}

func (g *Client) ChatCompletion(llmModel string, systemPrompt string, prompt string) (*string, error) {

	llm := llmModel

	if llmModel == "" {
		//default to llama8B
		llm = LLMModelLlama38b
	}
	groqMessages := make([]types.GroqMessage, 0)

	if systemPrompt != "" {
		systemMessage := types.GroqMessage{
			Role:    SYSTEM,
			Content: systemPrompt,
		}
		groqMessages = append(groqMessages, systemMessage)
	}

	if prompt != "" {
		userMessage := types.GroqMessage{
			Role:    USER,
			Content: prompt,
		}
		groqMessages = append(groqMessages, userMessage)
	} else {
		return nil, fmt.Errorf("prompt is required")
	}

	chatCompletionRequest := &types.ChatCompletionRequest{
		Messages:    groqMessages,
		Model:       llm,
		Temperature: 0,
		MaxTokens:   1024,
		TopP:        1,
		Stream:      false,
		Stop:        nil,
	}

	chatCompletionRequestJson, err := json.Marshal(chatCompletionRequest)
	if err != nil {
		return nil, err
	}

	//send http post request
	chatCompletionUrl := "/v1/chat/completions"
	finalUrl := fmt.Sprintf("%s%s", apiBaseUrl, chatCompletionUrl)

	req, err := http.NewRequest(http.MethodPost, finalUrl, bytes.NewBuffer(chatCompletionRequestJson))
	if err != nil {
		return nil, err
	}

	//set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", g.ApiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code: %d, reason: %s", resp.StatusCode, resp.Status)
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	chatCompletionResp := &types.ChatCompletionResponse{}

	err = json.Unmarshal(body, &chatCompletionResp)
	if err != nil {
		return nil, err
	}

	var content string
	if chatCompletionResp.Choices != nil && len(chatCompletionResp.Choices) > 0 {
		content = chatCompletionResp.Choices[0].Message.Content
	} else {
		return nil, fmt.Errorf("no choices")
	}

	return &content, nil
}
