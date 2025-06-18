package main

import "fmt"

func main() {
	word := "Hello Coders"
	for i, r := range word {
		fmt.Printf("Character %d: %c = %d\n", i, r, r)
	}
	fmt.Println(word)
}
