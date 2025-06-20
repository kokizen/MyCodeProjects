package piscine

import "github.com/01-edu/z01"

func PrintAlpha() {
	for i := 'a'; i <= 'z'; i++ {
		z01.PrintRune(i)
	}
}
