package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	openRouterBaseURL = "https://openrouter.ai/api/v1"
	referer           = "http://localhost"
	siteTitle         = "NeuroCoach"
)

// Список доступных моделей с приоритетами
var availableModels = []string{
	"mistralai/mistral-7b-instruct:free",      // Самый стабильный
	"google/gemma-7b-it:free",                 // От Google
	"openchat/openchat-7b:free",               // Оптимизирован для чата
	"anthropic/claude-3-haiku:free",           // Claude 3 (самая мощная из бесплатных)
	"meta-llama/llama-3-8b-instruct:free",     // LLaMA 3
	"nousresearch/nous-hermes-2-mixtral:free", // Mixtral
	"microsoft/phi-3-mini-128k-instruct:free", // Phi-3
}

type OpenRouterMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenRouterRequest struct {
	Model    string              `json:"model"`
	Messages []OpenRouterMessage `json:"messages"`
	Stream   bool                `json:"stream,omitempty"`
}

type OpenRouterResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Error struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"error"`
}

type OpenRouterClient struct {
	apiKey            string
	httpClient        *http.Client
	currentModelIndex int
}

func NewOpenRouterClient(apiKey string) *OpenRouterClient {
	return &OpenRouterClient{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: 90 * time.Second,
		},
		currentModelIndex: 0,
	}
}

func (c *OpenRouterClient) getCurrentModel() string {
	return availableModels[c.currentModelIndex]
}

func (c *OpenRouterClient) switchToNextModel() {
	c.currentModelIndex = (c.currentModelIndex + 1) % len(availableModels)
}

func (c *OpenRouterClient) isModelError(err error) bool {
	errMsg := strings.ToLower(err.Error())
	return strings.Contains(errMsg, "model") ||
		strings.Contains(errMsg, "provider") ||
		strings.Contains(errMsg, "unavailable") ||
		strings.Contains(errMsg, "503") ||
		strings.Contains(errMsg, "overloaded")
}

func (c *OpenRouterClient) CreateChatCompletion(ctx context.Context, messages []OpenRouterMessage, requireJSON bool) (string, error) {
	var response string
	var err error
	attempts := 0
	maxAttempts := len(availableModels) * 2 // Максимум 2 круга по всем моделям

	for attempts = 0; attempts < maxAttempts; attempts++ {
		response, err = c.sendRequest(ctx, messages, requireJSON)
		if err == nil {
			break
		}

		// Пробуем следующую модель при ошибке
		if c.isModelError(err) {
			c.switchToNextModel()
		}

		// Небольшая задержка перед повторной попыткой
		time.Sleep(1 * time.Second)
	}

	return response, err
}

func (c *OpenRouterClient) sendRequest(ctx context.Context, messages []OpenRouterMessage, requireJSON bool) (string, error) {
	// Если требуется JSON ответ, добавляем инструкцию в системное сообщение
	if requireJSON && len(messages) > 0 && messages[0].Role == "system" {
		messages[0].Content += "\n\nIMPORTANT: Respond ONLY with valid JSON. Do not include any explanation or additional text."
	}

	requestBody := OpenRouterRequest{
		Model:    c.getCurrentModel(),
		Messages: messages,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("encoding error: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", openRouterBaseURL+"/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("request creation failed: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("HTTP-Referer", referer)
	req.Header.Set("X-Title", siteTitle)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		// Пытаемся распарсить ошибку
		var errorResp struct {
			Error struct {
				Message string `json:"message"`
			} `json:"error"`
		}
		if json.Unmarshal(body, &errorResp) == nil && errorResp.Error.Message != "" {
			return "", fmt.Errorf("API error [%d]: %s", resp.StatusCode, errorResp.Error.Message)
		}
		return "", fmt.Errorf("API error [%d]: %s", resp.StatusCode, string(body))
	}

	var response OpenRouterResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("response parsing failed: %w", err)
	}

	if response.Error.Message != "" {
		return "", fmt.Errorf("model error: %s", response.Error.Message)
	}

	if len(response.Choices) == 0 {
		return "", fmt.Errorf("empty response from AI")
	}

	return response.Choices[0].Message.Content, nil
}
