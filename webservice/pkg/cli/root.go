package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dmarc-analyzer",
	Short: "dmarc commands",
	Long:  `dmarc commands`,
}

func init() {
	// Add commands
	rootCmd.AddCommand(exampleCmd)
	rootCmd.AddCommand(dmarcCmd)
}

func Execute() {
	// Execute command
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
