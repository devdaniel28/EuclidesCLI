package cmd

import (
	"fmt"
	"log"
	"strconv"
	"github.com/spf13/cobra"
)

var amount float64 //todo comfigura a flag dessa cmd abaixo

var simpleInterest = &cobra.Command{
	Use: "sint",
	Short: "A simple calculation about compound interest.",
	Run: func(cmd *cobra.Command, args []string) {
		capital, err := strconv.ParseFloat(args[0], 64)
		interestRate, err := strconv.ParseFloat(args[1], 64)
		time, err := strconv.ParseFloat(args[2], 64)

		

		if err != nil {
			log.Fatalln("Error: This is not a valid number.")
		}

		if interestRate > 0.1 {
			log.Fatalln("The interest rate cannot be greater than 0.01, and it must be written as a decimal below 0.01.")
		}

		fees := capital * interestRate * time
		fmt.Printf("The result of interest and %.f \n", fees)
	},
}