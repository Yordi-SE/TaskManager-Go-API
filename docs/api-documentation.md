
TaskManager API Documentation

Refer to the postman documentation
https://documenter.getpostman.com/view/37514043/2sA3rzKCjd

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
=======
postman doc => https://documenter.getpostman.com/view/37514043/2sA3rzKCjd


GET
Get Task By Id
http://localhost:8080/tasks/66b756c8f05e26f11dc95760
GET Specific Task
This endpoint retrieves a specific task by its task ID.

Request Body
This request does not require a request body.

Response Body
The response will contain the details of the specific task in JSON format. The task object includes the following properties:

id (string): The unique identifier of the task.
title (string): The title of the task.
description (string): The description of the task.
Example response body:

json
{
  "id": "0",
  "title": "",
  "description": ""
}


Example Request
Get Task By Id
curl
curl --location 'http://localhost:8080/tasks/1'
200 OK
Example Response
Body
Headers (3)
json
{
  "id": 1,
  "title": "Task 1",
  "description": "Description 1"
}
POST
create task
http://localhost:8080/tasks
create task endpoint
HTTP POST request to http://localhost:8080/tasks.

Request Body
Raw (application/json)
title (string, required)
description (string, required)
Response
The response of this request is a JSON schema with the following structure:

json
{
  "id": "number",
  "title": "string",
  "description": "string"
}


Body
raw (json)
json
{
    "title": "Task Managing clean architecture",
    "description": "This test task for task Managing clean arch"
}
Example Request
create task
curl
curl --location 'http://localhost:8080/tasks' \
--data '{
    "title": "Task Managing API",
    "description": "This test task for task Managing API"
}'
201 Created
Example Response
Body
Headers (3)
json
{
  "id": 4,
  "title": "Task Managing API",
  "description": "This test task for task Managing API"
}
PUT
update task endpoint
http://localhost:8080/tasks/66b3d76fa617bcbeef07057b
Update Task Details
This endpoint allows you to update the details of a specific task identified by its task ID.

Request Body
The request should include a JSON payload with the following parameters:
title (string): The updated title of the task.
description (string): The updated description of the task.
Example:

json
{
  "title": "Updated Title",
  "description": "Updated Description"
}
Response
Status: 200
Content-Type: application/json
The response will include a JSON object with the following fields:

MatchedCount (number): The number of matched documents.
ModifiedCount (number): The number of modified documents.
UpsertedCount (number): The number of upserted documents.
UpsertedID (null or string): The ID of the upserted document, if any.
Example response:

json
{
  "MatchedCount": 0,
  "ModifiedCount": 0,
  "UpsertedCount": 0,
  "UpsertedID": null
}
AUTHORIZATION
Bearer Token
Token
<token>

Body
raw (json)
View More
json
{
    "title": "let the it to my first task",
    "description": "let the description also be the description itself"
}
Example Request
update task endpoint
curl
curl --location --request PUT 'http://localhost:8080/tasks/4' \
--data '{
    "title": "let the it to my first task",
    "description": "let the description also be the description itself"
}'
200 OK
Example Response
Body
Headers (3)
json
{
  "id": 4,
  "title": "let the it to my first task",
  "description": "let the description also be the description itself"
}
DELETE
delete task endpoint
http://localhost:8080/tasks/66b756c8f05e26f11dc95760
The endpoint allows you to delete a specific task.

Response
The response for this request can be documented as a JSON schema:

json
{
    "type": "object",
    "properties": {
        "message": {
            "type": "string"
        }
    }
}
