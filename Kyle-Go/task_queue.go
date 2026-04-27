package main

import "fmt"

func main() {
	pqueue := PriorityQueue{}
	task1 := newTask("Laundry", 1, 10)
	task2 := newTask("Sweeping", 4, 20)
	task3 := newTask("Get Gas", 5, 15)
	pqueue.Push(task1)
	pqueue.Push(task2)
	pqueue.Push(task3)
	fmt.Println(pqueue)
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

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(index1, index2 int) bool {
	return pq[index1].taskPriority < pq[index2].taskPriority
}

func (pq PriorityQueue) Swap(index1, index2 int) {
	pq[index1], pq[index2] = pq[index2], pq[index1]
	pq[index1].index = index1
	pq[index2].index = index2
}

func (pq *PriorityQueue) Push(t any) {
	task := t.(*task)
	task.index = len(*pq)
	*pq = append(*pq, task)
}

func (pq *PriorityQueue) Pop() any {
	lastTask := (*pq)[len(*pq)-1]
	(*pq)[len(*pq)-1] = nil
	*pq = (*pq)[0 : len(*pq)-1]
	return lastTask
}
