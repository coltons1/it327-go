package main

import (
	"Substantial/Tree"
	"fmt"
)
//student struct for storing student information
type Student struct {
	name               string
	id                 int
	year               int
	major              string
	expectedGraduation string
}
//constructor fucntion to create a new student 
func NewStudent(studName string, studId int, studYear int, studMajor string, studExpGrad string) *Student {
	return &Student{
		name:               studName,
		id:                 studId,
		year:               studYear,
		major:              studMajor,
		expectedGraduation: studExpGrad,
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
	var john *Student = NewStudent("John Pork", 98765, 5, "Hamonomics", "Fall 2029")

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

	for {
		var SearchID int
		fmt.Print("\nEnter Student ID for lookup (Enter -1 to exit): ")
		fmt.Scanln(&SearchID)

		if SearchID == -1 {
			fmt.Print("\nExiting the program...")
			break
		}
		//check if ID is in BST 
		if tree.HasKey(SearchID) {
			student := studentRecords[SearchID]
			fmt.Println("\nStudent found: ")
			fmt.Println("ID:", student.id)
			fmt.Println("Name:", student.name)
			fmt.Println("Year:", student.year)
			fmt.Println("Major:", student.major)
			fmt.Println("Expected Graduation:", student.expectedGraduation)
		} else {
			fmt.Println("\nStudent not found!")
		}
	}

}
