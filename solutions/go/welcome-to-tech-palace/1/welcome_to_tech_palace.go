package techpalace

import (
	"regexp"
	"strings"
)

const name = "Simo"
const num int = 10

func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

func AddBorder(welcomeMsg string, n int) string {
	stars := strings.Repeat("*", n)
	return stars + "\n" + welcomeMsg + "\n" + stars
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	re := regexp.MustCompile(`[*+\n+]`)
	cleaned := re.ReplaceAllString(oldMsg, "")
	return strings.TrimSpace(cleaned)
}
