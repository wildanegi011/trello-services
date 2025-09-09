# --------------------------
# Stage 1: Builder
# --------------------------
FROM golang:1.25.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/trello-services ./cmd/server/main.go

# --------------------------
# Stage 2: Production
# --------------------------
FROM debian:bullseye-slim

WORKDIR /app

# Copy only the binary (not source code)
COPY --from=builder /app/bin/trello-services .

# Expose port (change if needed)
EXPOSE 8080

# Run the binary
CMD ["./trello-services"]    