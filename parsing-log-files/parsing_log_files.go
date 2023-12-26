package parsinglogfiles

import (
	"fmt"
	"regexp"
)

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
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	// had to make a new string slice because original one wasn't mutating
	new_lines := []string{}
	re := regexp.MustCompile(`User \s*\b[a-zA-Z]+\d+\b`)
	u_regexp := regexp.MustCompile(`\b[a-zA-Z]+\d+\b`)
	for _, line := range lines {
		if re.MatchString(line) {
			username := u_regexp.FindString(line)
			new_str := fmt.Sprintf("[USR] %s ", username)
			line = new_str + line
			fmt.Println("Line is: ", line)
			new_lines = append(new_lines, line)
		} else {
			new_lines = append(new_lines, line)
		}
	}
	return new_lines
}
