# Use Golang official image as the base
FROM golang:1.20

# Install MariaDB (MySQL alternative) and other dependencies
RUN apt-get update && apt-get install -y mariadb-server mariadb-client

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project into the container
COPY . .

# Initialize MariaDB
RUN service mysql start && \
    mysql -u root -e "CREATE DATABASE IF NOT EXISTS labms;" && \
    mysql -u root -e "GRANT ALL PRIVILEGES ON labms.* TO 'root'@'%';"

# Build the application
RUN go build -o main .

# Expose the port your server listens on
EXPOSE 8768

# Start MariaDB and your Golang application when the container starts
CMD service mysql start && ./main
