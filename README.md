# Student API

A simple API built with Go for managing student records. The API supports CRUD operations for students and integrates with the Ollama service to generate student summaries.

## Table of Contents
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Installation](#installation)
  - [Using Docker](#using-docker)
  - [Manual Installation](#manual-installation)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Concurrency Management](#concurrency-management)
- [Ollama Integration](#ollama-integration)
- [License](#license)

## Features
- Create, read, update, and delete student records.
- Generate summaries for student profiles using the Ollama service.

## Technologies Used
- Go (Golang)
- Docker
- JSON for data interchange

## Installation

### Using Docker
To run the application using Docker, follow these steps:

1. Ensure you have [Docker](https://docs.docker.com/get-docker/) installed on your machine.

2. Build the Docker image:
   ```bash
   docker build -t student-api .

3. Run the Docker container:
   ```bash
   docker run -p 8080:8080 --env-file .env student-api

### Using Docker
To run the application without Docker, follow these steps:
1. Ensure you have Go installed on your machine (version 1.20 or higher).
2. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/student-api.git
   cd student-api

3. Install dependencies:
   ```bash
   go mod tidy

4. Run the application:
   ```bash
   go run main.go
  The server will start on port 8080.


## Running the Application
Once the application is running, it listens for incoming HTTP requests on  `http://localhost:8080`.

## API Endpoints

### Students
1. `GET /students/ ` - Retrieve all students.
2. `POST /students/` - Create a new student.
3. `GET /students/{id}` - Retrieve a student by ID.
4. `PUT /students/{id}` - Update a student by ID.
5. `DELETE /students/{id}` - Delete a student by ID.
6. `GET /students/{id}/summary` - Get a summary of a student by ID (uses Ollama).


## Concurrency Management

Concurrency is handled in the application using Go's goroutines and mutexes. The `sync.Mutex` is used to ensure that only one goroutine can access or modify the student data at a time. This prevents race conditions and ensures data consistency when multiple requests are made simultaneously.

Key points include:

1. Use of `sync.Mutex` to lock access to critical sections of the code where student data is modified.
2. Functions such as `GetAllStudents`, `CreateStudent`, `GetStudentByID`, `UpdateStudent`, and `DeleteStudent` are all wrapped in mu.Lock() and mu.Unlock() to ensure thread safety.

## Ollama Integration

The application integrates with the Ollama service to generate student summaries. When a request is made to the `/students/{id}/summary` endpoint, the application:
1. Fetches the student data.
2. Constructs a prompt for the Ollama service.
3. Sends a request to the Ollama API and retrieves the summary.
4. Returns the summary as a response.

The integration is handled in the GenerateStudentSummary function located in the `services/ollama.go` file.

