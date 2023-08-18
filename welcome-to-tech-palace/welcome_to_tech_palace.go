package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	var message string = "Welcome to the Tech Palace, " + strings.ToUpper(customer)
	return message
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	repeatedStars := strings.Repeat("*", numStarsPerLine)
	return fmt.Sprintf("%s\n%s\n%s", repeatedStars, welcomeMsg, repeatedStars)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	// we want just the message with no leading or trailing white space!
	return strings.TrimSpace(strings.Replace(oldMsg, "*", "", -1))
}
