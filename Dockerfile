# Start from a Golang base image
FROM golang:1.20.4

# Set the working directory
WORKDIR /app

# Copy the source code to the working directory
COPY . /app

# Build the binary
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["/app/main"]