# Base Image
FROM golang:1.17.6-alpine3.15 as base

# Working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY src .

# Build the application
RUN go build -o server ./src/.

# Create master image
FROM alpine AS master

# Working directory
WORKDIR /app

# Copy execute file
COPY --from=base /app/server ./

# Set ENV to production
ENV GO_ENV production

# Expose port 3005
EXPOSE 3005

# Run the application
CMD ["./server"]