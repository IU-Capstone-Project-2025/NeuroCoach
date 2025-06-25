# NeuroCoach REST API Documentation

## Overview
The NeuroCoach API provides endpoints for user management, fitness profiles, AI coaching, and workout planning.

## Base URLs
- Development: `http://localhost:8080`
- Production: `https://api.neurocoach.com`

## Authentication
The API uses JWT Bearer token authentication. Include the token in the Authorization header:
```
Authorization: Bearer <your_jwt_token>
```

## Endpoints

### Health Check
Check if the API is running and database is connected.

```http
GET /health
```

**Response**
```json
{
  "status": "healthy",
  "version": "1.0"
}
```

### Authentication

#### Register User
Create a new user account.

```http
POST /register
```

**Request Body**
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response** (201 Created)
```json
{
  "token": "jwt_token_here",
  "email": "user@example.com"
}
```

#### Login
Authenticate user and get JWT token.

```http
POST /login
```

**Request Body**
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response** (200 OK)
```json
{
  "token": "jwt_token_here",
  "email": "user@example.com"
}
```

### Profile Management

#### Save Profile
Save or update user's fitness profile.

```http
POST /api/profile
Authorization: Bearer <token>
```

**Request Body**
```json
{
  "height": 175.5,
  "weight": 70.0,
  "age": 30,
  "goal": "weight_loss",
  "health_issues": ["knee_pain"],
  "timeframe": "3months",
  "fitness_level": "intermediate",
  "available_minutes": 60
}
```

**Response** (200 OK)
```json
{
  "message": "Profile saved successfully"
}
```

#### Get Profile
Retrieve user's fitness profile.

```http
GET /api/profile
Authorization: Bearer <token>
```

**Response** (200 OK)
```json
{
  "height": 175.5,
  "weight": 70.0,
  "age": 30,
  "goal": "weight_loss",
  "health_issues": ["knee_pain"],
  "timeframe": "3months",
  "fitness_level": "intermediate",
  "available_minutes": 60,
  "updated_at": "2024-03-20T10:00:00Z"
}
```

### AI Coach Chat

#### Send Message
Send a message to the AI coach.

```http
POST /api/chat
Authorization: Bearer <token>
```

**Request Body**
```json
{
  "message": "What's a good warm-up routine?"
}
```

**Response** (200 OK)
```json
{
  "response": "Here's a recommended warm-up routine..."
}
```

#### Get Chat History
Retrieve user's chat history with the AI coach.

```http
GET /api/chat/history
Authorization: Bearer <token>
```

**Response** (200 OK)
```json
{
  "messages": [
    {
      "id": 1,
      "message": "What's a good warm-up routine?",
      "response": "Here's a recommended warm-up routine...",
      "is_user": true,
      "created_at": "2024-03-20T10:00:00Z"
    }
  ]
}
```

### Workout Planning

#### Generate Plan
Generate a new workout plan based on user's profile.

```http
POST /api/generate-plan
Authorization: Bearer <token>
```

**Request Body**
```json
{
  "regenerate": false
}
```

**Response** (200 OK)
```json
{
  "plan": "Your personalized workout plan..."
}
```

#### Get Workout Plan
Retrieve user's current workout plan.

```http
GET /api/workout-plan
Authorization: Bearer <token>
```

**Response** (200 OK)
```json
{
  "id": "workout_plan_id",
  "user_id": 1,
  "title": "Personalized Fitness Plan",
  "workouts": [
    {
      "workout_id": "workout_id",
      "name": "Upper Body Strength",
      "description": "Focus on chest, back, and arms",
      "status": "planned",
      "exercises": [
        {
          "exercise_id": "exercise_id",
          "name": "Push-ups",
          "muscle_group": "chest",
          "sets": 3,
          "reps": 12,
          "rest_sec": 60,
          "notes": "Keep core tight",
          "technique": "Slow and controlled movement"
        }
      ]
    }
  ],
  "status": true,
  "created_at": "2024-03-20T10:00:00Z",
  "updated_at": "2024-03-20T10:00:00Z"
}
```

