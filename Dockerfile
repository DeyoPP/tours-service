# Use the official Golang image as a parent image
FROM golang:1.20 as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o tours-service

# Use a minimal base image for the final stage
FROM debian:bullseye-slim

# Set the working directory in the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/tours-service .

# Expose the port the app runs on
EXPOSE 8080

# Define the command to run the application
CMD [ "./tours-service" ]
