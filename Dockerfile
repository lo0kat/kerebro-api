# Stage 1: Builder
FROM golang:1.23-alpine AS builder

# Set the Current Working Directory
WORKDIR /app

# Copy dependency files first
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go binary with dynamic name
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o kerebro-api .

# Stage 2: Runner
FROM alpine:3.21.3

WORKDIR /app

# Copy only the binary
COPY --from=builder /app/kerebro-api .

# Add a non-root user for security
RUN adduser -D -g '' kerebro
USER kerebro

# Set entrypoint to run the binary
ENTRYPOINT ["./kerebro-api"]
