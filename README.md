# it327-go
A group project assignment covering an unfamilliar language, Go. 

## Programs

### Basic
1. Trivial Program
   #### Simulated Concurrent Database Calls
    This program was created to highlight concurrency in Go with a simulated Database call as well as safe concurrency practices.
    The program consists of a noncurrent 'control' and a concurrent test. A database call includes a 2 second delay to partly simulate a real call.
    The only differences between the calls is the concurrent call utilizes Wait groups and the 'go' keyword to make it concurrent.
    Times are also displayed after each process finishes.


### Intermediate
1. Brad
2. Colton
    #### CLI Expense Tracker
     This program is a command line expense tracker, built using Go's extensive external package collection. Specifically, Cobra-CLI to 
     make new commands, tracking and validating arguments, and all other command line features easy to implement. This tool allows you to
     add, read, update, and remove expenses consisting of description and cost. It simply uses a csv file for storing information such as 
     expense ID, description, cost, and date created. 
3. Kyle
4. Kirby
   #### Kitchen Simulator - "Ultimate Kitchen"
    This program is run in the command line and allows the user to choose as many menu items as they would like from the scrollable list and
    then, after choosing to finish their selection, the program simulates a real kitchen, where only so many meals are able to be worked on at
    once. In the case of this program, only 3 meals are able to be made at a time, and the program will only start making the fourth meal and
    onwards once the first, second, third, and so on, meals have been completed. This program was built with showcasing Go's ability to easily
    and effectively perform concurrency and is displayed via the print statements that show when each meal order item is completed and how long
    the order item took to complete in total.
6. Lucas

### Substantive
1. Substantive Program
   - A program making use of data structures implemented by us. 
