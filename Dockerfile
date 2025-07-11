# Use a minimal base image
FROM alpine:latest

# Install certificates (needed for HTTPS calls)
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /app

# Copy your prebuilt Go binary into the container
COPY inventory-api .

# Make sure it's executable
RUN chmod +x inventory-api

# Expose the port your app listens on
EXPOSE 8080

# Run the binary
CMD ["./inventory-api"]
