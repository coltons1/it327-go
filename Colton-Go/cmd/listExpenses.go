/*
Copyright © 2026 Colton Stanek cdstan1@ilstu.edu
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

// listExpensesCmd represents the listExpenses command
var listExpensesCmd = &cobra.Command{
	Use:   "listExpenses",
	Short: "Show all expenses currently recorded.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//THIS IS WHERE THE BODY OF THE FUNCTION WILL BE
		Setup()
		//printing expenses
		file, err := os.Open("expenses.csv")
		if err != nil {
			fmt.Print("could not open file.")
			return
		}
		filedata, err := csv.NewReader(file).ReadAll()

		w := tabwriter.NewWriter(os.Stdout, 0, 4, 4, ' ', 0)
		fmt.Fprintln(w, "ID\tDESC\tPRICE\tDATE")
		for i := 0; i < len(filedata); i++ {

			raw_time := filedata[i][3]
			if i := strings.Index(raw_time, " m="); i != -1 {
				raw_time = raw_time[:i]
			}

			// now just this:  2026-04-10 17:20:23.3072645 -0500 CDT
			layout := "2006-05-15 15:04:05.9999999 -0700 MST"
			t, err := time.Parse(layout, raw_time)
			if err != nil {
				panic(err)
			}

			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", filedata[i][0], filedata[i][1], filedata[i][2], timediff.TimeDiff(t))
		}
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listExpensesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listExpensesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listExpensesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
