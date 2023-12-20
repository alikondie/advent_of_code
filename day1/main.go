package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var digits = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	var firstDigit string
	var secondDigit string
	totalCount := 0
	startIndex := 0
	for index, char := range input {
		for key, digit := range digits {
			if strings.Contains(input[startIndex:index], key) {
				if firstDigit == "" {
					firstDigit = fmt.Sprint(digit)
				}
				secondDigit = fmt.Sprint(digit)
				startIndex = index - 1
			}
		}

		if unicode.IsDigit(char) {
			if firstDigit == "" {
				firstDigit = string(char)
			}
			secondDigit = string(char)
		}
		if char == '\n' {
			number, _ := strconv.Atoi(firstDigit + secondDigit)
			totalCount += number
			firstDigit = ""
			secondDigit = ""
			startIndex = index - 1
		}
	}
	fmt.Println(totalCount)
}
