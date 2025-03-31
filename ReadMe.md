# Microservices in Go

A learning project implementing microservices architecture using Go, following a Udemy course. [Course link](https://www.udemy.com/course/working-with-microservices-in-go/)

## Project Overview

This application demonstrates how microservices can communicate with each other using Go. It's designed as a learning tool to explore different aspects of microservice architecture and communication patterns.

### Services (Current & Planned)

- ✅ Broker Service - Central communication hub between services
- ✅ Authentication Service - Handles user authentication
- ✅ Logger Service - Records system events and activities
- ⏳ Mail Service - Manages email communications
- ⏳ gRPC Logger - Demonstrates gRPC communication

### Architecture

- **Frontend**: Simple gohtml interface with buttons to test each service
- **Backend**: RESTful API communication between services
- **Deployment**: Docker containerization for all services

## Service Interactions

The Broker Service acts as the central communication hub:

1. User clicks a button in the frontend
2. Frontend sends a POST request to the Broker Service
3. Broker Service forwards the request to the appropriate microservice
4. Microservice processes the request and returns a response
5. Broker Service returns the response to the frontend

## Database Architecture

| Service | Database |
|---------|----------|
| Authentication | PostgreSQL |
| Logger | MongoDB |

## Getting Started

### Prerequisites

- Docker and Docker Compose
- Make (optional but recommended)
  - [Make Installation](https://www.gnu.org/software/make/) - Official GNU Make website
- Go (for building services)

### Installation

#### Using Make (Recommended)
```bash
# Clone the repository
git clone [repository-url]
cd [project-directory]

# Build and start all services
cd compose-files
make up_build

# Start the frontend
make start
```

#### Without Make (Manual Steps)

1. **Build the service binaries**:
   ```bash
   # Build Broker Service
   cd broker-service
   set GOOS=linux
   set GOARCH=amd64
   set CGO_ENABLED=0
   go build -o brokerApp ./cmd/api
   
   # Build Authentication Service
   cd ../authentication-service
   set GOOS=linux
   set GOARCH=amd64
   set CGO_ENABLED=0
   go build -o authApp ./cmd/api
   
   # Build Logger Service
   cd ../logger-service
   set GOOS=linux
   set GOARCH=amd64
   set CGO_ENABLED=0
   go build -o loggerServiceApp ./cmd/api
   ```

2. **Start the services with Docker Compose**:
   ```bash
   cd compose-files
   docker-compose down
   docker-compose up --build -d
   ```

3. **Build and run the frontend**:
   ```bash
   cd ../front-end
   set CGO_ENABLED=0
   set GOOS=windows
   go build -o frontApp.exe ./cmd/web
   
   # Run the frontend
   frontApp.exe
   ```
> Since I have been using windows the builds are set with respect to windows but if you are a Linux user you would have to make a change in the Makefile and change the GOOS of front-end from "GOOS=windows" to "GOOS=linux"

## Development Status

- **Completed**: Broker, Authentication, and Logger services
- **In Progress**: Mail Service and gRPC Logger

## Technologies Used

- Go
- Docker
- PostgreSQL
- MongoDB
- RESTful API