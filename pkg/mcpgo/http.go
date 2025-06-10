package mcpgo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mark3labs/mcp-go/server"
)

// HTTPServerOptions configures HTTP transport
type HTTPServerOptions struct {
	Address      string
	EndpointPath string
	Stateless    bool
}

// HTTPTransportServer implements HTTP transport for MCP
type HTTPTransportServer struct {
	mcpHTTPServer *server.StreamableHTTPServer
	httpServer    *http.Server
	address       string
}

// NewHTTPServer creates HTTP transport server
func NewHTTPServer(mcpServer Server, opts HTTPServerOptions) (*HTTPTransportServer, error) {
	sImpl, ok := mcpServer.(*Mark3labsImpl)
	if !ok {
		return nil, fmt.Errorf("invalid server implementation: expected *Mark3labsImpl, got %T", mcpServer)
	}

	// Configure HTTP server options
	var serverOpts []server.StreamableHTTPOption
	
	if opts.EndpointPath != "" {
		serverOpts = append(serverOpts, server.WithEndpointPath(opts.EndpointPath))
	}
	
	if opts.Stateless {
		serverOpts = append(serverOpts, server.WithStateLess(true))
	}

	// Create the streamable HTTP server
	httpServer := server.NewStreamableHTTPServer(sImpl.GetMCPServer(), serverOpts...)
	
	return &HTTPTransportServer{
		mcpHTTPServer: httpServer,
		address:       opts.Address,
	}, nil
}

// Start begins HTTP server
func (h *HTTPTransportServer) Start() error {
	return h.mcpHTTPServer.Start(h.address)
}

// Shutdown gracefully stops the server
func (h *HTTPTransportServer) Shutdown(ctx context.Context) error {
	if h.httpServer != nil {
		return h.httpServer.Shutdown(ctx)
	}
	return nil
}

// ServeHTTP implements http.Handler interface
func (h *HTTPTransportServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mcpHTTPServer.ServeHTTP(w, r)
}
