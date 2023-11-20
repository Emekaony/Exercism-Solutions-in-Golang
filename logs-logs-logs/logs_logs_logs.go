package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {

	for _, char := range log {
		if char == '\U00002757' {
			return "recommendation"
		} else if char == '\U00002600' {
			return "weather"
		} else if char == '\U0001F50D' {
			return "search"
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	// convert the string to a slice of runes
	runes := []rune(log)
	// go thru the slice and find occurencies of old rune, replace them!
	for i, j := range runes {
		if j == oldRune {
			runes[i] = newRune
		}
	}
	// convert slice back to string and return it
	return string(runes)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	// functionality available in unicode/utf8
	// use the RuneCountInString function to get the actual rune length (character limit)
	return utf8.RuneCountInString(log) <= limit
}
