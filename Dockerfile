# ---------- STAGE 1: Builder ----------
FROM golang:1.23-bullseye as builder

# Set environment variables
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    SERVER_MODE=prod

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first for dependency caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go application
RUN go build -o main ./cmd


# ---------- STAGE 2: Final (slim Debian) ----------
FROM debian:bullseye-slim

RUN apt-get update && \
    apt-get install -y ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Add a non-root user
RUN useradd -m -s /bin/bash asp

# Set working directory
WORKDIR /app

# Copy the built binary from the builder
COPY --from=builder /app/main .

# Change ownership
RUN chown -R asp:asp /app

# Switch to non-root user
USER asp

# Expose application port
EXPOSE 8080

# Run the app
ENTRYPOINT ["./main"]
