# Use the official Go image
FROM golang:1.22-alpine AS builder

# Install bash
RUN apk add --no-cache bash

# Set the working directory
WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./

# Fetch dependencies
RUN go mod tidy

# Copy source code
COPY . .

# Build the Go application
RUN go build -o main ./cmd/main.go

# Copy the wait-for-it script
COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

# Expose the application port
EXPOSE 8080

# Run the application with wait-for-it.sh
CMD ["/wait-for-it.sh", "postgres:5432", "--", "./main"]
