# Start from the golang base image as the builder stage
FROM golang:1.22.1 as builder


# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build  -o geoip .

# Start a new stage from the alpine image
FROM alpine:3.19.1

WORKDIR /app/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/geoip .

# Create a user and group named "appuser"
RUN addgroup -S appuser && adduser -S appuser -G appuser

# Change the ownership of the /root directory to appuser
RUN chown -R appuser:appuser /app/

# Switch to "appuser"
USER appuser

# Command to run the executable
ENTRYPOINT ["./geoip"]