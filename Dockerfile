# Stage 1: Build the Go binary
FROM golang:1.23 AS builder

# Sets the working directory inside the container to /app.
WORKDIR /app

#Copies from your local project directory into the container's /app directory.
COPY go.mod go.sum ./
# Downloads all Go dependencies listed in go.mod.
RUN go mod download

# Copies everything from your local project directory into the container’s /app directory.
COPY . .

# Static build for Alpine
ENV CGO_ENABLED=0
# Builds your Go app from the ./cmd/server directory and outputs the binary as /app/inventory-api.
RUN go build -o inventory-api ./cmd/server

# Stage 2: Minimal runtime image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/inventory-api .
RUN chmod +x inventory-api
EXPOSE 8080

CMD ["./inventory-api"]
