
# Project Overview

This project is a scalable, high-performance Inventory Management REST API built using Golang, containerized with Docker, and deployed on AWS ECS (Fargate). It is designed to serve a mid-sized e-commerce platform, enabling efficient and reliable inventory operations.

# The API supports full CRUD functionality for managing product inventories, including:
- Adding new products
- Updating stock levels
- Retrieving product details
- Deleting discontinued items

# The system ensures:
- High concurrency handling using Go’s native Goroutines and channels
- Reliable data persistence via AWS RDS (PostgreSQL)
- Cloud-native deployment using AWS CloudFormation for infrastructure provisioning
- Monitoring and logging through AWS CloudWatch
- CI/CD automation using GitHub Actions
- This architecture embraces modern backend development practices, focusing on performance, scalability, and maintainability in a cloud-native environment.

# Project Folder Structure
    inventory-api/
    ├── cmd/
    │   └── server/              # Entry point for the application
    │       └── main.go
    ├── internal/
    │   ├── config/              # Configuration loading (env, AWS, DB)
    │   │   └── config.go
    │   ├── db/                  # Database connection and migrations
    │   │   ├── postgres.go
    │   │   └── models.go
    │   ├── product/             # Product domain logic
    │   │   ├── handler.go       # HTTP handlers
    │   │   ├── service.go       # Business logic
    │   │   └── repository.go    # DB operations
    │   ├── middleware/          # Logging, error handling, etc.
    │   │   └── logger.go
    │   └── utils/               # Utility functions
    │       └── response.go
    ├── pkg/                     # Shared packages (if needed)
    │   └── aws/                 # AWS SDK wrappers
    ├── api/
    │   └── routes.go            # Route definitions
    ├── Dockerfile               # Multi-stage Docker build
    ├── docker-compose.yml       # Local development setup
    ├── go.mod                   # Go module file
    ├── go.sum
    ├── cloudformation/          # AWS CloudFormation templates
    │   └── ecs-stack.yaml
    ├── .github/
    │   └── workflows/           # GitHub Actions CI/CD pipeline
    │       └── deploy.yml
    ├── docs/                    # All project documentation lives here
    ├── architecture.md          # System architecture explanation and diagrams
    ├── workflow.md              # API request lifecycle and service interactions
    ├── setup-guide.md           # Instructions for local setup, deployment, and CI/CD
    ├── monitoring.md            # CloudWatch setup, logging, and error tracking
    ├── diagrams/                # Folder for images like architecture diagrams
    │   ├── high-level-arch.png
    │   ├── request-flow.png
    │   └── multi-service-monitoring.png
    └── README.md                 # overview of what's in the docs folder

# Why This Structure?
- cmd/server
    Keeps the entry point clean and separate.
    Allows future expansion (e.g., CLI tools).
- internal/
    Encapsulates core logic and prevents external imports.
    Organized by concern: config, db, domain (product), middleware.
- product/
    Follows Clean Architecture principles:
    - handler.go: HTTP layer
    - service.go: Business logic
    - repository.go: Data access layer
- middleware/
    Centralized logging, error handling, request tracing.
- api/routes.go
    Keeps route definitions modular and readable.
- cloudformation/
    Infrastructure as Code for ECS, RDS, API Gateway, etc.
- .github/workflows/
    GitHub Actions for CI/CD: build, test, deploy.

