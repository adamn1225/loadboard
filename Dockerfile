# Use the official Go image as the base image
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port your application runs on
EXPOSE 8080

# Command to run the application
CMD ["./main"]