#### Regenerate Workout Plan
Regenerate workout plan based on user feedback.

```http
POST /api/regenerate-plan
Authorization: Bearer <token>
```

**Request Body**
```json
{
  "comments": "Make it more challenging and add more cardio"
}
```

**Response** (200 OK)
```json
{
  "id": "workout_plan_id",
  "title": "Updated Personalized Fitness Plan",
  "workouts": [...],
  "updated_at": "2024-03-20T11:00:00Z"
}
```

### Progress Tracking

#### Complete Workout
Mark a workout as completed and update user progress.

```http
POST /api/complete-workout
Authorization: Bearer <token>
```

**Request Body**
```json
{
  "workout_id": "workout_object_id_here"
}
```

**Response** (200 OK)
```json
{
  "message": "Workout completed successfully"
}
```

#### Get User Progress
Retrieve user's workout progress and statistics.

```http
GET /api/progress
Authorization: Bearer <token>
```

**Response** (200 OK)
```json
{
  "id": "progress_id",
  "user_id": 1,
  "total_workouts": 15,
  "consecutive_days": 5,
  "level": "Intermediate",
  "completed_workouts": ["Upper Body Strength", "Core Workout", "Leg Day"],
  "last_workout_date": "2024-03-20T10:00:00Z",
  "updated_at": "2024-03-20T10:00:00Z"
}
```

## Data Models

### Fitness Profile
| Field | Type | Description | Validation |
|-------|------|-------------|------------|
| height | number | User's height | > 0 |
| weight | number | User's weight | > 0 |
| age | integer | User's age | 13-120 |
| goal | string | Fitness goal | weight_loss, muscle_gain, endurance, flexibility, general_fitness |
| health_issues | string[] | List of health issues | - |
| timeframe | string | Goal timeframe | 1month, 3months, 6months, 1year |
| fitness_level | string | Current fitness level | beginner, intermediate, advanced |
| available_minutes | integer | Minutes per workout | 30-1000 |

### Chat Message
| Field | Type | Description |
|-------|------|-------------|
| id | string | Message ID |
| message | string | Message content |
| response | string | AI response |
| is_user | boolean | Whether from user |
| created_at | string | Timestamp |

### Workout Plan
| Field | Type | Description |
|-------|------|-------------|
| id | string | Plan ID |
| user_id | integer | User ID |
| title | string | Plan title |
| workouts | Workout[] | Array of workouts |
| status | boolean | Plan active status |
| created_at | string | Creation timestamp |
| updated_at | string | Update timestamp |

### Workout
| Field | Type | Description |
|-------|------|-------------|
| workout_id | string | Workout ID |
| name | string | Workout name |
| description | string | Workout description |
| status | string | Workout status |
| exercises | Exercise[] | Array of exercises |

### Exercise
| Field | Type | Description |
|-------|------|-------------|
| exercise_id | string | Exercise ID |
| name | string | Exercise name |
| muscle_group | string | Target muscle group |
| sets | integer | Number of sets |
| reps | integer | Repetitions per set |
| rest_sec | integer | Rest time in seconds |
| notes | string | Additional notes |
| technique | string | Technique instructions |

### User Progress
| Field | Type | Description |
|-------|------|-------------|
| id | string | Progress ID |
| user_id | integer | User ID |
| total_workouts | integer | Total completed workouts |
| consecutive_days | integer | Consecutive workout days |
| level | string | User fitness level |
| completed_workouts | string[] | Names of completed workouts |
| last_workout_date | string | Last workout timestamp |
| updated_at | string | Update timestamp |

### Level System
- **Beginner**: 0-4 completed workouts
- **Intermediate**: 5-19 completed workouts
- **Advanced**: 20-49 completed workouts
- **Expert**: 50+ completed workouts

### Error Response
```json
{
  "error": "Error Type",
  "message": "Detailed error message"
}
```

## Error Codes
- 400: Bad Request - Invalid input data
- 401: Unauthorized - Invalid or missing authentication
- 404: Not Found - Resource not found
- 503: Service Unavailable - Service health check failed
