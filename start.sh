#!/bin/bash

# Start script for Razorpay MCP Server
# Usage: ./start.sh [http|stdio] [additional_flags]

set -e

# Default values
TRANSPORT=${1:-http}
RAZORPAY_API_KEY=${RAZORPAY_API_KEY:-""}
RAZORPAY_API_SECRET=${RAZORPAY_API_SECRET:-""}
ADDRESS=${ADDRESS:-":8080"}
ENDPOINT_PATH=${ENDPOINT_PATH:-"/mcp"}
TOOLSETS=${TOOLSETS:-""}
READ_ONLY=${READ_ONLY:-"false"}

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}🚀 Starting Razorpay MCP Server${NC}"
echo -e "${YELLOW}Transport: ${TRANSPORT}${NC}"

# Check if binary exists
if [ ! -f "./razorpay-mcp-server" ] && [ ! -f "./server" ]; then
    echo -e "${RED}❌ Binary not found. Building...${NC}"
    go build -o razorpay-mcp-server ./cmd/razorpay-mcp-server
    echo -e "${GREEN}✅ Build complete${NC}"
fi

# Use the right binary name
BINARY="./razorpay-mcp-server"
if [ -f "./server" ]; then
    BINARY="./server"
fi

# Check API keys
if [ -z "$RAZORPAY_API_KEY" ] || [ -z "$RAZORPAY_API_SECRET" ]; then
    echo -e "${RED}❌ Error: RAZORPAY_API_KEY and RAZORPAY_API_SECRET must be set${NC}"
    echo "Export them as environment variables or pass as flags:"
    echo "  export RAZORPAY_API_KEY=your_key"
    echo "  export RAZORPAY_API_SECRET=your_secret"
    echo "  ./start.sh"
    echo ""
    echo "Or:"
    echo "  ./start.sh http --key your_key --secret your_secret"
    exit 1
fi

# Build command arguments
ARGS=""

if [ "$TRANSPORT" = "http" ]; then
    echo -e "${YELLOW}Starting HTTP server on ${ADDRESS}${ENDPOINT_PATH}${NC}"
    ARGS="$ARGS --address $ADDRESS --endpoint-path $ENDPOINT_PATH"
elif [ "$TRANSPORT" = "stdio" ]; then
    echo -e "${YELLOW}Starting stdio server${NC}"
    ARGS="stdio $ARGS"
else
    echo -e "${RED}❌ Invalid transport: $TRANSPORT. Use 'http' or 'stdio'${NC}"
    exit 1
fi

# Add API keys
ARGS="$ARGS --key $RAZORPAY_API_KEY --secret $RAZORPAY_API_SECRET"

# Add optional parameters
if [ -n "$TOOLSETS" ]; then
    ARGS="$ARGS --toolsets $TOOLSETS"
fi

if [ "$READ_ONLY" = "true" ]; then
    ARGS="$ARGS --read-only"
fi

# Add any additional arguments passed to script
shift 1 # Remove transport argument
ARGS="$ARGS $@"

echo -e "${GREEN}Command: $BINARY $ARGS${NC}"
echo -e "${YELLOW}Press Ctrl+C to stop${NC}"
echo ""

# Start the server
exec $BINARY $ARGS
