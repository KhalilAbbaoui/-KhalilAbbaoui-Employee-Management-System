# Use the official Golang image from the Docker registry
FROM golang:1.20-alpine

# Set the working directory
WORKDIR /app

# Install Git to fetch dependencies if needed
RUN apk add --no-cache git

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
