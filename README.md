# Task-Manager-REST-API
# A REST API for managing tasks, built using Go and MongoDB.

# Table of Contents
* Overview
* Features
* Project Structure
* Setup
* Usage
* License

# Overview
This project is a task manager REST API built with Go and MongoDB. It provides endpoints for users to manage their tasks, including creating, updating, and deleting tasks.

Features
Create, update, and delete tasks
Mark tasks as complete or pending
Get tasks based on different criteria (all tasks, completed tasks, pending tasks)
User authentication and management
Project Structure
.
├── controllers
│   ├── task.go
│   ├── user.go
├── middleware
│   ├── auth.go
├── model
├── services
├── main.go
├── go.mod
└── README.md
Setup
Clone the repository:
git clone https://github.com/SleepySlash/Task-Manager-REST-API.git
cd Task-Manager-REST-API
Set up your environment variables by creating a .env file in the project root:
MONGO_URI=<Your MongoDB URI>
SECRET_KEY=<Your Secret Key>
DATABASE=<Your Database Name>
USER_COLLECTION=<Your User Collection Name>
Install the dependencies:
go mod tidy
Run the application:
go run main.go
Usage
Endpoints
User Endpoints:

POST /register - Register a new user
POST /login - Login a user
PUT /user/update - Update user information
DELETE /user/delete - Delete a user
Task Endpoints:

POST /todo/new - Create a new task
POST /todo/newtasks - Create multiple new tasks
GET /todo/get/{name}/{date} - Get a specific task
GET /todo/gettasks - Get all pending tasks
GET /todo/getall - Get all tasks including done tasks
PUT /todo/update/{name}/{date} - Update a specific task
PUT /todo/mark/done - Mark given tasks as done
PUT /todo/mark/done/{name}/{date} - Mark a specific task as done
PUT /todo/mark/pending/{name}/{date} - Mark a specific task as pending
DELETE /todo/delete/{name}/{date} - Delete a specific task
DELETE /todo/deleteall - Delete all tasks
Contributing
Contributions are welcome! Please open an issue or submit a pull request if you have any improvements.

License
This project is licensed under the MIT License.
