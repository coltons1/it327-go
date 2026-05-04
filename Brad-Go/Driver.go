package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Driver struct {
	hotelSlice []*Hotel
}

/*
GenerateData method that concurrentyly makes Hotel structs and puts them in the hotel slice.
*/
func (driver *Driver) GenerateData() {
	//Slice literal of strings of hotel names.
	names := []string{
		"Ritz Carlton, Chicago\t",
		"Four Seasons, Chicago\t",
		"Beijing Hotel\t\t",
		"JW Marriot, Beijing\t",
		"Mandarin Oriental, Barcelona",
		"The Peninsula, Chicago\t",
		"Shangri La, Beijing\t",
		"Waldorf Astoria, Chicago",
	}

	//Allocate the slice of hotel pointers that are nil pointers for now, so the goroutine can write them in.
	driver.hotelSlice = make([]*Hotel, len(names))

	//Declare a WaitGroup, which acts as a boss/coordinator for the goroutines.
	var wg sync.WaitGroup

	//Iterate over the name slice.
	for i, name := range names {

		//Increment WaitGroup by 1 before starting goroutineto tell wg.Wait() that a goroutine isn't finished yet.
		wg.Add(1)
		go func(index int, hotelName string) {
			//Decrememnt WaitGroup once goruotine is finsihed to tell wg.wait() that the goroutine is finished.
			defer wg.Done()

			//Create a new Hotel and insert it into the slice index.
			driver.hotelSlice[index] = NewHotel(hotelName)
		}(i, name)
	}

	//GenerateData method wont be able to stop until WaitGroup reaches 0.
	wg.Wait()

}

/*
DisplayBanner method that prints the year banner and divider.
*/
func DisplayBanner() {
	fmt.Println("Hotel\t\t\t\t2025\t2024\t2023\t2022\t2021\t2020\t2019\t2018\t2017\t2016")
	fmt.Println("----------------------------------------------------------------------------------------------------------------")
}

/*
DisplayAllRatings method that concurrently builds fomratted hotel rating strings and prints them.
*/
func (driver *Driver) DisplayAllRatings() {

	fmt.Println("All ratings:")
	DisplayBanner()

	//Allocate the string slice of hotel info.
	results := make([]string, len(driver.hotelSlice))

	//Declare a WaitGroup, which acts as a boss/coordinator for the goroutines.
	var wg sync.WaitGroup

	for i, hotel := range driver.hotelSlice {

		//Incrememnt WaitGroup.
		wg.Add(1)
		//Both index and h are passed here so goroutines dont share the same i and hotel values.
		go func(index int, h *Hotel) {
			defer wg.Done()
			results[index] = h.ToString()
		}(i, hotel)
	}
	//Dont advance until all goroutines are finsihed.
	wg.Wait()

	for _, line := range results {
		fmt.Println(line)
	}
}

/*
FindBestRating concurrent function that finds the hotels with the best rating using channels.
*/
func (driver *Driver) FindBestRating() {

	//Make a channel of ints with a slot for each hotel.
	results := make(chan int, len(driver.hotelSlice))

	var wg sync.WaitGroup

	for i, hotel := range driver.hotelSlice {
		wg.Add(1)
		go func(index int, h *Hotel) {
			defer wg.Done()
			for _, rating := range h.ratings {

				//Send hotel's index into the channel and return it.
				if rating >= 95 {
					results <- index
					return
				}
			}
		}(i, hotel)
	}

	//Launch goroutine that waits for other goroutines to finish.
	//(Its a goroutine so the main flow is not blocked)
	go func() {
		wg.Wait()
		//Close the channels and stop blocking.
		close(results)
	}()

	highRating := make([]bool, len(driver.hotelSlice))
	highRatingFound := false
	for index := range results {
		highRating[index] = true
		highRatingFound = true
	}

	fmt.Println("\nThe Best Hotels:")
	if !highRatingFound {
		fmt.Println("Sorry, there's no hotels that recieved a 95 or above rating.")
		return
	}

	DisplayBanner()
	for index, isHigh := range highRating {
		if isHigh {
			fmt.Println(driver.hotelSlice[index].ToString())
		}
	}

}

/*
FindWorstRating method that finds the worst rated hotel sequentially.
*/
func (driver *Driver) FindWorstRating() {

	//Set to 100 because a rating of 100 should be checked aswell.
	lowestRating := 101
	hotelLowest := 0
	numLowRatings := 0
	for i, hotel := range driver.hotelSlice {
		//Dont need the index of rating, just the rating value.
		for _, rating := range hotel.ratings {
			//Skip a rating of 0 since the hotel was not open/operating during that time.
			if rating == 0 {
				//Skips everything below.
				continue
			}
			if rating < lowestRating {
				lowestRating = rating
				hotelLowest = i
				numLowRatings = 1
			} else if rating == lowestRating && i == hotelLowest {
				numLowRatings++
			}
		}
	}

	fmt.Printf("\nThe Worst Hotel:\nThe following hotel has recieved the worst rating of %v, %v time(s).\n", lowestRating, numLowRatings)
	DisplayBanner()
	fmt.Println(driver.hotelSlice[hotelLowest].ToString())

}

type Hotel struct {
	name    string
	ratings []int
}

/*
NewHotel method that makes and returns a new Hotel struct with random ratings slice.
*/
func NewHotel(name string) *Hotel {

	//Allocate a int slice for 10 years.
	ratingsTemp := make([]int, 10)

	//rand.IntN(6)+5 -> Random number between 5 and 10.
	//This makes it so the most years a hotel can be unavailible for early on is 4 years, which keeps rating at 0. (N/A rating)
	for i := 0; i < rand.IntN(6)+5; i++ {
		ratingsTemp[i] = rand.IntN(100) + 1
	}

	//Make and return the pointer to the hotel struct with the name and new ratings slice.
	//Returning a pointer so whoever calls NewHotel can reference and modify the same Hotel in memory.
	return &Hotel{name: name, ratings: ratingsTemp}
}

/*
ToString method that builds a string for a hotels name and ratings sequentially.
*/
func (hotel *Hotel) ToString() string {
	var strBuilder strings.Builder

	//Pre-allocates space for the string builder since we know how long the string will be.
	strBuilder.Grow(64)
	strBuilder.WriteString(hotel.name + "\t")
	for index := range hotel.ratings {
		if hotel.ratings[index] == 0 {
			strBuilder.WriteString("N/A\t")
		} else {
			//Convert the int rating to a string using strconv.Itoa.
			strBuilder.WriteString(strconv.Itoa(hotel.ratings[index]) + "\t")
		}

	}
	return strBuilder.String()
}

func main() {
	//Start recording time.
	start := time.Now()

	//Define an anonomuys function that executes when main returns.
	defer func() { fmt.Printf("Program took %s to run. \n", time.Since(start)) }()

	var driver Driver
	driver.GenerateData()
	driver.DisplayAllRatings()
	driver.FindBestRating()
	driver.FindWorstRating()
}
