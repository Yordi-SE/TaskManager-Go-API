TaskManager API Documentation
Overview
The TaskManager API provides functionality for managing tasks and user authentication. The API supports operations such as creating, reading, updating, and deleting tasks, as well as user registration, login, and role promotion.

Base URL
The base URL for all endpoints is:
http://<your-domain>/api

Authentication
This API uses JWT (JSON Web Token) for authentication. You need to include a valid JWT token in the Authorization header for protected routes.

Header format:
Authorization: Bearer <token>


Endpoints
1. User Endpoints
Register a New User
URL: /auth/register
Method: POST
Description: Registers a new user. The first registered user will have the "admin" role, and subsequent users will have the "user" role by default.
Request Body:
{
    "user_name": "string",
    "password": "string"
}


Response:
Success (200):
{
    "message": "User registered successfully",
    "user_id": "string"
}


Error (400/500):
{
    "message": "Error message"
}

Login
URL: /auth/login
Method: POST
Description: Authenticates a user and returns a JWT token.

Request Body:
{
    "user_name": "string",
    "password": "string"
}

Response:
Success (200):

{
    "message": "Login successful",
    "token": "string"
}

Error (400/403):
{
    "message": "Error message"
}

Promote User to Admin
URL: /auth/promote
Method: POST
Description: Promotes a user to the "admin" role. Only accessible by users with the "admin" role.

Request Body:
{
    "user_id": "string"
}

Response:
Success (200):
{
    "message": "Promotion successful"
}

Error (400/403):
{
    "message": "Error message"
}

2. Task Endpoints
Get All Tasks
URL: /tasks
Method: GET
Description: Retrieves a list of all tasks.

Response:
Success (200):
[
    {
        "_id": "string",
        "title": "string",
        "description": "string"
    }
]

Error (500):
{
    "message": "Error message"
}

Get Task by ID
URL: /tasks/:id
Method: GET
Description: Retrieves a specific task by its ID.
URL Parameters:
id (string): The ID of the task.

Response:
Success (200):
{
    "_id": "string",
    "title": "string",
    "description": "string"
}

Error (404/500):
{
    "message": "Error message"
}

Create a New Task
URL: /tasks
Method: POST
Description: Creates a new task.

Request Body:
{
    "title": "string",
    "description": "string"
}

Response:
Success (201):
{
    "message": "Task created successfully",
    "task_id": "string"
}

Error (400/500):
{
    "message": "Error message"
}


Update a Task
URL: /tasks/:id
Method: PUT
Description: Updates an existing task.
URL Parameters:
id (string): The ID of the task.

Request Body:
{
    "title": "string",
    "description": "string"
}

Response:
Success (200):
{
    "message": "Task updated successfully"
}

Error (400/404/500):
{
    "message": "Error message"
}


Delete a Task
URL: /tasks/:id
Method: DELETE
Description: Deletes a task by its ID.
URL Parameters:
id (string): The ID of the task.

Response:
Success (200):

{
    "message": "Task deleted successfully"
}

Error (404/500):

{
    "message": "Error message"
}

Error Handling
Errors are returned as JSON with an appropriate HTTP status code and a message describing the error.

Middleware
Authentication Middleware
The AuthMiddleware ensures that routes are accessible only to authenticated users by verifying the JWT token.

Admin Role Middleware
The Isadmin middleware restricts access to certain routes to users with the "admin" role.

Models
User
ID: ObjectId
UserName: string (required)
Password: string (required)
Role: string (default: "user")

Task
ID: ObjectId
Title: string (required)
Description: string