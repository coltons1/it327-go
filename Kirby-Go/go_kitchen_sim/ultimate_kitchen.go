package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/manifoldco/promptui"
)

type meal struct {
	MealName string
	CookTime time.Duration
}

// Constructor function for the meal struct
func createMeal(mName string, cookingTime time.Duration) meal {
	newMeal := meal{MealName: mName, CookTime: cookingTime}
	return newMeal
}

func main() {
	//Create a Wait Group to manage goroutines
	var waitGroup sync.WaitGroup
	//Create a Mutex that can be used to lock the output so that it fixes the race condition
	var printStabilizer sync.Mutex

	//Create the menu
	myMenu := buildMenu()

	//Create order slice to keep track of selected orders
	order := []meal{}

	//Testing TUI from promptui (example taken from and
	// modified from promptui's Github page github.com/manifoldco/promptui)
	menuTemplate := &promptui.SelectTemplates{
		Active:   "\u25B6 {{ printf \"%v - %v\" .MealName .CookTime}}",
		Inactive: " {{ printf \"%v - %v\" .MealName .CookTime}}",
		Selected: "\u2713 {{ printf \"You Chose %v, which takes %v seconds\"\n .MealName .CookTime}}",
	}

	for {
		prompt := promptui.Select{
			Label:     "Choose the Menu Item",
			Items:     myMenu,
			Templates: menuTemplate,
			Size:      12,
		}

		index, _, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		if myMenu[index].MealName == "Finish Selection" {
			break
		}

		order = append(order, myMenu[index])
	}

	//Print the order choices
	printOrder(order)

	//Open a channel that only allows three threads through,
	// in order to simulate only making three orders at a time.
	limitChannel := make(chan struct{}, 3)

	startTime := time.Now()

	for _, value := range order {
		waitGroup.Add(1)
		go cookOrder(value, &waitGroup, limitChannel, startTime, &printStabilizer)
	}

	waitGroup.Wait()
	fmt.Println("All orders are finished!")

}

func buildMenu() []meal {
	spaghetti := createMeal("Spaghetti", time.Second*12)
	pizza := createMeal("Pizza", time.Second*10)
	burger := createMeal("Hamburger", time.Second*8)
	taco := createMeal("Taco", time.Second*5)
	soup := createMeal("Soup", time.Second*15)
	salad := createMeal("Salad", time.Second*4)
	tiramisu := createMeal("Tiramisu", time.Second*18)
	falafel := createMeal("Falafel", time.Second*10)
	alligator := createMeal("Alligator", time.Second*8)
	chiliCheeseDog := createMeal("Chili Cheese Dog", time.Second*4)
	footCake := createMeal("Foot Cake", time.Second*10)
	breakCond := createMeal("Finish Selection", time.Second*0)

	menu := []meal{spaghetti, pizza, burger, taco, soup, salad, tiramisu, falafel, alligator, chiliCheeseDog, footCake, breakCond}
	return menu
}

// print function for the full menu
func printMenu(mealToPrint []meal) {
	for index, mealItem := range mealToPrint {
		fmt.Printf("   %d.) %v\n", index+1, mealItem.MealName)
	}
}

func printOrder(orderToPrint []meal) {
	//Print escape codes to clear console and move cursor to top of the output for clarity
	fmt.Printf("\033[2J\033[H")
	fmt.Printf("    You ordered: \n")
	for index, orderItem := range orderToPrint {
		fmt.Printf("    %d.) %v\n", index+1, orderItem.MealName)
	}
}

func cookOrder(orderItem meal, wg *sync.WaitGroup, channel chan struct{}, origStartTime time.Time, printFixer *sync.Mutex) {
	//Ensure that the program waits to end before all of this currently running function has finished
	defer wg.Done()
	channel <- struct{}{}

	//Lock output to fix bad output
	printFixer.Lock()

	//Fill channel to 'queue' up the orders
	fmt.Printf("The order '%v' is now cooking\n\n", orderItem.MealName)

	//Unlock output to continue
	printFixer.Unlock()

	//Wait to simulate the order being made
	time.Sleep(orderItem.CookTime)

	printFixer.Lock()

	fmt.Printf("-------------------------------------------------------------\n")
	//Record the time the order took from the beginning of the program
	fmt.Printf("    The %v finished cooking after %v\n", orderItem.MealName, time.Since(origStartTime).Round(time.Second))
	fmt.Printf("-------------------------------------------------------------\n\n")

	printFixer.Unlock()

	//Free up the channel by discarding the current thing (a struct in this case) once the loop is done
	// since the loop being done signifies the order being done and no longer needing to
	// to hold up the channel
	<-channel
}

// func firstfunc(numRuns int, wg *sync.WaitGroup) {
// 	fmt.Println("Printing i values:")
// 	fmt.Printf("The current value of i is: ")
// 	var max int = 0
// 	for i := 0; i < numRuns; i++ {
// 		//fmt.Printf("%d ", i)
// 		//Add together the current i value to an accumulator
// 		max += i
// 	}

// 	fmt.Println(max)
// 	defer wg.Done()
// }
