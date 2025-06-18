package main

import "fmt"

func main() {
	//a for-loop in go
	for index := 0; index < 21; index++ {
		if index%2 != 0 {
			fmt.Println(index)
		}
	}
}

// func main() {
// 	for index := "A"; index < "Z"; index++ {
// 		fmt.Println(index)
// 	}
// }
