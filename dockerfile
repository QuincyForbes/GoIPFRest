# Use the official Go image as a parent image
FROM golang:1.20 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY "./basicmodule" /app

# Download dependencies
RUN go mod download

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Use a lightweight image for deployment
FROM alpine:latest

# Expose port 8080 for the Go application
EXPOSE 8080

# Set necessary environment variables for better performance
ENV GIN_MODE=release

# Install necessary packages (like ca-certificates) for the application
RUN apk --no-cache add ca-certificates

# Set the working directory in docker
WORKDIR /root/

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/main .

# Command to run when starting the container
CMD ["./main"]
