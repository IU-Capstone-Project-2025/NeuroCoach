package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"rest-api/internal/models"
	"rest-api/internal/repository"

	"github.com/sashabaranov/go-openai"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AIService struct {
	BaseService
	Client *openai.Client
}

func NewAIService(repo repository.Repository, mongoRepo repository.MongoDBRep, apiKey string) *AIService {
	var client *openai.Client
	if apiKey != "" {
		client = openai.NewClient(apiKey)
	}
	return &AIService{
		BaseService: BaseService{Repo: repo, MongoDBRepo: mongoRepo},
		Client:      client,
	}
}

func (s *AIService) GenerateWorkoutPlan(ctx context.Context) (*models.WorkoutPlan, error) {
	userID, err := s.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	// Check if we have a recent plan
	plan, err := s.MongoDBRepo.GetWorkoutPlan(ctx, userID)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return nil, NewServiceError(
				http.StatusInternalServerError,
				"Failed to get existed workout plan",
				err,
			)
		}
	} else if plan != nil {
		return plan, nil
	}

	// Get user profile
	profile, err := s.Repo.GetFitnessProfile(ctx, userID)
	if err != nil {
		return nil, NewServiceError(
			http.StatusBadRequest,
			"Complete your profile first",
			err,
		)
	}

	// Generate new plan
	if s.Client == nil {
		return nil, NewServiceError(
			http.StatusServiceUnavailable,
			"AI service unavailable",
			nil,
		)
	}

	// Prepare system message with JSON schema
	systemMsg := openai.ChatCompletionMessage{
		Role: openai.ChatMessageRoleSystem,
		Content: `You are a fitness expert generating workout plans in JSON format. 
			Respond ONLY with valid JSON matching this structure:
			{
			"title": "string",
			"workouts": [
				{
				"name": "string",
				"description": "string",
				"status": "planned",
				"experience_gain": 100,
				"exercises": [
					{
					"name": "string",
					"muscle_group": "string",
					"sets": 3,
					"reps": 12,
					"rest_sec": 60,
					"notes": "string",
					"technique": "string"
					}
				]
				}
			]
			}`,
	}

	// Prepare user prompt with profile data
	userPrompt := s.formatWorkoutPrompt(profile)
	// Call AI with structured response requirement
	resp, err := s.Client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			systemMsg,
			{Role: openai.ChatMessageRoleUser, Content: userPrompt},
		},
		ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeJSONObject,
		},
	})

	if err != nil {
		return nil, NewServiceError(
			http.StatusInternalServerError,
			"AI request failed",
			err,
		)
	}

	if len(resp.Choices) == 0 {
		return nil, NewServiceError(
			http.StatusInternalServerError,
			"No response from AI",
			nil,
		)
	}

	content := resp.Choices[0].Message.Content

	var generatedData struct {
		Title    string           `json:"title"`
		Workouts []models.Workout `json:"workouts"`
	}

	if err := json.Unmarshal([]byte(content), &generatedData); err != nil {
		return nil, NewServiceError(
			http.StatusInternalServerError,
			"Failed to parse AI response",
			fmt.Errorf("JSON parse error: %v, content: %s", err, content),
		)
	}

	// Create full workout plan
	now := time.Now()
	workoutPlan := &models.WorkoutPlan{
		UserID:    userID,
		Title:     generatedData.Title,
		Workouts:  generatedData.Workouts,
		Status:    true,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Add IDs to nested objects
	for i := range workoutPlan.Workouts {
		workoutPlan.Workouts[i].WorkoutID = primitive.NewObjectID()

		for j := range workoutPlan.Workouts[i].Exercises {
			workoutPlan.Workouts[i].Exercises[j].ExerciseID = primitive.NewObjectID()
		}
	}

	// Save the generated plan
	if err := s.MongoDBRepo.SaveWorkoutPlan(ctx, workoutPlan); err != nil {
		// Log error but don't fail the request
		fmt.Printf("Failed to save workout plan: %v\n", err)
	}

	return workoutPlan, nil
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
	history, err := s.MongoDBRepo.GetChatHistory(ctx, userID)
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
	start := 0
	if len(history) > 10 {
		start = len(history) - 10
	}

	for i := start; i < len(history); i++ {
		msg := history[i]
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
	chatMsg := &models.ChatMessage{
		UserID:   userID,
		Message:  message,
		Response: response,
		IsUser:   true,
	}

	if err := s.MongoDBRepo.SaveChatMessage(ctx, chatMsg); err != nil {
		return "", NewServiceError(
			http.StatusInternalServerError,
			"Failed to save chat message",
			err,
		)
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

	return s.MongoDBRepo.GetChatHistory(ctx, userID)
}
