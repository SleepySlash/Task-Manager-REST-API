# Task-Manager-REST-API

A REST API for managing tasks, built using Go and MongoDB.

# Table of Contents

- Overview
- Features
- Project Structure
- Setup
- Usage
- License

## Overview

This project is a task manager REST API built with Go and MongoDB. It provides endpoints for users to manage their tasks, including creating, updating, and deleting tasks.

## Features

- Create, update, and delete tasks
- Mark tasks as complete or pending
- Get tasks based on different criteria (all tasks, completed tasks, pending tasks)
- User authentication and management

## Project Structure

    .
    ├── controllers
    │  ├── task.go
    │  ├── user.go
    │
    ├── middleware
    │  ├── auth.go
    │
    ├── model
    │  ├── task.go
    │  ├── taskmodel.go
    │  ├── user.go
    │  ├── usermodel.go
    │
    ├── services
    │  ├── taskservice.go
    │  ├── userservice.go
    │
    ├── main.go
    ├── go.mod
    └── README.md

## Setup

1.  Clone the repository:

    ```
    git clone https://github.com/SleepySlash/Task-Manager-REST-API.git
    cd Task-Manager-REST-API
    ```

2.  Set up your environment variables by creating a .env file in the project root:

    ```
    MONGO_URI=<Your MongoDB URI>
    SECRET_KEY=<Your Secret Key>
    DATABASE=<Your Database Name>
    USER_COLLECTION=<Your User Collection Name>
    ```

3.  Install the dependencies:

    ```
    go mod tidy
    Run the application:
    go run main.go
    ```

## Usage

### Endpoints

- **User Endpoints:**

* Register a new user

  - Method: POST
  - URL `/register`
  - Request Body (JSON):

    ````{
        "username": "exampleUser",
        "password": "examplePassword"
      }```

    ````

* Login a user

  - Method: POST
  - URL `/login`
  - Request Body (JSON):
    ````{
        "username": "exampleUser",
        "password": "examplePassword"
      }```
    ````

* Update user information

  - Method: PUT
  - URL `/user/update`
  - Request Body (JSON):
    ````{
        "username": "updatedUser",
        "password": "updatedPassword"
     }```
    ````

* Delete a user
  - Method: DELETE
  - URL `/user/delete``
  - Request Body (JSON):
    ````{
        "username": "updatedUser",
     }```
    ````

- **Task Endpoints:**

* `POST /todo/new` - Create a new task
* `POST /todo/newtasks` - Create multiple new tasks
* `GET /todo/get/{name}/{date}` - Get a specific task
* `GET /todo/gettasks` - Get all pending tasks
* `GET /todo/getall` - Get all tasks including done tasks
* `PUT /todo/update/{name}/{date}` - Update a specific task
* `PUT /todo/mark/done` - Mark given tasks as done
* `PUT /todo/mark/done/{name}/{date}` - Mark a specific task as done
* `PUT /todo/mark/pending/{name}/{date}` - Mark a specific task as pending
* `DELETE /todo/delete/{name}/{date}` - Delete a specific task
* `DELETE /todo/deleteall` - Delete all tasks

## License

This project is licensed under the MIT License.

---

Please let me know if there are any additional sections or specific details you would like to include in the README file.
