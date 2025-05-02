# API Endpoints Debug Scenarios

This document outlines all endpoints and their test scenarios for debugging purposes.

## Authentication Endpoints

### 1. User Registration - `POST /users`

#### Success Scenarios
- Valid registration with required fields:
  ```json
  {
    "username": "testuser",
    "email": "test@example.com",
    "password": "securepass123"
  }
  ```
  ```bash
  curl -X POST http://localhost:8080/users \
    -H "Content-Type: application/json" \
    -d '{
      "username": "testuser",
      "email": "test@example.com",
      "password": "securepass123"
    }'
  ```

- Valid registration with optional bio:
  ```json
  {
    "username": "testuser",
    "email": "test@example.com",
    "password": "securepass123",
    "bio": "Fitness enthusiast"
  }
  ```
  ```bash
  curl -X POST http://localhost:8080/users \
    -H "Content-Type: application/json" \
    -d '{
      "username": "testuser",
      "email": "test@example.com",
      "password": "securepass123",
      "bio": "Fitness enthusiast"
    }'
  ```

#### Error Scenarios
- Missing username
  ```bash
  curl -X POST http://localhost:8080/users \
    -H "Content-Type: application/json" \
    -d '{
      "email": "test@example.com",
      "password": "securepass123"
    }'
  ```
- Username too long (>50 characters)
  ```bash
  curl -X POST http://localhost:8080/users \
    -H "Content-Type: application/json" \
    -d '{
      "username": "thisIsAReallyLongUsernameThatExceedsFiftyCharactersAndShouldFail",
      "email": "test@example.com",
      "password": "securepass123"
    }'
  ```
- Missing email
- Invalid email format
- Missing password
- Invalid request payload format

### 2. Authentication - `POST /tokens/authentication`

#### Success Scenarios
- Valid credentials:
  ```json
  {
    "username": "testuser",
    "password": "securepass123"
  }
  ```
  ```bash
  curl -X POST http://localhost:8080/tokens/authentication \
    -H "Content-Type: application/json" \
    -d '{
      "username": "testuser",
      "password": "securepass123"
    }'
  ```

#### Error Scenarios
- Invalid username
  ```bash
  curl -X POST http://localhost:8080/tokens/authentication \
    -H "Content-Type: application/json" \
    -d '{
      "username": "nonexistentuser",
      "password": "securepass123"
    }'
  ```
- Invalid password
- Invalid request payload format
- Internal server errors

## Workout Endpoints

### 3. Get Workout - `GET /workouts/{id}`

#### Success Scenarios
- Valid workout ID with authenticated user
  ```bash
  curl -X GET http://localhost:8080/workouts/123 \
    -H "Authorization: Bearer <your-token>"
  ```

#### Error Scenarios
- Invalid workout ID format
  ```bash
  curl -X GET http://localhost:8080/workouts/invalid-id \
    -H "Authorization: Bearer <your-token>"
  ```
- Non-existent workout ID
- Unauthorized access (no authentication)
  ```bash
  curl -X GET http://localhost:8080/workouts/123
  ```
- Internal server errors

### 4. Create Workout - `POST /workouts`

#### Success Scenarios
- Create workout with valid data:
  ```json
  {
    "title": "Morning Run",
    "description": "5K run in the park",
    "duration_minutes": 30,
    "calories_burned": 300,
    "entries": [
      {
        "exercise": "Running",
        "duration": "30m",
        "distance": "5km"
      }
    ]
  }
  ```
  ```bash
  curl -X POST http://localhost:8080/workouts \
    -H "Authorization: Bearer <your-token>" \
    -H "Content-Type: application/json" \
    -d '{
      "title": "Morning Run",
      "description": "5K run in the park",
      "duration_minutes": 30,
      "calories_burned": 300,
      "entries": [
        {
          "exercise": "Running",
          "duration": "30m",
          "distance": "5km"
        }
      ]
    }'
  ```

#### Error Scenarios
- Invalid request payload
- Unauthorized access (no authentication)
  ```bash
  curl -X POST http://localhost:8080/workouts \
    -H "Content-Type: application/json" \
    -d '{
      "title": "Morning Run",
      "description": "5K run in the park",
      "duration_minutes": 30,
      "calories_burned": 300
    }'
  ```
- Internal server errors during creation

### 5. Update Workout - `PUT /workouts/{id}`

#### Success Scenarios
- Full update:
  ```json
  {
    "title": "Updated Morning Run",
    "description": "Extended 7K run",
    "duration_minutes": 45,
    "calories_burned": 450,
    "entries": [
      {
        "exercise": "Running",
        "duration": "45m",
        "distance": "7km"
      }
    ]
  }
  ```
  ```bash
  curl -X PUT http://localhost:8080/workouts/123 \
    -H "Authorization: Bearer <your-token>" \
    -H "Content-Type: application/json" \
    -d '{
      "title": "Updated Morning Run",
      "description": "Extended 7K run",
      "duration_minutes": 45,
      "calories_burned": 450,
      "entries": [
        {
          "exercise": "Running",
          "duration": "45m",
          "distance": "7km"
        }
      ]
    }'
  ```

- Partial update:
  ```json
  {
    "title": "Updated Morning Run",
    "duration_minutes": 45
  }
  ```
  ```bash
  curl -X PUT http://localhost:8080/workouts/123 \
    -H "Authorization: Bearer <your-token>" \
    -H "Content-Type: application/json" \
    -d '{
      "title": "Updated Morning Run",
      "duration_minutes": 45
    }'
  ```

#### Error Scenarios
- Invalid workout ID
- Non-existent workout
- Unauthorized access (no authentication)
  ```bash
  curl -X PUT http://localhost:8080/workouts/123 \
    -H "Content-Type: application/json" \
    -d '{
      "title": "Updated Morning Run",
      "duration_minutes": 45
    }'
  ```
- Forbidden access (not the workout owner)
- Invalid request payload
- Internal server errors

### 6. Delete Workout - `DELETE /workouts/{id}`

#### Success Scenarios
- Delete workout by owner
  ```bash
  curl -X DELETE http://localhost:8080/workouts/123 \
    -H "Authorization: Bearer <your-token>"
  ```

#### Error Scenarios
- Invalid workout ID
  ```bash
  curl -X DELETE http://localhost:8080/workouts/invalid-id \
    -H "Authorization: Bearer <your-token>"
  ```
- Non-existent workout
- Unauthorized access (no authentication)
  ```bash
  curl -X DELETE http://localhost:8080/workouts/123
  ```
- Forbidden access (not the workout owner)
- Internal server errors

### 7. Health Check - `GET /health`

#### Success Scenarios
- Server is running
  ```bash
  curl -X GET http://localhost:8080/health
  ```

#### Error Scenarios
- Server is not responding
  - Expected response: No response/timeout

## General Testing Considerations

### Authentication Testing
For each protected endpoint, test:
1. No authentication token
2. Invalid authentication token
3. Expired authentication token
4. Valid authentication token

### Request Validation
For each endpoint that accepts data, test:
1. Missing required fields
2. Invalid field formats
3. Extra/unknown fields
4. Invalid JSON format

### Business Logic
Test scenarios for:
1. Resource ownership
2. Resource existence
3. Data integrity

### Error Handling
Verify proper handling of:
1. Database errors
2. Network errors
3. Validation errors
4. Authorization errors

## Testing Tools
1. cURL or Postman for HTTP requests
2. Database client to verify data changes
3. Logging output for server-side issues
4. HTTP response codes and error messages 