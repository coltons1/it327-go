/*
Copyright © 2026 Colton Stanek cdstan1@ilstu.edu
*/
package cmd

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// rmExpenseCmd represents the rmExpense command
var rmExpenseCmd = &cobra.Command{
	Use:   "rmExpense",
	Short: "Remove one of your current recorded expenses.",
	Args: func(cmd *cobra.Command, args []string) error {
		numOfArgs := len(args)

		// Expect exactly one argument
		if numOfArgs == 0 {
			return errors.New("missing argument")
		}

		if numOfArgs > 1 {
			return errors.New("too many arguments provided")
		}

		// args[0] exists safely here
		if args[0] == "" {
			return errors.New("argument cannot be empty")
		}

		// Check if argument is a valid integer ID
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.New("id must be a valid integer")
		}

		// validation
		if id <= 0 {
			return errors.New("id must be a positive integer")
		}

		// No problems
		return nil

	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rmExpense called")
		//THIS IS WHERE THE BODY OF THE FUNCTION WILL BE
		file, err := os.Open("expenses.csv")
		if err != nil {
			fmt.Print("could not open file.")
			return
		}

		defer file.Close()

		filedata := csv.NewReader(file)

		tempFile, err2 := os.Create("temp.csv")
		if err2 != nil {
			return
		}
		defer tempFile.Close()

		writer := csv.NewWriter(tempFile)
		idColumnIndex := 0
		for {
			record, err := filedata.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return
			}

			// Only write the row if the ID does not match
			if record[idColumnIndex] != args[0] {
				if err := writer.Write(record); err != nil {
					return
				}
			}
		}

		writer.Flush()
		file.Close()
		tempFile.Close()

		if err := os.Remove("expenses.csv"); err != nil {
			fmt.Println("failed to remove original file:", err)
			return
		}

		if err := os.Rename("temp.csv", "expenses.csv"); err != nil {
			fmt.Println("rename failed:", err)
			return
		}

		fmt.Println("Expense removed successfully")
		listExpensesCmd.Run(cmd, []string{})
	},
}

func init() {
	rootCmd.AddCommand(rmExpenseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmExpenseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmExpenseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
