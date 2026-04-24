package main

func main() {

}

type task struct {
	taskName     string
	taskPriority int
	taskLength   int
	index        int
}

type PriorityQueue []*task

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
