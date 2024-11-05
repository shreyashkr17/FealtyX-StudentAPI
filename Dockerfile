# Start with a lightweight Go image
FROM golang:1.23.1-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o student-api

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./student-api"]
