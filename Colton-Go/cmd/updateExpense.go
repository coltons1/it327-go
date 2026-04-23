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

		if numOfArgs < 2 {
			return errors.New("Not enough arguments provided.")
		}
		if numOfArgs > 3 {
			return errors.New("Too many arguments provided: expected at most 3.")
		}

		if _, err := strconv.ParseInt(args[0], 10, 32); err != nil {
			return errors.New("first argument (id) needs to be a valid integer")
		}

		if numOfArgs == 3 {
			if _, err := strconv.ParseFloat(args[2], 64); err != nil {
				return fmt.Errorf("third argument must be a valid number.")
			}
		}
		return nil
		// command layout would be ./cli updateExpense [id] [newdesc] [newcost]
		// or [id] [newdesc]
		// or [id] [newcost]
	},
	Run: func(cmd *cobra.Command, args []string) {
		targetID := args[0]
		numOfArgs := len(args)

		file, err := os.Open(filePath)
		if err != nil {
			log.Fatalf("Couldn't open file: %v", err)
		}
		defer file.Close()

		reader := csv.NewReader(file)
		var records [][]string
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error reading file: %v", err)
			}
			records = append(records, record)
		}

		updated := false

		for _, record := range records {
			if len(record) < 3 {
				continue
			}
			if record[0] == targetID {
				switch numOfArgs {
				case 2:
					if newCost, err := strconv.ParseFloat(args[1], 64); err == nil {
						record[2] = strconv.FormatFloat(newCost, 'f', 2, 64)
					} else {
						record[1] = args[1]
					}

				case 3:
					record[1] = args[1]
					newCost, _ := strconv.ParseFloat(args[2], 64)
					record[2] = strconv.FormatFloat(newCost, 'f', 2, 64)
				}
				updated = true
			}

		}

		if !updated {
			fmt.Printf("No expense found with id %q\n", targetID)
			return
		}

		temp, err := os.CreateTemp("", "expenses-*.csv")
		if err != nil {
			log.Fatalf("Could not create temp file %v", err)
		}
		defer temp.Close()

		writer := csv.NewWriter(temp)
		if err := writer.WriteAll(records); err != nil {
			log.Fatalf("Error writing records: %v", err)
		}
		writer.Flush()
		if err := writer.Error(); err != nil {
			log.Fatalf("Error flushing writer %v", err)
		}
		temp.Close()

		input, err := os.ReadFile(temp.Name())
		if err != nil {
			log.Fatalf("Couldn't read temp file : %v", err)
		}
		if err := os.WriteFile(filePath, input, 0644); err != nil {
			log.Fatalf("Couldn't write to original file %v", err)
		}
		if err := os.Remove(temp.Name()); err != nil {
			log.Printf("warning: couldn't remove temp file %s : %v", temp.Name(), err)
		}

		fmt.Printf("Expense %q updated successfully.\n", targetID)

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
