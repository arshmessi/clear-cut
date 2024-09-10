# Clear-Cut: A Payment Splitting Application

## Overview

Clear-Cut is a payment splitting application that allows users to manage expenses and group them accordingly. Users can register, log in, manage groups, and record expenses.

## Project Structure

- **cmd/main.go**: Entry point of the application.
- **internal/auth**: Contains JWT-related functions.
- **internal/handlers**: Handles HTTP requests.
- **internal/models**: Defines data models.
- **internal/storage**: Manages database interactions and schema.
- **internal/services**: Contains in-memory data services (for testing).

## Running the Application

1. **Initialize the Database**:
   Before running the application, ensure that the database is initialized. You can do this by running the application, which will automatically set up the database and tables.

2. **Run the Application**:

   ```bash
   go run cmd/main.go
   ```

   The application will start and listen on port `8080` by default.

## API Endpoints

### User Registration

**Endpoint**: `POST /register`

**Description**: Registers a new user.

**Request**:

```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com", "password": "password123"}'
```

**Response**:

- `201 Created` on successful registration.

### User Login

**Endpoint**: `POST /login`

**Description**: Logs in a user and returns a JWT token.

**Request**:

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email": "john@example.com", "password": "password123"}'
```

**Response**:

- `200 OK` on successful login with a `Set-Cookie` header containing the JWT token.

### User Profile

**Endpoint**: `GET /profile`

**Description**: Retrieves the profile of the logged-in user.

**Request**:

```bash
curl -X GET http://localhost:8080/profile \
  -H "Cookie: token=your_jwt_token"
```

**Response**:

- `200 OK` with user profile details.

### Create Group

**Endpoint**: `POST /groups`

**Description**: Creates a new group.

**Request**:

```bash
curl -X POST http://localhost:8080/groups \
  -H "Content-Type: application/json" \
  -H "Cookie: token=your_jwt_token" \
  -d '{"name": "Family", "description": "Family expenses"}'
```

**Response**:

- `201 Created` with group details.

### Get Groups

**Endpoint**: `GET /groups`

**Description**: Retrieves all groups for the authenticated user.

**Request**:

```bash
curl -X GET http://localhost:8080/groups \
  -H "Cookie: token=your_jwt_token"
```

**Response**:

- `200 OK` with a list of groups.

### Add Expense

**Endpoint**: `POST /expenses`

**Description**: Adds a new expense.

**Request**:

```bash
curl -X POST http://localhost:8080/expenses \
  -H "Content-Type: application/json" \
  -H "Cookie: token=your_jwt_token" \
  -d '{"description": "Dinner", "amount": 50.00, "paid_by": "john@example.com", "group_id": 1}'
```

**Response**:

- `201 Created` with expense details.

### Get Expenses

**Endpoint**: `GET /expenses`

**Description**: Retrieves all expenses for a specific group.

**Request**:

```bash
curl -X GET "http://localhost:8080/expenses?group_id=1" \
  -H "Cookie: token=your_jwt_token"
```

**Response**:

- `200 OK` with a list of expenses.

## Running Tests

1. **Run Unit Tests**:

   ```bash
   go test ./...
   ```

2. **Run Integration Tests**:
   Make sure to initialize the database before running integration tests. This is usually handled in the test setup.

## Notes

- Ensure that you have Go installed and properly set up on your machine.
- Update the `token` in the curl requests with the actual JWT token you receive after login.

---

Feel free to adjust any specifics or add more information as needed!application will start and listen on port `8080` by default.

## API Endpoints

### User Registration

**Endpoint**: `POST /register`

**Description**: Registers a new user.

**Request**:

```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com", "password": "password123"}'
```

**Response**:

- `201 Created` on successful registration.

### User Login

**Endpoint**: `POST /login`

**Description**: Logs in a user and returns a JWT token.

**Request**:

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email": "john@example.com", "password": "password123"}'
```

**Response**:

- `200 OK` on successful login with a `Set-Cookie` header containing the JWT token.

### User Profile

**Endpoint**: `GET /profile`

**Description**: Retrieves the profile of the logged-in user.

**Request**:

```bash
curl -X GET http://localhost:8080/profile \
  -H "Cookie: token=your_jwt_token"
```

**Response**:

- `200 OK` with user profile details.

### Create Group

**Endpoint**: `POST /groups`

**Description**: Creates a new group.

**Request**:

```bash
curl -X POST http://localhost:8080/groups \
  -H "Content-Type: application/json" \
  -H "Cookie: token=your_jwt_token" \
  -d '{"name": "Family", "description": "Family expenses"}'
```

**Response**:

- `201 Created` with group details.

### Get Groups

**Endpoint**: `GET /groups`

**Description**: Retrieves all groups for the authenticated user.

**Request**:

```bash
curl -X GET http://localhost:8080/groups \
  -H "Cookie: token=your_jwt_token"
```

**Response**:

- `200 OK` with a list of groups.

### Add Expense

**Endpoint**: `POST /expenses`

**Description**: Adds a new expense.

**Request**:

```bash
curl -X POST http://localhost:8080/expenses \
  -H "Content-Type: application/json" \
  -H "Cookie: token=your_jwt_token" \
  -d '{"description": "Dinner", "amount": 50.00, "paid_by": "john@example.com", "group_id": 1}'
```

**Response**:

- `201 Created` with expense details.

### Get Expenses

**Endpoint**: `GET /expenses`

**Description**: Retrieves all expenses for a specific group.

**Request**:

```bash
curl -X GET "http://localhost:8080/expenses?group_id=1" \
  -H "Cookie: token=your_jwt_token"
```

**Response**:

- `200 OK` with a list of expenses.

## Running Tests

1. **Run Unit Tests**:

   ```bash
   go test ./...
   ```

2. **Run Integration Tests**:
   Make sure to initialize the database before running integration tests. This is usually handled in the test setup.

## Notes

- Ensure that you have Go installed and properly set up on your machine.
- Update the `token` in the curl requests with the actual JWT token you receive after login.

---
