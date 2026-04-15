/*
Copyright © 2026 Colton Stanek cdstan1@ilstu.edu
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updateExpenseCmd represents the updateExpense command
var updateExpenseCmd = &cobra.Command{
	Use:   "updateExpense",
	Short: "Update an existing expense.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("updateExpense called")
		//THIS IS WHERE THE BODY OF THE FUNCTION WILL BE
		fmt.Println(filedata)
	},
}

func init() {
	rootCmd.AddCommand(updateExpenseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateExpenseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateExpenseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
