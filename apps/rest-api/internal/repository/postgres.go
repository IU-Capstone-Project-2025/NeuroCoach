package repository

import (
	"context"
	"errors"
	"fmt"

	"rest-api/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(connString string) (*PostgresRepository, error) {
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	return &PostgresRepository{pool: pool}, nil
}

// User operations
func (r *PostgresRepository) CreateUser(ctx context.Context, email, passwordHash string) (int, error) {
	var id int
	err := r.pool.QueryRow(ctx,
		`INSERT INTO users (email, password_hash) 
		VALUES ($1, $2) 
		RETURNING id`,
		email, passwordHash).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("error creating user: %w", err)
	}
	return id, nil
}

func (r *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.pool.QueryRow(ctx,
		`SELECT id, email, password_hash, created_at 
		FROM users 
		WHERE email = $1`,
		email).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("error getting user: %w", err)
	}
	return &user, nil
}

// Fitness profile operations
func (r *PostgresRepository) SaveFitnessProfile(ctx context.Context, userID int, profile *models.FitnessProfile) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	// Upsert fitness profile
	_, err = tx.Exec(ctx,
		`INSERT INTO fitness_profiles 
			(user_id, height_cm, weight_kg, age, fitness_goal, timeframe, fitness_level, weekly_time_minutes)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (user_id) DO UPDATE SET
			height_cm = EXCLUDED.height_cm,
			weight_kg = EXCLUDED.weight_kg,
			age = EXCLUDED.age,
			fitness_goal = EXCLUDED.fitness_goal,
			timeframe = EXCLUDED.timeframe,
			fitness_level = EXCLUDED.fitness_level,
			weekly_time_minutes = EXCLUDED.weekly_time_minutes,
			updated_at = NOW()`,
		userID, profile.Height, profile.Weight, profile.Age,
		profile.Goal, profile.Timeframe, profile.FitnessLevel, profile.AvailableMinutes)

	if err != nil {
		return fmt.Errorf("error saving fitness profile: %w", err)
	}

	// Update health issues
	if _, err := tx.Exec(ctx, "DELETE FROM user_health_issues WHERE user_id = $1", userID); err != nil {
		return fmt.Errorf("error clearing health issues: %w", err)
	}

	for _, issue := range profile.HealthIssues {
		var issueID int
		err := tx.QueryRow(ctx,
			`INSERT INTO health_issues (name) 
			VALUES ($1) 
			ON CONFLICT (name) DO UPDATE SET name = EXCLUDED.name 
			RETURNING id`,
			issue).Scan(&issueID)

		if err != nil {
			return fmt.Errorf("error inserting health issue: %w", err)
		}

		if _, err := tx.Exec(ctx,
			"INSERT INTO user_health_issues (user_id, issue_id) VALUES ($1, $2)",
			userID, issueID); err != nil {
			return fmt.Errorf("error linking health issue: %w", err)
		}
	}

	return tx.Commit(ctx)
}

func (r *PostgresRepository) GetFitnessProfile(ctx context.Context, userID int) (*models.FitnessProfile, error) {
	var profile models.FitnessProfile
	err := r.pool.QueryRow(ctx,
		`SELECT height_cm, weight_kg, age, fitness_goal, timeframe, 
				fitness_level, weekly_time_minutes, updated_at
		FROM fitness_profiles 
		WHERE user_id = $1`,
		userID).Scan(
		&profile.Height, &profile.Weight, &profile.Age,
		&profile.Goal, &profile.Timeframe, &profile.FitnessLevel,
		&profile.AvailableMinutes, &profile.UpdatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("error getting fitness profile: %w", err)
	}

	// Get health issues
	rows, err := r.pool.Query(ctx,
		`SELECT hi.name 
		FROM health_issues hi
		JOIN user_health_issues uhi ON hi.id = uhi.issue_id
		WHERE uhi.user_id = $1`,
		userID)

	if err != nil {
		return nil, fmt.Errorf("error getting health issues: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var issue string
		if err := rows.Scan(&issue); err != nil {
			return nil, fmt.Errorf("error scanning health issue: %w", err)
		}
		profile.HealthIssues = append(profile.HealthIssues, issue)
	}

	return &profile, nil
}

// Workout plan operations
// func (r *PostgresRepository) SaveWorkoutPlan(ctx context.Context, userID int, plan *models.WorkoutPlan) error {
// 	_, err := r.pool.Exec(ctx,
// 		`INSERT INTO workout_plans (user_id, content)
// 		VALUES ($1, $2)
// 		ON CONFLICT (user_id) DO UPDATE SET
// 			content = EXCLUDED.content,
// 			updated_at = NOW()`,
// 		userID, plan.Content)

// 	return err
// }

// func (r *PostgresRepository) GetWorkoutPlan(ctx context.Context, userID int) (*models.WorkoutPlan, error) {
// 	var plan models.WorkoutPlan
// 	err := r.pool.QueryRow(ctx,
// 		`SELECT id, content, created_at, updated_at
// 		FROM workout_plans
// 		WHERE user_id = $1`,
// 		userID).Scan(&plan.PlanID, &plan.Content, &plan.CreatedAt, &plan.UpdatedAt)

// 	if err != nil {
// 		if errors.Is(err, pgx.ErrNoRows) {
// 			return nil, ErrNotFound
// 		}
// 		return nil, fmt.Errorf("error getting workout plan: %w", err)
// 	}
// 	return &plan, nil
// }

// Chat operations
// func (r *PostgresRepository) SaveChatMessage(ctx context.Context, message *models.ChatMessage) error {
// 	_, err := r.pool.Exec(ctx,
// 		`INSERT INTO chat_messages
// 			(user_id, message, response, is_user)
// 		VALUES ($1, $2, $3, $4)`,
// 		message.UserID, message.Message, message.Response, message.IsUser)

// 	return err
// }

// func (r *PostgresRepository) GetChatHistory(ctx context.Context, userID int, limit int) ([]models.ChatMessage, error) {
// 	rows, err := r.pool.Query(ctx,
// 		`SELECT id, message, response, is_user, created_at
// 		FROM chat_messages
// 		WHERE user_id = $1
// 		ORDER BY created_at DESC
// 		LIMIT $2`,
// 		userID, limit)

// 	if err != nil {
// 		return nil, fmt.Errorf("error getting chat history: %w", err)
// 	}
// 	defer rows.Close()

// 	var messages []models.ChatMessage
// 	for rows.Next() {
// 		var msg models.ChatMessage
// 		if err := rows.Scan(
// 			&msg.ID, &msg.Message, &msg.Response,
// 			&msg.IsUser, &msg.CreatedAt); err != nil {
// 			return nil, fmt.Errorf("error scanning chat message: %w", err)
// 		}
// 		messages = append(messages, msg)
// 	}

// 	// Reverse the order to return oldest first
// 	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
// 		messages[i], messages[j] = messages[j], messages[i]
// 	}

// 	return messages, nil
// }

// Health check
func (r *PostgresRepository) Ping(ctx context.Context) error {
	return r.pool.Ping(ctx)
}

func (r *PostgresRepository) Close() error {
	r.pool.Close()
	return nil
}

var ErrNotFound = errors.New("record not found")
