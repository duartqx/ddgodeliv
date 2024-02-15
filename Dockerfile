# Uses golang 1.21 alpine image as base
FROM golang:1.22.0-alpine3.19

# Sets the container pwd
WORKDIR /app

# Copy all project files over to the container
COPY src/ ./

# Project compilation
RUN go build -o app -mod=vendor

# Exposes the port to the host
EXPOSE 8080

# Starts the server
CMD ["./app"]
