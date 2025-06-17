package main

import "fmt"

func main() {
	var age int
	var name string
	var funFact string

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("your name is ", name)
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)
	fmt.Println("your age is ", age)
	fmt.Print("Enter a fun fact about yourself: ")
	fmt.Scanln(&funFact)
	fmt.Println("A fun fact about you is: ", funFact)
	fmt.Println("Thank you for sharing!")
}
