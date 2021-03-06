# === #

# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Vilbert Gunawan <vilbertgunawan@gmail.com>"

# Environment
ENV ENV="staging"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main cmd/http/main.go

# Expose port 8888 to the outside world
EXPOSE 8888

# Command to run the executable
CMD ["./main"]