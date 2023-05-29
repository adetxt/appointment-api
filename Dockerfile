FROM golang:alpine

# Set the working directory
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o main .

# Expose the port that the application listens on
EXPOSE ${APP_PORT}

# Run the application
CMD ["./main"]
