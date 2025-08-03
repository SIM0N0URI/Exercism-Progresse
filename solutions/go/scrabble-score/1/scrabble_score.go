package scrabble

import "strings"
func Score(word string) int {
	count := 0
	word = strings.ToUpper(word)
	for _, char := range word {
		if char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' || char == 'L' || char == 'N' || char == 'R' || char == 'S' || char == 'T' {
			count += 1
		}
		if char == 'D' || char == 'G' {
			count += 2
		}
		if char == 'B' || char == 'C' || char == 'M' || char == 'P' {
			count += 3
		}
		if char == 'F' || char == 'H' || char == 'V' || char == 'W' || char == 'Y' {
			count += 4
		}
		if char == 'K' {
			count += 5
		}
		if char == 'J' || char == 'X' {
			count += 8
		}
		if char == 'Q' || char == 'Z' {
			count += 10
		}
	}
	return count
}
