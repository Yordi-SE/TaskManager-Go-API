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
AUTHORIZATION
Bearer Token
Token
<token>

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
AUTHORIZATION
Bearer Token
Token
<token>

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
