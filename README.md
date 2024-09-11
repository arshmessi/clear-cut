---

# Clear-Cut: A Payment Splitting Application

## Overview

Clear-Cut is a split payment application with a React frontend and Go backend. This project allows users to register, log in, manage groups, and record expenses.

## Project Structure

- **cmd/main.go**: Entry point of the Go application.
- **internal/auth**: Contains JWT-related functions.
- **internal/handlers**: Handles HTTP requests.
- **internal/models**: Defines data models.
- **internal/storage**: Manages database interactions and schema.
- **internal/services**: Contains in-memory data services (for testing).
- **src/**: Contains the React frontend code.
- **src/context/AuthContext.tsx**: Handles authentication context.
- **src/services/api.ts**: Provides API request functions.
- **src/pages/**: Contains React components for different pages.

## Setup Instructions

### Backend

1. **Navigate to the backend directory:**

   ```sh
   cd clear-cut/backend
   ```

2. **Set up environment variables:**

   Create a `.env` file and set `JWT_SECRET_KEY`:

   ```env
   JWT_SECRET_KEY=your_secret_key
   ```

3. **Install dependencies:**

   ```sh
   go mod tidy
   ```

4. **Run the backend server:**

   ```sh
   go run cmd/main.go
   ```

   The server will start and listen on port `8080` by default.

### Frontend

1. **Navigate to the frontend directory:**

   ```sh
   cd clear-cut/frontend
   ```

2. **Install dependencies:**

   ```sh
   npm install
   ```

3. **Run the frontend app:**

   ```sh
   npm start
   ```

   The React application will start and be available at `http://localhost:3000` by default.

## API Endpoints

### Register

**Endpoint**: `POST /register`

**Description**: Registers a new user.

**Request**:
```sh
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com", "password": "password123"}'
```

### Login

**Endpoint**: `POST /login`

**Description**: Logs in a user and returns a JWT token.

**Request**:
```sh
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email": "john@example.com", "password": "password123"}'
```

### Fetch Profile

**Endpoint**: `GET /profile`

**Description**: Retrieves the profile of the logged-in user.

**Request**:
```sh
curl -X GET http://localhost:8080/profile \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### Create Group

**Endpoint**: `POST /groups`

**Description**: Creates a new group.

**Request**:
```sh
curl -X POST http://localhost:8080/groups \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{"name": "Family", "description": "Family expenses"}'
```

### Get Groups

**Endpoint**: `GET /groups`

**Description**: Retrieves all groups for the authenticated user.

**Request**:
```sh
curl -X GET http://localhost:8080/groups \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### Add Expense

**Endpoint**: `POST /expenses`

**Description**: Adds a new expense.

**Request**:
```sh
curl -X POST http://localhost:8080/expenses \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{"description": "Dinner", "amount": 50.00, "paid_by": "john@example.com", "group_id": 1}'
```

### Get Expenses

**Endpoint**: `GET /expenses`

**Description**: Retrieves all expenses for a specific group.

**Request**:
```sh
curl -X GET "http://localhost:8080/expenses?group_id=1" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Running Tests

1. **Run Unit Tests:**

   ```sh
   go test ./...
   ```

2. **Run Integration Tests:**

   Ensure that the database is initialized before running integration tests. This is usually handled in the test setup.

## Development

### Frontend

- **State Management:** Utilize the React Context API to manage global state efficiently across the application. This helps maintain user authentication state, user profiles, and other global data seamlessly.
  
- **UI/UX Design:** Leverage Material UI to build a modern and responsive user interface. Material UI’s pre-designed components are highly customizable, ensuring a consistent and polished look throughout the application.

- **TypeScript:** Use TypeScript to enhance code quality and reliability. TypeScript’s static type checking helps catch potential errors during development and provides robust tooling support.

- **API Integration:** Employ the `fetch` API to handle HTTP requests to the backend. Ensure proper management of API calls and implement robust error handling for a smooth user experience.

### Backend

- **API Services:** Develop RESTful API services using Go to manage user authentication, group management, and expense tracking. Ensure that endpoints are secure, efficient, and thoroughly documented.

- **Database Management:** Implement SQLite for straightforward and lightweight database management. Structure the database to support efficient queries and maintain data integrity.

- **JWT Authentication:** Utilize JWT tokens for secure user authentication. Tokens should be stored in session storage and validated on each request to safeguard user data.

- **Testing:** Write comprehensive unit and integration tests to ensure backend reliability. Use Go’s testing framework to validate your application’s logic and API endpoints thoroughly.

## Final Check

- **Integration Testing:** Verify that the frontend and backend work together seamlessly. Test the entire user workflow from registration to login, profile fetching, group management, and expense tracking. Ensure all data is synchronized and accurately displayed.

- **User Experience:** Focus on delivering an intuitive and engaging user experience. Test the application’s responsiveness across various devices and screen sizes, and ensure smooth navigation and interactions.

- **Error Handling:** Confirm that the application handles errors gracefully. Users should receive clear and actionable error messages when issues arise, enhancing overall usability.

- **Performance:** Assess the performance of both frontend and backend components. Optimize API calls and reduce load times to deliver a fast and responsive user experience.

- **Security:** Conduct thorough security checks to protect against common vulnerabilities. Ensure sensitive data is handled securely and that authentication mechanisms are robust and reliable.

---
