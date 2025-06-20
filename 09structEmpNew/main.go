package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	var pers4 Employee

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

	// Pers4 specifications
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the name of the employee: ")
	nameInput, _ := reader.ReadString('\n')
	nameInput = strings.TrimSpace(nameInput)
	fmt.Print("Enter the age of the employee: ")
	ageInput, _ := reader.ReadString('\n')
	ageInput = strings.TrimSpace(ageInput)
	fmt.Print("Enter the job of the employee: ")
	jobInput, _ := reader.ReadString('\n')
	jobInput = strings.TrimSpace(jobInput)
	fmt.Print("Enter the salary of the employee: ")
	salaryInput, _ := reader.ReadString('\n')
	salaryInput = strings.TrimSpace(salaryInput)

	// Convert age and salary to int [ANSWER FROM COPILOT]
	age, _ := strconv.Atoi(ageInput)
	salary, _ := strconv.Atoi(salaryInput)

	// Assign values to Pers4 [ANSWER FROM COPILOT]
	pers4.name = nameInput
	pers4.age = age
	pers4.job = jobInput
	pers4.salary = salary

	// Print Pers1 info by calling a function
	printPerson(pers1)

	// Print Pers2 info by calling a function
	printPerson(pers2)

	// Print Pers3 info by calling a function
	printPerson(pers3)

	// Print Pers4 info by calling a function
	printPerson(pers4)

}

func printPerson(pers Employee) {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
	fmt.Println("-------------------------")
}
