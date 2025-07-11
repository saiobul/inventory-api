# Stage 1: Build the Go binary
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# ✅ Static build for Alpine
ENV CGO_ENABLED=0
RUN go build -o inventory-api ./cmd/server

# Stage 2: Minimal runtime image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/inventory-api .
RUN chmod +x inventory-api
EXPOSE 8080

CMD ["./inventory-api"]
