package cmd

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"github.com/spf13/cobra"
)

var squareRootCmd = &cobra.Command{
	Use: "sqrt",
	Short: "Calculate the square root of the given number.",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		source, err := strconv.ParseFloat(args[0], 64)

		if err != nil {
			log.Fatalln("Error: This is not a valid number.\n", args[0])
		}

		if source < 0 {
			log.Fatalln("Error: It is not possible to perform tasks with negative numbers.")
		}

		result := math.Sqrt(source)
		fmt.Printf("The square root of %.2f is %.4f \n", source, result)
	},
}

var exponentiation = &cobra.Command{
	Use: "exp",
	Short: "A command that generates the exponentiation of numbers.",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		base, err := strconv.ParseFloat(args[0], 64)
		exponent, err1 := strconv.ParseFloat(args[1], 64)

		if err != nil {
			log.Fatalf("Error: Wrong base number %.f \n try: exp <base> <exponent> \n", base)
		}

		if err1 != nil {
			log.Fatalf("Error: Wrong exponent number %.f \n try: exp <base> <exponent> \n", exponent)
		}

		if base == 0 || exponent <= 0 {
			log.Fatalln("It is not possible to perform exponentiation with negative numbers or numbers equal to 0.")
		}

		var result = math.Pow(base, exponent)
		fmt.Printf("The result is %.f raised to the power of %.f, which equals %.f \n", base, exponent, result)
		fmt.Printf("The formula is: %.f ^ %.f = %.f \n", base, exponent, result)
	},
}

var percentage = &cobra.Command{
	Use: "perc",
	Short: "Calculate the percentage of a given number.",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		number, err := strconv.ParseFloat(args[0], 64)
		percent, err := strconv.ParseFloat(args[1], 64)

		if err != nil{
			log.Fatalln("Error: This is not a valid number.")
		}

		if number > 32000000000 {
			log.Fatalln("Error: Number too large, try another.")
		}

		result := (percent / 100) * number
		percentStr := strconv.FormatFloat(percent, 'f', 2, 64) + "%"

		fmt.Printf("%s of %.2f equals %.2f \n", percentStr, number, result)
	},
}

var simpleRuleOfThree = &cobra.Command{
	Use: "srot",
	Short: "This is the rule of three.",
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		numericalRelation, err := strconv.ParseFloat(args[0], 64)
		numericalRelation2, err := strconv.ParseFloat(args[1], 64)
		numericalRelation3, err := strconv.ParseFloat(args[2], 64)

		if err != nil {
			log.Fatalln("Error: This is not a valid number.")
		}

		 valueX := numericalRelation * numericalRelation2 / numericalRelation3
		 fmt.Printf("The value of X is %.f \n", valueX)
	},
}