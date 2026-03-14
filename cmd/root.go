package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func Run() {
	fmt.Println("Calculating...")
	rootCmd.Execute()
}