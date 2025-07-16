# NeuroCoach REST API Documentation

## Overview
The NeuroCoach API provides endpoints for user management, fitness profiles, AI coaching, and workout planning.

## Base URLs
- Development: `http://localhost:8080`
- Production: `https://api.neurocoach.com`

## Authentication
The API uses JWT Bearer token authentication with access/refresh token pattern:
- **Access Token**: Short-lived (15 minutes) for API requests
- **Refresh Token**: Long-lived (7 days) to get new access tokens

Include the access token in the Authorization header:
```
Authorization: Bearer <your_access_token>
```

## Frontend Token Management

### Token Storage
```javascript
// Store tokens after login/register
localStorage.setItem('access_token', response.access_token);
localStorage.setItem('refresh_token', response.refresh_token);
localStorage.setItem('token_expires_at', Date.now() + (response.expires_in * 1000));
```

### Token Validation (Proactive)
```javascript
function isTokenValid() {
  const expiresAt = localStorage.getItem('token_expires_at');
  const buffer = 60000; // 1 minute buffer
  return Date.now() < (parseInt(expiresAt) - buffer);
}

async function getValidToken() {
  if (isTokenValid()) {
    return localStorage.getItem('access_token');
  }
  return await refreshToken();
}
```

### Token Refresh
```javascript
async function refreshToken() {
  const refreshToken = localStorage.getItem('refresh_token');
  
  const response = await fetch('/refresh', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ refresh_token: refreshToken })
  });
  
  if (response.ok) {
    const data = await response.json();
    localStorage.setItem('access_token', data.access_token);
    localStorage.setItem('refresh_token', data.refresh_token);
    localStorage.setItem('token_expires_at', Date.now() + (data.expires_in * 1000));
    return data.access_token;
  }
  
  // Refresh failed - redirect to login
  window.location.href = '/login';
  return null;
}
```

### API Call with Auto-Refresh
```javascript
async function apiCall(url, options = {}) {
  const token = await getValidToken();
  
  const response = await fetch(url, {
    ...options,
    headers: {
      'Authorization': `Bearer ${token}`,
      ...options.headers
    }
  });
  
  // Handle 401 (token expired)
  if (response.status === 401) {
    const newToken = await refreshToken();
    if (newToken) {
      return fetch(url, {
        ...options,
        headers: {
          'Authorization': `Bearer ${newToken}`,
          ...options.headers
        }
      });
    }
  }
  
  return response;
}
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
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "token_type": "Bearer",
  "expires_in": 900,
  "email": "user@example.com"
}
```

#### Login
Authenticate user and get JWT tokens.

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
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "token_type": "Bearer",
  "expires_in": 900,
  "email": "user@example.com"
}
```

#### Refresh Token
Get new access token using refresh token.

```http
POST /refresh
```

**Request Body**
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Response** (200 OK)
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "token_type": "Bearer",
  "expires_in": 900,
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

#### Get Rating
Retrieve user rating leaderboard sorted by score.

```http
GET /api/rating
Authorization: Bearer <token>
```

**Response** (200 OK)
```json
[
  {
    "user_id": 1,
    "total_workouts": 25,
    "max_consecutive": 7,
    "score": 32
  },
  {
    "user_id": 2,
    "total_workouts": 15,
    "max_consecutive": 5,
    "score": 20
  }
]
```

### Exercise Media Management

#### Save Exercise Media
Save an exercise image with name and description.

```http
POST /api/media
Authorization: Bearer <token>
```

**Request Body**
```json
{
  "name": "Squat Exercise",
  "image_url": "https://example.com/images/squat.jpg",
  "description": "Starting position for squat exercise"
}
```

**Response** (201 Created)
```json
{
  "message": "Exercise media saved successfully"
}
```

#### Get All Exercise Media
Retrieve all exercise media (images with descriptions).

```http
GET /api/media
Authorization: Bearer <token>
```

**Response** (200 OK)
```json
[
  {
    "id": "media_id_1",
    "name": "Squat Exercise",
    "image_url": "https://example.com/images/squat.jpg",
    "description": "Starting position for squat exercise",
    "created_at": "2024-03-20T10:00:00Z"
  },
  {
    "id": "media_id_2",
    "name": "Push-up Exercise",
    "image_url": "https://example.com/images/pushup.jpg",
    "description": "Proper push-up form demonstration",
    "created_at": "2024-03-20T10:05:00Z"
  }
]
```

#### Delete Exercise Media
Delete a specific exercise media item.

```http
DELETE /api/media/{media_id}
Authorization: Bearer <token>
```

**Response** (200 OK)
```json
{
  "message": "Exercise media deleted successfully"
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

### Exercise Media
| Field | Type | Description |
|-------|------|-------------|
| id | string | Media ID |
| name | string | Exercise name |
| image_url | string | URL to exercise image |
| description | string | Description of the image |
| created_at | string | Creation timestamp |

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

### User Rating
| Field | Type | Description |
|-------|------|-------------|
| user_id | integer | User ID |
| total_workouts | integer | Total completed workouts |
| max_consecutive | integer | Max consecutive workout days |
| score | integer | Total score (total_workouts + max_consecutive) |

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
