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
	Args: func(cmd *cobra.Command, args []string) error {
		// numOfArgs := len(args)

		// command layout would be ./cli updateExpense [id] [newdesc] [newcost]
		// or [id] [newdesc]
		// or [id] [newcost]
		return fmt.Errorf("ball out")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("updateExpense called")

		// need to take in some parameters, need ID, what item they want to edit, and what they want to change.

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
