FROM golang:1.24.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ARG VERSION="dev"
ARG COMMIT="unknown"
ARG BUILD_DATE="unknown"

WORKDIR /app/cmd/razorpay-mcp-server
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.version=${VERSION} -X main.commit=${COMMIT} -X main.date=${BUILD_DATE}" -o ../../server .

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
ENTRYPOINT ["./server"]
CMD ["--key", "", "--secret", "", "--address", ":8080", "--endpoint-path", "/mcp"]
