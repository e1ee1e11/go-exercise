# Stage 1: Build the Go binary
FROM golang:1.21-alpine as builder

# Set the working directory
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o myapp

# Stage 2: Create a minimal runtime image
FROM alpine:3.19

# Set the working directory in the minimal image
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/myapp .

# Expose 8080 port
EXPOSE 8080

# Command to run the application
CMD ["./myapp"]
