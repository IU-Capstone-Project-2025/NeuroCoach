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

	// Calculate required workouts
	workoutsPerWeek := profile.AvailableMinutes / 50
	if workoutsPerWeek < 2 {
		workoutsPerWeek = 2
	}
	if workoutsPerWeek > 6 {
		workoutsPerWeek = 6
	}

	// Prepare system message with JSON schema
	systemContent := fmt.Sprintf(`You are a fitness expert. Generate a workout plan with EXACTLY %d workouts and respond with ONLY valid JSON.

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
}

IMPORTANT: Create EXACTLY %d different workouts in the workouts array.`, workoutsPerWeek, workoutsPerWeek)

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
		fmt.Printf("AI REQUEST FAILED during generate plan due to: %s", err)
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

	// Fill missing fields with default values
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

	// Trim workouts to required count
	if len(generatedData.Workouts) > workoutsPerWeek {
		generatedData.Workouts = generatedData.Workouts[:workoutsPerWeek]
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

	// Generate full schedule for timeframe
	totalWeeks := s.getWeeksFromTimeframe(profile.Timeframe)
	totalWorkouts := workoutsPerWeek * totalWeeks
	fullSchedule := s.generateFullSchedule(generatedData.Workouts, workoutsPerWeek, totalWorkouts)

	// Replace workouts with full schedule
	workoutPlan.Workouts = fullSchedule

	// Save both short and full plans
	shortPlan := &models.ShortWorkoutPlan{
		UserID:          userID,
		Title:           generatedData.Title,
		BaseWorkouts:    generatedData.Workouts,
		Timeframe:       profile.Timeframe,
		WorkoutsPerWeek: workoutsPerWeek,
		Status:          true,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	if err := s.MongoDBRepo.SaveShortPlan(ctx, shortPlan); err != nil {
		fmt.Printf("Failed to save short plan: %v\n", err)
	}

	if err := s.MongoDBRepo.SaveWorkoutPlan(ctx, workoutPlan); err != nil {
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

	// Get user's fitness profile to check if they're a beginner
	profile, err := s.Repo.GetFitnessProfile(ctx, userID)
	isBeginner := false
	if err == nil && profile != nil && profile.FitnessLevel == "beginner" {
		isBeginner = true
	}

	// Build conversation context with beginner mode if needed
	systemContent := "You are a helpful fitness assistant. Provide concise and helpful responses about fitness, nutrition, and health."
	if isBeginner {
		systemContent += " IMPORTANT: The user is a beginner with limited fitness knowledge. Explain concepts in very simple terms as if explaining to a kid. Avoid technical jargon, use basic language, and include extra safety tips."
	}

	messages := []OpenRouterMessage{
		{
			Role:    "system",
			Content: systemContent,
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
		fmt.Printf("ERROR: AI REQUEST FAILED in Chat: %v\n", err)
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

	// Add timeframe-specific guidance
	sb.WriteString("\n")
	sb.WriteString(s.getTimeframeGuidance(profile.Timeframe, profile.AvailableMinutes))
	sb.WriteString("\n")

	sb.WriteString("\nThe plan should include:\n")
	sb.WriteString("1. Weekly schedule with specific exercises\n")
	sb.WriteString("2. Sets, reps, and rest periods\n")
	sb.WriteString("3. Progression plan\n")
	sb.WriteString("4. Safety considerations\n")
	sb.WriteString("5. Format in JSON\n")

	// Add beginner mode instructions if user is a beginner
	if profile.FitnessLevel == "beginner" {
		sb.WriteString("\nIMPORTANT: This user is a beginner. Please explain all exercises in very simple terms as if explaining to someone with no fitness experience. Use basic language, avoid technical jargon, and include extra safety tips. Provide detailed step-by-step instructions for each exercise.\n")
	}

	return sb.String()
}

func (s *AIService) getTimeframeGuidance(timeframe string, availableMinutes int) string {
	// Calculate workouts per week (assuming 45-60 min per workout)
	workoutsPerWeek := availableMinutes / 50
	if workoutsPerWeek < 1 {
		workoutsPerWeek = 1
	}
	if workoutsPerWeek > 6 {
		workoutsPerWeek = 6
	}

	// Calculate total weeks and workouts
	totalWeeks := s.getWeeksFromTimeframe(timeframe)
	totalWorkouts := workoutsPerWeek * totalWeeks

	// Generate workout schedule with dates
	schedule := s.generateWorkoutSchedule(workoutsPerWeek, totalWeeks)

	return fmt.Sprintf("PLAN: %d workouts per week for %d weeks (%d total workouts).\nSCHEDULE: %s\nFOCUS: %s",
		workoutsPerWeek, totalWeeks, totalWorkouts, schedule, s.getFocusByTimeframe(timeframe))
}

func (s *AIService) getWeeksFromTimeframe(timeframe string) int {
	switch timeframe {
	case "1month":
		return 4
	case "3months":
		return 12
	case "6months":
		return 24
	case "1year":
		return 52
	default:
		return 12
	}
}

func (s *AIService) getFocusByTimeframe(timeframe string) string {
	switch timeframe {
	case "1month":
		return "Foundation building and form mastery"
	case "3months":
		return "Skill development and strength building"
	case "6months":
		return "Progressive overload and body transformation"
	case "1year":
		return "Complete fitness transformation with periodization"
	default:
		return "General fitness improvement"
	}
}

func (s *AIService) generateWorkoutSchedule(workoutsPerWeek, totalWeeks int) string {
	now := time.Now()
	var schedule []string

	// Generate first week as example
	for i := 0; i < workoutsPerWeek && i < 7; i++ {
		workoutDate := now.AddDate(0, 0, i*2) // Every other day
		schedule = append(schedule, workoutDate.Format("Jan 2"))
	}

	return fmt.Sprintf("Week 1: %s (pattern repeats for %d weeks)", strings.Join(schedule, ", "), totalWeeks)
}

func (s *AIService) fixJSONWithAI(ctx context.Context, content, errorMsg string) (string, error) {
	messages := []OpenRouterMessage{
		{Role: "system", Content: "Fix this JSON to be valid. Return only the corrected JSON without any additional comments."},
		{Role: "user", Content: fmt.Sprintf("Error: %s\nJSON: %s\nFix this JSON", errorMsg, content)},
	}
	response, err := s.Client.CreateChatCompletion(ctx, messages, false)
	if err != nil {
		fmt.Printf("ERROR: AI request failed in fixJSONWithAI: %v\n", err)
	}
	return response, err
}

func (s *AIService) createFallbackWorkouts(count int) []models.Workout {
	workouts := []models.Workout{
		{Name: "Upper Body", Description: "Chest, back, shoulders", Status: "planned", Exercises: []models.Exercise{{Name: "Push-ups", MuscleGroup: "Chest", Sets: 3, Reps: 12, RestSec: 60}}},
		{Name: "Lower Body", Description: "Legs and glutes", Status: "planned", Exercises: []models.Exercise{{Name: "Squats", MuscleGroup: "Legs", Sets: 3, Reps: 15, RestSec: 60}}},
		{Name: "Core", Description: "Abs and core", Status: "planned", Exercises: []models.Exercise{{Name: "Plank", MuscleGroup: "Core", Sets: 3, Reps: 1, RestSec: 60}}},
		{Name: "Cardio", Description: "Cardiovascular training", Status: "planned", Exercises: []models.Exercise{{Name: "Walking", MuscleGroup: "Cardio", Sets: 1, Reps: 1, RestSec: 0}}},
	}

	if count > len(workouts) {
		count = len(workouts)
	}
	return workouts[:count]
}

func (s *AIService) calculateWorkoutDate(workoutIndex, workoutsPerWeek int) time.Time {
	now := time.Now()
	weekNumber := workoutIndex / workoutsPerWeek
	positionInWeek := workoutIndex % workoutsPerWeek

	// Calculate days within week based on workouts per week
	var dayInWeek int
	if workoutsPerWeek <= 3 {
		dayInWeek = positionInWeek * 2 // Every other day
	} else if workoutsPerWeek <= 5 {
		dayInWeek = positionInWeek + (positionInWeek / 2) // Mon, Tue, Thu, Fri, Sat
	} else {
		dayInWeek = positionInWeek // 6 days: Mon-Sat, Sunday rest
	}

	daysFromStart := weekNumber*7 + dayInWeek
	workoutDate := now.AddDate(0, 0, daysFromStart)
	// Return date without time (start of day)
	return time.Date(workoutDate.Year(), workoutDate.Month(), workoutDate.Day(), 0, 0, 0, 0, workoutDate.Location())
}

func (s *AIService) generateFullSchedule(baseWorkouts []models.Workout, workoutsPerWeek, totalWorkouts int) []models.Workout {
	var fullSchedule []models.Workout

	for i := 0; i < totalWorkouts; i++ {
		// Cycle through base workouts
		baseIndex := i % len(baseWorkouts)
		workout := baseWorkouts[baseIndex]

		// Create new workout with unique ID and date
		scheduledWorkout := models.Workout{
			WorkoutID:     primitive.NewObjectID(),
			Name:          fmt.Sprintf("%s - Week %d", workout.Name, (i/workoutsPerWeek)+1),
			Description:   workout.Description,
			Status:        "planned",
			ScheduledDate: s.calculateWorkoutDate(i, workoutsPerWeek),
			Exercises:     make([]models.Exercise, len(workout.Exercises)),
		}

		// Copy exercises with new IDs
		for j, exercise := range workout.Exercises {
			scheduledWorkout.Exercises[j] = models.Exercise{
				ExerciseID:  primitive.NewObjectID(),
				Name:        exercise.Name,
				MuscleGroup: exercise.MuscleGroup,
				Sets:        exercise.Sets,
				Reps:        exercise.Reps,
				RestSec:     exercise.RestSec,
				Notes:       exercise.Notes,
				Technique:   exercise.Technique,
			}
		}

		fullSchedule = append(fullSchedule, scheduledWorkout)
	}

	return fullSchedule
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

	// Get short plan for context
	currentShortPlan, err := s.MongoDBRepo.GetShortPlan(ctx, userID)
	if err != nil {
		return nil, NewServiceError(
			http.StatusNotFound,
			"No existing workout plan found",
			err,
		)
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

	// Calculate required workouts for regeneration
	workoutsPerWeek := profile.AvailableMinutes / 50
	if workoutsPerWeek < 2 {
		workoutsPerWeek = 2
	}
	if workoutsPerWeek > 6 {
		workoutsPerWeek = 6
	}

	// Prepare system message
	systemContent := fmt.Sprintf(`You are a fitness expert. You MUST follow user feedback exactly. Create EXACTLY %d workouts and respond with ONLY valid JSON.

CRITICAL: User feedback in the prompt is MANDATORY and must be implemented precisely. Do not ignore any user requirements.

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
}

IMPORTANT: Create EXACTLY %d different workouts. Follow ALL user requirements from the prompt.`, workoutsPerWeek, workoutsPerWeek)

	// Prepare prompt with short plan and comments
	if currentShortPlan == nil {
		currentShortPlan = &models.ShortWorkoutPlan{
			UserID:          userID,
			Title:           "Basic Workout Plan",
			BaseWorkouts:    s.createFallbackWorkouts(workoutsPerWeek),
			Timeframe:       profile.Timeframe,
			WorkoutsPerWeek: workoutsPerWeek,
			Status:          true,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}
	}
	userPrompt := s.formatRegeneratePrompt(profile, currentShortPlan, userComments, workoutsPerWeek)

	// Prepare messages for OpenRouter
	messages := []OpenRouterMessage{
		{Role: "system", Content: systemContent},
		{Role: "user", Content: userPrompt},
	}

	// Call AI
	fmt.Printf("Starting AI request...\n")
	content, err := s.Client.CreateChatCompletion(ctx, messages, true)
	if err != nil {
		fmt.Printf("AI REQUEST FAILED when regenerating plan: %v\n", err)
		return nil, NewServiceError(
			http.StatusInternalServerError,
			"AI request failed",
			err,
		)
	}
	fmt.Printf("AI request completed\n")

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
		res, err := s.fixJSONWithAI(ctx, cleanContent, err.Error())
		fmt.Printf("%s", res)
		if err != nil {
			return nil, NewServiceError(
				http.StatusInternalServerError,
				"Failed to parse AI response",
				fmt.Errorf("JSON parse error: %v, cleaned content: %s", err, cleanContent),
			)
		}
		if err := json.Unmarshal([]byte(res), &generatedData); err != nil {
			return nil, NewServiceError(
				http.StatusInternalServerError,
				"Failed to parse AI response",
				fmt.Errorf("JSON parse error: %v, cleaned content: %s", err, cleanContent),
			)
		}
	}

	// Fill missing fields with default values
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

	// Update short plan
	now := time.Now()
	currentShortPlan.Title = generatedData.Title
	currentShortPlan.BaseWorkouts = generatedData.Workouts
	currentShortPlan.UpdatedAt = now

	// Save updated short plan
	if err := s.MongoDBRepo.SaveShortPlan(ctx, currentShortPlan); err != nil {
		fmt.Printf("Failed to save updated short plan: %v\n", err)
	}

	// Create full plan
	updatedPlan := &models.WorkoutPlan{
		UserID:    userID,
		Title:     generatedData.Title,
		Workouts:  generatedData.Workouts,
		Status:    true,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Generate full schedule for timeframe
	totalWeeks := s.getWeeksFromTimeframe(profile.Timeframe)
	totalWorkouts := workoutsPerWeek * totalWeeks
	fullSchedule := s.generateFullSchedule(generatedData.Workouts, workoutsPerWeek, totalWorkouts)

	// Replace workouts with full schedule
	updatedPlan.Workouts = fullSchedule

	// Save updated plan
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

func (s *AIService) GetRating(ctx context.Context) ([]models.UserRating, error) {
	return s.MongoDBRepo.GetRating(ctx)
}

func (s *AIService) GenerateMotivationalMessage(ctx context.Context) (string, error) {
	userID, err := s.GetUserIDFromContext(ctx)
	if err != nil {
		return "", err
	}

	progress, err := s.MongoDBRepo.GetUserProgress(ctx, userID)
	if err != nil {
		return "Keep pushing forward! Every workout counts.", nil
	}

	if s.Client == nil {
		return "You're doing amazing! Keep up the great work!", nil
	}

	messages := []OpenRouterMessage{
		{Role: "system", Content: "Generate a short motivational fitness message. Be encouraging and specific."},
		{Role: "user", Content: fmt.Sprintf("User: %d workouts, %d consecutive days, %s level. Motivate them!", progress.TotalWorkouts, progress.ConsecutiveDays, progress.Level)},
	}

	response, err := s.Client.CreateChatCompletion(ctx, messages, false)
	if err != nil {
		fmt.Printf("ERROR: AI request failed in GenerateMotivationalMessage: %v\n", err)
		return "You're crushing it! Keep up the excellent work!", nil
	}

	return response, nil
}

func (s *AIService) formatRegeneratePrompt(profile *models.FitnessProfile, currentShortPlan *models.ShortWorkoutPlan, userComments string, requiredWorkouts int) string {
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

	sb.WriteString("\nCurrent Base Workouts:\n")
	fmt.Fprintf(&sb, "Title: %s\n", currentShortPlan.Title)
	for i, workout := range currentShortPlan.BaseWorkouts {
		fmt.Fprintf(&sb, "Workout %d: %s - %s\n", i+1, workout.Name, workout.Description)
		for j, exercise := range workout.Exercises {
			fmt.Fprintf(&sb, "  Exercise %d: %s (%s) - %d sets x %d reps\n",
				j+1, exercise.Name, exercise.MuscleGroup, exercise.Sets, exercise.Reps)
		}
	}

	sb.WriteString("\n\n=== CRITICAL USER REQUIREMENTS ===\n")
	sb.WriteString("MUST FOLLOW THESE COMMENTS EXACTLY:\n")
	sb.WriteString(userComments)
	sb.WriteString("\n=== END CRITICAL REQUIREMENTS ===\n\n")
	sb.WriteString("The above user comments are MANDATORY and must be implemented precisely.\n")

	// Add timeframe-specific guidance
	sb.WriteString("\n")
	sb.WriteString(s.getTimeframeGuidance(profile.Timeframe, profile.AvailableMinutes))
	sb.WriteString("\n")

	sb.WriteString(fmt.Sprintf("\n\nPlease update the workout plan with EXACTLY %d workouts based on the user's feedback while maintaining:", requiredWorkouts))
	sb.WriteString("\n1. Appropriate difficulty for their fitness level")
	sb.WriteString("\n2. Alignment with their fitness goals")
	sb.WriteString("\n3. Consideration of their health issues")
	sb.WriteString("\n4. Time constraints")
	sb.WriteString("\n5. Progressive overload principles")

	// Add beginner mode instructions if user is a beginner
	if profile.FitnessLevel == "beginner" {
		sb.WriteString("\n\nIMPORTANT: This user is a beginner. Please explain all exercises in very simple terms as if explaining to someone with no fitness experience. Use basic language, avoid technical jargon, and include extra safety tips. Provide detailed step-by-step instructions for each exercise.\n")
	}

	return sb.String()
}
