package cmd

func init() {
	rootCmd.AddCommand(squareRootCmd)
	rootCmd.AddCommand(exponentiation)
	rootCmd.AddCommand(percentage)
	rootCmd.AddCommand(simpleRuleOfThree)
	rootCmd.AddCommand(simpleInterest)

	simpleInterest.Flags().Float64VarP(&amount, "amount", "a", 0, "Amount value")
}