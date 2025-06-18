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
  "plan_id": 1,
  "content": "Your personalized workout plan...",
  "created_at": "2024-03-20T10:00:00Z",
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
| id | integer | Message ID |
| message | string | Message content |
| response | string | AI response |
| is_user | boolean | Whether from user |
| created_at | string | Timestamp |

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

