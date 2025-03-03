# Use the official Golang image as the base image
FROM golang:1.24 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -tags=viper_bind_struct -o build/server main.go

# Use a minimal base image to run the application
FROM debian:12-slim

# Install root certificates
RUN apt-get update && apt-get install -y ca-certificates

# Set the working directory inside the container
WORKDIR /app

# Copy the built application from the builder stage
COPY --from=builder /app/build/server /app/server
COPY --from=builder /app/config.hcl /app/config.hcl
EXPOSE 8080

# Command to run the application
CMD ["/app/server"]