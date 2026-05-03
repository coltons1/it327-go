package main

import (
	"container/heap"
	"fmt"
)

func main() {
	pqueue := PriorityQueue{}
	task1 := newTask("Laundry", 1, 10)
	task2 := newTask("Sweeping", 4, 20)
	task3 := newTask("Get Gas", 5, 15)
	task4 := newTask("Make Breakfast", 2, 8)
	task5 := newTask("Make Dinner", 10, 50)
	task6 := newTask("Push Code", 8, 2)
	heap.Push(&pqueue, task1)
	heap.Push(&pqueue, task2)
	heap.Push(&pqueue, task3)
	heap.Push(&pqueue, task4)
	heap.Push(&pqueue, task5)
	heap.Push(&pqueue, task6)
	for _, value := range pqueue {
		fmt.Println(*value)
	}
}

type task struct {
	taskName     string
	taskPriority int
	taskLength   int
	index        int
}

type PriorityQueue []*task

// Creates a new task to put into the priority queue.
// The arguments are the name of the task, the priority it has, and the length it takes to complete.
// Note that the length of the task should always just be the amount of minutes it will take.
func newTask(name string, priority int, length int) *task {
	tName := task{taskName: name, taskPriority: priority, taskLength: length}
	return &tName
}

// Implements the Len function needed to get the length of the slice used.
func (pq PriorityQueue) Len() int {
	return len(pq)
}

// Implements the Less function which defines the priority needed for the priority queue to
// know whether it's a minHeap or a maxHeap
func (pq PriorityQueue) Less(index1, index2 int) bool {
	return pq[index1].taskPriority < pq[index2].taskPriority
}

// Implements the Swap function needed to swap values in the queue
func (pq PriorityQueue) Swap(index1, index2 int) {
	pq[index1], pq[index2] = pq[index2], pq[index1]
	pq[index1].index = index1
	pq[index2].index = index2
}

// Implements the Push function that will be called by the heap interface
func (pq *PriorityQueue) Push(t any) {
	task := t.(*task)
	task.index = len(*pq)
	*pq = append(*pq, task)
}

// Implements the Pop funciton that is called by the heap interface
func (pq *PriorityQueue) Pop() any {
	lastTask := (*pq)[len(*pq)-1]
	(*pq)[len(*pq)-1] = nil
	*pq = (*pq)[0 : len(*pq)-1]
	return lastTask
}
