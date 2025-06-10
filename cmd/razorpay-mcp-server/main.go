//nolint:lll
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	version = "version"
	commit  = "commit"
	date    = "date"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:     "server",
	Short:   "Razorpay MCP Server",
	Version: fmt.Sprintf("%s\ncommit %s\ndate %s", version, commit, date),
	Run: func(cmd *cobra.Command, args []string) {
		// Default to HTTP command when no subcommand is specified
		httpCmd.Run(cmd, args)
	},
}

// Execute runs the root command and handles any errors
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// flags will be available for all subcommands
	rootCmd.PersistentFlags().StringP("key", "k", "", "your razorpay api key")
	rootCmd.PersistentFlags().StringP("secret", "s", "", "your razorpay api secret")
	rootCmd.PersistentFlags().StringP("log-file", "l", "", "path to the log file")
	rootCmd.PersistentFlags().StringSliceP("toolsets", "t", []string{}, "comma-separated list of toolsets to enable")
	rootCmd.PersistentFlags().Bool("read-only", false, "run server in read-only mode")
	rootCmd.PersistentFlags().StringP("address", "a", ":8080", "address to listen on for HTTP transport")
	rootCmd.PersistentFlags().String("endpoint-path", "/mcp", "endpoint path for MCP requests")
	rootCmd.PersistentFlags().Bool("stateless", false, "run in stateless mode (no session management)")
	// bind flags to viper
	_ = viper.BindPFlag("key", rootCmd.PersistentFlags().Lookup("key"))
	_ = viper.BindPFlag("secret", rootCmd.PersistentFlags().Lookup("secret"))
	_ = viper.BindPFlag("log_file", rootCmd.PersistentFlags().Lookup("log-file"))
	_ = viper.BindPFlag("toolsets", rootCmd.PersistentFlags().Lookup("toolsets"))
	_ = viper.BindPFlag("read_only", rootCmd.PersistentFlags().Lookup("read-only"))
	_ = viper.BindPFlag("address", rootCmd.PersistentFlags().Lookup("address"))
	_ = viper.BindPFlag("endpoint_path", rootCmd.PersistentFlags().Lookup("endpoint-path"))
	_ = viper.BindPFlag("stateless", rootCmd.PersistentFlags().Lookup("stateless"))
	// Set environment variable mappings
	_ = viper.BindEnv("key", "RAZORPAY_KEY_ID", "RAZORPAY_API_KEY")        // Maps RAZORPAY_KEY_ID or RAZORPAY_API_KEY to key
	_ = viper.BindEnv("secret", "RAZORPAY_KEY_SECRET", "RAZORPAY_API_SECRET") // Maps RAZORPAY_KEY_SECRET or RAZORPAY_API_SECRET to secret
	_ = viper.BindEnv("address", "ADDRESS")                               // Maps ADDRESS to address
	_ = viper.BindEnv("toolsets", "TOOLSETS")                            // Maps TOOLSETS to toolsets
	_ = viper.BindEnv("read_only", "READ_ONLY")                          // Maps READ_ONLY to read_only
	_ = viper.BindEnv("endpoint_path", "ENDPOINT_PATH")                  // Maps ENDPOINT_PATH to endpoint_path

	// Enable environment variable reading
	viper.AutomaticEnv()
	// subcommands
	rootCmd.AddCommand(httpCmd)
	rootCmd.AddCommand(stdioCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".razorpay-mcp-server")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
