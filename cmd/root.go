package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func test() {
	fmt.Println("testets")
}

var rootCmd = &cobra.Command{
	Use: "euclides",
	Short: "A math cli",
	Run: func(cmd *cobra.Command, args []string) {
		test()
	},
}

func Execute() {
	rootCmd.Execute()
}
