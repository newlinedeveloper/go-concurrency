# Use a minimal Alpine Linux image as the base
FROM golang:alpine

# Set the working directory
WORKDIR /app

# Copy the application source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Set the application as the entrypoint for the container
ENTRYPOINT ["./main"]
