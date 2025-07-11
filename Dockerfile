# Use official Golang image
FROM golang:1.23

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build and run the app
RUN go build -o app ./cmd/server

EXPOSE 8080

CMD ["./app"]
