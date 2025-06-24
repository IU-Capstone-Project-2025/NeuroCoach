package repository

import (
	"context"
	"fmt"
	"time"

	"rest-api/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBRepository struct {
	chatCollection    *mongo.Collection
	workoutCollection *mongo.Collection
}

func NewMongoDBRepository(uri, dbName string) (MongoDBRep, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	db := client.Database(dbName)
	return &MongoDBRepository{
		chatCollection:    db.Collection("chat_messages"),
		workoutCollection: db.Collection("workout_plans"),
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
