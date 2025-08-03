package pangram

import "strings"

func IsPangram(input string) bool {
    lowinput := strings.ToLower(input)
	seen := make(map[rune]bool)
    for _, ch := range lowinput {
        if ch >= 'a' && ch <= 'z' {
            seen[ch] = true
        }
    }
    return len(seen) == 26
}
