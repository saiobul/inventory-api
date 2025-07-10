# Stage 1: Build the Go binary
FROM golang:1.20 AS builder

WORKDIR /app

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN go build -o inventory-api ./cmd/server

# Stage 2: Create a minimal runtime image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/inventory-api .

# Expose the port the app runs on
EXPOSE 8080

# Run the binary
CMD ["./inventory-api"]
