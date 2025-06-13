package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "loganalyzer",
	Short: "Analyze log files",
	Long:  "Analyze is a tool to process and analyze log files.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error : %v\n", err)
		os.Exit(1)
	}
}
