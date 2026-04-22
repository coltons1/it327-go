//Author: Lucas Paul
//Date: 4/9/26

package main

import (
	"bufio"   //used for reading user input
	"fmt"     //used for printing to console
	"math"    //used for power function
	"os"      //used for standard input
	"strconv" //used to convert string to float
	"strings" //used to clean input (remove newline)
)

func main() {
	//create a reader to take input from the user
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Go Calculator")
	fmt.Println("---------------------------")

	//loop continuously until user exits
	for {
		fmt.Println("\nEnter a calculation (or type 'exit' to quit)")

		//read first number as string
		num1Str := readLine(reader, "Enter first number: ")

		//check if user wants to exit
		if num1Str == "exit" {
			fmt.Println("Goodbye!")
			return
		}

		//convert first input to float
		num1, err := strconv.ParseFloat(num1Str, 64)
		if err != nil {
			fmt.Println("Error: first input is not a valid number")
			continue
		}

		//read operator from user
		operator := readLine(reader, "Enter operator (+, -, *, /, %, ^): ")

		//validate operator
		if !isValidOperator(operator) {
			fmt.Println("Error: invalid operator")
			continue
		}

		//read second number
		num2Str := readLine(reader, "Enter second number: ")

		//convert second input to float
		num2, err := strconv.ParseFloat(num2Str, 64)
		if err != nil {
			fmt.Println("Error: second input is not a valid number")
			continue
		}

		//perform calculation
		result, errMsg := calculate(num1, num2, operator)

		//handle calculation errors (like divide by zero)
		if errMsg != "" {
			fmt.Println("Error:", errMsg)
			continue
		}

		//print result
		fmt.Printf("Result: %.2f\n", result)
	}
}

// performs the calculation based on operator the user inputs
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
		// use math.Pow for exponentiation
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

// reads a line of input and removes extra whitespace/newlines
func readLine(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
