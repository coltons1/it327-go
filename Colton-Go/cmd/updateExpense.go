/*
Copyright © 2026 Colton Stanek cdstan1@ilstu.edu
*/
package cmd

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// updateExpenseCmd represents the updateExpense command
var updateExpenseCmd = &cobra.Command{
	Use:   "updateExpense",
	Short: "Update an existing expense.",
	Long:  "updateExpense [id] [newdesc] [newcost],  updateExpense [id] [newdesc], or updateExpense [id] [newcost]",
	Args: func(cmd *cobra.Command, args []string) error {
		numOfArgs := len(args)

		if numOfArgs <= 1 {
			return errors.New("Not enough arguments provided.")
		} else if numOfArgs == 2 {
			// this case there is the correct amount but need to check if the second arg is description or cost
			selecID, err := strconv.ParseInt(args[0], 10, 32)
			if err != nil {
				// if there is an error
				return errors.New("Could not parse the selected ID.")
			} else {
				fmt.Printf("Arg[0] parsed as %v", selecID)
			}

			// need to check if the second arg is able to be parsed as a float or not
			newCost, err := strconv.ParseFloat(args[1], 32)
			if err != nil {
				// is likley a description then.
			} else {
				fmt.Printf("Arg[1] parsed as %v", newCost)
			}

			// if args[1] case fails use it as a description.

		} else if numOfArgs == 3 {
			// this case is total update.
			// check that the first is an float, it not error.
			selecID, err := strconv.ParseInt(args[0], 10, 32)
			if err != nil {
				// if there is an error
				return errors.New("Could not parse the selected ID.")
			} else {
				fmt.Printf("Arg[0] parsed as %v", selecID)
			}

			newCost, err := strconv.ParseFloat(args[1], 32)
			if err != nil {
				// is likley a description then.
			} else {
				fmt.Printf("Arg[1] parsed as %v", newCost)
			}

			// if all pass, we good.

		} else if numOfArgs >= 4 {
			return errors.New("Too many arguments provided.")
		}

		return nil
		// command layout would be ./cli updateExpense [id] [newdesc] [newcost]
		// or [id] [newdesc]
		// or [id] [newcost]
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("updateExpense called")

		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal("Couldn't access file.")
		}
		writer := csv.NewWriter(file)
		filedata := csv.NewReader(file)
		if err != nil {
			fmt.Println("Couldn't read file.")
		}

		for {
			record, err := filedata.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return
			}

			// Only write the row if the ID matches
			if record[0] == args[0] {
				// replace the record..
				_, err := strconv.ParseFloat(args[1], 32)
				if err != nil {
					// is likley a description then.
					record[1] = args[1] // description
				} else {
					newCost, _ := strconv.ParseFloat(args[2], 32)
					record[2] = strconv.FormatFloat(newCost, 'f', 2, 32) // cost
				}
			}
			writer.Write(record)
			writer.Flush()
		}
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
