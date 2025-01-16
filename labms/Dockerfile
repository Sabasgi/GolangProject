# Use Golang official image as the base
FROM golang:1.20

# Install MySQL client and server
RUN apt-get update && \
    apt-get install -y mysql-server mysql-client

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project into the container
COPY . .

# Initialize MySQL (if needed, create a database and set up)
RUN service mysql start && \
    mysql -e "CREATE DATABASE IF NOT EXISTS labms;" && \
    mysql -e "GRANT ALL PRIVILEGES ON labms.* TO 'root'@'%' IDENTIFIED BY 'root';"

# Expose the port your server listens on
EXPOSE 8768

# Build the application
RUN go build -o main .

# Command to run MySQL and the Go application
CMD service mysql start && ./main
