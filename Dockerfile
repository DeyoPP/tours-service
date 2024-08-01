# Use the official Golang image for Go 1.20
FROM golang:1.20 as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod download

# Copy the rest of the application code
COPY . .

# Print directory contents for debugging
RUN ls -la

# Build the Go application with detailed error output
RUN go build -v -o tours-service

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
