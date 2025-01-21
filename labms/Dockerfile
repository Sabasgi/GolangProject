# Use Golang official image as the base
FROM golang:1.20
# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project into the container
COPY . .

# Build the application
RUN go build -o main .

# Expose the port your server listens on
EXPOSE 8768

# Command to run the application
CMD ["./main"]