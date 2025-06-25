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

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AIService struct {
	BaseService
	Client *OpenRouterClient
}

func NewAIService(repo repository.Repository, mongoRepo repository.MongoDBRep, openrouterKey string) *AIService {
	var client *OpenRouterClient

	if openrouterKey != "" {
		client = NewOpenRouterClient(openrouterKey)
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
	systemContent := `You are a fitness expert. Generate a workout plan and respond with ONLY valid JSON. No explanations, no markdown, no code blocks - just pure JSON.

JSON structure:
{
  "title": "Workout Plan Title",
  "workouts": [
    {
      "name": "Workout Name",
      "description": "Brief description",
      "status": "planned",
      "exercises": [
        {
          "name": "Exercise Name",
          "muscle_group": "Target Muscle",
          "sets": 3,
          "reps": 12,
          "rest_sec": 60,
          "notes": "Form tips",
          "technique": "How to perform"
        }
      ]
    }
  ]
}`

	// Prepare user prompt with profile data
	userPrompt := s.formatWorkoutPrompt(profile)

	// Prepare messages for OpenRouter
	messages := []OpenRouterMessage{
		{Role: "system", Content: systemContent},
		{Role: "user", Content: userPrompt},
	}

	// Call AI with structured response requirement
	content, err := s.Client.CreateChatCompletion(ctx, messages, true)
	if err != nil {
		return nil, NewServiceError(
			http.StatusInternalServerError,
			"AI request failed",
			err,
		)
	}

	// Log raw response for debugging
	fmt.Printf("Raw AI response: %s\n", content)

	// Clean response - remove markdown code blocks if present
	cleanContent := strings.TrimSpace(content)
	if strings.HasPrefix(cleanContent, "```json") {
		cleanContent = strings.TrimPrefix(cleanContent, "```json")
		cleanContent = strings.TrimSuffix(cleanContent, "```")
		cleanContent = strings.TrimSpace(cleanContent)
	}
	if strings.HasPrefix(cleanContent, "```") {
		cleanContent = strings.TrimPrefix(cleanContent, "```")
		cleanContent = strings.TrimSuffix(cleanContent, "```")
		cleanContent = strings.TrimSpace(cleanContent)
	}

	var generatedData struct {
		Title    string           `json:"title"`
		Workouts []models.Workout `json:"workouts"`
	}

	if err := json.Unmarshal([]byte(cleanContent), &generatedData); err != nil {
		return nil, NewServiceError(
			http.StatusInternalServerError,
			"Failed to parse AI response",
			fmt.Errorf("JSON parse error: %v, cleaned content: %s", err, cleanContent),
		)
	}

	// Заполняем отсутствующие поля значениями по умолчанию
	for i := range generatedData.Workouts {
		if generatedData.Workouts[i].Status == "" {
			generatedData.Workouts[i].Status = "planned"
		}
		for j := range generatedData.Workouts[i].Exercises {
			if generatedData.Workouts[i].Exercises[j].RestSec == 0 {
				generatedData.Workouts[i].Exercises[j].RestSec = 60
			}
		}
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
	messages := []OpenRouterMessage{
		{
			Role: "system",
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
		messages = append(messages, OpenRouterMessage{
			Role:    "user",
			Content: msg.Message,
		})
		messages = append(messages, OpenRouterMessage{
			Role:    "assistant",
			Content: msg.Response,
		})
	}

	// Add current message
	messages = append(messages, OpenRouterMessage{
		Role:    "user",
		Content: message,
	})

	// Call AI
	response, err := s.Client.CreateChatCompletion(ctx, messages, false)
	if err != nil {
		return "", NewServiceError(
			http.StatusInternalServerError,
			"AI request failed",
			err,
		)
	}

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

func (s *AIService) RegenerateWorkoutPlan(ctx context.Context, userComments string) (*models.WorkoutPlan, error) {
	userID, err := s.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if s.Client == nil {
		return nil, NewServiceError(
			http.StatusServiceUnavailable,
			"AI service unavailable",
			nil,
		)
	}

	// Получаем текущий план
	currentPlan, err := s.MongoDBRepo.GetWorkoutPlan(ctx, userID)
	if err != nil {
		return nil, NewServiceError(
			http.StatusNotFound,
			"No existing workout plan found",
			err,
		)
	}

	// Получаем профиль пользователя
	profile, err := s.Repo.GetFitnessProfile(ctx, userID)
	if err != nil {
		return nil, NewServiceError(
			http.StatusBadRequest,
			"Complete your profile first",
			err,
		)
	}

	// Подготавливаем системное сообщение
	systemContent := `You are a fitness expert updating workout plans. Respond with ONLY valid JSON. No explanations, no markdown, no code blocks - just pure JSON.

JSON structure:
{
  "title": "Updated Plan Title",
  "workouts": [
    {
      "name": "Workout Name",
      "description": "Brief description",
      "status": "planned",
      "exercises": [
        {
          "name": "Exercise Name",
          "muscle_group": "Target Muscle",
          "sets": 3,
          "reps": 12,
          "rest_sec": 60,
          "notes": "Form tips",
          "technique": "How to perform"
        }
      ]
    }
  ]
}`

	// Подготавливаем промпт с текущим планом и комментариями
	userPrompt := s.formatRegeneratePrompt(profile, currentPlan, userComments)

	// Подготавливаем сообщения для OpenRouter
	messages := []OpenRouterMessage{
		{Role: "system", Content: systemContent},
		{Role: "user", Content: userPrompt},
	}

	// Вызываем ИИ
	content, err := s.Client.CreateChatCompletion(ctx, messages, true)
	if err != nil {
		return nil, NewServiceError(
			http.StatusInternalServerError,
			"AI request failed",
			err,
		)
	}

	// Log raw response for debugging
	fmt.Printf("Raw AI response: %s\n", content)

	// Clean response - remove markdown code blocks if present
	cleanContent := strings.TrimSpace(content)
	if strings.HasPrefix(cleanContent, "```json") {
		cleanContent = strings.TrimPrefix(cleanContent, "```json")
		cleanContent = strings.TrimSuffix(cleanContent, "```")
		cleanContent = strings.TrimSpace(cleanContent)
	}
	if strings.HasPrefix(cleanContent, "```") {
		cleanContent = strings.TrimPrefix(cleanContent, "```")
		cleanContent = strings.TrimSuffix(cleanContent, "```")
		cleanContent = strings.TrimSpace(cleanContent)
	}

	var generatedData struct {
		Title    string           `json:"title"`
		Workouts []models.Workout `json:"workouts"`
	}

	if err := json.Unmarshal([]byte(cleanContent), &generatedData); err != nil {
		return nil, NewServiceError(
			http.StatusInternalServerError,
			"Failed to parse AI response",
			fmt.Errorf("JSON parse error: %v, cleaned content: %s", err, cleanContent),
		)
	}

	// Заполняем отсутствующие поля значениями по умолчанию
	for i := range generatedData.Workouts {
		if generatedData.Workouts[i].Status == "" {
			generatedData.Workouts[i].Status = "planned"
		}
		for j := range generatedData.Workouts[i].Exercises {
			if generatedData.Workouts[i].Exercises[j].RestSec == 0 {
				generatedData.Workouts[i].Exercises[j].RestSec = 60
			}
		}
	}

	// Создаем обновленный план
	now := time.Now()
	updatedPlan := &models.WorkoutPlan{
		ID:        currentPlan.ID, // Сохраняем ID существующего плана
		UserID:    userID,
		Title:     generatedData.Title,
		Workouts:  generatedData.Workouts,
		Status:    true,
		CreatedAt: currentPlan.CreatedAt, // Сохраняем дату создания
		UpdatedAt: now,
	}

	// Добавляем ID к вложенным объектам
	for i := range updatedPlan.Workouts {
		updatedPlan.Workouts[i].WorkoutID = primitive.NewObjectID()

		for j := range updatedPlan.Workouts[i].Exercises {
			updatedPlan.Workouts[i].Exercises[j].ExerciseID = primitive.NewObjectID()
		}
	}

	// Сохраняем обновленный план
	if err := s.MongoDBRepo.SaveWorkoutPlan(ctx, updatedPlan); err != nil {
		return nil, NewServiceError(
			http.StatusInternalServerError,
			"Failed to save updated workout plan",
			err,
		)
	}

	return updatedPlan, nil
}

func (s *AIService) GetWorkoutByID(ctx context.Context, workoutID string) (*models.Workout, error) {
	userID, err := s.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	workout, err := s.MongoDBRepo.GetWorkoutByID(ctx, userID, workoutID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, NewServiceError(
				http.StatusNotFound,
				"Workout not found",
				err,
			)
		}
		return nil, NewServiceError(
			http.StatusInternalServerError,
			"Failed to get workout",
			err,
		)
	}

	return workout, nil
}

func (s *AIService) CompleteWorkout(ctx context.Context, workoutID string) error {
	userID, err := s.GetUserIDFromContext(ctx)
	if err != nil {
		return err
	}

	return s.MongoDBRepo.CompleteWorkout(ctx, userID, workoutID)
}

func (s *AIService) GetUserProgress(ctx context.Context) (*models.UserProgress, error) {
	userID, err := s.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return s.MongoDBRepo.GetUserProgress(ctx, userID)
}

func (s *AIService) formatRegeneratePrompt(profile *models.FitnessProfile, currentPlan *models.WorkoutPlan, userComments string) string {
	var sb strings.Builder

	sb.WriteString("Update the existing workout plan based on user feedback.\n\n")
	sb.WriteString("User Profile:\n")
	fmt.Fprintf(&sb, "- Age: %d\n", profile.Age)
	fmt.Fprintf(&sb, "- Height: %.1f cm\n", profile.Height)
	fmt.Fprintf(&sb, "- Weight: %.1f kg\n", profile.Weight)
	fmt.Fprintf(&sb, "- Fitness Goal: %s\n", profile.Goal)
	fmt.Fprintf(&sb, "- Fitness Level: %s\n", profile.FitnessLevel)
	fmt.Fprintf(&sb, "- Available Time: %d minutes per week\n", profile.AvailableMinutes)

	if len(profile.HealthIssues) > 0 {
		sb.WriteString("- Health Issues: ")
		sb.WriteString(strings.Join(profile.HealthIssues, ", "))
		sb.WriteString("\n")
	}

	sb.WriteString("\nCurrent Workout Plan:\n")
	fmt.Fprintf(&sb, "Title: %s\n", currentPlan.Title)
	for i, workout := range currentPlan.Workouts {
		fmt.Fprintf(&sb, "Workout %d: %s - %s\n", i+1, workout.Name, workout.Description)
		for j, exercise := range workout.Exercises {
			fmt.Fprintf(&sb, "  Exercise %d: %s (%s) - %d sets x %d reps, %d sec rest\n",
				j+1, exercise.Name, exercise.MuscleGroup, exercise.Sets, exercise.Reps, exercise.RestSec)
		}
	}

	sb.WriteString("\nUser Comments/Feedback:\n")
	sb.WriteString(userComments)

	sb.WriteString("\n\nPlease update the workout plan based on the user's feedback while maintaining:")
	sb.WriteString("\n1. Appropriate difficulty for their fitness level")
	sb.WriteString("\n2. Alignment with their fitness goals")
	sb.WriteString("\n3. Consideration of their health issues")
	sb.WriteString("\n4. Time constraints")
	sb.WriteString("\n5. Progressive overload principles")

	return sb.String()
}
