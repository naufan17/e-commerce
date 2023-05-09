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

# Set the environment variables for the database
# ENV DB_USER=root
# ENV DB_PASSWORD=
# ENV DB_HOST=localhost
# ENV DB_PORT=3306
# ENV DB_NAME=ecommerce

# Wait for the database to be ready
# CMD ["sh", "-c", "while ! nc -z $DB_HOST $DB_PORT; do sleep 1; done; ./main"]

# Run the application
CMD ["/app/main"]