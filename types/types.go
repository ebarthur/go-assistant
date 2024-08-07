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

type GenerateBody struct {
	Text string `json:"text" binding:"required"`
}

type EvaluateBody struct {
	Text string `json:"text" binding:"required"`
}

type SummarizeBody struct {
	Text string `json:"text" binding:"required"`
}

type ConverseBody struct {
	Text string `json:"text" binding:"required"`
}

type LoginRequestBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type SignupRequestBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type ChangePasswordRequestBody struct {
	Password string `json:"password" binding:"required"`
	NewPass  string `json:"new_pass" binding:"required,min=8"`
}

type HistoryResponse struct {
	ID       uint   `json:"id"`
	Request  string `json:"request"`
	Response string `json:"response"`
	Endpoint string `json:"endpoint"`
}

type UserHistoryResponse struct {
	UserID    uint              `json:"user_id"`
	Histories []HistoryResponse `json:"histories"`
}

type APIErrorMessage struct {
	ErrorMessage string `json:"error"`
}

type APISuccessMessage struct {
	SuccessMessage string `json:"success"`
	Data           any    `json:"data"`
}
