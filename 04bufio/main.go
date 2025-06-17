package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Who is your favourite book author? ")
	name, _ := reader.ReadString('\n') //includes newline character
	fmt.Println("Your favourite author is: ", name)
	// fmt.Println(reflect.TypeOf(name))
	fmt.Print("What is a book written by this author? ")
	favAuthor, _ := reader.ReadString('\n') //includes newline character
	fmt.Println("A book written by your favourite author is: ", favAuthor)

	var yearPublished int

	fmt.Print("When was the book published? ")
	fmt.Scanln(&yearPublished)
	fmt.Println("The book from your favourite author was published ", yearPublished)
}
