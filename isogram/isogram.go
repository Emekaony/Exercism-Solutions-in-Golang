package isogram

import "strings"

func IsIsogram(word string) bool {
	mp := make(map[string]int)
	isValid := true
	for i := 1; i <= len(word); i++ {
		// found a nice hack to getting substring
		curr := strings.ToLower(word[i-1 : i])
		_, exists := mp[curr]
		if curr == "-" || curr == " " {
			continue
		} else if exists {
			isValid = false
			break
		} else {
			mp[curr] = 1
		}
	}
	return isValid
}
