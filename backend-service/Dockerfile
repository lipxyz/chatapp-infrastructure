# Stage 1: build
FROM golang:1.24-alpine AS builder
WORKDIR /app

# Copy Go modules first (for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy all source files
COPY . .

# Build the Go backend
RUN CGO_ENABLED=0 GOOS=linux go build -o backend main.go

# Stage 2: runtime
FROM alpine:3.20
WORKDIR /app

COPY --from=builder /app/backend .

COPY sql ./sql
EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/healthz


CMD ["./backend"]