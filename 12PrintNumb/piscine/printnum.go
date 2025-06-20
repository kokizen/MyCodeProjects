package piscine

import "github.com/01-edu/z01"

func Printnum(num int) {
	if num == 0 {
		z01.PrintRune('0')
	}
	if num < 0 {
		z01.PrintRune('-')
		num = -num
	}
	digits := []rune{}
	for num > 0 {
		digit := rune(num%10 + '0')
		digits = append([]rune{digit}, digits...)
		num /= 10
	}
	for _, digit := range digits {
		z01.PrintRune(digit)
	}
}
