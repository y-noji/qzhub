# Use an official Golang runtime as a parent image
FROM golang:1.20-alpine

# Set working directory
WORKDIR /app

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Install air for hot reloading
RUN go install github.com/cosmtrek/air@latest

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies.
# Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Expose port to the outside
EXPOSE 8080

# Command to run the executable
CMD ["air"]
