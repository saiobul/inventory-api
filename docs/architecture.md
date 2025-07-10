# Architecture Overview
This document outlines the high-level system architecture for the Inventory Management REST API.

## Components
- **Client**: Web or mobile frontend API tools that interacts with the API.
- **API Gateway**: Entry point for all HTTP requests, handles routing and basic validation.
- **Load Balancer**: Distributes traffic across ECS containers for high availability.
- **ECS (Fargate)**: Hosts the Dockerized Golang REST API service.
- **RDS (PostgreSQL)**: Stores inventory data with ACID-compliant transactions.
- **CloudWatch**: Collects logs and metrics for monitoring and alerting.
- **CloudFormation**: Provisions infrastructure as code for consistent deployment.

## Architecture Flow
    Client → API Gateway → Load Balancer → ECS/EKS (Dockerized Go App) → RDS/DynamoDB
                                                ↘︎ CloudWatch (Monitoring & Logging)

## Technology Justification
- **Golang**: High performance and native concurrency support.
- **Docker**: Simplifies containerization and deployment.
- **AWS ECS (Fargate)**: Serverless container orchestration.
- **RDS**: Reliable managed database service.
- **CloudFormation**: Automates infrastructure provisioning.
- **CloudWatch**: Centralized monitoring and logging.
