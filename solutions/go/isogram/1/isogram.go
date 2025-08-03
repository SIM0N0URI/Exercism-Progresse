package isogram


import 
	"strings"

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	for i := 0; i < len(word); i++ {
		if !(word[i] >= 'a' && word[i] <= 'z') {
			continue
		}
		for j := i + 1; j < len(word); j++ {
			if !(word[j] >= 'a' && word[j] <= 'z') {
				continue
			}
			if word[i] == word[j] {
				return false
			}
		}
	}
	return true
}

