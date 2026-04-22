package main

import (
	"fmt"
	"sync"
	"time"
)

/*----------------------+-----------------------*/
// Beginner Go Program 							//
// Colton, Kyle, Kirby, Brad, Lucas				//
// 2 April 2026									//
/*----------------------+-----------------------*/

// Program demonstrating concurrency in Go.

var m = sync.Mutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5", "id6", "id7", "id8", "id9", "id10"}
var results = []string{}

func main() {
	fmt.Println("\t-----------+ Non-Concurrent Test +-----------")
	nonconcMain()

	results = nil
	results = []string{}

	fmt.Println("\n\t-------------+ Concurrent Test +-------------")
	concMain()

}

func dbCall(i int) {

	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is: ", dbData[i])
	m.Lock()
	results = append(results, dbData[i])
	m.Unlock()
	wg.Done()
}

func concMain() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal Execution Time: %v", time.Since(t0))
	fmt.Printf("\nThe results are %v", results)

}

func nonconcMain() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		nonConcdbCall(i)
	}
	fmt.Printf("\nTotal Execution Time: %v", time.Since(t0))
	fmt.Printf("\nThe results are %v", results)
}

func nonConcdbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is: ", dbData[i])
	results = append(results, dbData[i])
}
