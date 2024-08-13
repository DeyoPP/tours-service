# Use the official Golang image for the build stage
FROM golang:1.20-alpine as builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o tours-service

# Use a minimal base image for the final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/tours-service .

EXPOSE 8080

# Define environment variables (optional: set these dynamically when running the container)
ENV DB_HOST=explorer-db.cscjwnksxo6a.eu-central-1.rds.amazonaws.com
ENV DB_USER=dejo
ENV DB_PASSWORD=sE7F0UY*8oDz-cbJ
ENV DB_NAME=explorer-db
ENV DB_PORT=5432

# Define the command to run the application
CMD [ "./tours-service" ]
