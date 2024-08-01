# Use the official Golang image for the build stage
FROM golang:1.20 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application with detailed output
RUN go build -v -o tours-service

# Use the same base image for the final stage to match glibc version
FROM debian:bullseye-slim

# Set the working directory in the container
WORKDIR /app

# Copy the binary from the build stage
COPY --from=builder /app/tours-service .

# Ensure that libc6 is installed (should be available in slim images)
RUN apt-get update && apt-get install -y libc6

# Expose the port the app runs on
EXPOSE 8080

# Define the command to run the application
CMD [ "./tours-service", "--help" ]
