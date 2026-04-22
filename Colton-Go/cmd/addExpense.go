/*
Copyright © 2026 Colton Stanek cdstan1@ilstu.edu
*/
package cmd

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// addExpenseCmd represents the addExpense command
var addExpenseCmd = &cobra.Command{
	Use:   "addExpense",
	Short: "Add a new expense to your tracker. | addExpense [description] [expense]",
	Args: func(cmd *cobra.Command, args []string) error {
		var numOfArgs int = len(args)
		if numOfArgs <= 1 {
			//if the user enters 0-1 arguments, return error for not enough args
			return errors.New("Not enough arguments provided.")
		} else if numOfArgs > 2 {
			//if the user enters too many args, return error for too many args
			return errors.New("Too many arguments provided.")
		} else if args[0] != "" && args[1] != "" {
			//check if args[1] can be turned into float.
			f, err := strconv.ParseFloat(args[1], 32)
			if err != nil {
				fmt.Printf("Argument 2 parsed as: %v\n", f)
				return errors.New("Price cannot be recognized. Is it your second argument?")
			}
		}
		//no problems
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		Setup()

		//add a new expense with a description, price, and date.
		type Expense struct {
			expID       int64
			description string
			price       float32
			date        time.Time
		}

		//variable to access struct
		expenseCount = readExpenseCount(keyValFile)
		var exp Expense = Expense{expID: expenseCount, description: "", price: 0.0, date: time.Now()}
		exp.expID = expenseCount + 1
		fmt.Println(exp.description)

		//use positional arguments to take in description and price.
		var inDesc string = args[0]

		f, err := strconv.ParseFloat(args[1], 32)
		if err != nil {
			fmt.Println("error converting..", err)
			return
		}

		//converting type to float32
		var inPrice float32 = float32(f)
		exp.description = inDesc
		exp.price = inPrice
		exp.date = time.Now()

		fmt.Printf("ID is %v, Description is %v, Price is %v, Date is %v\n", exp.expID, exp.description, exp.price, exp.date)

		//need to write to expenses file.
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal("Couldn't access file.")
		}

		var record []string = []string{strconv.FormatInt(int64(exp.expID), 10), exp.description, strconv.FormatFloat(float64(exp.price), 'f', 2, 32), exp.date.String()}
		writer := csv.NewWriter(file)

		writeExpenseCount(keyValFile, exp.expID)

		if err := writer.Write(record); err != nil {
			log.Fatal("Couldn't write to file:", err)
		}

		writer.Flush()
		if err := writer.Error(); err != nil {
			log.Fatal("Flush failed:", err)
		}

		fmt.Println("Wrote to file.")

	},
}

func readExpenseCount(path string) int64 {
	file, err := os.Open(path)
	if err != nil {
		return 0
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil || len(records) == 0 || len(records[0]) == 0 {
		return 0
	}

	count, err := strconv.ParseInt(records[0][0], 10, 32)
	if err != nil {
		return 0
	}

	return count
}

func writeExpenseCount(path string, count int64) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Write([]string{strconv.FormatInt(count, 10)})
	writer.Flush()
}

func init() {
	rootCmd.AddCommand(addExpenseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addExpenseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addExpenseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
