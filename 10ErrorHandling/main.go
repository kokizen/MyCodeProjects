package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("example.txt")
	if err != nil {
		// File might not exist, permission denied, etc.
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("File content:", string(data))
}
