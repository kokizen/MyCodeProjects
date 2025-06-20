package main

import "fmt"

// Define a struct for a student
type Student struct {
	Name   string
	School string
	Year   int
}

func main() {
	// Create a new student
	student1 := Student{
		Name:   "Alex",
		School: "Founders High School",
		Year:   2024,
	}

	student2 := Student{
		Name:   "Kobi",
		School: "01 Founders",
		Year:   2025,
	}
	student3 := Student{
		Name:   "Franck",
		School: "01 Founders",
		Year:   2025,
	}
	// Print student info
	fmt.Println("Student Information:")
	fmt.Println("Name:", student1.Name)
	fmt.Println("School:", student1.School)
	fmt.Println("Year:", student1.Year)
	// Print student info
	fmt.Println("Student Information:")
	fmt.Println("Name:", student2.Name)
	fmt.Println("School:", student2.School)
	fmt.Println("Year:", student2.Year)
	// Print student info
	fmt.Println("Student Information:")
	fmt.Println("Name:", student3.Name)
	fmt.Println("School:", student3.School)
	fmt.Println("Year:", student3.Year)
}
