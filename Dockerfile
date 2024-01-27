# Uses golang 1.21 alpine image as base
FROM golang:1.21.6-alpine3.19

# Sets the container pwd
WORKDIR /app

# Copy all project files over to the container
COPY src/ ./

# Downloads the project dependencies
RUN go mod download

# Project compilation
RUN go build -o app

# Exposes the port to the host
EXPOSE 8080

# Starts the server
CMD ["./app"]
