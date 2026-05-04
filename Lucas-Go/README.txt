README - Go Calculator
Author: Lucas Paul
Date: 4/9/26

Description:
This program is a command-line calculator written in Go. It allows the user to enter two numbers and an operator to perform calculations. The program runs continuously until the user types "exit". It also includes a queue-based history feature that stores the most recent calculations (up to 5) and allows the user to view them by typing "history".

Supported Operators:
+   addition
-   subtraction
*   multiplication
/   division
%   modulus
^   exponent/power

Features:
- Continuous execution until user exits
- Input validation for numbers and operators
- Error handling (division/modulus by zero)
- Queue data structure to store recent calculations
- "history" command to display recent calculations

How to Run:
1. Make sure Go is installed.
2. Open a terminal in the folder containing the program.
3. Run:

go run main.go

How to Use:
1. Enter the first number (or type "exit" to quit, "history" to view recent calculations).
2. Enter an operator.
3. Enter the second number.
4. The program prints the result.
5. Repeat as needed.

Commands:
- exit     → ends the program
- history  → shows the last 5 calculations

Example:
Go Calculator
---------------------------
Enter a calculation
Type 'exit' to quit
Type 'history' to view recent calculations

Enter first number: 5
Enter operator (+, -, *, /, %, ^): *
Enter second number: 3
Result: 15.00

Recent Calculations:
5.00 * 3.00 = 15.00

Error Handling:
The program checks for:
- Invalid numeric input
- Invalid operators
- Division by zero
- Modulus by zero
