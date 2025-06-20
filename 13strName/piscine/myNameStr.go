package piscine

import "github.com/01-edu/z01"

func PrintMyName(name string) {
	for _, ch := range name {
		z01.PrintRune(ch)
		//z01.PrintRune('\n')
	}
}
