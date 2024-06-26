package types

type GroqMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatCompletionRequest struct {
	Messages    []GroqMessage `json:"messages"`
	Model       string        `json:"model"`
	Temperature int           `json:"temperature"`
	MaxTokens   int           `json:"max_tokens"`
	TopP        int           `json:"top_p"`
	Stream      bool          `json:"stream"`
	Stop        interface{}   `json:"stop"`
}

type ChatCompletionResponse struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int     `json:"prompt_tokens"`
		PromptTime       float64 `json:"prompt_time"`
		CompletionTokens int     `json:"completion_tokens"`
		CompletionTime   float64 `json:"completion_time"`
		TotalTokens      int     `json:"total_tokens"`
		TotalTime        float64 `json:"total_time"`
	} `json:"usage"`
	SystemFingerprint string `json:"system_fingerprint"`
	XGroq             struct {
		Id string `json:"id"`
	} `json:"x_groq"`
}

// Response from users

type TranslateBody struct {
	Text string `json:"text" binding:"required"`
}

type OpenBody struct {
	Request string `json:"request" binding:"required"`
}
