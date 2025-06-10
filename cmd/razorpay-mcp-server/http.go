package main

import (
	"context"
	"fmt"
	stdlog "log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	rzpsdk "github.com/razorpay/razorpay-go"

	"github.com/razorpay/razorpay-mcp-server/pkg/log"
	"github.com/razorpay/razorpay-mcp-server/pkg/mcpgo"
	"github.com/razorpay/razorpay-mcp-server/pkg/razorpay"
)

// httpCmd starts the mcp server in HTTP transport mode
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "start the HTTP server (default)",
	Run: func(cmd *cobra.Command, args []string) {
		logPath := viper.GetString("log_file")
		log, close, err := log.New(logPath)
		if err != nil {
			stdlog.Fatalf("create logger: %v", err)
		}
		defer close()

		key := viper.GetString("key")
		secret := viper.GetString("secret")
		client := rzpsdk.NewClient(key, secret)

		client.SetUserAgent("razorpay-mcp/" + version + "/http")
		// Get address from flags or config, with PORT environment variable support for Render
		address := viper.GetString("address")
		if address == "" {
			address = ":8080"
		}

		// Handle Render's PORT environment variable
		if port := os.Getenv("PORT"); port != "" {
			address = ":" + port
		}

		// Get endpoint path from flags or config, default to /mcp
		endpointPath := viper.GetString("endpoint_path")
		if endpointPath == "" {
			endpointPath = "/mcp"
		}

		// Get toolsets to enable from config
		enabledToolsets := viper.GetStringSlice("toolsets")

		// Get read-only mode from config
		readOnly := viper.GetBool("read_only")

		// Get stateless mode from config
		stateless := viper.GetBool("stateless")

		err = runHTTPServer(log, client, address, endpointPath, enabledToolsets, readOnly, stateless)
		if err != nil {
			log.Error("error running HTTP server", "error", err)
			stdlog.Fatalf("failed to run HTTP server: %v", err)
		}
	},
}

func runHTTPServer(
	log *slog.Logger,
	client *rzpsdk.Client,
	address string,
	endpointPath string,
	enabledToolsets []string,
	readOnly bool,
	stateless bool,
) error {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	// Create the Razorpay server
	srv, err := razorpay.NewServer(
		log,
		client,
		version,
		enabledToolsets,
		readOnly,
	)
	if err != nil {
		return fmt.Errorf("failed to create server: %w", err)
	}

	// Create HTTP transport server
	httpSrv, err := mcpgo.NewHTTPServer(srv.GetMCPServer(), mcpgo.HTTPServerOptions{
		Address:      address,
		EndpointPath: endpointPath,
		Stateless:    stateless,
	})
	if err != nil {
		return fmt.Errorf("failed to create HTTP server: %w", err)
	}

	// Start server in a goroutine
	errC := make(chan error, 1)
	go func() {
		log.Info("starting HTTP server",
			"address", address,
			"endpoint", endpointPath,
			"stateless", stateless)
		errC <- httpSrv.Start()
	}()

	_, _ = fmt.Fprintf(
		os.Stderr,
		"Razorpay MCP Server running on %s%s\n",
		address,
		endpointPath,
	)

	// Wait for shutdown signal or error
	select {
	case <-ctx.Done():
		log.Info("shutting down HTTP server...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 30)
		defer cancel()
		return httpSrv.Shutdown(shutdownCtx)
	case err := <-errC:
		if err != nil {
			log.Error("HTTP server error", "error", err)
			return err
		}
		return nil
	}
}
