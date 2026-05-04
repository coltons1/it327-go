package main

import (
	"Substantial/Tree"
	"fmt"
	"sync"
)

// student struct for storing student information
type Student struct {
	name               string
	id                 int
	year               int
	major              string
	expectedGraduation string
}

// constructor fucntion to create a new student
func NewStudent(studName string, studId int, studYear int, studMajor string, studExpGrad string) *Student {
	return &Student{
		name:               studName,
		id:                 studId,
		year:               studYear,
		major:              studMajor,
		expectedGraduation: studExpGrad,
	}
}

func searchForStudent(tree *Tree.BST, studentRec map[int]*Student, id int, wg *sync.WaitGroup) {
	// program will wait until the wg is done to finish executing.
	defer wg.Done()

	// search tree for student based on given id.
	foundStudent := tree.HasKey(id)
	if foundStudent {
		student := studentRec[id]
		fmt.Printf("\nFOUND ID: %v | Name: %v | Major: %v | Year: %v | Graduates: %v\n",
			student.id, student.name, student.major, student.year, student.expectedGraduation)
	} else {
		fmt.Printf("\nNO ID: %v EXISTS\n", id)

	}

}

func main() {
	fmt.Print("-------- STUDENT ID LOOKUP SYSTEM --------")
	//create a new BST
	var tree *Tree.BST = Tree.NewTree()
	//create student objects
	var kirb *Student = NewStudent("Kirby Conrad", 15482, 3, "Comp Sci", "Fall 2027")
	var lucas *Student = NewStudent("Lucas Paul ", 47382, 4, "Comp Sci", "Spring 2026")
	var brad *Student = NewStudent("Brad Wickert", 94857, 3, "Comp Sci", "Spring 2027")
	var colton *Student = NewStudent("Colton Stanek", 19734, 2, "Comp Sci", "Spring 2028")
	var kyle *Student = NewStudent("Kyle Conrad", 39485, 3, "Comp Sci", "Spring 2027")
	var jimBob *Student = NewStudent("Jim Bob", 12345, 1, "Economics", "Fall 2030")
	var mango *Student = NewStudent("Mangolika Bhattacharya", 99999, 0, "Electrical Engineering", "Spring 2023")
	var john *Student = NewStudent("John Smith", 98765, 5, "Environmental Science", "Fall 2029")

	//slice
	students := []*Student{kirb, lucas, brad, colton, kyle, jimBob, mango, john}

	//map for student records
	studentRecords := make(map[int]*Student)

	//Insert ID's into BST and store student info in a map
	for _, student := range students {
		tree.Insert(student.id)
		studentRecords[student.id] = student
	}
	//print all student ID's in sorted order of the BST
	fmt.Println("\nStudent ID's in sorted order:")
	tree.PrintTree()

	fmt.Println("\nAll Students (Unordered): ")

	for _, student := range students {
		fmt.Println("---------")
		fmt.Printf("Name: %v, ID: %v\n", student.name, student.id)
	}

	fmt.Println("\nAll Students, Ordered on ID")

	orderedArr := tree.InorderArr()

	for _, node := range orderedArr {
		fmt.Println("---------")
		id, err := node.GetData()
		if err != nil {
			return
		}

		student := studentRecords[id]
		fmt.Printf("Name: %v, ID: %v\n", student.name, student.id)
	}

	// Concurrent Searches with real ID's and fake ID's
	targetIDs := []int{19734, 99999, 00000, 47382, 11111, 98765, 12345}

	fmt.Println("\n---------+--------| Concurrent Student Look-Up |---------+---------")
	var wg sync.WaitGroup
	for _, id := range targetIDs {
		wg.Add(1)
		go searchForStudent(tree, studentRecords, id, &wg)
	}

	wg.Wait()
	fmt.Println("\n---------+--------| All lookups completed. |---------+---------")
}
