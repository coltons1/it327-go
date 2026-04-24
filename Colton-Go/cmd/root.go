/*
Copyright © 2026 Colton Stanek cdstan1@ilstu.edu
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "it327go-personal.git",
	Short: "A CLI tool to track your expenses.",
	Long: `This CLI program helps you track personal expenses by letting you create, view, update, and delete entries that include a description, price, and date. 
	You can list all expenses or filter specific fields using command‑line flags, and quickly calculate total spending or costs within custom time ranges..`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// declaring for global use
var filedata [][]string
var expenseCount int64

const filePath string = "expenses.csv"
const keyValFile string = "keyVals.csv"

func Setup() {
	//check if there is a already an existing .csv file to track the expenses
	//if so, do nothing
	//if there is not one, create one.

	//attempt to open the file at the filepath, if the file does not exist log the error.
	file, err := os.Open(filePath)
	if err != nil {
		//fmt.Println("File doesn't exist, creating new file.")
		//handle the file not existing here
		data := []byte("")
		err = os.WriteFile(filePath, data, 0644)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		//file does exist!!
		_, err := csv.NewReader(file).ReadAll()
		if err != nil {
			log.Fatal(err)
		}
	}

	keyVals, err := os.Open(keyValFile)
	if err != nil {
		fmt.Println("File doesn't exist, creating new keyvals.")
		//handle the file not existing here
		data := []byte("0")
		err = os.WriteFile(keyValFile, data, 0644)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		//file does exist!!
		keyValData, err := csv.NewReader(keyVals).ReadAll()
		if err != nil {
			log.Fatal(err)
		} else {
			//fmt.Println("File does exist, printing key val data.")
			//fmt.Println(keyValData)
			expenseCount, err = strconv.ParseInt(keyValData[0][0], 10, 32)
			//fmt.Print("\n")
		}
	}

	// NEED TO UPDATE KEYVALS WITH THE VARIABLE
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.it327go-personal.git.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
