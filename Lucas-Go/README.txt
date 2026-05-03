README - Go Calculator
Author: Lucas Paul
Date: 4/9/26

Description:
This program is a simple calculator written in Go. It allows the user to enter two numbers and choose an operator to perform a calculation. The program keeps running until the user types "exit".

Supported Operators:
+   addition
-   subtraction
*   multiplication
/   division
%   modulus
^   exponent/power

How to Run:
1. Make sure Go is installed.
2. Open a terminal in the folder containing the program.
3. Run:

go run main.go

How to Use:
1. Enter the first number.
2. Enter an operator.
3. Enter the second number.
4. The program prints the result.
5. Type "exit" when asked for the first number to quit.

Example:
Go Calculator
---------------------------
Enter a calculation (or type 'exit' to quit)
Enter first number: 5
Enter operator (+, -, *, /, %, ^): +
Enter second number: 3
Result: 8.00

Error Handling:
The program checks for invalid numbers, invalid operators, division by zero, and modulus by zero.
