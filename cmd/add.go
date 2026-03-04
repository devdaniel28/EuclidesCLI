package cmd

import (
	// "fmt"
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(squareRootCmd)
}

var squareRootCmd = &cobra.Command{
	Use: "sqrt",
	Short: "Calculate the square root of the given number.",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		source, err := strconv.ParseFloat(args[0], 64)

		if err != nil {
			log.Fatalln("Error: This is not a valid number.\n", args[0])
			return
		}

		if source < 0 {
			log.Fatalln("Error: It is not possible to perform tasks with negative numbers.")
		}

		result := math.Sqrt(source)
		fmt.Printf("The square root of %.2f is %.4f \n", source, result)
	},
}