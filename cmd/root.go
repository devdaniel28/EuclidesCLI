package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "euc",
	Short: "A basic math app",
}

func Run() {
	rootCmd.Execute()
}