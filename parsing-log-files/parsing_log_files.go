package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	// the .* at the end is key. This matches any character
	re := regexp.MustCompile(`^(\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\]|) .*`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`.*".*(?i)\bpassword\b.*".*`)
	count := 0
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	panic("Please implement the RemoveEndOfLineText function")
}

func TagWithUserName(lines []string) []string {
	panic("Please implement the TagWithUserName function")
}
