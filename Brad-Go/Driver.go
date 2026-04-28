package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"

	//"sync"
	"time"
)

type Driver struct {
	hotelList []*Hotel
}

func (driver *Driver) GenerateData() {
	//Preallocate slice
	driver.hotelList = make([]*Hotel, 0, 8)

	driver.hotelList = append(driver.hotelList, NewHotel("Ritz Carlton, Chicago\t"))
	driver.hotelList = append(driver.hotelList, NewHotel("Four Seasons, Chicago\t"))
	driver.hotelList = append(driver.hotelList, NewHotel("Beijing Hotel\t\t"))
	driver.hotelList = append(driver.hotelList, NewHotel("JW Marriot, Beijing\t"))
	driver.hotelList = append(driver.hotelList, NewHotel("Mandarin Oriental, Barcelona"))
	driver.hotelList = append(driver.hotelList, NewHotel("The Peninsula, Chicago\t"))
	driver.hotelList = append(driver.hotelList, NewHotel("Shangri La, Beijing\t"))
	driver.hotelList = append(driver.hotelList, NewHotel("Waldorf Astoria, Chicago"))
}

func DisplayBanner() {
	fmt.Println("Hotel\t\t\t\t2025\t2024\t2023\t2022\t2021\t2020\t2019\t2018\t2017\t2016")
	fmt.Println("----------------------------------------------------------------------------------------------------------------")
}

func (driver *Driver) DisplayAllRatings() {
	fmt.Println("All ratings:")
	DisplayBanner()

	for _, hotel := range driver.hotelList {
		fmt.Println(hotel.ToString())
	}
}

func (driver *Driver) FindBestRating() {

	highRating := make([]bool, 8)
	highRatingFound := false
	fmt.Println("\nThe Best Hotels:")
	for i, hotel := range driver.hotelList {
		for _, rating := range hotel.ratings {
			if rating >= 95 {
				highRating[i] = true
				highRatingFound = true
			}
		}

	}

	if highRatingFound {
		DisplayBanner()
	} else {
		fmt.Println("Sorry, there's no hotels that recieved a 95 or above rating.")
	}

	for index := range driver.hotelList {
		if highRating[index] == true {
			fmt.Println(driver.hotelList[index].ToString())
		}
	}

}

func (driver *Driver) FindWorstRating() {

	//Set to 100 because a rating of 100 should be checked aswell.
	lowestRating := 101
	hotelLowest := 0
	numLowRatings := 0
	for i, hotel := range driver.hotelList {
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
	fmt.Println(driver.hotelList[hotelLowest].ToString())

}

type Hotel struct {
	name    string
	ratings []int
}

func NewHotel(name string) *Hotel {
	ratingsTemp := make([]int, 10)
	for i := 0; i < rand.IntN(6)+5; i++ {
		ratingsTemp[i] = rand.IntN(100) + 1
	}
	return &Hotel{name: name, ratings: ratingsTemp}
}

func (hotel *Hotel) ToString() string {
	var strBuilder strings.Builder
	strBuilder.Grow(64)
	strBuilder.WriteString(hotel.name + "\t")
	for index := range hotel.ratings {
		if hotel.ratings[index] == 0 {
			strBuilder.WriteString("N/A\t")
		} else {
			strBuilder.WriteString(strconv.Itoa(hotel.ratings[index]) + "\t")
		}

	}
	return strBuilder.String()
}

func main() {
	start := time.Now()
	var driver Driver
	driver.GenerateData()
	driver.DisplayAllRatings()
	driver.FindBestRating()
	driver.FindWorstRating()
	elapsed := time.Since(start)
	fmt.Printf("Program took %s to run.\n", elapsed)
	os.Exit(3)

}
