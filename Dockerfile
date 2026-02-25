# Stage 1: Build the Go application
FROM golang:1.24-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go application
# CGO_ENABLED=0 is important for a statically linked binary
# -o /app/main builds the binary into the /app directory with the name "main"
# ./cmd/main.go is the entrypoint of the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main ./cmd/main.go

# Stage 2: Create a minimal final image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Copy the config directory
COPY config ./config

# Expose the port the application runs on
EXPOSE 8000

# Command to run the application
CMD ["/app/main"]
