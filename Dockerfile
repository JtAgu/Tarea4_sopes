# Use the official Go base image
FROM golang:1.19-alpine

# Set the working directory inside the container
WORKDIR /back-container

# Copy the Go modules files and download dependencies
COPY . .

RUN go mod download

EXPOSE 8000

# run
CMD ["go","run","main.go"]