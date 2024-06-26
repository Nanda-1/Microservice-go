# Build stage
FROM golang:1.21 AS builder

# Create and set working directory
RUN mkdir /app
WORKDIR /app

# Copy the source code to the working directory
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 go build -o authApp ./cmd/api

# Final stage
FROM alpine:latest

# Create application directory
RUN mkdir /app

# Copy the binary from the builder stage
COPY --from=builder /app/authApp /app/

# Ensure the binary has execute permissions
RUN chmod +x /app/authApp

# Set the command to run the application
CMD [ "/app/authApp"]
