# API Request Workflow
This document describes the lifecycle of an API request from client to database and back.

## Request Flow
1. **Client** sends an HTTP request (e.g., POST /products).
2. **API Gateway** receives and forwards the request to the Load Balancer.
3. **Load Balancer** routes the request to an ECS container.
4. **ECS (Golang API)** processes the request:
   - Validates input
   - Executes business logic
   - Interacts with RDS for data persistence
5. **Response** is returned via Load Balancer and API Gateway to the client.

## Error Handling Flow
- If an error occurs (e.g., DB failure, invalid input):
  - ECS logs the error using structured logging (Zap/Logrus).
  - Logs are forwarded to CloudWatch Logs.
  - CloudWatch Alarms can trigger notifications based on error metrics.
  - ECS returns an appropriate error response to the client.

## Inter-Service Communication (if applicable)
- Services communicate via internal APIs or message queues.
- Each service logs independently to CloudWatch.

## Visual Reference
Refer to `docs/diagrams/request-flow.png` for a visual representation.
