# Setup Guide
This guide explains how to set up the project locally and deploy it to AWS.

## Local Development

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/inventory-api.git
   cd inventory-api
   ```
2. Run the API locally using Docker Compose:
    ```bash
    docker-compose up
    ```

3. Access the API at:
    http://localhost:8080

# Environment Configuration
- Set environment variables for DB connection, AWS credentials, etc.
- Use .env file or Docker secrets for local development.

# AWS Deployment
- Use CloudFormation templates in cloudformation/ to provision:
    - VPC, ECS Cluster, RDS, Load Balancer, API Gateway
- Build and push Docker image to Amazon ECR:
    ```bash
    docker build -t inventory-api .
    docker tag inventory-api:latest <your-ecr-repo-url>
    docker push <your-ecr-repo-url>
    ```
- Deploy ECS service using the pushed image.

# CI/CD Pipeline
- GitHub Actions workflow in .github/workflows/deploy.yml:
- Lints and tests the Go app
- Builds and pushes Docker image to ECR
- Deploys to ECS using CloudFormation

---
