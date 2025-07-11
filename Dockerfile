# Stage 1: Build the Go binary
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o inventory-api ./cmd/server
RUN ls -l /app/inventory-api  # ✅ Debug: confirm binary exists

# Stage 2: Create a minimal runtime image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/inventory-api .
RUN chmod +x inventory-api     # Ensure it's executable
RUN ls -l                      # Debug: confirm binary is here

EXPOSE 8080

CMD ["./inventory-api"]
