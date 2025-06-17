package services

import (
	"context"
	"fmt"
	"strings"
	"net/http"

	"github.com/sashabaranov/go-openai"
	"rest-api/internal/models"
	"rest-api/internal/repository"
)

type AIService struct {
	BaseService
	Client *openai.Client
}

func NewAIService(repo repository.Repository, apiKey string) *AIService {
	var client *openai.Client
	if apiKey != "" {
		client = openai.NewClient(apiKey)
	}
	return &AIService{
		BaseService: BaseService{Repo: repo},
		Client:      client,
	}
}

func (s *AIService) GenerateWorkoutPlan(ctx context.Context) (string, error) {
	userID, err := s.GetUserIDFromContext(ctx)
	if err != nil {
		return "", err
	}

	// Check if we have a recent plan
	plan, err := s.Repo.GetWorkoutPlan(ctx, userID)
	if err == nil && plan != nil {
		return plan.Content, nil
	}

	// Get user profile
	profile, err := s.Repo.GetFitnessProfile(ctx, userID)
	if err != nil {
		return "", NewServiceError(
			http.StatusBadRequest, 
			"Complete your profile first", 
			err,
		)
	}

	// Generate new plan
	if s.Client == nil {
		return "", NewServiceError(
			http.StatusServiceUnavailable, 
			"AI service unavailable", 
			nil,
		)
	}

	prompt := s.formatWorkoutPrompt(profile)
	resp, err := s.Client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleUser, Content: prompt},
		},
	})

	if err != nil {
		return "", NewServiceError(
			http.StatusInternalServerError, 
			"AI request failed", 
			err,
		)
	}

	if len(resp.Choices) == 0 {
		return "", NewServiceError(
			http.StatusInternalServerError, 
			"No response from AI", 
			nil,
		)
	}

	content := resp.Choices[0].Message.Content

	// Save the generated plan
	if err := s.Repo.SaveWorkoutPlan(ctx, userID, &models.WorkoutPlan{
		Content: content,
	}); err != nil {
		// Log error but don't fail the request
		fmt.Printf("Failed to save workout plan: %v\n", err)
	}

	return content, nil
}

func (s *AIService) Chat(ctx context.Context, message string) (string, error) {
	userID, err := s.GetUserIDFromContext(ctx)
	if err != nil {
		return "", err
	}

	if s.Client == nil {
		return "", NewServiceError(
			http.StatusServiceUnavailable, 
			"AI service unavailable", 
			nil,
		)
	}

	// Get chat history
	history, err := s.Repo.GetChatHistory(ctx, userID, 10)
	if err != nil {
		return "", NewServiceError(
			http.StatusInternalServerError, 
			"Failed to get chat history", 
			err,
		)
	}

	// Build conversation context
	messages := []openai.ChatCompletionMessage{
		{
			Role: openai.ChatMessageRoleSystem,
			Content: "You are a helpful fitness assistant. " +
				"Provide concise and helpful responses about fitness, nutrition, and health.",
		},
	}

	// Add history to context
	for _, msg := range history {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: msg.Message,
		})
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: msg.Response,
		})
	}

	// Add current message
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: message,
	})

	// Call AI
	resp, err := s.Client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo,
		Messages: messages,
	})
	if err != nil {
		return "", NewServiceError(
			http.StatusInternalServerError, 
			"AI request failed", 
			err,
		)
	}

	if len(resp.Choices) == 0 {
		return "", NewServiceError(
			http.StatusInternalServerError, 
			"No response from AI", 
			nil,
		)
	}

	response := resp.Choices[0].Message.Content

	// Save chat message
	if err := s.Repo.SaveChatMessage(ctx, &models.ChatMessage{
		UserID:   userID,
		Message:  message,
		Response: response,
		IsUser:   true,
	}); err != nil {
		// Log error but don't fail the request
		fmt.Printf("Failed to save chat message: %v\n", err)
	}

	return response, nil
}

func (s *AIService) formatWorkoutPrompt(profile *models.FitnessProfile) string {
	var sb strings.Builder

	sb.WriteString("Create a personalized workout plan with the following specifications:\n")
	fmt.Fprintf(&sb, "- Age: %d\n", profile.Age)
	fmt.Fprintf(&sb, "- Height: %.1f cm\n", profile.Height)
	fmt.Fprintf(&sb, "- Weight: %.1f kg\n", profile.Weight)
	fmt.Fprintf(&sb, "- Fitness Goal: %s\n", profile.Goal)
	fmt.Fprintf(&sb, "- Timeframe: %s\n", profile.Timeframe)
	fmt.Fprintf(&sb, "- Fitness Level: %s\n", profile.FitnessLevel)
	fmt.Fprintf(&sb, "- Available Time: %d minutes per week\n", profile.AvailableMinutes)

	if len(profile.HealthIssues) > 0 {
		sb.WriteString("- Health Issues: ")
		sb.WriteString(strings.Join(profile.HealthIssues, ", "))
		sb.WriteString("\n")
	}

	sb.WriteString("\nThe plan should include:\n")
	sb.WriteString("1. Weekly schedule with specific exercises\n")
	sb.WriteString("2. Sets, reps, and rest periods\n")
	sb.WriteString("3. Progression plan\n")
	sb.WriteString("4. Safety considerations\n")
	sb.WriteString("5. Format in markdown\n")

	return sb.String()
}

func (s *AIService) GetChatHistory(ctx context.Context) ([]models.ChatMessage, error) {
	userID, err := s.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	// Get last 20 messages
	messages, err := s.Repo.GetChatHistory(ctx, userID, 20)
	if err != nil {
		return nil, NewServiceError(
			http.StatusInternalServerError, 
			"Failed to get chat history", 
			err,
		)
	}

	return messages, nil
}