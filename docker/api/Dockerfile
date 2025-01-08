# Use the official Go image as the base image
FROM golang:alpine AS builder

# Set the working directory inside the container
WORKDIR /api/golang

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o app ./cmd/api

# Use a minimal base image for production
FROM alpine:3.18

# Set the working directory in the minimal image
WORKDIR /api/golang

# Copy the compiled binary from th  e builder stage
COPY --from=builder /api/golang/app /api/golang/app

# Expose the port your app listens on
EXPOSE 8000

# Command to run the application
CMD ["./app"]
