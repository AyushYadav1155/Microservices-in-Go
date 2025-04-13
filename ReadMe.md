# Microservices in Go

A learning project implementing microservices architecture using Go, following a Udemy course. [Course link](https://www.udemy.com/course/working-with-microservices-in-go/)

## Project Overview

This application demonstrates how microservices can communicate with each other using Go. It's designed as a learning tool to explore different aspects of microservice architecture and communication patterns.

### Services (Current & Planned)

- ✅ Broker Service - Central communication hub between services
- ✅ Authentication Service - Handles user authentication
- ✅ Logger Service - Records system events and activities
- ✅ Mail Service - Manages email communications
- ⏳ RabbitMQ - Message broker for service communication
- ⏳ gRPC Logger - Demonstrates gRPC communication

### Architecture

- **Frontend**: Simple gohtml interface with buttons to test each service
- **Backend**: RESTful API communication between services
- **Deployment**: Docker containerization for all services

## Service Details

### Broker Service
- Acts as the central entry point for all requests
- Determines which service to forward requests to based on payload data
- Exposed on host machine at port 8086

### Authentication Service
- Authenticates users based on user ID and password
- Uses Postgres (postgres:14.2) as database
- Securely handles credential verification

### Logger Service
- Responsible for logging various communications within the microservices architecture
- Uses MongoDB (mongo:4.2.16-bionic) as database
- Provides centralized logging functionality

### Mail Service
- Sends emails to specified email addresses
- Utilizes Mailhog for email testing and development
- Mailhog interface is accessible at port 8025 on the host machine

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
| Authentication | PostgreSQL (14.2) |
| Logger | MongoDB (4.2.16-bionic) |
| Mail | Mailhog |

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
   
   # Build Mail Service
   cd ../mail-service
   set GOOS=linux
   set GOARCH=amd64
   set CGO_ENABLED=0
   go build -o mailerApp ./cmd/api
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

- **Completed**: Broker, Authentication, Logger, and Mail services
- **In Progress**: RabbitMQ integration
- **Planned**: gRPC Logger

## Technologies Used

- Go
- Docker
- PostgreSQL
- MongoDB
- Mailhog
- RESTful API