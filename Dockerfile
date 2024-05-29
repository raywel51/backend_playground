# Use an official Golang runtime as a parent image
FROM golang:1.21.3

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Install the "air" tool for hot reloading
RUN go get -u github.com/cosmtrek/air && \
    go install github.com/cosmtrek/air

# Expose port 1662 to the outside world (if needed)
EXPOSE 1662

# Command to run the "air" tool for hot reloading
CMD ["air"]
