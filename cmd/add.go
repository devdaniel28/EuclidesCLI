package cmd

func init() {
	rootCmd.AddCommand(squareRootCmd)
	rootCmd.AddCommand(exponentiation)
	rootCmd.AddCommand(percentage)
	rootCmd.AddCommand(simpleRuleOfThree)
}