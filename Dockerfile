FROM golang:1.24.2-alpine AS builder

# Install git for build info
RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ARG VERSION="dev"

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.version=${VERSION} -X main.commit=$(git rev-parse HEAD || echo 'unknown') -X main.date=$(date -u +%Y-%m-%dT%H:%M:%SZ)" -o server ./cmd/razorpay-mcp-server

FROM alpine:latest

RUN apk --no-cache add ca-certificates wget

# Create a non-root user to run the application
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/server .

# Change ownership of the application to the non-root user
RUN chown -R appuser:appgroup /app

# Environment variables for production
ENV RAZORPAY_API_KEY="" \
    RAZORPAY_API_SECRET="" \
    PORT=8080 \
    TOOLSETS="" \
    READ_ONLY="" \
    ENDPOINT_PATH="/mcp"

# Switch to the non-root user
USER appuser

# Expose the port
EXPOSE 8080

# Health check for Render
HEALTHCHECK --interval=30s --timeout=10s --start-period=30s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:${PORT}/mcp || exit 1

# Start the HTTP server
CMD ["sh", "-c", "./server --key ${RAZORPAY_API_KEY} --secret ${RAZORPAY_API_SECRET} --address :${PORT} --endpoint-path ${ENDPOINT_PATH} ${TOOLSETS:+--toolsets ${TOOLSETS}} ${READ_ONLY:+--read-only}"]
