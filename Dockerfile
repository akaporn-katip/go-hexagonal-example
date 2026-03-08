# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod tidy

# Copy source code
COPY . .

# Log copy operation and list files
RUN echo "Source code copied successfully" && ls -la


# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/main ./main.go

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/config/config.yaml ./config/
COPY --from=builder /app/migrate/sql    ./migrate/sql
COPY --from=builder /app/migrate/mongo ./migrate/mongo
COPY --from=builder /app/bin/main .

# Expose port (adjust as needed)
EXPOSE 8080

# Run the application
ENTRYPOINT ["./main"]
CMD ["serve"]