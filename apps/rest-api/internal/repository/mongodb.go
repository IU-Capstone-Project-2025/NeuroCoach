package repository

import (
	"context"
	"fmt"
	"time"

	"rest-api/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBRepository struct {
	chatCollection       *mongo.Collection
	workoutCollection    *mongo.Collection
	completionCollection *mongo.Collection
	progressCollection   *mongo.Collection
}

func NewMongoDBRepository(uri, dbName string) (MongoDBRep, error) {
	fmt.Printf("Connecting to MongoDB with URI: %s, DB: %s\n", uri, dbName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Test the connection
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}
	fmt.Println("MongoDB connection successful")

	db := client.Database(dbName)
	return &MongoDBRepository{
		chatCollection:       db.Collection("chat_messages"),
		workoutCollection:    db.Collection("workout_plans"),
		completionCollection: db.Collection("workout_completions"),
		progressCollection:   db.Collection("user_progress"),
	}, nil
}

func (m *MongoDBRepository) SaveChatMessage(ctx context.Context, msg *models.ChatMessage) error {
	_, err := m.chatCollection.InsertOne(ctx, bson.M{
		"user_id":    msg.UserID,
		"message":    msg.Message,
		"response":   msg.Response,
		"is_user":    msg.IsUser,
		"created_at": time.Now(),
	})
	return err
}

func (m *MongoDBRepository) GetChatHistory(ctx context.Context, userID int) ([]models.ChatMessage, error) {
	filter := bson.M{"user_id": userID}
	cursor, err := m.chatCollection.Find(ctx, filter, options.Find().SetSort(bson.M{"created_at": 1}))
	if err != nil {
		return nil, err
	}

	var messages []models.ChatMessage
	if err := cursor.All(ctx, &messages); err != nil {
		return nil, err
	}
	return messages, nil
}

func (m *MongoDBRepository) SaveWorkoutPlan(ctx context.Context, plan *models.WorkoutPlan) error {
	_, err := m.workoutCollection.UpdateOne(
		ctx,
		bson.M{"user_id": plan.UserID},
		bson.M{"$set": plan},
		options.Update().SetUpsert(true),
	)
	return err
}

func (m *MongoDBRepository) GetWorkoutPlan(ctx context.Context, userID int) (*models.WorkoutPlan, error) {
	var plan models.WorkoutPlan
	err := m.workoutCollection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&plan)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &plan, err
}

func (m *MongoDBRepository) GetWorkoutByID(ctx context.Context, userID int, workoutID string) (*models.Workout, error) {
	var plan models.WorkoutPlan
	err := m.workoutCollection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&plan)
	if err != nil {
		return nil, err
	}

	for _, workout := range plan.Workouts {
		if workout.WorkoutID.Hex() == workoutID {
			return &workout, nil
		}
	}
	return nil, mongo.ErrNoDocuments
}

func (m *MongoDBRepository) CompleteWorkout(ctx context.Context, userID int, workoutID string) error {
	objID, err := primitive.ObjectIDFromHex(workoutID)
	if err != nil {
		return err
	}

	// Get workout name
	workoutName, err := m.getWorkoutName(ctx, userID, workoutID)
	if err != nil {
		return err
	}

	// Save completion record
	completion := &models.WorkoutCompletion{
		UserID:      userID,
		WorkoutID:   objID,
		CompletedAt: time.Now(),
	}
	_, err = m.completionCollection.InsertOne(ctx, completion)
	if err != nil {
		return err
	}

	// Update user progress
	return m.updateUserProgress(ctx, userID, workoutName)
}

func (m *MongoDBRepository) GetUserProgress(ctx context.Context, userID int) (*models.UserProgress, error) {
	var progress models.UserProgress
	err := m.progressCollection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&progress)
	if err == mongo.ErrNoDocuments {
		return &models.UserProgress{
			UserID:           userID,
			TotalWorkouts:    0,
			ConsecutiveDays:  0,
			Level:            "Beginner",
			CompletedWorkouts: []string{},
		}, nil
	}
	return &progress, err
}

func (m *MongoDBRepository) getWorkoutName(ctx context.Context, userID int, workoutID string) (string, error) {
	workout, err := m.GetWorkoutByID(ctx, userID, workoutID)
	if err != nil {
		return "", err
	}
	return workout.Name, nil
}

func (m *MongoDBRepository) updateUserProgress(ctx context.Context, userID int, workoutName string) error {
	// Count total workouts
	totalWorkouts, err := m.completionCollection.CountDocuments(ctx, bson.M{"user_id": userID})
	if err != nil {
		return err
	}

	// Calculate consecutive days
	consecutiveDays := m.calculateConsecutiveDays(ctx, userID)

	// Determine level
	level := m.calculateLevel(int(totalWorkouts))

	now := time.Now()
	
	// Get existing progress to append workout name
	var existingProgress models.UserProgress
	completedWorkouts := []string{}
	
	err = m.progressCollection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&existingProgress)
	if err == nil {
		completedWorkouts = existingProgress.CompletedWorkouts
	}
	
	// Add new workout name if not already in list
	if workoutName != "" {
		found := false
		for _, name := range completedWorkouts {
			if name == workoutName {
				found = true
				break
			}
		}
		if !found {
			completedWorkouts = append(completedWorkouts, workoutName)
		}
	}
	
	progress := &models.UserProgress{
		UserID:           userID,
		TotalWorkouts:    int(totalWorkouts),
		ConsecutiveDays:  consecutiveDays,
		Level:            level,
		CompletedWorkouts: completedWorkouts,
		LastWorkoutDate:  now,
		UpdatedAt:        now,
	}

	_, err = m.progressCollection.UpdateOne(
		ctx,
		bson.M{"user_id": userID},
		bson.M{"$set": progress},
		options.Update().SetUpsert(true),
	)
	return err
}

func (m *MongoDBRepository) calculateConsecutiveDays(ctx context.Context, userID int) int {
	// Get completions sorted by date descending
	cursor, err := m.completionCollection.Find(
		ctx,
		bson.M{"user_id": userID},
		options.Find().SetSort(bson.M{"completed_at": -1}),
	)
	if err != nil {
		return 0
	}
	defer cursor.Close(ctx)

	var completions []models.WorkoutCompletion
	if err := cursor.All(ctx, &completions); err != nil {
		return 0
	}

	if len(completions) == 0 {
		return 0
	}

	consecutiveDays := 1
	lastDate := completions[0].CompletedAt.Truncate(24 * time.Hour)

	for i := 1; i < len(completions); i++ {
		currentDate := completions[i].CompletedAt.Truncate(24 * time.Hour)
		if lastDate.Sub(currentDate) == 24*time.Hour {
			consecutiveDays++
			lastDate = currentDate
		} else {
			break
		}
	}

	return consecutiveDays
}

func (m *MongoDBRepository) calculateLevel(totalWorkouts int) string {
	if totalWorkouts >= 50 {
		return "Expert"
	} else if totalWorkouts >= 20 {
		return "Advanced"
	} else if totalWorkouts >= 5 {
		return "Intermediate"
	}
	return "Beginner"
}
