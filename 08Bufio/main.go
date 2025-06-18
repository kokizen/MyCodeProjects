package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func greet(name string) string {
	return "Hello, " + name + "!"
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your name?")
	nameInput, _ := reader.ReadString('\n')
	nameInput = strings.TrimSpace(nameInput)
	message := greet(nameInput)
	fmt.Println(message)
}
