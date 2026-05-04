//Author: Lucas Paul
//Date: 4/9/26

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Struct to store a single calculation
type Calculation struct {
	num1   float64 // first number
	num2   float64 // second number
	op     string  // operator (+, -, etc.)
	result float64 // result of calculation
}

func main() {
	//reate reader to take input from user
	reader := bufio.NewReader(os.Stdin)

	//slice used as a queue to store recent calculations
	var recentCalculations []Calculation

	//maximum number of calculations stored in queue
	const maxQueueSize = 5

	fmt.Println("Go Calculator")
	fmt.Println("---------------------------")

	//loop continuously until user exits
	for {
		fmt.Println("\nEnter a calculation")
		fmt.Println("Type 'exit' to quit")
		fmt.Println("Type 'history' to view recent calculations")

		//read first input
		num1Str := readLine(reader, "Enter first number: ")

		//exit condition
		if num1Str == "exit" {
			fmt.Println("Exiting program...")
			return
		}
		// show history if user types "history"
		if num1Str == "history" {
			printHistory(recentCalculations)
			continue
		}
		// convert first input to float
		num1, err := strconv.ParseFloat(num1Str, 64)
		if err != nil {
			fmt.Println("Error: first input is not a valid number")
			continue
		}
		// read operator
		operator := readLine(reader, "Enter operator (+, -, *, /, %, ^): ")
		// validate operator
		if !isValidOperator(operator) {
			fmt.Println("Error: invalid operator")
			continue
		}

		// read second number
		num2Str := readLine(reader, "Enter second number: ")

		// convert second input to float
		num2, err := strconv.ParseFloat(num2Str, 64)
		if err != nil {
			fmt.Println("Error: second input is not a valid number")
			continue
		}

		// perform calculation
		result, errMsg := calculate(num1, num2, operator)

		// handle errors (ex: divide by zero)
		if errMsg != "" {
			fmt.Println("Error:", errMsg)
			continue
		}

		// print result
		fmt.Printf("Result: %.2f\n", result)

		// add new calculation to queue (slice)
		recentCalculations = append(recentCalculations, Calculation{
			num1:   num1,
			num2:   num2,
			op:     operator,
			result: result,
		})

		// if queue exceeds max size, remove oldest element (FIFO behavior)
		if len(recentCalculations) > maxQueueSize {
			recentCalculations = recentCalculations[1:]
		}
	}
}

// performs calculation based on operator
func calculate(a, b float64, op string) (float64, string) {
	switch op {
	case "+":
		return a + b, ""
	case "-":
		return a - b, ""
	case "*":
		return a * b, ""
	case "/":
		// check for division by zero
		if b == 0 {
			return 0, "cannot divide by zero"
		}
		return a / b, ""
	case "%":
		// check for modulus by zero
		if b == 0 {
			return 0, "cannot mod by zero"
		}
		return float64(int(a) % int(b)), ""
	case "^":
		// exponentiation
		return math.Pow(a, b), ""
	default:
		return 0, "invalid operator"
	}
}

// checks if the operator is valid
func isValidOperator(op string) bool {
	switch op {
	case "+", "-", "*", "/", "%", "^":
		return true
	default:
		return false
	}
}

// reads input from user and removes extra whitespace/newlines
func readLine(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// prints the recent calculation history (queue)
func printHistory(queue []Calculation) {
	// check if queue is emptyg
	if len(queue) == 0 {
		fmt.Println("No recent calculations.")
		return
	}

	fmt.Println("\nRecent Calculations:")

	// loop through queue and print each calculation
	for _, calc := range queue {
		fmt.Printf("%.2f %s %.2f = %.2f\n",
			calc.num1, calc.op, calc.num2, calc.result)
	}
}
