package main

import (
	"fmt"
)

type Employee struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var pers1 Employee
	var pers2 Employee
	var pers3 Employee

	// Pers1 speciications
	pers1.name = "Yasin"
	pers1.age = 20
	pers1.job = "Pastry Chef"
	pers1.salary = 2000

	// Pers2 specifications
	pers2.name = "Cigdem"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary = 2500

	// Pers3 specifications
	pers3.name = "Sat"
	pers3.age = 32
	pers3.job = "Executive Chef"
	pers3.salary = 3500

	// Print Pers1 info by calling a function
	printPerson(pers1)

	// Print Pers2 info by calling a function
	printPerson(pers2)

	// Print Pers3 info by calling a function
	printPerson(pers3)
}

func printPerson(pers Employee) {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
	fmt.Println("-------------------------")
}
